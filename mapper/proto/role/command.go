package roleprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
)

type roleCommandProtoMapper struct {
}

// NewRoleCommandProtoMapper returns a new instance of roleCommandProtoMapper,
// which provides methods to map RoleRecord domain models to RoleResponse API-compatible
// response types for command operations.
func NewRoleCommandProtoMapper() RoleCommandProtoMapper {
	return &roleCommandProtoMapper{}
}

// ToProtoResponseRole converts a single role response to a protobuf message.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A pointer to the role response to be mapped.
//
// Returns:
//   - A pointer to a pb.ApiResponseRole containing the mapped data.
func (s *roleCommandProtoMapper) ToProtoResponseRole(status string, message string, pbResponse *response.RoleResponse) *pb.ApiResponseRole {
	return &pb.ApiResponseRole{
		Status:  status,
		Message: message,
		Data:    s.mapResponseRole(pbResponse),
	}
}

func (s *roleCommandProtoMapper) ToProtoResponseRoleDeleteAt(status string, message string, pbResponse *response.RoleResponseDeleteAt) *pb.ApiResponseRoleDeleteAt {
	return &pb.ApiResponseRoleDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseRoleDeleteAt(pbResponse),
	}
}

// ToProtoResponseRoleAll returns a protobuf message for all roles without pagination.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//
// Returns:
//   - A pointer to a pb.ApiResponseRoleAll containing the mapped data.
func (s *roleCommandProtoMapper) ToProtoResponseRoleAll(status string, message string) *pb.ApiResponseRoleAll {
	return &pb.ApiResponseRoleAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseRoleDelete returns a protobuf message indicating a role has been deleted.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//
// Returns:
//   - A pointer to a pb.ApiResponseRoleDelete containing the status and message.
func (s *roleCommandProtoMapper) ToProtoResponseRoleDelete(status string, message string) *pb.ApiResponseRoleDelete {
	return &pb.ApiResponseRoleDelete{
		Status:  status,
		Message: message,
	}
}

// mapResponseRole maps a single role response to a protobuf message.
//
// Args:
//   - role: A pointer to the role response to be mapped.
//
// Returns:
//   - A pointer to a pb.RoleResponse containing the mapped data.
func (s *roleCommandProtoMapper) mapResponseRole(role *response.RoleResponse) *pb.RoleResponse {
	return &pb.RoleResponse{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (s *roleCommandProtoMapper) mapResponseRoleDeleteAt(role *response.RoleResponseDeleteAt) *pb.RoleResponseDeleteAt {
	res := &pb.RoleResponseDeleteAt{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}

	if role.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(role.DeletedAt)
	}

	return res
}
