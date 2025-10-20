package transferstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferStatisticAmountRecordMapper maps transfer amount (monthly & yearly) to domain models.
type transferStatisticAmountRecordMapper struct{}

// NewTransferStatisticAmountRecordMapper returns a new instance of TransferStatisticAmountRecordMapper
// which provides methods to map GetMonthlyTransferAmountsRow and GetYearlyTransferAmountsRow
// database rows to TransferMonthAmount and TransferYearAmount domain models.
func NewTransferStatisticAmountRecordMapper() TransferStatisticAmountRecordMapper {
	return &transferStatisticAmountRecordMapper{}
}

// ToTransferMonthAmount maps a GetMonthlyTransferAmountsRow database row
// to a TransferMonthAmount domain model.
//
// Parameters:
//   - ss: A pointer to a GetMonthlyTransferAmountsRow representing the database row.
//
// Returns:
//   - A pointer to a TransferMonthAmount containing the mapped data, including Month
//     and TotalAmount.
func (s *transferStatisticAmountRecordMapper) ToTransferMonthAmount(ss *db.GetMonthlyTransferAmountsRow) *record.TransferMonthAmount {
	return &record.TransferMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferMonthAmounts maps a slice of GetMonthlyTransferAmountsRow database rows
// to a slice of TransferMonthAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyTransferAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month
//     and TotalAmount.
func (s *transferStatisticAmountRecordMapper) ToTransferMonthAmounts(ss []*db.GetMonthlyTransferAmountsRow) []*record.TransferMonthAmount {
	var transferRecords []*record.TransferMonthAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferMonthAmount(transfer))
	}

	return transferRecords
}

// ToTransferYearAmount maps a GetYearlyTransferAmountsRow database row to a TransferYearAmount domain model.
//
// Parameters:
//   - ss: A pointer to a GetYearlyTransferAmountsRow representing the database row.
//
// Returns:
//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferStatisticAmountRecordMapper) ToTransferYearAmount(ss *db.GetYearlyTransferAmountsRow) *record.TransferYearAmount {
	return &record.TransferYearAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferYearAmounts maps a slice of GetYearlyTransferAmountsRow database rows
// to a slice of TransferYearAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyTransferAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferStatisticAmountRecordMapper) ToTransferYearAmounts(ss []*db.GetYearlyTransferAmountsRow) []*record.TransferYearAmount {
	var transferRecords []*record.TransferYearAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferYearAmount(transfer))
	}

	return transferRecords
}
