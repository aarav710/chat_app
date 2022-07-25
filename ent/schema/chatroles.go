package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatRoles holds the schema definition for the ChatRoles entity.
type ChatRoles struct {
	ent.Schema
}

// Fields of the ChatRoles.
func (ChatRoles) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").Values("ROLE_ADMIN", "ROLE_PARTICIPANT"),
	}
}

// Edges of the ChatRoles.
func (ChatRoles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chat", Chat.Type).Ref("chat_roles").Unique(),
		edge.From("user", User.Type).Ref("roles_in_chats").Unique(),
	}
}
