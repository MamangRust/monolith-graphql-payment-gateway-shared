package roleservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// RoleCommandResponseMapper provides methods to map RoleRecord domain models to RoleResponse API-compatible response types for command operations.
type roleCommandResponseMapper struct {
}

// NewRoleCommandResponseMapper returns a new instance of roleCommandResponseMapper,
// which provides methods to map RoleRecord domain models to RoleResponse API-compatible
// response types for command operations.
func NewRoleCommandResponseMapper() RoleCommandResponseMapper {
	return &roleCommandResponseMapper{}
}

// ToRoleResponse converts a single RoleRecord into RoleResponse.
//
// Args:
//   - role: A pointer to the RoleRecord to be mapped.
//
// Returns:
//   - A pointer to a RoleResponse containing the mapped data, with fields ID, Name,
//     CreatedAt, and UpdatedAt extracted from the RoleRecord.
func (s *roleCommandResponseMapper) ToRoleResponse(role *record.RoleRecord) *response.RoleResponse {
	return &response.RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// ToRoleResponseDeleteAt converts a RoleRecord with deletion information
// into a RoleResponseDeleteAt structure.
//
// Args:
//   - role: A pointer to the RoleRecord to be mapped.
//
// Returns:
//   - A pointer to a RoleResponseDeleteAt containing the mapped data, with fields ID, Name,
//     CreatedAt, UpdatedAt, and DeletedAt extracted from the RoleRecord.
func (s *roleCommandResponseMapper) ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt {
	return &response.RoleResponseDeleteAt{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}
}
