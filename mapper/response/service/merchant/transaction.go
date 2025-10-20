package merchantservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// merchantTransactionResponse is a struct that provides methods to map MerchantTransactionsRecord domain models to MerchantTransactionResponse API-compatible response types.
type merchantTransactionResponse struct {
}

// NewMerchantTransactionResponseMapper returns a new instance of merchantTransactionResponse,
// which provides methods to map MerchantTransactionsRecord domain models to
// MerchantTransactionResponse API-compatible response types.
func NewMerchantTransactionResponseMapper() MerchantTransactionResponseMapper {
	return &merchantTransactionResponse{}
}

// ToMerchantTransactionResponse maps a single MerchantTransactionsRecord to a MerchantTransactionResponse API-compatible response.
//
// Args:
//   - merchant: A pointer to a MerchantTransactionsRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantTransactionResponse containing the mapped data, including
//     ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt.
func (m *merchantTransactionResponse) ToMerchantTransactionResponse(merchant *record.MerchantTransactionsRecord) *response.MerchantTransactionResponse {

	return &response.MerchantTransactionResponse{
		ID:              int(merchant.TransactionID),
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime.Format("2006-01-02"),
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

// ToMerchantsTransactionResponse maps multiple MerchantTransactionsRecords to a slice of MerchantTransactionResponse API-compatible responses.
// Args:
//   - merchants: A slice of pointers to MerchantTransactionsRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantTransactionResponse containing the mapped data, including
//     ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt.
func (m *merchantTransactionResponse) ToMerchantsTransactionResponse(merchants []*record.MerchantTransactionsRecord) []*response.MerchantTransactionResponse {
	var records []*response.MerchantTransactionResponse
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionResponse(merchant))
	}
	return records
}
