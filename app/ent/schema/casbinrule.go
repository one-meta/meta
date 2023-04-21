package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CasbinRule holds the schema definition for the CasbinRule entity.
type CasbinRule struct {
	ent.Schema
}

// Mixin of the CasbinRule schema.
func (CasbinRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		//TimeMixin{},
	}
}

// Fields of the CasbinRule.
func (CasbinRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").StorageKey("ptype").NotEmpty(),
		field.String("sub").StorageKey("v0").NotEmpty(),
		field.String("dom").StorageKey("v1").NotEmpty(),
		field.String("obj").StorageKey("v2").NotEmpty(),
		field.String("act").StorageKey("v3"),
		field.String("v4").StorageKey("v4").Sensitive(),
		field.String("v5").StorageKey("v5").Sensitive(),
	}
}

// Edges of the CasbinRule.
func (CasbinRule) Edges() []ent.Edge {
	return nil
}

// Policy defines the privacy policy of the CasbinRule.
func (CasbinRule) Policy() ent.Policy {
	// Privacy policy defined in the BaseMixin.
	return nil
}
