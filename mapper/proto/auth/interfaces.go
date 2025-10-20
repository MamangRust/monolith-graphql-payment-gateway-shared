package authprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/auth"
)

type AuthProtoMapper interface {
	// ToProtoResponseVerifyCode converts a verification code response to a protobuf message.
	ToProtoResponseVerifyCode(status string, message string) *pb.ApiResponseVerifyCode

	// ToProtoResponseForgotPassword converts a forgot password response to a protobuf message.
	ToProtoResponseForgotPassword(status string, message string) *pb.ApiResponseForgotPassword

	// ToProtoResponseResetPassword converts a reset password response to a protobuf message.
	ToProtoResponseResetPassword(status string, message string) *pb.ApiResponseResetPassword

	// ToProtoResponseLogin converts a login response with token details to a protobuf message.
	ToProtoResponseLogin(status string, message string, response *response.TokenResponse) *pb.ApiResponseLogin

	// ToProtoResponseRegister converts a user registration response to a protobuf message.
	ToProtoResponseRegister(status string, message string, response *response.UserResponse) *pb.ApiResponseRegister

	// ToProtoResponseRefreshToken converts a token refresh response to a protobuf message.
	ToProtoResponseRefreshToken(status string, message string, response *response.TokenResponse) *pb.ApiResponseRefreshToken

	// ToProtoResponseGetMe converts a user profile fetch ("get me") response to a protobuf message.
	ToProtoResponseGetMe(status string, message string, response *response.UserResponse) *pb.ApiResponseGetMe
}
