package generator

import (
	"github.com/one-meta/meta/app/entity"
	"github.com/one-meta/meta/app/entity/config"
	"github.com/one-meta/meta/pkg/auth"
	"github.com/one-meta/meta/pkg/common"
	"github.com/one-meta/meta/pkg/jwt"

	"github.com/google/wire"

	"github.com/gofiber/fiber/v2"
)

var Set = wire.NewSet(wire.Struct(new(Gen), "*"))

type Gen struct {
	AuthEnt *auth.Entx
}

func (a *Gen) IntId() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return common.NewResult(c, err)
		}
		c.Locals("id", id)
		return c.Next()
	}
}
func (a *Gen) AuthCtx() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := a.AuthEnt.GetContext(c)
		if ctx == nil {
			return common.NewErrorWithStatusCode(c, "Access denied", fiber.StatusForbidden)
		}
		c.Locals("ctx", ctx)
		return c.Next()
	}
}
func (a *Gen) SaCtx() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := a.AuthEnt.GetSAContext()
		c.Locals("ctx", ctx)
		return c.Next()
	}
}

func (a *Gen) SaCasbin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := "all"
		casbinUser := &entity.CasbinUser{
			UserName: config.CFG.Stage.User,
			Project:  project,
		}
		c.Locals("casbinUser", casbinUser)
		return c.Next()
	}
}
func (a *Gen) SaJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtConfig := config.CFG.Auth.JWT
		authorization, _ := jwt.CreateJWT("", config.CFG.Stage.User, jwtConfig.Key, jwtConfig.TTL, jwtConfig.Hmac)
		c.Locals("Authorization", authorization)
		return c.Next()
	}
}
