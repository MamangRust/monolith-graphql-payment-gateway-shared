package response

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewErrorResponse creates a new ErrorResponse struct with standardized error format.
// This is used to maintain consistent error responses across the API.
//
// Parameters:
//   - message: Human-readable error description
//   - code: HTTP status code
//
// Returns:
//   - *ErrorResponse: Pointer to the constructed error response
func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{
		Status:  "error",
		Message: message,
		Code:    code,
	}
}

// NewApiErrorResponse creates and sends a JSON error response using Echo framework.
// This is a convenience wrapper for returning standardized error responses in HTTP handlers.
//
// Parameters:
//   - c: Echo context
//   - statusText: Status category ("error", "fail", etc.)
//   - message: Human-readable error description
//   - code: HTTP status code
//
// Returns:
//   - error: Echo error that will trigger the JSON response
func NewApiErrorResponse(c echo.Context, statusText string, message string, code int) error {
	return c.JSON(code, ErrorResponse{
		Status:  statusText,
		Message: message,
		Code:    code,
	})
}

// ToGrpcErrorFromErrorResponse converts an ErrorResponse to a gRPC error.
// This enables consistent error handling between HTTP and gRPC interfaces.
//
// Parameters:
//   - err: ErrorResponse to convert
//
// Returns:
//   - error: gRPC status error with embedded error details
//   - nil: if input error is nil
func ToGrpcErrorFromErrorResponse(err *ErrorResponse) error {
	if err == nil {
		return nil
	}
	return status.Errorf(codes.Code(err.Code),
		"%s", errors.GrpcErrorToJson(&pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
			Code:    int32(err.Code),
		}),
	)
}

// NewGrpcError creates a new gRPC error with standardized format.
// This is used to maintain consistent error responses across gRPC services.
//
// Parameters:
//   - statusText: Status category ("error", "fail", etc.)
//   - message: Human-readable error description
//   - code: gRPC status code
//
// Returns:
//   - error: gRPC status error with embedded error details
func NewGrpcError(statusText string, message string, code int) error {
	return status.Errorf(codes.Code(code),
		"%s", errors.GrpcErrorToJson(&pb.ErrorResponse{
			Status:  statusText,
			Message: message,
			Code:    int32(code),
		}),
	)
}
