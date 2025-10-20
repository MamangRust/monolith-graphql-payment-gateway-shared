package saldoservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// SaldoQueryResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types for query operations.
type saldoQueryResponseMapper struct {
}

// NewSaldoQueryResponseMapper returns a new instance of SaldoQueryResponseMapper which provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.
func NewSaldoQueryResponseMapper() SaldoQueryResponseMapper {
	return &saldoQueryResponseMapper{}
}

// ToSaldoResponse maps a single SaldoRecord to a SaldoResponse.
//
// Args:
//   - saldo: A pointer to a SaldoRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoResponse containing the mapped data.
func (s *saldoQueryResponseMapper) ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse {
	return &response.SaldoResponse{
		ID:             saldo.ID,
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

// ToSaldoResponses maps a list of SaldoRecord into a list of SaldoResponse.
//
// Args:
//
//	saldos: A slice of pointers to SaldoRecord containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoResponse containing the mapped data.
func (s *saldoQueryResponseMapper) ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse {
	var responses []*response.SaldoResponse

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoResponse(response))
	}

	return responses
}

// ToSaldoResponseDeleteAt maps a single SaldoRecord to a SaldoResponseDeleteAt.
//
// Args:
//   - saldo: A pointer to a SaldoRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoResponseDeleteAt containing the mapped data.
func (s *saldoQueryResponseMapper) ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt {
	return &response.SaldoResponseDeleteAt{
		ID:             saldo.ID,
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      saldo.DeletedAt,
	}
}

// ToSaldoResponsesDeleteAt maps a list of SaldoRecord into a list of SaldoResponseDeleteAt.
//
// Args:
//
//	saldos: A slice of pointers to SaldoRecord containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoResponseDeleteAt containing the mapped data.
func (s *saldoQueryResponseMapper) ToSaldoResponsesDeleteAt(saldos []*record.SaldoRecord) []*response.SaldoResponseDeleteAt {
	var responses []*response.SaldoResponseDeleteAt

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoResponseDeleteAt(response))
	}

	return responses
}
