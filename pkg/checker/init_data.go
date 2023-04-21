package checker

import (
	"context"
	"github.com/one-meta/meta/app/ent"
	"github.com/one-meta/meta/app/ent/casbinrule"
	"github.com/one-meta/meta/app/service"
)

func CheckAndInitUser(r *service.UserService, ctx context.Context, userName string, pwd string, saUser bool) (*ent.User, bool) {
	queryUser, _ := r.QueryByUserName(ctx, userName)
	if queryUser == nil {
		u := &ent.User{
			Valid:      true,
			Name:       userName,
			Password:   pwd,
			SuperAdmin: saUser,
		}
		r.Create(ctx, u)
		return u, true
	}
	return queryUser, false
}

func CheckAndInitPermission(r *service.CasbinRuleService, role, project, recourse, httpMethod string, ctx context.Context) {
	queryRulePermission, _ := r.Dao.CasbinRule.Query().Where(casbinrule.Type("p"), casbinrule.Sub(role), casbinrule.Dom(project), casbinrule.Obj(recourse), casbinrule.Act(httpMethod)).Only(ctx)
	if queryRulePermission == nil {
		rule := &ent.CasbinRule{
			Type: "p",
			Sub:  role,
			Dom:  project,
			Obj:  recourse,
			Act:  httpMethod,
		}
		r.Create(ctx, rule)
	}
}

func CheckAndInitRole(r *service.CasbinRuleService, userName string, role, project string, ctx context.Context) {
	queryRuleRole, _ := r.Dao.CasbinRule.Query().Where(casbinrule.Type("g"), casbinrule.Sub(userName), casbinrule.Dom(role), casbinrule.Obj(project)).Only(ctx)
	if queryRuleRole == nil {
		rule := &ent.CasbinRule{
			Type: "g",
			Sub:  userName,
			Dom:  role,
			Obj:  project,
		}
		r.Create(ctx, rule)
	}
}
