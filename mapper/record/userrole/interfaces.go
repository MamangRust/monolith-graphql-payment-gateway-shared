package userrolerecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// UserRoleRecordMapping defines a mapping function from a UserRole database row to a UserRoleRecord domain model.
type UserRoleRecordMapping interface {
	// ToUserRoleRecord maps a UserRole database row to a UserRoleRecord domain model.
	//
	// Parameters:
	//   - user: A pointer to a UserRole representing the database row.
	//
	// Returns:
	//   - A pointer to a UserRoleRecord containing the mapped data, including
	//     UserRoleID, UserID, RoleID, CreatedAt, and UpdatedAt.
	ToUserRoleRecord(userRole *db.UserRole) *record.UserRoleRecord
}
