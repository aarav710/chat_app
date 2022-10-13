// Code generated by ent, DO NOT EDIT.

package chat

import (
	"time"
)

const (
	// Label holds the string label denoting the chat type in the database.
	Label = "chat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDisplayPictureURL holds the string denoting the display_picture_url field in the database.
	FieldDisplayPictureURL = "display_picture_url"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeChatRoles holds the string denoting the chat_roles edge name in mutations.
	EdgeChatRoles = "chat_roles"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the chat in the database.
	Table = "chats"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "chat_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// ChatRolesTable is the table that holds the chat_roles relation/edge.
	ChatRolesTable = "chat_roles"
	// ChatRolesInverseTable is the table name for the ChatRoles entity.
	// It exists in this package in order to avoid circular dependency with the "chatroles" package.
	ChatRolesInverseTable = "chat_roles"
	// ChatRolesColumn is the table column denoting the chat_roles relation/edge.
	ChatRolesColumn = "chat_chat_roles"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "chat_messages"
)

// Columns holds all SQL columns for chat fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldCreatedAt,
	FieldDescription,
	FieldDisplayPictureURL,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"chat_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
