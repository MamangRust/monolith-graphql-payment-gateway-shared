package saldorecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// SaldoCommandRecordMapping provides methods to map Saldo database rows to SaldoRecord domain models for command operations.
type saldoCommandRecordMapper struct{}

// NewSaldoCommandRecordMapper returns a new instance of SaldoCommandRecordMapping which provides methods to map Saldo database rows to SaldoRecord domain models for command operations.
func NewSaldoCommandRecordMapper() SaldoCommandRecordMapping {
	return &saldoCommandRecordMapper{}
}

// ToSaldoRecord maps a Saldo database row to a SaldoRecord domain model.
//
// Parameters:
//   - saldo: A pointer to a Saldo representing the database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoCommandRecordMapper) ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord {
	var deletedAt *string

	if saldo.DeletedAt.Valid {
		formatedDeletedAt := saldo.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.SaldoRecord{
		ID:             int(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawAmount: int(saldo.WithdrawAmount.Int32),
		WithdrawTime:   saldo.WithdrawTime.Time.Format("2006-01-02"),
		CreatedAt:      saldo.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:      saldo.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:      deletedAt,
	}
}
