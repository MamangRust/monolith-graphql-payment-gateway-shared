package saldorecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// SaldoQueryRecordMapping provides methods to map Saldo database rows to SaldoRecord domain models for command operations.
type saldoQueryRecordMapper struct{}

// NewSaldoQueryRecordMapper returns a new instance of SaldoQueryRecordMapping which provides methods to map Saldo database rows to SaldoRecord domain models for command operations.
func NewSaldoQueryRecordMapper() SaldoQueryRecordMapping {
	return &saldoQueryRecordMapper{}
}

// ToSaldoRecord maps a Saldo database row to a SaldoRecord domain model.
//
// Parameters:
//   - saldo: A pointer to a Saldo representing the database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoQueryRecordMapper) ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord {
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

// ToSaldoRecordAll maps a GetSaldosRow database row to a SaldoRecord domain model.
//
// Parameters:
//   - saldo: A pointer to a GetSaldosRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoQueryRecordMapper) ToSaldoRecordAll(saldo *db.GetSaldosRow) *record.SaldoRecord {
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

// ToSaldosRecordAll maps a slice of GetSaldosRow database rows to a slice of SaldoRecord domain models.
//
// Parameters:
//   - saldos: A slice of pointers to GetSaldosRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoQueryRecordMapper) ToSaldosRecordAll(saldos []*db.GetSaldosRow) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecordAll(saldo))
	}
	return saldoRecords
}

// ToSaldoRecordActive maps a GetActiveSaldosRow database row to a SaldoRecord domain model.
//
// Parameters:
//   - saldo: A pointer to a GetActiveSaldosRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoQueryRecordMapper) ToSaldoRecordActive(saldo *db.GetActiveSaldosRow) *record.SaldoRecord {
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

// ToSaldosRecordActive maps a slice of GetActiveSaldosRow database rows to a slice of SaldoRecord domain models.
//
// Parameters:
//   - saldos: A slice of pointers to GetActiveSaldosRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoQueryRecordMapper) ToSaldosRecordActive(saldos []*db.GetActiveSaldosRow) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecordActive(saldo))
	}
	return saldoRecords
}

// ToSaldoRecordTrashed maps a GetTrashedSaldosRow database row to a SaldoRecord domain model.
//
// Parameters:
//   - saldo: A pointer to a GetTrashedSaldosRow representing the trashed saldo database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including ID, CardNumber,
//     TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoQueryRecordMapper) ToSaldoRecordTrashed(saldo *db.GetTrashedSaldosRow) *record.SaldoRecord {
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

// ToSaldosRecordTrashed maps a slice of GetTrashedSaldosRow database rows to a slice of SaldoRecord domain models.
//
// Parameters:
//   - saldos: A slice of pointers to GetTrashedSaldosRow representing the trashed saldo database rows.
//
// Returns:
//   - A slice of pointers to SaldoRecord containing the mapped data, including ID, CardNumber,
//     TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoQueryRecordMapper) ToSaldosRecordTrashed(saldos []*db.GetTrashedSaldosRow) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecordTrashed(saldo))
	}
	return saldoRecords
}
