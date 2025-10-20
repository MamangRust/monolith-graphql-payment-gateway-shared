package merchantdocumentapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiValidateCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create merchant document request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant document request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant document status request", http.StatusBadRequest)
}
