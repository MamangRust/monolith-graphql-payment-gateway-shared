package merchantdocumentapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiInvalidMerchantId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "validation_error", "Invalid merchant ID", http.StatusBadRequest)
}

var ErrApiDocumentTypeRequired = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "validation_error", "Document type is required", http.StatusBadRequest)
}

var ErrApiDocumentFileRequired = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "validation_error", "Document file is required", http.StatusBadRequest)
}

var ErrApiStatusRequired = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "validation_error", "Status is required", http.StatusBadRequest)
}

var ErrApiNoteRequired = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "validation_error", "Note is required", http.StatusBadRequest)
}

var ErrApiFailedCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create merchant document", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant document", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant document status", http.StatusInternalServerError)
}

var ErrApiFailedTrashMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash merchant document", http.StatusInternalServerError)
}

var ErrApiFailedRestoreMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore merchant document", http.StatusInternalServerError)
}

var ErrApiFailedDeleteMerchantDocumentPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete merchant document", http.StatusInternalServerError)
}

var ErrApiFailedRestoreAllMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all merchant documents", http.StatusInternalServerError)
}

var ErrApiFailedDeleteAllMerchantDocumentsPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all merchant documents", http.StatusInternalServerError)
}
