package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Time("created_at").Default(time.Now),
		field.String("description"),
		field.String("display_picture_url"),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("users", User.Type),
		edge.To("chat_roles", ChatRoles.Type),
		edge.To("messages", Message.Type),
	}
}
