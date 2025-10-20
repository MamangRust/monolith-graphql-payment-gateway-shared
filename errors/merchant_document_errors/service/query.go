package merchantdocumentserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrMerchantDocumentNotFoundRes is returned when the requested merchant document cannot be found.
var ErrMerchantDocumentNotFoundRes = response.NewErrorResponse("Merchant Document not found", http.StatusNotFound)

// ErrFailedFindAllMerchantDocuments is returned when failing to fetch all merchant documents.
var ErrFailedFindAllMerchantDocuments = response.NewErrorResponse("Failed to fetch Merchant Documents", http.StatusInternalServerError)

// ErrFailedFindActiveMerchantDocuments is returned when failing to fetch active merchant documents.
var ErrFailedFindActiveMerchantDocuments = response.NewErrorResponse("Failed to fetch active Merchant Documents", http.StatusInternalServerError)

// ErrFailedFindTrashedMerchantDocuments is returned when failing to fetch trashed merchant documents.
var ErrFailedFindTrashedMerchantDocuments = response.NewErrorResponse("Failed to fetch trashed Merchant Documents", http.StatusInternalServerError)

// ErrFailedFindMerchantDocumentById is returned when failing to find a merchant document by ID.
var ErrFailedFindMerchantDocumentById = response.NewErrorResponse("Failed to find Merchant Document by ID", http.StatusInternalServerError)
