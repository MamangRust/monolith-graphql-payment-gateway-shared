package transferstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferStatisticReceiverAmountByCardRecordMapper maps transfer amounts filtered by receiver card number.
type transferStatisticReceiverAmountByCardRecordMapper struct {
}

// NewTransferStatisticReceiverAmountByCardRecordMapper returns a new instance of
// TransferStatisticReceiverAmountByCardRecordMapper which provides methods to map
// transfer amounts filtered by receiver card number.
func NewTransferStatisticReceiverAmountByCardRecordMapper() TransferStatisticReceiverAmountByCardRecordMapper {
	return &transferStatisticReceiverAmountByCardRecordMapper{}
}

// ToTransferMonthAmountReceiver maps a GetMonthlyTransferAmountsByReceiverCardNumberRow database row
// to a TransferMonthAmount domain model.
//
// Parameters:
//   - ss: A pointer to a GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transferStatisticReceiverAmountByCardRecordMapper) ToTransferMonthAmountReceiver(ss *db.GetMonthlyTransferAmountsByReceiverCardNumberRow) *record.TransferMonthAmount {
	return &record.TransferMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferMonthAmountsReceiver maps a slice of GetMonthlyTransferAmountsByReceiverCardNumberRow database rows
// to a slice of TransferMonthAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transferStatisticReceiverAmountByCardRecordMapper) ToTransferMonthAmountsReceiver(ss []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) []*record.TransferMonthAmount {
	var transferRecords []*record.TransferMonthAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferMonthAmountReceiver(transfer))
	}

	return transferRecords
}

// ToTransferYearAmountReceiver maps a GetYearlyTransferAmountsByReceiverCardNumberRow database row
// to a TransferYearAmount domain model.
//
// Parameters:
//   - ss: A pointer to a GetYearlyTransferAmountsByReceiverCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferStatisticReceiverAmountByCardRecordMapper) ToTransferYearAmountReceiver(ss *db.GetYearlyTransferAmountsByReceiverCardNumberRow) *record.TransferYearAmount {
	return &record.TransferYearAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferYearAmountsReceiver maps a slice of GetYearlyTransferAmountsByReceiverCardNumberRow database rows
// to a slice of TransferYearAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyTransferAmountsByReceiverCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferStatisticReceiverAmountByCardRecordMapper) ToTransferYearAmountsReceiver(ss []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) []*record.TransferYearAmount {
	var transferRecords []*record.TransferYearAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferYearAmountReceiver(transfer))
	}

	return transferRecords
}
