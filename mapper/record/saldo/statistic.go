package saldorecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// SaldoStatisticRecordMapping provides methods to map Saldo database rows to SaldoRecord domain models for command operations.
type saldoStatisticRecordMapper struct{}

// NewSaldoStatisticRecordMapper returns a new instance of SaldoStatisticRecordMapping which provides methods to map Saldo database rows to SaldoRecord domain models for command operations.
func NewSaldoStatisticRecordMapper() SaldoStatisticRecordMapping {
	return &saldoStatisticRecordMapper{}
}

// ToSaldoMonthTotalBalance maps a GetMonthlyTotalSaldoBalanceRow database row to a SaldoMonthTotalBalance domain model.
//
// Parameters:
//   - ss: A pointer to a GetMonthlyTotalSaldoBalanceRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoMonthTotalBalance containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoMonthTotalBalance(ss *db.GetMonthlyTotalSaldoBalanceRow) *record.SaldoMonthTotalBalance {
	totalBalance := 0
	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &record.SaldoMonthTotalBalance{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// ToSaldoMonthTotalBalances maps a slice of GetMonthlyTotalSaldoBalanceRow database rows to a slice of SaldoMonthTotalBalance domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyTotalSaldoBalanceRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoMonthTotalBalance containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoMonthTotalBalances(ss []*db.GetMonthlyTotalSaldoBalanceRow) []*record.SaldoMonthTotalBalance {
	var saldoRecords []*record.SaldoMonthTotalBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoMonthTotalBalance(saldo))
	}
	return saldoRecords
}

// ToSaldoYearTotalBalance maps a GetYearlyTotalSaldoBalancesRow database row to a SaldoYearTotalBalance domain model.
//
// Parameters:
//   - ss: A pointer to a GetYearlyTotalSaldoBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoYearTotalBalance containing the mapped data, including Year, and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoYearTotalBalance(ss *db.GetYearlyTotalSaldoBalancesRow) *record.SaldoYearTotalBalance {
	totalBalance := 0
	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &record.SaldoYearTotalBalance{
		Year:         ss.Year,
		TotalBalance: int(totalBalance),
	}
}

// ToSaldoYearTotalBalances maps a slice of GetYearlyTotalSaldoBalancesRow database rows to a slice of SaldoYearTotalBalance domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyTotalSaldoBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoYearTotalBalance containing the mapped data, including Year, and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoYearTotalBalances(ss []*db.GetYearlyTotalSaldoBalancesRow) []*record.SaldoYearTotalBalance {
	var saldoRecords []*record.SaldoYearTotalBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoYearTotalBalance(saldo))
	}
	return saldoRecords
}

// ToSaldoMonthBalance maps a MonthlySaldoBalancesRow database row to a SaldoMonthSaldoBalance domain model.
//
// Parameters:
//   - ss: A pointer to a MonthlySaldoBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoMonthSaldoBalance containing the mapped data, including Month, and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoMonthBalance(ss *db.GetMonthlySaldoBalancesRow) *record.SaldoMonthSaldoBalance {
	return &record.SaldoMonthSaldoBalance{
		Month:        ss.Month,
		TotalBalance: int(ss.TotalBalance),
	}
}

// ToSaldoMonthBalances maps a slice of MonthlySaldoBalancesRow database rows to a slice of SaldoMonthSaldoBalance domain models.
//
// Parameters:
//   - ss: A slice of pointers to MonthlySaldoBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoMonthSaldoBalance containing the mapped data, including Month, and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoMonthBalances(ss []*db.GetMonthlySaldoBalancesRow) []*record.SaldoMonthSaldoBalance {
	var saldoRecords []*record.SaldoMonthSaldoBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoMonthBalance(saldo))
	}
	return saldoRecords
}

// ToSaldoYearSaldoBalance maps a GetYearlySaldoBalancesRow database row to a SaldoYearSaldoBalance domain model.
//
// Parameters:
//   - ss: A pointer to a GetYearlySaldoBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoYearSaldoBalance containing the mapped data, including Year and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoYearSaldoBalance(ss *db.GetYearlySaldoBalancesRow) *record.SaldoYearSaldoBalance {
	return &record.SaldoYearSaldoBalance{
		Year:         ss.Year,
		TotalBalance: int(ss.TotalBalance),
	}
}

// ToSaldoYearSaldoBalances maps a slice of GetYearlySaldoBalancesRow database rows to a slice of SaldoYearSaldoBalance domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlySaldoBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoYearSaldoBalance containing the mapped data, including Year and TotalBalance.
func (s *saldoStatisticRecordMapper) ToSaldoYearSaldoBalances(ss []*db.GetYearlySaldoBalancesRow) []*record.SaldoYearSaldoBalance {
	var saldoRecords []*record.SaldoYearSaldoBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoYearSaldoBalance(saldo))
	}
	return saldoRecords
}
