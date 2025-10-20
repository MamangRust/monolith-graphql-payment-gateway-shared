package record

// SaldoRecord represents a balance/withdrawal record for a card.
// Tracks the current balance and withdrawal transactions for a specific card.
type SaldoRecord struct {
	ID             int     `json:"id"`              // Unique identifier for the record
	CardNumber     string  `json:"card_number"`     // Card number associated with the balance (typically masked)
	TotalBalance   int     `json:"total_balance"`   // Current available balance in smallest currency unit (e.g., cents)
	WithdrawAmount int     `json:"withdraw_amount"` // Amount withdrawn in smallest currency unit (e.g., cents)
	WithdrawTime   string  `json:"withdraw_time"`   // Timestamp of when the withdrawal occurred
	CreatedAt      string  `json:"created_at"`      // Timestamp when the record was created
	UpdatedAt      string  `json:"updated_at"`      // Timestamp when the record was last updated
	DeletedAt      *string `json:"deleted_at"`      // Timestamp when record was soft-deleted (nil if active)
}

// SaldoMonthTotalBalance represents the total balance aggregated by month and year.
// Used for monthly balance reporting and analytics.
type SaldoMonthTotalBalance struct {
	Year         string `json:"year"`          // Year in YYYY format
	Month        string `json:"month"`         // Month in MM format
	TotalBalance int    `json:"total_balance"` // Aggregate balance for the month in smallest currency unit
}

// SaldoYearTotalBalance represents the total balance aggregated by year.
// Used for annual balance reporting and analytics.
type SaldoYearTotalBalance struct {
	Year         string `json:"year"`          // Year in YYYY format
	TotalBalance int    `json:"total_balance"` // Aggregate balance for the year in smallest currency unit
}

// SaldoMonthSaldoBalance represents monthly balance information.
// Similar to SaldoMonthTotalBalance but with different aggregation context.
type SaldoMonthSaldoBalance struct {
	Month        string `json:"month"`         // Month in YYYY-MM format
	TotalBalance int    `json:"total_balance"` // Aggregate balance for the month in smallest currency unit
}

// SaldoYearSaldoBalance represents yearly balance information.
// Similar to SaldoYearTotalBalance but with different aggregation context.
type SaldoYearSaldoBalance struct {
	Year         string `json:"year"`          // Year in YYYY format
	TotalBalance int    `json:"total_balance"` // Aggregate balance for the year in smallest currency unit
}
