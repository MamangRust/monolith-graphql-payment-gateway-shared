package record

// MerchantDocumentRecord represents a document associated with a merchant.
// This is typically used for verification, compliance, or contractual documents.
type MerchantDocumentRecord struct {
	ID           int     `json:"id"`            // Unique identifier for the document record
	MerchantID   int     `json:"merchant_id"`   // ID of the merchant this document belongs to
	DocumentType string  `json:"document_type"` // Type of document (e.g., "license", "contract", "identity")
	DocumentURL  string  `json:"document_url"`  // URL or path where the document is stored
	Status       string  `json:"status"`        // Current status of the document (e.g., "pending", "approved", "rejected")
	Note         string  `json:"note"`          // Additional notes or comments about the document
	CreatedAt    string  `json:"created_at"`    // Timestamp when the record was created
	UpdatedAt    string  `json:"updated_at"`    // Timestamp when the record was last updated
	DeletedAt    *string `json:"deleted_at"`    // Timestamp when the record was soft-deleted (nil if not deleted)
}
