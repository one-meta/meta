package middleware

import (
	"github.com/one-meta/meta/app/entity"
	"github.com/one-meta/meta/app/entity/config"
	"github.com/one-meta/meta/pkg/common"
	"github.com/one-meta/meta/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var JWTSet = wire.NewSet(wire.Struct(new(JWT), "*"))

type JWT struct {
	Logger *zap.Logger
}

func (a *JWT) AuthJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return common.NewErrorWithStatusCode(c, "Permission denied", fiber.StatusUnauthorized)
		}
		claims, err := jwt.ParseJWT(authorization, config.CFG.Auth.JWT.Key)
		if err != nil {
			a.Logger.Sugar().Error(err)
			return common.NewErrorWithStatusCode(c, err.Error(), fiber.StatusUnauthorized)
		}
		project := c.Get("Auth-Project")
		casbinUser := &entity.CasbinUser{
			UserName: claims.Username,
			Project:  project,
		}
		c.Locals("casbinUser", casbinUser)
		return c.Next()
	}
}
