package wireset

import (
	"github.com/google/wire"
	v1 "github.com/one-meta/meta/api/v1"
	"github.com/one-meta/meta/pkg/register"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Router *v1.Router
	RegisterInjector *register.Runner
}
