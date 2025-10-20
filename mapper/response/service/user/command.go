package userservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type userCommandResponseMapper struct{}

func NewUserCommandResponseMapper() UserCommandResponseMapper {
	return &userCommandResponseMapper{}
}

// ToUserResponse maps a UserRecord domain model to a UserResponse API-compatible format.
//
// Args:
//   - user: A pointer to a UserRecord containing the user data to be mapped.
//
// Returns:
//   - A pointer to a UserResponse with the mapped user data, including ID, FirstName, LastName,
//     Email, IsVerified, CreatedAt, and UpdatedAt.
func (s *userCommandResponseMapper) ToUserResponse(user *record.UserRecord) *response.UserResponse {
	return &response.UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

// ToUserResponseDeleteAt converts a UserRecord domain model into a UserResponseDeleteAt API-compatible format.
//
// Args:
//   - user: A pointer to a UserRecord containing the user data to be mapped.
//
// Returns:
//   - A pointer to a UserResponseDeleteAt with the mapped user data, including ID, FirstName, LastName,
//     Email, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userCommandResponseMapper) ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt {
	return &response.UserResponseDeleteAt{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt,
	}
}
