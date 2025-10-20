package userrolerecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// userRoleRecordMapper provides methods to map UserRole database rows to UserRoleRecord domain models.
type userRoleRecordMapper struct {
}

// NewUserRoleRecordMapper creates a new UserRoleRecordMapper instance.
func NewUserRoleRecordMapper() UserRoleRecordMapping {
	return &userRoleRecordMapper{}
}

// ToUserRoleRecord maps a UserRole database row to a UserRoleRecord domain model.
//
// Parameters:
//   - user: A pointer to a UserRole representing the database row.
//
// Returns:
//   - A pointer to a UserRoleRecord containing the mapped data, including
//     UserRoleID, UserID, RoleID, CreatedAt, and UpdatedAt.
func (t *userRoleRecordMapper) ToUserRoleRecord(userRole *db.UserRole) *record.UserRoleRecord {
	return &record.UserRoleRecord{
		UserRoleID: int32(userRole.UserRoleID),
		UserID:     int32(userRole.UserID),
		RoleID:     int32(userRole.RoleID),
		CreatedAt:  userRole.CreatedAt.Time,
		UpdatedAt:  userRole.UpdatedAt.Time,
	}
}
