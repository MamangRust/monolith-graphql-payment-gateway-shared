package cardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardCommandRecordMapper provides methods to map Card database rows to CardRecord domain models for command operations.
type cardCommandRecordMapper struct {
}

// NewCardCommandRecordMapper returns a new instance of CardCommandRecordMapper which provides methods to map Card database rows to CardRecord domain models for command operations.
func NewCardCommandRecordMapper() CardCommandRecordMapper {
	return &cardCommandRecordMapper{}
}

// ToCardRecord maps a Card database row to a CardRecord domain model.
//
// Parameters:
//   - card: A pointer to a Card representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID, CardNumber,
//     CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardCommandRecordMapper) ToCardRecord(card *db.Card) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formatedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}
