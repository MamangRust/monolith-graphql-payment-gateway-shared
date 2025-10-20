package saldoservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// SaldoCommandResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types for command operations.
type saldoCommandResponseMapper struct {
}

// NewSaldoCommandResponseMapper returns a new instance of SaldoCommandResponseMapper which provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types for command operations.
func NewSaldoCommandResponseMapper() SaldoCommandResponseMapper {
	return &saldoCommandResponseMapper{}
}

// ToSaldoResponse maps a single SaldoRecord to a SaldoResponse.
//
// Args:
//   - saldo: A pointer to a SaldoRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoResponse containing the mapped data.
func (s *saldoCommandResponseMapper) ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse {
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

// ToSaldoResponseDeleteAt maps a single SaldoRecord to a SaldoResponseDeleteAt.
//
// Args:
//   - saldo: A pointer to a SaldoRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoResponseDeleteAt containing the mapped data.
func (s *saldoCommandResponseMapper) ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt {
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
