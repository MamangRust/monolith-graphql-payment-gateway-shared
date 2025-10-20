package merchantdocumentservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// MerchantDocumentBaseResponseMapper defines methods for converting merchant document records into API-compatible response types.
type MerchantDocumentBaseResponseMapper interface {
	// Converts a single merchant document record into a MerchantDocumentResponse.
	ToMerchantDocumentResponse(document *record.MerchantDocumentRecord) *response.MerchantDocumentResponse
}

// MerchantDocumentQueryResponseMapper defines methods for converting merchant document records into API-compatible response types.
type MerchantDocumentQueryResponseMapper interface {
	MerchantDocumentBaseResponseMapper

	// Converts multiple MerchantDocumentRecords to a slice of MerchantDocumentResponse.
	ToMerchantDocumentsResponse(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponse

	// Converts multiple soft-deleted MerchantDocumentRecords to their corresponding responses.
	ToMerchantDocumentsResponseDeleteAt(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponseDeleteAt
}

// MerchantDocumentCommandResponseMapper defines methods for converting merchant document records into API-compatible response types.
type MerchantDocumentCommandResponseMapper interface {
	MerchantDocumentBaseResponseMapper

	// Converts a soft-deleted MerchantDocumentRecord to its corresponding response.
	ToMerchantDocumentResponseDeleteAt(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponseDeleteAt
}
