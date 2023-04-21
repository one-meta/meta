package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User schema.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ValidMixin{},
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		// 增加 query的StructTag，解决url中的参数带下划线时（两个单词带下划线）无法反序列化,
		field.Bool("super_admin").Default(false).Comment("超级管理员，可操作所有租户数据").StructTag(`query:"super_admin,omitempty"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	// Privacy policy defined in the BaseMixin and TenantMixin.
	return nil
}
