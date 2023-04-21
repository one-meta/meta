package util

import (
	"github.com/gofiber/fiber/v2"
	"log"
	v1 "github.com/one-meta/meta/api/v1"
	"github.com/one-meta/meta/app"
	"github.com/one-meta/meta/app/entity/config"
	"github.com/one-meta/meta/pkg/middleware"
	"github.com/one-meta/meta/pkg/register"
	"os"
	"os/signal"
)

func RegisterRouter(router *v1.Router, fiberApp *fiber.App) {
	//Swagger api
	if config.CFG.Swagger.Enable {
		router.Swagger(fiberApp)
	}
	//中间件
	middleware.Middleware(fiberApp)
	//公共路由
	router.Public(fiberApp)
	//私有路由
	router.Private(fiberApp)
	//404
	router.NotFound(fiberApp)
}

func InjectApp() (*fiber.App, *v1.Router, *register.Runner, func()) {
	fiberApp := fiber.New(config.CFG.Fiber.NewConfig())

	//依赖注入
	injector, f, err := app.Init()
	if err != nil {
		log.Fatal(err)
	}
	// Routes.
	router := injector.Router
	registerInjector := injector.RegisterInjector
	return fiberApp, router, registerInjector, f
}

func LoadConfig(path string) error {
	return config.LoadConfig(path)
}

// StartProdServer 生产，优雅关闭
func StartProdServer(app *fiber.App) {
	done := make(chan struct{})
	quit := make(chan os.Signal, 1)
	//os.Interrupt: syscall.SIGINT
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Server is shutting down...")
		if err := app.Shutdown(); err != nil {
			log.Printf("Could not gracefully shutdown the server: %v", err)
		}

		close(done)
	}()
	ListenServer(app)

	<-done
	log.Println("Done.")
}

// StartDevServer 开发测试
func StartDevServer(app *fiber.App) {
	ListenServer(app)
}

func ListenServer(app *fiber.App) {
	if err := app.Listen(config.CFG.Fiber.Url()); err != nil {
		log.Printf("Server start failed: %v", err)
	}
}
