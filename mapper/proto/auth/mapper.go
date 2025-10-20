package authprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/auth"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
)

type authProtoMapper struct {
}

// NewAuthProtoMapper returns a new instance of authProtoMapper.
func NewAuthProtoMapper() AuthProtoMapper {
	return &authProtoMapper{}
}

// ToProtoResponseVerifyCode maps the response to the proto type.
//
// Args:
//
//	status: The status of the response.
//	message: The message of the response.
//
// Returns:
//
//	A pointer to the proto type.
func (s *authProtoMapper) ToProtoResponseVerifyCode(status string, message string) *pb.ApiResponseVerifyCode {
	return &pb.ApiResponseVerifyCode{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseForgotPassword maps the response to the proto type.
//
// Args:
//
//	status: The status of the response.
//	message: The message of the response.
//
// Returns:
//
//	A pointer to the proto type.
func (s *authProtoMapper) ToProtoResponseForgotPassword(status string, message string) *pb.ApiResponseForgotPassword {
	return &pb.ApiResponseForgotPassword{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseResetPassword maps the response to the proto type.
//
// Args:
//
//	status: The status of the response.
//	message: The message of the response.
//
// Returns:
//
//	A pointer to the proto type.
func (s *authProtoMapper) ToProtoResponseResetPassword(status string, message string) *pb.ApiResponseResetPassword {
	return &pb.ApiResponseResetPassword{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseLogin maps the response to the proto type.
//
// Args:
//
//	status: The status of the response.
//	message: The message of the response.
//	response: The response containing the access and refresh tokens.
//
// Returns:
//
//	A pointer to the proto type.
func (s *authProtoMapper) ToProtoResponseLogin(status string, message string, response *response.TokenResponse) *pb.ApiResponseLogin {
	return &pb.ApiResponseLogin{
		Status:  status,
		Message: message,
		Data: &pb.TokenResponse{
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		},
	}
}

// ToProtoResponseRegister maps a user registration response to the protobuf format.
//
// Args:
//
//	status: The status of the registration process.
//	message: The message associated with the registration response.
//	response: The user response containing the registered user's details.
//
// Returns:
//
//	A pointer to the ApiResponseRegister protobuf type containing the status, message, and user data.

func (s *authProtoMapper) ToProtoResponseRegister(status string, message string, response *response.UserResponse) *pb.ApiResponseRegister {
	return &pb.ApiResponseRegister{
		Status:  status,
		Message: message,
		Data: &pbuser.UserResponse{
			Id:        int32(response.ID),
			Firstname: response.FirstName,
			Lastname:  response.LastName,
			Email:     response.Email,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
}

// ToProtoResponseRefreshToken maps a token refresh operation response to the protobuf format.
//
// Args:
//
//	status: The status of the refresh operation.
//	message: The message associated with the refresh operation response.
//	response: The TokenResponse containing the new access and refresh tokens.
//
// Returns:
//
//	A pointer to the ApiResponseRefreshToken protobuf type containing the status, message, and token data.

func (s *authProtoMapper) ToProtoResponseRefreshToken(status string, message string, response *response.TokenResponse) *pb.ApiResponseRefreshToken {
	return &pb.ApiResponseRefreshToken{
		Status:  status,
		Message: message,
		Data: &pb.TokenResponse{
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		},
	}
}

// ToProtoResponseGetMe maps a user profile response to the protobuf format.
//
// Args:
//
//	status: The status of the response.
//	message: The message associated with the response.
//	response: The UserResponse containing the user's details.
//
// Returns:
//
//	A pointer to the ApiResponseGetMe protobuf type containing the status, message, and user data.
func (s *authProtoMapper) ToProtoResponseGetMe(status string, message string, response *response.UserResponse) *pb.ApiResponseGetMe {
	return &pb.ApiResponseGetMe{
		Status:  status,
		Message: message,
		Data: &pbuser.UserResponse{
			Id:        int32(response.ID),
			Firstname: response.FirstName,
			Lastname:  response.LastName,
			Email:     response.Email,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
}
