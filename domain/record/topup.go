package record

// TopupRecord represents a single top-up transaction in the system.
// Records the details of funds being added to a card/account.
type TopupRecord struct {
	ID          int     `json:"id"`           // Unique identifier for the top-up record
	CardNumber  string  `json:"card_number"`  // Card number being topped up (typically masked)
	TopupNo     string  `json:"topup_no"`     // Unique reference number for the top-up transaction
	TopupAmount int     `json:"topup_amount"` // Amount added in smallest currency unit (e.g., cents)
	TopupMethod string  `json:"topup_method"` // Payment method used (e.g., "bank_transfer", "credit_card")
	TopupTime   string  `json:"topup_time"`   // Timestamp when the top-up was processed
	CreatedAt   string  `json:"created_at"`   // Timestamp when record was created
	UpdatedAt   string  `json:"updated_at"`   // Timestamp when record was last updated
	DeletedAt   *string `json:"deleted_at"`   // Timestamp when record was soft-deleted (nil if active)
}

// TopupRecordMonthStatusSuccess represents aggregated successful top-ups by month.
// Used for monthly success rate reporting and analytics.
type TopupRecordMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	Month        string `json:"month"`         // Month in MM format
	TotalSuccess int    `json:"total_success"` // Count of successful top-ups in the month
	TotalAmount  int    `json:"total_amount"`  // Total amount of successful top-ups in smallest currency unit
}

// TopupRecordYearStatusSuccess represents aggregated successful top-ups by year.
// Used for annual success rate reporting and analytics.
type TopupRecordYearStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	TotalSuccess int    `json:"total_success"` // Count of successful top-ups in the year
	TotalAmount  int    `json:"total_amount"`  // Total amount of successful top-ups in smallest currency unit
}

// TopupRecordMonthStatusFailed represents aggregated failed top-ups by month.
// Used for monthly failure rate reporting and analytics.
type TopupRecordMonthStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	Month       string `json:"month"`        // Month in MM format
	TotalFailed int    `json:"total_failed"` // Count of failed top-ups in the month
	TotalAmount int    `json:"total_amount"` // Total amount of failed top-up attempts in smallest currency unit
}

// TopupRecordYearStatusFailed represents aggregated failed top-ups by year.
// Used for annual failure rate reporting and analytics.
type TopupRecordYearStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalFailed int    `json:"total_failed"` // Count of failed top-ups in the year
	TotalAmount int    `json:"total_amount"` // Total amount of failed top-up attempts in smallest currency unit
}

// TopupMonthMethod represents top-ups aggregated by month and payment method.
// Used for analyzing payment method trends and preferences.
type TopupMonthMethod struct {
	Month       string `json:"month"`        // Month in YYYY-MM format
	TopupMethod string `json:"topup_method"` // Payment method used (e.g., "bank_transfer")
	TotalTopups int    `json:"total_topups"` // Number of top-ups using this method in the month
	TotalAmount int    `json:"total_amount"` // Total amount processed via this method in smallest currency unit
}

// TopupYearlyMethod represents top-ups aggregated by year and payment method.
// Used for annual analysis of payment method trends.
type TopupYearlyMethod struct {
	Year        string `json:"year"`         // Year in YYYY format
	TopupMethod string `json:"topup_method"` // Payment method used (e.g., "credit_card")
	TotalTopups int    `json:"total_topups"` // Number of top-ups using this method in the year
	TotalAmount int    `json:"total_amount"` // Total amount processed via this method in smallest currency unit
}

// TopupMonthAmount represents total top-up amounts by month.
// Used for monthly financial reporting and analysis.
type TopupMonthAmount struct {
	Month       string `json:"month"`        // Month in YYYY-MM format
	TotalAmount int    `json:"total_amount"` // Total top-up amount for the month in smallest currency unit
}

// TopupYearlyAmount represents total top-up amounts by year.
// Used for annual financial reporting and analysis.
type TopupYearlyAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalAmount int    `json:"total_amount"` // Total top-up amount for the year in smallest currency unit
}
