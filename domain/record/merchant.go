package record

import (
	"time"
)

// MerchantRecord represents a merchant entity in the system.
// Merchants are businesses or service providers that can process transactions.
type MerchantRecord struct {
	ID        int     `json:"id"`         // Unique identifier for the merchant
	Name      string  `json:"name"`       // Legal name of the merchant
	ApiKey    string  `json:"api_key"`    // API key used for merchant authentication (typically masked)
	UserID    int     `json:"user_id"`    // ID of the user who owns/administers this merchant account
	Status    string  `json:"status"`     // Current status of the merchant (e.g., "active", "suspended")
	CreatedAt string  `json:"created_at"` // Timestamp when the merchant was created
	UpdatedAt string  `json:"updated_at"` // Timestamp when the merchant was last updated
	DeletedAt *string `json:"deleted_at"` // Timestamp when the merchant was soft-deleted (nil if not deleted)
}

// MerchantTransactionsRecord represents a transaction processed by a merchant.
type MerchantTransactionsRecord struct {
	TransactionID   int32     `json:"transaction_id"`   // Unique identifier for the transaction
	CardNumber      string    `json:"card_number"`      // Masked card number used for the transaction
	Amount          int32     `json:"amount"`           // Transaction amount in smallest currency unit (e.g., cents)
	PaymentMethod   string    `json:"payment_method"`   // Payment method used (e.g., "credit", "debit")
	MerchantID      int32     `json:"merchant_id"`      // ID of the merchant who processed the transaction
	MerchantName    string    `json:"merchant_name"`    // Name of the merchant at time of transaction
	TransactionTime time.Time `json:"transaction_time"` // Timestamp when the transaction occurred
	CreatedAt       string    `json:"created_at"`       // Timestamp when the record was created
	UpdatedAt       string    `json:"updated_at"`       // Timestamp when the record was last updated
	DeletedAt       *string   `json:"deleted_at"`       // Timestamp when the record was soft-deleted (nil if not deleted)
}

// MerchantYearlyPaymentMethod represents yearly transaction totals by payment method for a merchant.
// Used for financial reporting and analytics.
type MerchantYearlyPaymentMethod struct {
	Year          string `json:"year"`           // Year in YYYY format
	PaymentMethod string `json:"payment_method"` // Type of payment method (e.g., "credit", "debit")
	TotalAmount   int    `json:"total_amount"`   // Total amount processed for this payment method in the year
}

// MerchantMonthlyPaymentMethod represents monthly transaction totals by payment method for a merchant.
// Used for financial reporting and analytics.
type MerchantMonthlyPaymentMethod struct {
	Month         string `json:"month"`          // Month in YYYY-MM format
	PaymentMethod string `json:"payment_method"` // Type of payment method (e.g., "credit", "debit")
	TotalAmount   int    `json:"total_amount"`   // Total amount processed for this payment method in the month
}

// MerchantMonthlyAmount represents the total transaction amount for a merchant in a specific month.
type MerchantMonthlyAmount struct {
	Month       string `json:"month"`        // Month in YYYY-MM format
	TotalAmount int    `json:"total_amount"` // Total transaction amount for the month in smallest currency unit
}

// MerchantYearlyAmount represents the total transaction amount for a merchant in a specific year.
type MerchantYearlyAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalAmount int    `json:"total_amount"` // Total transaction amount for the year in smallest currency unit
}

// MerchantMonthlyTotalAmount represents the total transaction amount for a merchant in a specific month of a year.
// This is similar to MerchantMonthlyAmount but includes the year explicitly.
type MerchantMonthlyTotalAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	Month       string `json:"month"`        // Month in MM format
	TotalAmount int    `json:"total_amount"` // Total transaction amount for the month in smallest currency unit
}

// MerchantYearlyTotalAmount represents the total transaction amount for a merchant in a specific year.
// This is an alternative to MerchantYearlyAmount with potentially different aggregation.
type MerchantYearlyTotalAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalAmount int    `json:"total_amount"` // Total transaction amount for the year in smallest currency unit
}
