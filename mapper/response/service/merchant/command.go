package merchantservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// MerchantCommandResponseMapper provides methods to map MerchantRecord domain models to MerchantResponse API-compatible response types for command operations.
type merchantCommandResponseMapper struct{}

// NewMerchantCommandResponseMapper returns a new instance of merchantCommandResponseMapper, which provides methods to map
// MerchantRecord domain models to MerchantResponse API-compatible response types for command operations.
func NewMerchantCommandResponseMapper() MerchantCommandResponseMapper {
	return &merchantCommandResponseMapper{}
}

// ToMerchantResponse maps a single MerchantRecord to a MerchantResponse API-compatible response.
// Args:
//   - merchant: A pointer to a MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponse containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.
func (s *merchantCommandResponseMapper) ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse {
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
func (s *merchantCommandResponseMapper) ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt {
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
