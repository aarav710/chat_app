// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chatapp/backend/ent/chat"
	"chatapp/backend/ent/chatroles"
	"chatapp/backend/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ChatRoles is the model entity for the ChatRoles schema.
type ChatRoles struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role chatroles.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChatRolesQuery when eager-loading is set.
	Edges               ChatRolesEdges `json:"edges"`
	chat_chat_roles     *int
	user_roles_in_chats *int
}

// ChatRolesEdges holds the relations/edges for other nodes in the graph.
type ChatRolesEdges struct {
	// Chat holds the value of the chat edge.
	Chat *Chat `json:"chat,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChatOrErr returns the Chat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatRolesEdges) ChatOrErr() (*Chat, error) {
	if e.loadedTypes[0] {
		if e.Chat == nil {
			// The edge chat was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: chat.Label}
		}
		return e.Chat, nil
	}
	return nil, &NotLoadedError{edge: "chat"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatRolesEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChatRoles) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case chatroles.FieldID:
			values[i] = new(sql.NullInt64)
		case chatroles.FieldRole:
			values[i] = new(sql.NullString)
		case chatroles.ForeignKeys[0]: // chat_chat_roles
			values[i] = new(sql.NullInt64)
		case chatroles.ForeignKeys[1]: // user_roles_in_chats
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChatRoles", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChatRoles fields.
func (cr *ChatRoles) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chatroles.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cr.ID = int(value.Int64)
		case chatroles.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				cr.Role = chatroles.Role(value.String)
			}
		case chatroles.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field chat_chat_roles", value)
			} else if value.Valid {
				cr.chat_chat_roles = new(int)
				*cr.chat_chat_roles = int(value.Int64)
			}
		case chatroles.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_roles_in_chats", value)
			} else if value.Valid {
				cr.user_roles_in_chats = new(int)
				*cr.user_roles_in_chats = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryChat queries the "chat" edge of the ChatRoles entity.
func (cr *ChatRoles) QueryChat() *ChatQuery {
	return (&ChatRolesClient{config: cr.config}).QueryChat(cr)
}

// QueryUser queries the "user" edge of the ChatRoles entity.
func (cr *ChatRoles) QueryUser() *UserQuery {
	return (&ChatRolesClient{config: cr.config}).QueryUser(cr)
}

// Update returns a builder for updating this ChatRoles.
// Note that you need to call ChatRoles.Unwrap() before calling this method if this ChatRoles
// was returned from a transaction, and the transaction was committed or rolled back.
func (cr *ChatRoles) Update() *ChatRolesUpdateOne {
	return (&ChatRolesClient{config: cr.config}).UpdateOne(cr)
}

// Unwrap unwraps the ChatRoles entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cr *ChatRoles) Unwrap() *ChatRoles {
	_tx, ok := cr.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChatRoles is not a transactional entity")
	}
	cr.config.driver = _tx.drv
	return cr
}

// String implements the fmt.Stringer.
func (cr *ChatRoles) String() string {
	var builder strings.Builder
	builder.WriteString("ChatRoles(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cr.ID))
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", cr.Role))
	builder.WriteByte(')')
	return builder.String()
}

// ChatRolesSlice is a parsable slice of ChatRoles.
type ChatRolesSlice []*ChatRoles

func (cr ChatRolesSlice) config(cfg config) {
	for _i := range cr {
		cr[_i].config = cfg
	}
}