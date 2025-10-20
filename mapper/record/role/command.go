package rolerecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// roleCommandMapper provides methods to map Role database rows to RoleRecord domain models.
type roleCommandMapper struct {
}

// NewRoleRecordMapper returns a new instance of roleCommandMapper which provides methods to map Role database rows to RoleRecord domain models.
func NewRoleCommandRecordMapping() *roleCommandMapper {
	return &roleCommandMapper{}
}

// ToRoleRecord maps a Role database row to a RoleRecord domain model.
//
// Parameters:
//   - role: A pointer to a Role representing the database row.
//
// Returns:
//   - A pointer to a RoleRecord containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleCommandMapper) ToRoleRecord(role *db.Role) *record.RoleRecord {
	deletedAt := role.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.RoleRecord{
		ID:        int(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}
