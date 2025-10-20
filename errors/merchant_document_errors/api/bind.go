package merchantdocumentapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateMerchantDocument is returned when binding the create request payload fails.
var ErrApiBindCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create merchant document request", http.StatusBadRequest)
}

// ErrApiBindUpdateMerchantDocument is returned when binding the update request payload fails.
var ErrApiBindUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant document request", http.StatusBadRequest)
}

// ErrApiBindUpdateMerchantDocumentStatus is returned when binding the update status request payload fails.
var ErrApiBindUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant document status request", http.StatusBadRequest)
}
