package cardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardBaseRecordMapper provides methods to map Card database rows to CardRecord domain models.
type CardBaseRecordMapper interface {
	// ToCardRecord maps a Card database row to a CardRecord domain model.
	//
	// Parameters:
	//   - card: A pointer to a Card representing the database row.
	//
	// Returns:
	//   - A pointer to a CardRecord containing the mapped data, including ID, UserID, CardNumber,
	//     CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardRecord(card *db.Card) *record.CardRecord
}

// CardCommandRecordMapper provides methods to map Card database rows to CardRecord domain models.
type CardCommandRecordMapper interface {
	CardBaseRecordMapper
}

// CardQueryRecordMapper provides methods to map Card database rows to CardRecord domain models.
type CardQueryRecordMapper interface {
	CardBaseRecordMapper

	// ToCardEmailRecord maps a GetUserEmailByCardNumberRow database row to a CardEmailRecord domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetUserEmailByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardEmailRecord containing the mapped data, including ID, UserID,
	//     Email, CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
	ToCardEmailRecord(card *db.GetUserEmailByCardNumberRow) *record.CardEmailRecord

	// ToCardGetAll maps a GetCardsRow database row to a CardRecord domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetCardsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardGetAll(card *db.GetCardsRow) *record.CardRecord
	// ToCardsRecord maps a slice of GetCardsRow database rows to a slice of CardRecord domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetCardsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardsRecord(cards []*db.GetCardsRow) []*record.CardRecord
	// ToCardRecordActive maps a GetActiveCardsWithCountRow database row to a CardRecord domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetActiveCardsWithCountRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardRecordActive(card *db.GetActiveCardsWithCountRow) *record.CardRecord
	// ToCardRecordsActive maps a slice of GetActiveCardsWithCountRow database rows to a slice of CardRecord domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetActiveCardsWithCountRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardRecordsActive(cards []*db.GetActiveCardsWithCountRow) []*record.CardRecord

	// ToCardRecordTrashed maps a GetTrashedCardsWithCountRow database row to a CardRecord domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetTrashedCardsWithCountRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardRecordTrashed(card *db.GetTrashedCardsWithCountRow) *record.CardRecord
	// ToCardRecordsTrashed maps a slice of GetTrashedCardsWithCountRow database rows to a slice of CardRecord domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetTrashedCardsWithCountRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardRecordsTrashed(cards []*db.GetTrashedCardsWithCountRow) []*record.CardRecord
}
