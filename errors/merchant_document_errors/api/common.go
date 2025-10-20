package merchantdocumentapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiInvalidMerchantDocumentID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid merchant document ID", http.StatusBadRequest)
}
