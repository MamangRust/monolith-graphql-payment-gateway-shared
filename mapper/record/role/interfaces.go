package rolerecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// RoleBaseRecordMapper provides methods to map Role database rows to RoleRecord domain models.
type RoleBaseRecordMapper interface {
	// ToRoleRecord maps a Role database row to a RoleRecord domain model.
	//
	// Parameters:
	//   - role: A pointer to a Role representing the database row.
	//
	// Returns:
	//   - A pointer to a RoleRecord containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRoleRecord(role *db.Role) *record.RoleRecord
}

// RoleQueryRecordMapping maps query result rows to RoleRecord domain models.
type RoleQueryRecordMapping interface {
	RoleBaseRecordMapper

	// ToRolesRecord maps a slice of Role database rows to a slice of RoleRecord domain models.
	//
	// Parameters:
	//   - roles: A slice of pointers to Role structs representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to RoleRecord structs containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRolesRecord(roles []*db.Role) []*record.RoleRecord

	// ToRoleRecordAll maps a GetRolesRow to a RoleRecord domain model.
	//
	// Parameters:
	//   - role: A pointer to a GetRolesRow representing the database row.
	//
	// Returns:
	//   - A pointer to a RoleRecord containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRoleRecordAll(role *db.GetRolesRow) *record.RoleRecord

	// ToRolesRecordAll maps a slice of GetRolesRow to a slice of RoleRecord domain models.
	//
	// Parameters:
	//   - roles: A slice of pointers to GetRolesRow structs representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to RoleRecord structs containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRolesRecordAll(roles []*db.GetRolesRow) []*record.RoleRecord

	// ToRoleRecordActive maps a GetActiveRolesRow to a RoleRecord domain model.
	//
	// Parameters:
	//   - role: A pointer to a GetActiveRolesRow representing the database row.
	//
	// Returns:
	//   - A pointer to a RoleRecord containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRoleRecordActive(role *db.GetActiveRolesRow) *record.RoleRecord

	// ToRolesRecordActive maps a slice of GetActiveRolesRow to a slice of RoleRecord domain models.
	//
	// Parameters:
	//   - roles: A slice of pointers to GetActiveRolesRow structs representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to RoleRecord structs containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRolesRecordActive(roles []*db.GetActiveRolesRow) []*record.RoleRecord

	// ToRoleRecordTrashed maps a GetTrashedRolesRow to a RoleRecord domain model.
	//
	// Parameters:
	//   - role: A pointer to a GetTrashedRolesRow representing the database row.
	//
	// Returns:
	//   - A pointer to a RoleRecord containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRoleRecordTrashed(role *db.GetTrashedRolesRow) *record.RoleRecord

	// ToRolesRecordTrashed maps a slice of GetTrashedRolesRow to a slice of RoleRecord domain models.
	//
	// Parameters:
	//   - roles: A slice of pointers to GetTrashedRolesRow structs representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to RoleRecord structs containing the mapped data, including
	//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
	ToRolesRecordTrashed(roles []*db.GetTrashedRolesRow) []*record.RoleRecord
}

// RoleCommandRecordMapping maps basic Role rows to domain models for command operations.
type RoleCommandRecordMapping interface {
	RoleBaseRecordMapper
}
