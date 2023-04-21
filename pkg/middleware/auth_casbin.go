package middleware

import (
	"github.com/one-meta/meta/app/ent/user"
	"github.com/one-meta/meta/app/entity"
	"github.com/one-meta/meta/app/service"
	"github.com/one-meta/meta/pkg/auth"
	"github.com/one-meta/meta/pkg/common"
	"strings"

	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

var CasbinxSet = wire.NewSet(wire.Struct(new(Casbinx), "*"))

type Casbinx struct {
	UserService *service.UserService
	Logger      *zap.Logger
	AuthEnt     *auth.Entx
	Enf         *casbin.Enforcer
}

func (b *Casbinx) AuthCasbin(enf *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		obj := string(c.Request().RequestURI())
		act := c.Method()
		casbinUser := c.Locals("casbinUser").(*entity.CasbinUser)

		project := casbinUser.Project
		username := casbinUser.UserName

		//sa用户不需要经过casbin
		adminCtx := b.AuthEnt.GetSAContext()
		u, err := b.UserService.Dao.User.Query().Where(user.Name(username)).Only(adminCtx)
		if u == nil {
			return common.NewErrorWithStatusCode(c, "User invalid", fiber.StatusForbidden)
		}
		// sa用户且有效
		if u.Valid && u.SuperAdmin {
			return c.Next()
		}

		//未带Auth-Project头，从casbin表中获取一个项目给用户
		//获取当前用户的域（租户/项目）
		if project == "" {
			userDomains, _ := b.Enf.GetDomainsForUser(username)
			if len(userDomains) > 0 {
				project = userDomains[0]
				casbinUser := &entity.CasbinUser{
					UserName: username,
					Project:  project,
				}
				c.Locals("casbinUser", casbinUser)
			}
		}
		//带?查询链接转换
		if strings.Contains(obj, "?") {
			obj = strings.Split(obj, "?")[0]
		}

		ok, _ := enf.Enforce(username, project, obj, act)

		if err != nil || !u.Valid {
			if err == nil {
				b.Logger.Sugar().Errorf("User: %s is invalid", username)
			} else {
				b.Logger.Sugar().Error(err)
			}
			return common.NewErrorWithStatusCode(c, "User invalid", fiber.StatusForbidden)
		}

		if ok {
			return c.Next()
		} else {
			return common.NewErrorWithStatusCode(c, "Access denied", fiber.StatusForbidden)
		}
	}
}
