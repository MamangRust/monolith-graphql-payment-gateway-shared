package userprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type UserBaseProtoMapper interface {
	// ToProtoResponseUser converts a single user response to a protobuf message.
	ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser
}

type UserQueryProtoMapper interface {
	UserBaseProtoMapper

	// ToProtoResponsePaginationUserDeleteAt returns a paginated protobuf message for soft-deleted users.
	ToProtoResponsePaginationUserDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) *pb.ApiResponsePaginationUserDeleteAt

	// ToProtoResponsePaginationUser returns a paginated protobuf message for active users.
	ToProtoResponsePaginationUser(pagination *pbhelpers.PaginationMeta, status string, message string, users []*response.UserResponse) *pb.ApiResponsePaginationUser
}

// UserCommandProtoMapper is an interface for mapping user commands to protobuf messages.
type UserCommandProtoMapper interface {
	UserBaseProtoMapper

	// ToProtoResponseUserDeleteAt converts a soft-deleted user response to a protobuf message.
	ToProtoResponseUserDeleteAt(status string, message string, pbResponse *response.UserResponseDeleteAt) *pb.ApiResponseUserDeleteAt

	// ToProtoResponseUserDelete returns a protobuf message indicating a user has been permanently deleted.
	ToProtoResponseUserDelete(status string, message string) *pb.ApiResponseUserDelete

	// ToProtoResponseUserAll returns a protobuf message for all users without pagination.
	ToProtoResponseUserAll(status string, message string) *pb.ApiResponseUserAll
}
