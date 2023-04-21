package register

import "github.com/google/wire"

var (
	Set         = wire.NewSet(RegisterSet)
	RegisterSet = wire.NewSet(wire.Struct(new(Runner), "*"))
)
