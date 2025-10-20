package rolerecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// roleQueryRecordMapper is a struct that implements the RoleQueryRecordMapping interface.
type roleQueryRecordMapper struct{}

// NewRoleQueryRecordMapping creates a new instance of roleQueryRecordMapper, which implements the
// RoleQueryRecordMapping interface. This mapper is responsible for mapping role query results
// from the database to domain models, facilitating data transformation and access in the application.
func NewRoleQueryRecordMapping() RoleQueryRecordMapping {
	return &roleQueryRecordMapper{}
}

// ToRoleRecord maps a Role database row to a RoleRecord domain model.
//
// Parameters:
//   - role: A pointer to a Role representing the database row.
//
// Returns:
//   - A pointer to a RoleRecord containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRoleRecord(role *db.Role) *record.RoleRecord {
	deletedAt := role.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.RoleRecord{
		ID:        int(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

// ToRolesRecord maps a slice of Role database rows to a slice of RoleRecord domain models.
//
// Parameters:
//   - roles: A slice of pointers to Role structs representing the database rows.
//
// Returns:
//   - A slice of pointers to RoleRecord structs containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRolesRecord(roles []*db.Role) []*record.RoleRecord {
	var result []*record.RoleRecord

	for _, role := range roles {
		result = append(result, s.ToRoleRecord(role))
	}

	return result
}

// ToRoleRecordAll maps a GetRolesRow to a RoleRecord domain model.
//
// Parameters:
//   - role: A pointer to a GetRolesRow representing the database row.
//
// Returns:
//   - A pointer to a RoleRecord containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRoleRecordAll(role *db.GetRolesRow) *record.RoleRecord {
	deletedAt := role.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.RoleRecord{
		ID:        int(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

// ToRolesRecordAll maps a slice of GetRolesRow to a slice of RoleRecord domain models.
//
// Parameters:
//   - roles: A slice of pointers to GetRolesRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to RoleRecord structs containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRolesRecordAll(roles []*db.GetRolesRow) []*record.RoleRecord {
	var result []*record.RoleRecord

	for _, role := range roles {
		result = append(result, s.ToRoleRecordAll(role))
	}

	return result
}

// ToRoleRecordActive maps a GetActiveRolesRow to a RoleRecord domain model.
//
// Parameters:
//   - role: A pointer to a GetActiveRolesRow representing the database row.
//
// Returns:
//   - A pointer to a RoleRecord containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRoleRecordActive(role *db.GetActiveRolesRow) *record.RoleRecord {
	deletedAt := role.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.RoleRecord{
		ID:        int(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

// ToRolesRecordActive maps a slice of GetActiveRolesRow to a slice of RoleRecord domain models.
//
// Parameters:
//   - roles: A slice of pointers to GetActiveRolesRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to RoleRecord structs containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRolesRecordActive(roles []*db.GetActiveRolesRow) []*record.RoleRecord {
	var result []*record.RoleRecord

	for _, role := range roles {
		result = append(result, s.ToRoleRecordActive(role))
	}

	return result
}

// ToRoleRecordTrashed maps a GetTrashedRolesRow to a RoleRecord domain model.
//
// Parameters:
//   - role: A pointer to a GetTrashedRolesRow representing the database row.
//
// Returns:
//   - A pointer to a RoleRecord containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRoleRecordTrashed(role *db.GetTrashedRolesRow) *record.RoleRecord {
	deletedAt := role.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.RoleRecord{
		ID:        int(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

// ToRolesRecordTrashed maps a slice of GetTrashedRolesRow to a slice of RoleRecord domain models.
//
// Parameters:
//   - roles: A slice of pointers to GetTrashedRolesRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to RoleRecord structs containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and DeletedAt.
func (s *roleQueryRecordMapper) ToRolesRecordTrashed(roles []*db.GetTrashedRolesRow) []*record.RoleRecord {
	var result []*record.RoleRecord

	for _, role := range roles {
		result = append(result, s.ToRoleRecordTrashed(role))
	}

	return result
}
