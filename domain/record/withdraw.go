package record

// WithdrawRecord represents a cash withdrawal transaction from a card/account.
// Records the details of funds being withdrawn from the system.
type WithdrawRecord struct {
	ID             int     `json:"id"`              // Unique identifier for the withdrawal
	WithdrawNo     string  `json:"withdraw_no"`     // Unique withdrawal reference number
	CardNumber     string  `json:"card_number"`     // Card number used (typically masked)
	WithdrawAmount int     `json:"withdraw_amount"` // Amount withdrawn in smallest currency unit (e.g., cents)
	WithdrawTime   string  `json:"withdraw_time"`   // Timestamp when withdrawal was processed
	CreatedAt      string  `json:"created_at"`      // Timestamp when record was created
	UpdatedAt      string  `json:"updated_at"`      // Timestamp when record was last updated
	DeletedAt      *string `json:"deleted_at"`      // Timestamp when record was soft-deleted (nil if active)
}

// WithdrawRecordMonthStatusSuccess represents successful withdrawals aggregated by month.
// Used for monthly success rate reporting and cash flow analysis.
type WithdrawRecordMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	Month        string `json:"month"`         // Month in MM format
	TotalSuccess int    `json:"total_success"` // Count of successful withdrawals
	TotalAmount  int    `json:"total_amount"`  // Total amount withdrawn successfully
}

// WithdrawRecordYearStatusSuccess represents successful withdrawals aggregated by year.
// Used for annual success rate reporting and financial analysis.
type WithdrawRecordYearStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	TotalSuccess int    `json:"total_success"` // Count of successful withdrawals
	TotalAmount  int    `json:"total_amount"`  // Total amount withdrawn successfully
}

// WithdrawRecordMonthStatusFailed represents failed withdrawal attempts aggregated by month.
// Used for monthly failure analysis and system health monitoring.
type WithdrawRecordMonthStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	Month       string `json:"month"`        // Month in MM format
	TotalFailed int    `json:"total_failed"` // Count of failed withdrawal attempts
	TotalAmount int    `json:"total_amount"` // Total amount of failed withdrawals
}

// WithdrawRecordYearStatusFailed represents failed withdrawal attempts aggregated by year.
// Used for annual failure analysis and system reliability reporting.
type WithdrawRecordYearStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalFailed int    `json:"total_failed"` // Count of failed withdrawal attempts
	TotalAmount int    `json:"total_amount"` // Total amount of failed withdrawals
}

// WithdrawMonthlyAmount represents total withdrawal amounts by month.
// Used for monthly cash flow analysis and liquidity reporting.
type WithdrawMonthlyAmount struct {
	Month       string `json:"month"`        // Month in YYYY-MM format
	TotalAmount int    `json:"total_amount"` // Total amount withdrawn
}

// WithdrawYearlyAmount represents total withdrawal amounts by year.
// Used for annual financial reporting and cash flow analysis.
type WithdrawYearlyAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalAmount int    `json:"total_amount"` // Total amount withdrawn
}
