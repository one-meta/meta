package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SystemApi holds the schema definition for the SystemApi entity.
type SystemApi struct {
	ent.Schema
}

// Mixin of the SystemApi.
func (SystemApi) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TimeMixin{},
		RemarkMixin{},
	}
}

// Fields of the SystemApi.
func (SystemApi) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("path").Comment("路径"),
		field.String("http_method").Comment("http方法").StructTag(`query:"http_method,omitempty"`),
		field.JSON("roles", &[]string{}).Comment("角色"),
		field.Bool("public").Comment("是否公共接口，所有用户可以访问"),
		field.Bool("sa").Comment("是否sa接口，sa接口仅sa用户可操作（如果public，则普通用户可以访问(get）"),
	}
}

// Edges of the SystemApi.
func (SystemApi) Edges() []ent.Edge {
	return nil
}
