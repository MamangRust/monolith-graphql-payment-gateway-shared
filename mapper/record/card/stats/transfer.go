package cardstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticTransferRecordMapper provides methods to map database rows to CardStatisticTransferRecord domain models.
type cardStatisticTransferRecordMapper struct {
}

// NewCardStatisticTransferRecordMapper returns a new instance of CardStatisticTransferRecordMapper which provides methods
// to map database rows to CardStatisticTransferRecord domain models.
func NewCardStatisticTransferRecordMapper() CardStatisticTransferRecordMapper {
	return &cardStatisticTransferRecordMapper{}
}

// ToMonthlyTransferSenderAmount maps a GetMonthlyTransferAmountSenderRow database row to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTransferAmountSenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferSenderAmount(card *db.GetMonthlyTransferAmountSenderRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalSentAmount),
	}
}

// ToMonthlyTransferSenderAmounts maps a slice of GetMonthlyTransferAmountSenderRow database rows to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTransferAmountSenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferSenderAmounts(cards []*db.GetMonthlyTransferAmountSenderRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferSenderAmount(card))
	}
	return records
}

// ToYearlyTransferSenderAmount maps a GetYearlyTransferAmountSenderRow database row to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTransferAmountSenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferSenderAmount(card *db.GetYearlyTransferAmountSenderRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalSentAmount,
	}
}

// ToYearlyTransferSenderAmounts maps a slice of GetYearlyTransferAmountSenderRow database rows to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTransferAmountSenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferSenderAmounts(cards []*db.GetYearlyTransferAmountSenderRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferSenderAmount(card))
	}
	return records
}

// ToMonthlyTransferReceiverAmount maps a GetMonthlyTransferAmountReceiverRow database row to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTransferAmountReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferReceiverAmount(card *db.GetMonthlyTransferAmountReceiverRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalReceivedAmount),
	}
}

// ToMonthlyTransferReceiverAmounts maps a slice of GetMonthlyTransferAmountReceiverRow database rows to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTransferAmountReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToMonthlyTransferReceiverAmounts(cards []*db.GetMonthlyTransferAmountReceiverRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferReceiverAmount(card))
	}
	return records
}

// ToYearlyTransferReceiverAmount maps a GetYearlyTransferAmountReceiverRow database row to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTransferAmountReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferReceiverAmount(card *db.GetYearlyTransferAmountReceiverRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalReceivedAmount,
	}
}

// ToYearlyTransferReceiverAmounts maps a slice of GetYearlyTransferAmountReceiverRow database rows to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTransferAmountReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransferRecordMapper) ToYearlyTransferReceiverAmounts(cards []*db.GetYearlyTransferAmountReceiverRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferReceiverAmount(card))
	}
	return records
}
