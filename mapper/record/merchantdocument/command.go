package merchantdocumentrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantDocumentCommandMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
type merchantDocumentCommandMapper struct{}

// NewMerchantDocumentCommandMapper returns a new instance of
// merchantDocumentCommandMapper which provides methods to map
// MerchantDocument database rows to MerchantDocumentRecord domain models.
func NewMerchantDocumentCommandMapper() MerchantDocumentCommandMapper {
	return &merchantDocumentCommandMapper{}
}

// MerchantDocumentRecord maps a MerchantDocument database row to a MerchantDocumentRecord domain model.
//
// Args:
//   - merchantdocument: A pointer to a MerchantDocument representing the database row.
//
// Returns:
//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentCommandMapper) ToGetMerchantDocument(doc *db.MerchantDocument) *record.MerchantDocumentRecord {
	return &record.MerchantDocumentRecord{
		ID:           int(doc.DocumentID),
		MerchantID:   int(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note.String,
		CreatedAt:    doc.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    doc.UpdatedAt.Time.Format("2006-01-02"),
	}
}
