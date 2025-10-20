package merchantservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// MerchantQueryResponseMapper provides methods to map MerchantRecord domain models to MerchantResponse and MerchantResponseDeleteAt API-compatible responses for query operations.
type merchantQueryResponseMapper struct{}

// NewMerchantQueryResponseMapper returns a new instance of merchantQueryResponseMapper,
// which provides methods to map MerchantRecord domain models to API-compatible
// MerchantResponse types for query operations.
func NewMerchantQueryResponseMapper() MerchantQueryResponseMapper {
	return &merchantQueryResponseMapper{}
}

// ToMerchantResponse maps a single MerchantRecord to a MerchantResponse API-compatible response.
// Args:
//   - merchant: A pointer to a MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponse containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.
func (s *merchantQueryResponseMapper) ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse {
	return &response.MerchantResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		UserID:    merchant.UserID,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

// ToMerchantsResponse maps multiple MerchantRecords to a slice of MerchantResponse API-compatible responses.
// Args:
//   - merchants: A slice of pointers to MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponse containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.
func (s *merchantQueryResponseMapper) ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse {
	var response []*response.MerchantResponse
	for _, merchant := range merchants {
		response = append(response, s.ToMerchantResponse(merchant))
	}
	return response
}

// ToMerchantResponseDeleteAt maps a MerchantRecord to a MerchantResponseDeleteAt,
// which includes additional deletion data.
// This function is useful for handling soft-deleted merchants where the deletion
// timestamp is relevant.
//
// Args:
//   - merchant: A pointer to a MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseDeleteAt containing the mapped data,
//     including ID, Name, UserID, Status, ApiKey, CreatedAt, UpdatedAt, and DeletedAt.
func (s *merchantQueryResponseMapper) ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt {
	return &response.MerchantResponseDeleteAt{
		ID:        merchant.ID,
		Name:      merchant.Name,
		UserID:    merchant.UserID,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: merchant.DeletedAt,
	}
}

// ToMerchantsResponseDeleteAt maps multiple soft-deleted MerchantRecords to their corresponding responses.
// Args:
//   - merchants: A slice of pointers to MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseDeleteAt containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, UpdatedAt, and DeletedAt.
func (s *merchantQueryResponseMapper) ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt {
	var response []*response.MerchantResponseDeleteAt
	for _, merchant := range merchants {
		response = append(response, s.ToMerchantResponseDeleteAt(merchant))
	}
	return response
}
