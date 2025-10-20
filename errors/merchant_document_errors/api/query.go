package merchantdocumentapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiFailedFindAllMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchant documents", http.StatusInternalServerError)
}

var ErrApiFailedFindByIdMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant document by ID", http.StatusInternalServerError)
}

var ErrApiFailedFindAllActiveMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all active merchant documents", http.StatusInternalServerError)
}

var ErrApiFailedFindAllTrashedMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all trashed merchant documents", http.StatusInternalServerError)
}
