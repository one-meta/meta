package register

import (
	"fmt"
	"math/rand"
	"github.com/one-meta/meta/app/ent"
	"github.com/one-meta/meta/app/ent/systemapi"
	"github.com/one-meta/meta/app/ent/tenant"
	"github.com/one-meta/meta/app/entity/config"
	"github.com/one-meta/meta/app/service"
	"github.com/one-meta/meta/pkg/auth"
	"github.com/one-meta/meta/pkg/checker"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sethvargo/go-password/password"
	"go.uber.org/zap"
)

type Runner struct {
	Logger            *zap.Logger
	AuthEnt           *auth.Entx
	UserService       *service.UserService
	TenantService     *service.TenantService
	CasbinRuleService *service.CasbinRuleService
	SystemApiService  *service.SystemApiService
}

func (r *Runner) ScheduleTask1() {
	r.Logger.Sugar().Info("running schedule")
	fmt.Println("running schedule")
}

func (r *Runner) RegisterSAUser() {
	ctx := r.AuthEnt.GetSAContext()
	userName := config.CFG.Stage.User
	pwd := config.CFG.Stage.Password
	if userName == "" {
		userName = "github.com/one-meta/meta"
	}
	if pwd == "" {
		// Generate a pwd that is 16 characters long with 5 digits, 5 symbols,
		// allowing upper and lower case letters, disallowing repeat characters.
		res, err := password.Generate(16, 5, 5, false, false)
		if err != nil {
			pwd = fmt.Sprintf("%d%v", 20, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100))
		}
		pwd = res
	}
	_, create := checker.CheckAndInitUser(r.UserService, ctx, userName, pwd, true)
	if create {
		fmt.Printf("Create sa user %s:%s\n", userName, pwd)
	}
	// 设置默认项目
	var project *ent.Tenant
	queryTenant, _ := r.TenantService.Dao.Tenant.Query().Where(tenant.Name("all"), tenant.Code("all")).Only(ctx)
	if queryTenant == nil {
		project = &ent.Tenant{
			Name: "all",
			Code: "all",
		}
		r.TenantService.Create(ctx, project)
	} else {
		project = queryTenant
	}
	// 设置casbin权限
	// 角色，增加all项目的查询权限（为了前端能够显示项目而已
	checker.CheckAndInitRole(r.CasbinRuleService, userName, "query", "all", ctx)
	// 资源，增加 /* 的查询权限（其实也没必要
	checker.CheckAndInitPermission(r.CasbinRuleService, "query", "all", "/*", "GET", ctx)
}

func (r *Runner) RegisterSystemApi(routes []fiber.Route) {
	ctx := r.AuthEnt.GetSAContext()
	for _, v := range routes {
		method := v.Method
		path := v.Path
		if path == "/" || path == "/api" {
			continue
		}
		var name string
		split := strings.Split(path, "/")
		if len(split) > 3 {
			name = split[3]
		} else if len(split) > 1 {
			name = split[1]
		} else {
			name = path
		}

		queryApi, _ := r.SystemApiService.Dao.SystemApi.Query().Where(systemapi.Name(name), systemapi.Path(path), systemapi.HTTPMethod(method)).Only(ctx)
		publicApiFlag := false
		var roles []string
		publicGetPath := config.CFG.Stage.Api.PublicGetPath
		if method == "GET" {
			for _, v := range publicGetPath {
				if path == v {
					publicApiFlag = true
				}
			}
		}
		// 查询，查看，query,view
		if method == "GET" && !strings.HasSuffix(path, "/:id") {
			roles = []string{"query", "view"}
		}
		// 查看详情，viewDetail
		if method == "GET" && strings.HasSuffix(path, "/:id") {
			roles = []string{"viewDetail"}
		}
		// 编辑，edit
		if method == "PUT" && strings.HasSuffix(path, "/:id") {
			roles = []string{"edit"}
		}
		//删除，delete
		//&& strings.HasSuffix(path, "/:id")
		if method == "DELETE" {
			roles = []string{"delete"}
		}
		// 批量删除，bulkDelete
		if method == "POST" && strings.HasSuffix(path, "/bulk/delete") {
			roles = []string{"bulkDelete"}
		}
		// 新建，new
		if method == "POST" && !strings.HasSuffix(path, "/bulk") {
			roles = []string{"new"}
		}
		// 批量创建，bulkCreate
		if method == "POST" && strings.HasSuffix(path, "/bulk") {
			roles = []string{"bulkCreate"}
		}

		// 区分sa路由
		saFlag := false
		saPathPrefix := config.CFG.Stage.Api.SaPathPrefix
		for _, v := range saPathPrefix {
			if strings.HasPrefix(path, v) {
				if !publicApiFlag {
					saFlag = true
					break
				}
			}
		}

		if queryApi == nil {
			createApi := &ent.SystemApi{
				Name: name,
				Path: path,
				HTTPMethod: method,
				Public:     publicApiFlag,
				Roles:      &roles,
				Sa:         saFlag,
			}
			r.SystemApiService.Create(ctx, createApi)
		}
	}
}
