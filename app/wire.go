//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	router "github.com/one-meta/meta/api/v1"
	"github.com/one-meta/meta/app/controller"
	"github.com/one-meta/meta/app/service"
	"github.com/one-meta/meta/app/wireset"
	"github.com/one-meta/meta/pkg"
)

func Init() (*wireset.Injector, func(), error) {
	// 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
	panic(
		wire.Build(
			wireset.InitLog,
			wireset.CacheProvider,
			wireset.InitEnt,
			wireset.InitCasbin,
			service.Set,
			controller.Set,
			router.Set,
			pkg.Set,
			wireset.InjectorSet,
		),
	)
}
