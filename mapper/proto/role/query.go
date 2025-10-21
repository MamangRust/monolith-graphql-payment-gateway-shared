package roleprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type roleQueryProtoMapper struct {
}

// NewRoleProtoMapper returns a new instance of roleQueryProtoMapper.
func NewRoleQueryProtoMapper() RoleQueryProtoMapper {
	return &roleQueryProtoMapper{}
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
func (s *roleQueryProtoMapper) ToProtoResponseRole(status string, message string, pbResponse *response.RoleResponse) *pb.ApiResponseRole {
	return &pb.ApiResponseRole{
		Status:  status,
		Message: message,
		Data:    s.mapResponseRole(pbResponse),
	}
}

// ToProtoResponsesRole converts a list of role responses to a protobuf message.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A slice of pointers to RoleResponse to be mapped.
//
// Returns:
//   - A pointer to a pb.ApiResponsesRole containing the mapped data.
func (s *roleQueryProtoMapper) ToProtoResponsesRole(status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsesRole {
	return &pb.ApiResponsesRole{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesRole(pbResponse),
	}
}

// ToProtoResponsePaginationRole maps paginated role responses to a protobuf message.
//
// Args:
//   - pagination: A pointer to the pagination metadata for the response.
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A slice of pointers to RoleResponse to be mapped.
//
// Returns:
//   - A pointer to a pb.ApiResponsePaginationRole containing the mapped data.
func (s *roleQueryProtoMapper) ToProtoResponsePaginationRole(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsePaginationRole {
	return &pb.ApiResponsePaginationRole{
		Status:         status,
		Message:        message,
		Data:           s.mapResponsesRole(pbResponse),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationRoleDeleteAt maps paginated soft-deleted role responses to a protobuf message.
//
// Args:
//   - pagination: A pointer to the pagination metadata for the response.
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A slice of pointers to RoleResponseDeleteAt to be mapped.
//
// Returns:
//   - A pointer to a pb.ApiResponsePaginationRoleDeleteAt containing the mapped data.
func (s *roleQueryProtoMapper) ToProtoResponsePaginationRoleDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.RoleResponseDeleteAt) *pb.ApiResponsePaginationRoleDeleteAt {
	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:         status,
		Message:        message,
		Data:           s.mapResponsesRoleDeleteAt(pbResponse),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// mapResponseRole maps a single role response to a protobuf message.
//
// Args:
//   - role: A pointer to the role response to be mapped.
//
// Returns:
//   - A pointer to a pb.RoleResponse containing the mapped data.
func (s *roleQueryProtoMapper) mapResponseRole(role *response.RoleResponse) *pb.RoleResponse {
	return &pb.RoleResponse{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// mapResponsesRole maps a list of *response.RoleResponse to a list of
// *pb.RoleResponse proto responses.
//
// It iterates over each RoleResponse in the input slice, converting
// them to their protobuf equivalent using the mapResponseRole function.
// This function is used to generate the response data for role RPC methods.
func (s *roleQueryProtoMapper) mapResponsesRole(roles []*response.RoleResponse) []*pb.RoleResponse {
	var responseRoles []*pb.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRole(role))
	}

	return responseRoles
}

// mapResponseRoleDeleteAt maps a RoleResponseDeleteAt to its protobuf representation.
//
// Args:
//   - role: A pointer to a RoleResponseDeleteAt containing the details of the role.
//
// Returns:
//   - A pointer to a pb.RoleResponseDeleteAt containing the mapped data, including
//     ID, Name, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.
func (s *roleQueryProtoMapper) mapResponseRoleDeleteAt(role *response.RoleResponseDeleteAt) *pb.RoleResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if role.DeletedAt != nil {
		deletedAt = wrapperspb.String(*role.DeletedAt)
	}

	return &pb.RoleResponseDeleteAt{
		Id:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

// mapResponsesRoleDeleteAt maps a list of *response.RoleResponseDeleteAt to a list of
// *pb.RoleResponseDeleteAt proto responses.
//
// It iterates over each RoleResponseDeleteAt in the input slice, converting
// them to their protobuf equivalent using the mapResponseRoleDeleteAt function.
// This function is used to generate the response data for the RoleService.ListDeletedRole rpc method.
func (s *roleQueryProtoMapper) mapResponsesRoleDeleteAt(roles []*response.RoleResponseDeleteAt) []*pb.RoleResponseDeleteAt {
	var responseRoles []*pb.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRoleDeleteAt(role))
	}

	return responseRoles
}
