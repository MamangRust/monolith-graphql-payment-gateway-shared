package record

// TransactionRecord represents a financial transaction in the system.
// Records details of payments made from a card to a merchant.
type TransactionRecord struct {
	ID              int     `json:"id"`               // Unique identifier for the transaction
	CardNumber      string  `json:"card_number"`      // Card number used (typically masked)
	TransactionNo   string  `json:"transaction_no"`   // Unique transaction reference number
	Amount          int     `json:"amount"`           // Transaction amount in smallest currency unit (e.g., cents)
	PaymentMethod   string  `json:"payment_method"`   // Method used (e.g., "credit", "debit", "wallet")
	MerchantID      int     `json:"merchant_id"`      // ID of the merchant receiving payment
	TransactionTime string  `json:"transaction_time"` // Timestamp when transaction occurred
	CreatedAt       string  `json:"created_at"`       // Timestamp when record was created
	UpdatedAt       string  `json:"updated_at"`       // Timestamp when record was last updated
	DeletedAt       *string `json:"deleted_at"`       // Timestamp when record was soft-deleted (nil if active)
}

// TransactionRecordMonthStatusSuccess represents successful transactions aggregated by month.
// Used for monthly success rate analysis and reporting.
type TransactionRecordMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	Month        string `json:"month"`         // Month in MM format
	TotalSuccess int    `json:"total_success"` // Count of successful transactions
	TotalAmount  int    `json:"total_amount"`  // Total amount processed successfully
}

// TransactionRecordYearStatusSuccess represents successful transactions aggregated by year.
// Used for annual success rate analysis and reporting.
type TransactionRecordYearStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	TotalSuccess int    `json:"total_success"` // Count of successful transactions
	TotalAmount  int    `json:"total_amount"`  // Total amount processed successfully
}

// TransactionRecordMonthStatusFailed represents failed transactions aggregated by month.
// Used for monthly failure analysis and reporting.
type TransactionRecordMonthStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	Month       string `json:"month"`        // Month in MM format
	TotalFailed int    `json:"total_failed"` // Count of failed transactions
	TotalAmount int    `json:"total_amount"` // Total amount of failed attempts
}

// TransactionRecordYearStatusFailed represents failed transactions aggregated by year.
// Used for annual failure analysis and reporting.
type TransactionRecordYearStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalFailed int    `json:"total_failed"` // Count of failed transactions
	TotalAmount int    `json:"total_amount"` // Total amount of failed attempts
}

// TransactionMonthMethod represents transactions aggregated by month and payment method.
// Used for analyzing payment method trends by month.
type TransactionMonthMethod struct {
	Month             string `json:"month"`              // Month in YYYY-MM format
	PaymentMethod     string `json:"payment_method"`     // Payment method used
	TotalTransactions int    `json:"total_transactions"` // Number of transactions
	TotalAmount       int    `json:"total_amount"`       // Total amount processed
}

// TransactionYearMethod represents transactions aggregated by year and payment method.
// Used for analyzing annual payment method trends.
type TransactionYearMethod struct {
	Year              string `json:"year"`               // Year in YYYY format
	PaymentMethod     string `json:"payment_method"`     // Payment method used
	TotalTransactions int    `json:"total_transactions"` // Number of transactions
	TotalAmount       int    `json:"total_amount"`       // Total amount processed
}

// TransactionMonthAmount represents total transaction amounts by month.
// Used for monthly financial reporting.
type TransactionMonthAmount struct {
	Month       string `json:"month"`        // Month in YYYY-MM format
	TotalAmount int    `json:"total_amount"` // Total transaction amount
}

// TransactionYearlyAmount represents total transaction amounts by year.
// Used for annual financial reporting.
type TransactionYearlyAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalAmount int    `json:"total_amount"` // Total transaction amount
}
