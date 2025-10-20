package roleservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// RoleQueryResponseMapper provides methods to map RoleRecord domain models to RoleResponse API-compatible responses for query operations.
type roleQueryResponseMapper struct{}

// NewRoleQueryResponseMapper creates a new instance of RoleQueryResponseMapper.
// This mapper is responsible for converting role query domain models into
// API-compatible response types, facilitating the transformation of RoleRecord
// data to RoleResponse structures.
func NewRoleQueryResponseMapper() RoleQueryResponseMapper {
	return &roleQueryResponseMapper{}
}

// ToRoleResponse converts a single RoleRecord into RoleResponse.
//
// Args:
//   - role: A pointer to the RoleRecord to be mapped.
//
// Returns:
//   - A pointer to a RoleResponse containing the mapped data, with fields ID, Name,
//     CreatedAt, and UpdatedAt extracted from the RoleRecord.
func (s *roleQueryResponseMapper) ToRoleResponse(role *record.RoleRecord) *response.RoleResponse {
	return &response.RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// ToRolesResponse maps a slice of RoleRecord domain models to a slice of RoleResponse API-compatible response types.
//
// Args:
//   - roles: A slice of pointers to RoleRecord domain models to be mapped.
//
// Returns:
//   - A slice of pointers to RoleResponse, each containing the mapped data, with fields ID, Name,
//     CreatedAt, and UpdatedAt extracted from the RoleRecord.
func (s *roleQueryResponseMapper) ToRolesResponse(roles []*record.RoleRecord) []*response.RoleResponse {
	var responseRoles []*response.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.ToRoleResponse(role))
	}

	return responseRoles
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
func (s *roleQueryResponseMapper) ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt {
	return &response.RoleResponseDeleteAt{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}
}

// ToRolesResponseDeleteAt maps a slice of RoleRecord domain models with deletion information
// to a slice of RoleResponseDeleteAt API-compatible response types.
//
// Args:
//   - roles: A slice of pointers to RoleRecord domain models to be mapped,
//     which must have deletion information (DeletedAt != nil).
//
// Returns:
//   - A slice of pointers to RoleResponseDeleteAt, each containing the mapped data, with fields ID, Name,
//     CreatedAt, UpdatedAt, and DeletedAt extracted from the RoleRecord.
func (s *roleQueryResponseMapper) ToRolesResponseDeleteAt(roles []*record.RoleRecord) []*response.RoleResponseDeleteAt {
	var responseRoles []*response.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.ToRoleResponseDeleteAt(role))
	}

	return responseRoles
}
