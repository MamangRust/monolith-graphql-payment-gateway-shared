package auth_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiVerifyCode returns an API error response when code verification fails.
var ErrApiVerifyCode = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to verify code", http.StatusBadRequest)
}

// ErrApiForgotPassword returns an API error response when forgot password request fails.
var ErrApiForgotPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to forgot password", http.StatusBadRequest)
}

// ErrApiResetPassword returns an API error response when reset password request fails.
var ErrApiResetPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to reset password", http.StatusBadRequest)
}

// ErrApiLogin returns an API error response when login fails due to invalid arguments.
var ErrApiLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "login failed: invalid argument provided", http.StatusBadRequest)
}

// ErrApiRefreshToken returns an API error response when access token is invalid during refresh token flow.
var ErrApiRefreshToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "refresh-token failed: invalid access token", http.StatusBadRequest)
}

// ErrApiGetMe returns an API error response when unauthenticated user tries to access their info.
var ErrApiGetMe = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "get user info failed: unauthenticated", http.StatusUnauthorized)
}

// ErrValidateLogin returns an API error response for invalid login request validation.
var ErrValidateLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid login request", http.StatusBadRequest)
}

// ErrValidateRegister returns an API error response for invalid register request validation.
var ErrValidateRegister = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid register request", http.StatusBadRequest)
}

// ErrValidateRefreshToken returns an API error response for invalid refresh token request validation.
var ErrValidateRefreshToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid refresh-token request", http.StatusBadRequest)
}

// ErrValidateForgotPassword returns an API error response for invalid forgot password request validation.
var ErrValidateForgotPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid forgot-password request", http.StatusBadRequest)
}

// ErrValidateResetPassword returns an API error response for invalid reset password request validation.
var ErrValidateResetPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid reset-password request", http.StatusBadRequest)
}

// ErrBindForgotPassword returns an API error response for binding failure on forgot password request payload.
var ErrBindForgotPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid forgot password request payload", http.StatusBadRequest)
}

// ErrBindResetPassword returns an API error response for binding failure on reset password request payload.
var ErrBindResetPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid reset password request payload", http.StatusBadRequest)
}

// ErrBindLogin returns an API error response for binding failure on login request payload.
var ErrBindLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid login request payload", http.StatusBadRequest)
}

// ErrBindRefreshToken returns an API error response for binding failure on refresh token request payload.
var ErrBindRefreshToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid refresh token request payload", http.StatusBadRequest)
}

// ErrBindRegister returns an API error response for binding failure on register request payload.
var ErrBindRegister = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid register request payload", http.StatusBadRequest)
}

// ErrInvalidLogin returns an API error response for incorrect email or password.
var ErrInvalidLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid email or password", http.StatusUnauthorized)
}

// ErrInvalidAccessToken returns an API error response for an invalid access token.
var ErrInvalidAccessToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid access token", http.StatusInternalServerError)
}

// ErrApiRegister returns an API error response when register request fails due to invalid arguments.
var ErrApiRegister = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "register failed: invalid argument", http.StatusBadRequest)
}
