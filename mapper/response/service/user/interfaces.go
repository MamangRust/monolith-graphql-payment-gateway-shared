package userservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// UserBaseResponseMapper provides methods to map UserRecord domain models to UserResponse API-compatible response types.
type UserBaseResponseMapper interface {
	// Converts a single user record into a UserResponse.
	ToUserResponse(user *record.UserRecord) *response.UserResponse
}

// UserQueryResponseMapper provides methods to map UserRecord domain models to UserResponse API-compatible response types.
type UserQueryResponseMapper interface {
	UserBaseResponseMapper

	// Converts a list of user records into a list of UserResponse.
	ToUsersResponse(users []*record.UserRecord) []*response.UserResponse

	// Converts a list of soft-deleted user records into a list of UserResponseDeleteAt.
	ToUsersResponseDeleteAt(users []*record.UserRecord) []*response.UserResponseDeleteAt
}

// UserCommandResponseMapper provides methods to map UserRecord domain models to UserResponse API-compatible response types.
type UserCommandResponseMapper interface {
	UserBaseResponseMapper

	// Converts a soft-deleted user record into a UserResponseDeleteAt.
	ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt
}
