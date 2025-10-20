package errorhandler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/MamangRust/monolith-graphql-payment-gateway-pkg/logger"
	traceunic "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/trace_unic"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// handleErrorTemplate is a generic function for handling errors with predefined error response templates.
// Parameters:
//   - logger: LoggerInterface instance for structured logging
//   - err: The error that occurred
//   - method: Name of the method where error occurred (e.g., "GetUser")
//   - tracePrefix: Prefix for trace ID generation (e.g., "GET_USER")
//   - errorMessage: Descriptive error message for logging
//   - span: OpenTelemetry span for distributed tracing
//   - status: Pointer to status string that will be updated
//   - errorResp: Predefined error response template
//   - fields: Additional zap fields for contextual logging
//
// Returns:
//   - Zero value of type T
//   - Pointer to response.ErrorResponse
func HandleErrorTemplate[T any](
	logger logger.LoggerInterface,
	err error,
	method, tracePrefix, errorMessage string,
	span trace.Span,
	status *string,
	errorResp *response.ErrorResponse,
	fields ...zap.Field,
) (T, *response.ErrorResponse) {
	traceID := traceunic.GenerateTraceID(tracePrefix)
	logMsg := fmt.Sprintf("%s in %s", errorMessage, method)

	allFields := append(fields,
		zap.Error(err),
		zap.String("trace.id", traceID),
	)

	logger.Error(logMsg, allFields...)

	span.SetAttributes(attribute.String("trace.id", traceID))
	span.RecordError(err)
	span.AddEvent(logMsg)
	span.SetStatus(codes.Error, logMsg)

	*status = fmt.Sprintf("%s_error_%s", toSnakeCase(method), toSnakeCase(errorMessage))

	var zero T
	return zero, errorResp
}

// toSnakeCase converts a camelCase string to a snake_case string.
//
// Parameters:
//   - s: A string in camelCase format.
//
// Returns:
//   - The equivalent string in snake_case format.
func toSnakeCase(s string) string {
	re := regexp.MustCompile("([a-z])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}
