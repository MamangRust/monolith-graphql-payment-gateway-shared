package cardstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticTransferByCardRecordMapper provides methods to map database rows to CardStatisticTransferByCardRecord domain models.
type cardStatisticTransferRecordMapper struct {
}

// NewCardStatisticTransferRecordMapper returns a new instance of CardStatisticTransferByCardRecordMapper
// which provides methods to map GetMonthlyTransferAmountBySenderRow and GetYearlyTransferAmountBySenderRow
// database rows to CardMonthAmount and CardYearAmount domain models.
func NewCardStatisticTransferRecordMapper() CardStatisticTransferByCardRecordMapper {
	return &cardStatisticTransferRecordMapper{}
}

// ToMonthlyTransferSenderAmountByCardNumber maps a GetMonthlyTransferAmountBySenderRow database row to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTransferAmountBySenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferSenderAmountByCardNumber(card *db.GetMonthlyTransferAmountBySenderRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalSentAmount),
	}
}

// ToMonthlyTransferSenderAmountsByCardNumber maps a slice of GetMonthlyTransferAmountBySenderRow database rows
// to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTransferAmountBySenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferSenderAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountBySenderRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferSenderAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTransferSenderAmountByCardNumber maps a GetYearlyTransferAmountBySenderRow database row
// to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTransferAmountBySenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferSenderAmountByCardNumber(card *db.GetYearlyTransferAmountBySenderRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalSentAmount,
	}
}

// ToYearlyTransferSenderAmountsByCardNumber maps a slice of GetYearlyTransferAmountBySenderRow database rows
// to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTransferAmountBySenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferSenderAmountsByCardNumber(cards []*db.GetYearlyTransferAmountBySenderRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferSenderAmountByCardNumber(card))
	}
	return records
}

// ToMonthlyTransferReceiverAmountByCardNumber maps a GetMonthlyTransferAmountByReceiverRow database row
// to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTransferAmountByReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferReceiverAmountByCardNumber(card *db.GetMonthlyTransferAmountByReceiverRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalReceivedAmount),
	}
}

// ToMonthlyTransferReceiverAmountsByCardNumber maps a slice of GetMonthlyTransferAmountByReceiverRow database rows
// to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTransferAmountByReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferReceiverAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountByReceiverRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferReceiverAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTransferReceiverAmountByCardNumber maps a GetYearlyTransferAmountByReceiverRow database row
// to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTransferAmountByReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferReceiverAmountByCardNumber(card *db.GetYearlyTransferAmountByReceiverRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalReceivedAmount,
	}
}

// ToYearlyTransferReceiverAmountsByCardNumber maps a slice of GetYearlyTransferAmountByReceiverRow database rows
// to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTransferAmountByReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferReceiverAmountsByCardNumber(cards []*db.GetYearlyTransferAmountByReceiverRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferReceiverAmountByCardNumber(card))
	}
	return records
}
