package userprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type userCommandProtoMapper struct {
}

func NewUserCommandProtoMapper() UserCommandProtoMapper {
	return &userCommandProtoMapper{}
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
func (u *userCommandProtoMapper) ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser {
	return &pb.ApiResponseUser{
		Status:  status,
		Message: message,
		Data:    u.mapResponseUser(pbResponse),
	}
}

// ToProtoResponseUserDeleteAt maps a UserResponseDeleteAt to a pb.ApiResponseUserDeleteAt proto message.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A pointer to the UserResponseDeleteAt to be mapped.
//
// Returns:
//   - A pointer to a pb.ApiResponseUserDeleteAt containing the mapped data.
//   - If pbResponse is nil, returns nil.
func (u *userCommandProtoMapper) ToProtoResponseUserDeleteAt(status string, message string, pbResponse *response.UserResponseDeleteAt) *pb.ApiResponseUserDeleteAt {
	return &pb.ApiResponseUserDeleteAt{
		Status:  status,
		Message: message,
		Data:    u.mapResponseUserDelete(pbResponse),
	}
}

// ToProtoResponseUserDelete maps a UserResponseDelete to a pb.ApiResponseUserDelete proto message.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//
// Returns:
//   - A pointer to a pb.ApiResponseUserDelete containing the mapped data.
func (u *userCommandProtoMapper) ToProtoResponseUserDelete(status string, message string) *pb.ApiResponseUserDelete {
	return &pb.ApiResponseUserDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseUserAll returns a protobuf message for all users without pagination.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//
// Returns:
//   - A pointer to a pb.ApiResponseUserAll containing the mapped data.
func (u *userCommandProtoMapper) ToProtoResponseUserAll(status string, message string) *pb.ApiResponseUserAll {
	return &pb.ApiResponseUserAll{
		Status:  status,
		Message: message,
	}
}

// mapResponseUser maps a UserResponse to a protobuf UserResponse.
//
// Args:
//   - user: A pointer to a UserResponse to be mapped.
//
// Returns:
//   - A pointer to a protobuf UserResponse containing the mapped data.
func (u *userCommandProtoMapper) mapResponseUser(user *response.UserResponse) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// mapResponseUserDelete maps a UserResponseDeleteAt to a protobuf UserResponseDeleteAt.
//
// Args:
//   - user: A pointer to a UserResponseDeleteAt containing the user's details.
//
// Returns:
//   - A pointer to a protobuf UserResponseDeleteAt containing the mapped data, including
//     ID, Firstname, Lastname, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.
func (u *userCommandProtoMapper) mapResponseUserDelete(user *response.UserResponseDeleteAt) *pb.UserResponseDeleteAt {
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
