package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Login holds the schema definition for the Login entity.
type Login struct {
	ent.Schema
}

// Fields of the Login.
func (Login) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("uuid").Unique(),
		field.Time("created_at").Default(time.Now),
		field.Enum("status").Values("USER", "INCOMPLETE_REGISTRATION"),
	}
}

// Edges of the Login.
func (Login) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
}
}
