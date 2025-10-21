package userprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type userQueryProtoMapper struct {
}

func NewUserQueryProtoMapper() UserQueryProtoMapper {
	return &userQueryProtoMapper{}
}

// ToProtoResponseUser maps a UserResponse to a pb.UserResponse
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A pointer to the UserResponse to be mapped.
//
// Returns:
//   - A pointer to a pb.UserResponse containing the mapped data.
//   - If pbResponse is nil, returns nil.
func (u *userQueryProtoMapper) ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser {
	return &pb.ApiResponseUser{
		Status:  status,
		Message: message,
		Data:    u.mapResponseUser(pbResponse),
	}
}

// ToProtoResponsePaginationUserDeleteAt maps a list of UserResponseDeleteAt and pagination metadata
// to a protobuf ApiResponsePaginationUserDeleteAt. It includes the status and message for the API response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - users: A list of UserResponseDeleteAt objects to be included in the response.
//
// Returns:
//
//	A pointer to ApiResponsePaginationUserDeleteAt containing the status, message, user data, and pagination data.
func (u *userQueryProtoMapper) ToProtoResponsePaginationUserDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) *pb.ApiResponsePaginationUserDeleteAt {
	return &pb.ApiResponsePaginationUserDeleteAt{
		Status:         status,
		Message:        message,
		Data:           u.mapResponsesUserDeleteAt(users),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationUser maps a list of UserResponse and pagination metadata
// to a protobuf ApiResponsePaginationUser. It includes the status and message for the API response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - users: A list of UserResponse objects to be included in the response.
//
// Returns:
//
//	A pointer to ApiResponsePaginationUser containing the status, message, user data, and pagination data.
func (u *userQueryProtoMapper) ToProtoResponsePaginationUser(pagination *pbhelpers.PaginationMeta, status string, message string, users []*response.UserResponse) *pb.ApiResponsePaginationUser {
	return &pb.ApiResponsePaginationUser{
		Status:         status,
		Message:        message,
		Data:           u.mapResponsesUser(users),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// mapResponseUser maps a UserResponse to a protobuf UserResponse.
//
// Args:
//   - user: A pointer to a UserResponse to be mapped.
//
// Returns:
//   - A pointer to a protobuf UserResponse containing the mapped data.
func (u *userQueryProtoMapper) mapResponseUser(user *response.UserResponse) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// mapResponsesUser maps a slice of UserResponse to a slice of protobuf UserResponse.
//
// Args:
//   - users: A slice of pointers to UserResponse to be mapped.
//
// Returns:
//   - A slice of pointers to protobuf UserResponse containing the mapped data for each user.
func (u *userQueryProtoMapper) mapResponsesUser(users []*response.UserResponse) []*pb.UserResponse {
	var mappedUsers []*pb.UserResponse

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.mapResponseUser(user))
	}

	return mappedUsers
}

// mapResponseUserDelete maps a UserResponseDeleteAt to a protobuf UserResponseDeleteAt.
//
// Args:
//   - user: A pointer to a UserResponseDeleteAt containing the user's details.
//
// Returns:
//   - A pointer to a protobuf UserResponseDeleteAt containing the mapped data, including
//     ID, Firstname, Lastname, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.
func (u *userQueryProtoMapper) mapResponseUserDelete(user *response.UserResponseDeleteAt) *pb.UserResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if user.DeletedAt != nil {
		deletedAt = wrapperspb.String(*user.DeletedAt)
	}

	return &pb.UserResponseDeleteAt{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

// mapResponsesUserDeleteAt maps a slice of UserResponseDeleteAt to a slice of protobuf UserResponseDeleteAt.
//
// Args:
//   - users: A slice of pointers to UserResponseDeleteAt to be mapped.
//
// Returns:
//   - A slice of pointers to protobuf UserResponseDeleteAt containing the mapped data for each user, including
//     ID, Firstname, Lastname, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.
func (u *userQueryProtoMapper) mapResponsesUserDeleteAt(users []*response.UserResponseDeleteAt) []*pb.UserResponseDeleteAt {
	var mappedUsers []*pb.UserResponseDeleteAt

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.mapResponseUserDelete(user))
	}

	return mappedUsers
}
