package roleprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type RoleBaseProtoMapper interface {
	// ToProtoResponseRole converts a single role response to a protobuf message.
	ToProtoResponseRole(status string, message string, pbResponse *response.RoleResponse) *pb.ApiResponseRole
}

type RoleQueryProtoMapper interface {
	RoleBaseProtoMapper

	// ToProtoResponsesRole converts a list of role responses to a protobuf message.
	ToProtoResponsesRole(status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsesRole

	// ToProtoResponsePaginationRole returns a paginated protobuf message for roles.
	ToProtoResponsePaginationRole(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsePaginationRole

	// ToProtoResponsePaginationRoleDeleteAt returns a paginated protobuf message for soft-deleted roles.
	ToProtoResponsePaginationRoleDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.RoleResponseDeleteAt) *pb.ApiResponsePaginationRoleDeleteAt
}

type RoleCommandProtoMapper interface {
	RoleBaseProtoMapper

	ToProtoResponseRoleDeleteAt(status string, message string, pbResponse *response.RoleResponseDeleteAt) *pb.ApiResponseRoleDeleteAt

	// ToProtoResponseRoleAll returns a protobuf message for all roles without pagination.
	ToProtoResponseRoleAll(status string, message string) *pb.ApiResponseRoleAll

	// ToProtoResponseRoleDelete returns a protobuf message indicating a role has been deleted.
	ToProtoResponseRoleDelete(status string, message string) *pb.ApiResponseRoleDelete
}
