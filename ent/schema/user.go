package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("bio").Unique(),
  }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("login", Login.Type).Ref("user").Unique().Required(),
    edge.To("messages", Message.Type),
		edge.From("chats", Chat.Type).Ref("users"),
		edge.To("roles_in_chats", ChatRoles.Type),
	}
}
