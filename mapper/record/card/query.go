package cardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardQueryRecordMapper provides methods to map Card database rows to CardRecord domain models for query operations.
type cardQueryRecordMapper struct {
}

// NewCardQueryRecordMapper creates a new instance of CardQueryRecordMapper which provides
// methods to map Card database rows to CardRecord domain models. This function returns
// a CardQueryRecordMapper interface that can be used for query operations.
func NewCardQueryRecordMapper() CardQueryRecordMapper {
	return &cardQueryRecordMapper{}
}

// ToCardRecord maps a Card database row to a CardRecord domain model.
//
// Parameters:
//   - card: A pointer to a Card representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID, CardNumber,
//     CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryRecordMapper) ToCardRecord(card *db.Card) *record.CardRecord {
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

// ToCardEmailRecord maps a GetUserEmailByCardNumberRow database row to a CardEmailRecord domain model.
//
// Parameters:
//   - card: A pointer to a GetUserEmailByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardEmailRecord containing the mapped data, including ID, UserID,
//     Email, CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardQueryRecordMapper) ToCardEmailRecord(card *db.GetUserEmailByCardNumberRow) *record.CardEmailRecord {

	return &record.CardEmailRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		Email:        card.Email,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
	}
}

// ToCardGetAll maps a GetCardsRow database row to a CardRecord domain model.
//
// Parameters:
//   - card: A pointer to a GetCardsRow representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryRecordMapper) ToCardGetAll(card *db.GetCardsRow) *record.CardRecord {
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

// ToCardsRecord maps a slice of GetCardsRow database rows to a slice of CardRecord domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetCardsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryRecordMapper) ToCardsRecord(cards []*db.GetCardsRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardGetAll(card))
	}
	return records
}

// ToCardRecordActive maps a GetActiveCardsWithCountRow database row to a CardRecord domain model.
//
// Parameters:
//   - card: A pointer to a GetActiveCardsWithCountRow representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryRecordMapper) ToCardRecordActive(card *db.GetActiveCardsWithCountRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formattedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
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

// ToCardRecordsActive maps a slice of GetActiveCardsWithCountRow database rows to a slice of CardRecord domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetActiveCardsWithCountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryRecordMapper) ToCardRecordsActive(cards []*db.GetActiveCardsWithCountRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardRecordActive(card))
	}
	return records
}

// ToCardRecordTrashed maps a GetTrashedCardsWithCountRow database row to a CardRecord domain model.
//
// Parameters:
//   - card: A pointer to a GetTrashedCardsWithCountRow representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryRecordMapper) ToCardRecordTrashed(card *db.GetTrashedCardsWithCountRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formattedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
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

// ToCardRecordsTrashed maps a slice of GetTrashedCardsWithCountRow database rows to a slice of CardRecord domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetTrashedCardsWithCountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryRecordMapper) ToCardRecordsTrashed(cards []*db.GetTrashedCardsWithCountRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardRecordTrashed(card))
	}
	return records
}
