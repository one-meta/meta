package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/one-meta/meta/app/ent/privacy"
	"github.com/one-meta/meta/pkg/ent/rule"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Mixin of the Tenant schema.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().Unique(),
		field.String("code").NotEmpty().Unique(),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	// return []ent.Edge{
	// 	edge.To("users", User.Type),
	// }
	return nil
}

// Policy defines the privacy policy of the User.
func (Tenant) Policy() ent.Policy {
	return privacy.Policy{
		//仅管理员可写，其他用户可读
		Mutation: privacy.MutationPolicy{
			// For Tenant type, we only allow admin users to mutate
			// the tenant information and deny otherwise.
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
