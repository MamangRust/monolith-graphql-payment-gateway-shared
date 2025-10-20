package record

// RoleRecord represents a role in the system that can be assigned to users.
// Roles are used to define permissions and access levels for different user types.
type RoleRecord struct {
	ID        int     `json:"id"`         // Unique identifier for the role
	Name      string  `json:"name"`       // Name of the role (e.g., "admin", "user", "manager")
	CreatedAt string  `json:"created_at"` // Timestamp when the role was created
	UpdatedAt string  `json:"updated_at"` // Timestamp when the role was last updated
	DeletedAt *string `json:"deleted_at"` // Timestamp when the role was soft-deleted (nil if active)
}
