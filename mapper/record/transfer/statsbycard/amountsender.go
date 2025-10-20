package transferstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferStatisticSenderAmountByCardRecordMapper maps transfer amounts filtered by receiver card number.
type transferStatisticSenderAmountByCardRecordMapper struct {
}

// NewTransferStatisticSenderAmountByCardRecordMapper returns a new instance of TransferStatisticSenderAmountByCardRecordMapper
// which provides methods to map transfer amounts filtered by receiver card number.
func NewTransferStatisticSenderAmountByCardRecordMapper() TransferStatisticSenderAmountByCardRecordMapper {
	return &transferStatisticSenderAmountByCardRecordMapper{}
}

// ToTransferMonthAmountSender maps a GetMonthlyTransferAmountsBySenderCardNumberRow database row
// to a TransferMonthAmount domain model.
//
// Parameters:
//   - ss: A pointer to a GetMonthlyTransferAmountsBySenderCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferMonthAmount containing the mapped data, including Month
//     and TotalAmount.
func (s *transferStatisticSenderAmountByCardRecordMapper) ToTransferMonthAmountSender(ss *db.GetMonthlyTransferAmountsBySenderCardNumberRow) *record.TransferMonthAmount {
	return &record.TransferMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferMonthAmountsSender maps a slice of GetMonthlyTransferAmountsBySenderCardNumberRow
// database rows to a slice of TransferMonthAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyTransferAmountsBySenderCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transferStatisticSenderAmountByCardRecordMapper) ToTransferMonthAmountsSender(ss []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) []*record.TransferMonthAmount {
	var transferRecords []*record.TransferMonthAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferMonthAmountSender(transfer))
	}

	return transferRecords
}

// ToTransferYearAmountSender maps a GetYearlyTransferAmountsBySenderCardNumberRow database row
// to a TransferYearAmount domain model.
//
// Parameters:
//   - ss: A pointer to a GetYearlyTransferAmountsBySenderCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferStatisticSenderAmountByCardRecordMapper) ToTransferYearAmountSender(ss *db.GetYearlyTransferAmountsBySenderCardNumberRow) *record.TransferYearAmount {
	return &record.TransferYearAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferYearAmountsSender maps a slice of GetYearlyTransferAmountsBySenderCardNumberRow database rows
// to a slice of TransferYearAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyTransferAmountsBySenderCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferStatisticSenderAmountByCardRecordMapper) ToTransferYearAmountsSender(ss []*db.GetYearlyTransferAmountsBySenderCardNumberRow) []*record.TransferYearAmount {
	var transferRecords []*record.TransferYearAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferYearAmountSender(transfer))
	}

	return transferRecords
}
