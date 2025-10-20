package record

import "time"

// UserRoleRecord represents the association between a user and a role.
// This is a join table that defines which roles are assigned to which users.
type UserRoleRecord struct {
	UserRoleID int32      `json:"user_role_id"`         // Unique identifier for the user-role association
	UserID     int32      `json:"user_id"`              // ID of the user being assigned a role
	RoleID     int32      `json:"role_id"`              // ID of the role being assigned
	RoleName   string     `json:"role_name,omitempty"`  // Name of the role (optional, may be populated in joins)
	CreatedAt  time.Time  `json:"created_at"`           // Timestamp when the association was created
	UpdatedAt  time.Time  `json:"updated_at"`           // Timestamp when the association was last updated
	DeletedAt  *time.Time `json:"deleted_at,omitempty"` // Timestamp when association was removed (nil if active)
}
