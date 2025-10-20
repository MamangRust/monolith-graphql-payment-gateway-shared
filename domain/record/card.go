package record

// CardRecord represents a credit/debit card record in the system.
// It contains all the essential card information along with timestamps for tracking.
type CardRecord struct {
	ID           int     `json:"id"`            // Unique identifier for the card record
	UserID       int     `json:"user_id"`       // ID of the user who owns this card
	CardNumber   string  `json:"card_number"`   // Card number (typically masked in outputs)
	CardType     string  `json:"card_type"`     // Type of card (e.g., Visa, MasterCard)
	ExpireDate   string  `json:"expire_date"`   // Card expiration date in MM/YY format
	CVV          string  `json:"cvv"`           // Card verification value (typically masked)
	CardProvider string  `json:"card_provider"` // Provider or issuing bank of the card
	CreatedAt    string  `json:"created_at"`    // Timestamp when the record was created
	UpdatedAt    string  `json:"updated_at"`    // Timestamp when the record was last updated
	DeletedAt    *string `json:"deleted_at"`    // Timestamp when the record was soft-deleted (nil if not deleted)
}

// CardEmailRecord represents a card record with associated email information.
// This is typically used for transactions or notifications where email is required.
type CardEmailRecord struct {
	ID           int    `json:"id"`            // Unique identifier for the record
	Email        string `json:"email"`         // Email address associated with the card
	UserID       int    `json:"user_id"`       // ID of the user who owns this card
	CardNumber   string `json:"card_number"`   // Card number (typically masked)
	CardType     string `json:"card_type"`     // Type of card (e.g., Visa, MasterCard)
	ExpireDate   string `json:"expire_date"`   // Card expiration date in MM/YY format
	CVV          string `json:"cvv"`           // Card verification value (typically masked)
	CardProvider string `json:"card_provider"` // Provider or issuing bank of the card
	CreatedAt    string `json:"created_at"`    // Timestamp when the record was created
	UpdatedAt    string `json:"updated_at"`    // Timestamp when the record was last updated
}

// CardMonthBalance represents the total balance for a card for a specific month.
// Used for monthly balance tracking and reporting.
type CardMonthBalance struct {
	Month        string `json:"month"`         // Month in YYYY-MM format
	TotalBalance int64  `json:"total_balance"` // Total balance for the month in smallest currency unit (e.g., cents)
}

// CardYearlyBalance represents the total balance for a card for a specific year.
// Used for annual balance tracking and reporting.
type CardYearlyBalance struct {
	Year         string `json:"year"`          // Year in YYYY format
	TotalBalance int64  `json:"total_balance"` // Total balance for the year in smallest currency unit (e.g., cents)
}

// CardMonthAmount represents the total transaction amount for a card for a specific month.
// Used for monthly spending analysis and reporting.
type CardMonthAmount struct {
	Month       string `json:"month"`        // Month in YYYY-MM format
	TotalAmount int64  `json:"total_amount"` // Total transaction amount for the month in smallest currency unit
}

// CardYearAmount represents the total transaction amount for a card for a specific year.
// Used for annual spending analysis and reporting.
type CardYearAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalAmount int64  `json:"total_amount"` // Total transaction amount for the year in smallest currency unit
}
