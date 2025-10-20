package saldorecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// SaldoBaseRecordMapper provides methods to map Saldo database rows to SaldoRecord domain models.
type SaldoBaseRecordMapper interface {
	// ToSaldoRecord maps a Saldo database row to a SaldoRecord domain model.
	//
	// Parameters:
	//   - saldo: A pointer to a Saldo representing the database row.
	//
	// Returns:
	//   - A pointer to a SaldoRecord containing the mapped data, including
	//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord
}

// SaldoCommandRecordMapping maps raw Saldo rows from DB to Saldo domain models for command operations.
type SaldoCommandRecordMapping interface {
	SaldoBaseRecordMapper
}

// SaldoQueryRecordMapping maps query results to Saldo domain models.
type SaldoQueryRecordMapping interface {
	SaldoBaseRecordMapper

	// ToSaldoRecordAll maps a GetSaldosRow database row to a SaldoRecord domain model.
	//
	// Parameters:
	//   - saldo: A pointer to a GetSaldosRow representing the database row.
	//
	// Returns:
	//   - A pointer to a SaldoRecord containing the mapped data, including
	//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToSaldoRecordAll(saldo *db.GetSaldosRow) *record.SaldoRecord

	// ToSaldosRecordAll maps a slice of GetSaldosRow database rows to a slice of SaldoRecord domain models.
	//
	// Parameters:
	//   - saldos: A slice of pointers to GetSaldosRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to SaldoRecord containing the mapped data, including
	//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToSaldosRecordAll(saldos []*db.GetSaldosRow) []*record.SaldoRecord

	// ToSaldoRecordActive maps a GetActiveSaldosRow database row to a SaldoRecord domain model.
	//
	// Parameters:
	//   - saldo: A pointer to a GetActiveSaldosRow representing the database row.
	//
	// Returns:
	//   - A pointer to a SaldoRecord containing the mapped data, including
	//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToSaldoRecordActive(saldo *db.GetActiveSaldosRow) *record.SaldoRecord

	// ToSaldosRecordActive maps a slice of GetActiveSaldosRow database rows to a slice of SaldoRecord domain models.
	//
	// Parameters:
	//   - saldos: A slice of pointers to GetActiveSaldosRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to SaldoRecord containing the mapped data, including
	//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToSaldosRecordActive(saldos []*db.GetActiveSaldosRow) []*record.SaldoRecord

	// ToSaldoRecordTrashed maps a GetTrashedSaldosRow database row to a SaldoRecord domain model.
	//
	// Parameters:
	//   - saldo: A pointer to a GetTrashedSaldosRow representing the trashed saldo database row.
	//
	// Returns:
	//   - A pointer to a SaldoRecord containing the mapped data, including ID, CardNumber,
	//     TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToSaldoRecordTrashed(saldo *db.GetTrashedSaldosRow) *record.SaldoRecord

	// ToSaldosRecordTrashed maps a slice of GetTrashedSaldosRow database rows to a slice of SaldoRecord domain models.
	//
	// Parameters:
	//   - saldos: A slice of pointers to GetTrashedSaldosRow representing the trashed saldo database rows.
	//
	// Returns:
	//   - A slice of pointers to SaldoRecord containing the mapped data, including ID, CardNumber,
	//     TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToSaldosRecordTrashed(saldos []*db.GetTrashedSaldosRow) []*record.SaldoRecord
}

// SaldoStatisticRecordMapping maps statistical saldo rows to domain models.
type SaldoStatisticRecordMapping interface {
	// ToSaldoMonthTotalBalance maps a MonthlyTotalSaldoBalanceRow to a SaldoMonthTotalBalance domain model.
	ToSaldoMonthTotalBalance(ss *db.GetMonthlyTotalSaldoBalanceRow) *record.SaldoMonthTotalBalance

	// ToSaldoMonthTotalBalances maps a slice of MonthlyTotalSaldoBalanceRow to a slice of SaldoMonthTotalBalance domain models.
	ToSaldoMonthTotalBalances(ss []*db.GetMonthlyTotalSaldoBalanceRow) []*record.SaldoMonthTotalBalance

	// ToSaldoYearTotalBalance maps a YearlyTotalSaldoBalancesRow to a SaldoYearTotalBalance domain model.
	ToSaldoYearTotalBalance(ss *db.GetYearlyTotalSaldoBalancesRow) *record.SaldoYearTotalBalance

	// ToSaldoYearTotalBalances maps a slice of YearlyTotalSaldoBalancesRow to a slice of SaldoYearTotalBalance domain models.
	ToSaldoYearTotalBalances(ss []*db.GetYearlyTotalSaldoBalancesRow) []*record.SaldoYearTotalBalance

	// ToSaldoMonthBalance maps a MonthlySaldoBalancesRow to a SaldoMonthSaldoBalance domain model.
	ToSaldoMonthBalance(ss *db.GetMonthlySaldoBalancesRow) *record.SaldoMonthSaldoBalance

	// ToSaldoMonthBalances maps a slice of MonthlySaldoBalancesRow to a slice of SaldoMonthSaldoBalance domain models.
	ToSaldoMonthBalances(ss []*db.GetMonthlySaldoBalancesRow) []*record.SaldoMonthSaldoBalance

	// ToSaldoYearSaldoBalance maps a YearlySaldoBalancesRow to a SaldoYearSaldoBalance domain model.
	ToSaldoYearSaldoBalance(ss *db.GetYearlySaldoBalancesRow) *record.SaldoYearSaldoBalance

	// ToSaldoYearSaldoBalances maps a slice of YearlySaldoBalancesRow to a slice of SaldoYearSaldoBalance domain models.
	ToSaldoYearSaldoBalances(ss []*db.GetYearlySaldoBalancesRow) []*record.SaldoYearSaldoBalance
}
