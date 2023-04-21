package middleware

import (
	"github.com/google/wire"
	"github.com/one-meta/meta/pkg/middleware/generator"
)

var Set = wire.NewSet(
	CasbinxSet,
	JWTSet,
	generator.Set,
)
