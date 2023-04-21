package pkg

import (
	"github.com/google/wire"
	"github.com/one-meta/meta/pkg/auth"
	"github.com/one-meta/meta/pkg/middleware"
	"github.com/one-meta/meta/pkg/register"
)

var Set = wire.NewSet(
	auth.Set,
	middleware.Set,
	register.Set,
)
