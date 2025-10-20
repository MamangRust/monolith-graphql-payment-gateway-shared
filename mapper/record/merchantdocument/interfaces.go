package merchantdocumentrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantDocumentBaseMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
type MerchantDocumentBaseMapper interface {
	// MerchantDocumentRecord maps a MerchantDocument database row to a MerchantDocumentRecord domain model.
	//
	// Parameters:
	//   - merchantdocument: A pointer to a MerchantDocument representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
	//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
	ToGetMerchantDocument(doc *db.MerchantDocument) *record.MerchantDocumentRecord
}

// MerchantDocumentCommandMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
type MerchantDocumentCommandMapper interface {
	MerchantDocumentBaseMapper
}

// MerchantDocumentQueryMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
type MerchantDocumentQueryMapper interface {
	MerchantDocumentBaseMapper

	// ToMerchantDocumentRecord maps a GetMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.
	//
	// Parameters:
	//   - doc: A pointer to a GetMerchantDocumentsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
	//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantDocumentRecord(doc *db.GetMerchantDocumentsRow) *record.MerchantDocumentRecord
	// ToMerchantDocumentsRecord maps a slice of GetMerchantDocumentsRow database rows to a slice of MerchantDocumentRecord domain models.
	//
	// Parameters:
	//   - docs: A slice of pointers to GetMerchantDocumentsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
	//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantDocumentsRecord(docs []*db.GetMerchantDocumentsRow) []*record.MerchantDocumentRecord

	// ToMerchantDocumentActiveRecord maps a GetActiveMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.
	//
	// Parameters:
	//   - doc: A pointer to a GetActiveMerchantDocumentsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
	//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantDocumentActiveRecord(doc *db.GetActiveMerchantDocumentsRow) *record.MerchantDocumentRecord
	// ToMerchantDocumentsActiveRecord maps a slice of GetActiveMerchantDocumentsRow database rows to a slice of MerchantDocumentRecord domain models.
	//
	// Parameters:
	//   - docs: A slice of pointers to GetActiveMerchantDocumentsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
	//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantDocumentsActiveRecord(docs []*db.GetActiveMerchantDocumentsRow) []*record.MerchantDocumentRecord

	// ToMerchantDocumentTrashedRecord maps a GetTrashedMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.
	//
	// Parameters:
	//   - doc: A pointer to a GetTrashedMerchantDocumentsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
	//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantDocumentTrashedRecord(doc *db.GetTrashedMerchantDocumentsRow) *record.MerchantDocumentRecord
	// ToMerchantDocumentsTrashedRecord maps a slice of GetTrashedMerchantDocumentsRow database rows to a slice of
	// MerchantDocumentRecord domain models.
	//
	// Parameters:
	//   - docs: A slice of pointers to GetTrashedMerchantDocumentsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
	//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantDocumentsTrashedRecord(docs []*db.GetTrashedMerchantDocumentsRow) []*record.MerchantDocumentRecord
}
