package schema

import (
	"time"

	"github.com/one-meta/meta/app/ent/privacy"
	"github.com/one-meta/meta/pkg/ent/rule"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

// Policy defines the privacy policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			// Deny any operation in case there is no "viewer context".
			rule.DenyIfNoViewer(),
			// Allow admins to query any information.
			rule.AllowIfAdmin(),
		},
		Mutation: privacy.MutationPolicy{
			// Deny any operation in case there is no "viewer context".
			rule.DenyIfNoViewer(),
			// Allow admins to mutate any information.
			rule.AllowIfAdmin(),
		},
	}
}

// TenantMixin for embedding the tenant info in different schemas.
type TenantMixin struct {
	mixin.Schema
}

// Fields for all schemas that embed TenantMixin.
func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id").Optional().Comment("租户Id，可选").Nillable(),
	}
}

// Edges for all schemas that embed TenantMixin.
func (TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant", Tenant.Type).
			Field("tenant_id").
			Unique(),
	}
}

// Policy for all schemas that embed TenantMixin.
func (TenantMixin) Policy() ent.Policy {
	return rule.FilterTenantRule()
}

// TimeMixin 创建时间、更新时间
type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// CreatedTimeMixin 创建时间
type CreatedTimeMixin struct {
	mixin.Schema
}

func (CreatedTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// UpdatedTimeMixin 更新时间
type UpdatedTimeMixin struct {
	mixin.Schema
}

func (UpdatedTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// RemarkMixin 备注
type RemarkMixin struct {
	mixin.Schema
}

func (RemarkMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("remark").Comment("备注").MaxLen(500).Optional().Nillable(),
	}
}

// ValidMixin 有效
type ValidMixin struct {
	mixin.Schema
}

func (ValidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("valid").Default(true).Comment("有效"),
	}
}
