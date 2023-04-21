package main

import (
	"github.com/one-meta/meta/app/entity/config"
	"github.com/one-meta/meta/pkg/util"
	"log"

	_ "github.com/one-meta/meta/app/ent/runtime"
	_ "github.com/one-meta/meta/docs" // load API Docs files (Swagger)
)

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as abow.

// @title API
// @version 0.1
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := util.LoadConfig("resource")
	if err != nil {
		log.Fatal(err)
	}
	fiberApp, router, register, f := util.InjectApp()
	defer f()

	util.RegisterRouter(router, fiberApp)
	//获取api 路径
	routes := fiberApp.GetRoutes()
	go register.Inject(routes)

	// 启动服务
	if config.CFG.Stage.Status == "dev" {
		util.StartDevServer(fiberApp)
	} else {
		util.StartProdServer(fiberApp)
	}
}
