package auth

import (
	"context"
	"github.com/one-meta/meta/app/entity/config"

	"github.com/one-meta/meta/app/ent/tenant"
	"github.com/one-meta/meta/app/ent/user"
	"github.com/one-meta/meta/app/entity"
	"github.com/one-meta/meta/app/service"
	"github.com/one-meta/meta/pkg/ent/viewer"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var Set = wire.NewSet(wire.Struct(new(Entx), "*"))

type Entx struct {
	UserService   *service.UserService
	TenantService *service.TenantService
	Logger        *zap.Logger
}

// GetSAContext SA视图，可以查看所有项目数据，用于public路由访问有视图权限的表
func (b *Entx) GetSAContext() context.Context {
	return viewer.NewContext(context.Background(), viewer.UserViewer{Role: viewer.Admin})
}
func (b *Entx) GetContext(c *fiber.Ctx) context.Context {
	ctx := context.Background()
	//没有启用认证和授权直接返回一个SA context
	if !config.CFG.Auth.Enable && config.CFG.Stage.Status == "dev" {
		return b.GetSAContext()
	}
	domain, username := getUserFromCtx(c)
	if username == "" {
		return nil
	}
	adminCtx := b.GetSAContext()
	u, err := b.UserService.Dao.User.Query().Where(user.Name(username)).Only(adminCtx)
	if err != nil {
		b.Logger.Sugar().Error(err)
		return nil
	}
	// sa用户
	if u.SuperAdmin {
		ctx = viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	} else {
		//普通租户用户
		tenantUser, err := b.TenantService.Dao.Tenant.Query().Where(tenant.Name(domain)).Only(adminCtx)
		if err != nil {
			b.Logger.Sugar().Error(err)
		}
		ctx = viewer.NewContext(ctx, viewer.UserViewer{T: tenantUser})
	}
	return ctx
}

func getUserFromCtx(c *fiber.Ctx) (string, string) {
	var domain, username string
	if c.Locals("casbinUser") != nil {
		casbinUser := c.Locals("casbinUser").(*entity.CasbinUser)
		username = casbinUser.UserName
		domain = c.Get("Auth-Domain", casbinUser.Project)
	}
	return domain, username
}
