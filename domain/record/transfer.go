package record

// TransferRecord represents a funds transfer transaction between accounts.
// Records the movement of money from one account/card to another.
type TransferRecord struct {
	ID             int     `json:"id"`              // Unique identifier for the transfer
	TransferNo     string  `json:"transfer_no"`     // Unique transaction reference number
	TransferFrom   string  `json:"transfer_from"`   // Source account/card (typically masked)
	TransferTo     string  `json:"transfer_to"`     // Destination account/card (typically masked)
	TransferAmount int     `json:"transfer_amount"` // Amount transferred in smallest currency unit (e.g., cents)
	TransferTime   string  `json:"transfer_time"`   // Timestamp when transfer was executed
	CreatedAt      string  `json:"created_at"`      // Timestamp when record was created
	UpdatedAt      string  `json:"updated_at"`      // Timestamp when record was last updated
	DeletedAt      *string `json:"deleted_at"`      // Timestamp when record was soft-deleted (nil if active)
}

// TransferRecordMonthStatusSuccess represents successful transfers aggregated by month.
// Used for monthly success rate analysis and reporting.
type TransferRecordMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	Month        string `json:"month"`         // Month in MM format
	TotalSuccess int    `json:"total_success"` // Count of successful transfers
	TotalAmount  int    `json:"total_amount"`  // Total amount transferred successfully
}

// TransferRecordYearStatusSuccess represents successful transfers aggregated by year.
// Used for annual success rate analysis and reporting.
type TransferRecordYearStatusSuccess struct {
	Year         string `json:"year"`          // Year in YYYY format
	TotalSuccess int    `json:"total_success"` // Count of successful transfers
	TotalAmount  int    `json:"total_amount"`  // Total amount transferred successfully
}

// TransferRecordMonthStatusFailed represents failed transfers aggregated by month.
// Used for monthly failure analysis and reporting.
type TransferRecordMonthStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	Month       string `json:"month"`        // Month in MM format
	TotalFailed int    `json:"total_failed"` // Count of failed transfer attempts
	TotalAmount int    `json:"total_amount"` // Total amount of failed transfers
}

// TransferRecordYearStatusFailed represents failed transfers aggregated by year.
// Used for annual failure analysis and reporting.
type TransferRecordYearStatusFailed struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalFailed int    `json:"total_failed"` // Count of failed transfer attempts
	TotalAmount int    `json:"total_amount"` // Total amount of failed transfers
}

// TransferMonthAmount represents total transfer amounts by month.
// Used for monthly financial reporting and analysis.
type TransferMonthAmount struct {
	Month       string `json:"month"`        // Month in YYYY-MM format
	TotalAmount int    `json:"total_amount"` // Total amount transferred
}

// TransferYearAmount represents total transfer amounts by year.
// Used for annual financial reporting and analysis.
type TransferYearAmount struct {
	Year        string `json:"year"`         // Year in YYYY format
	TotalAmount int    `json:"total_amount"` // Total amount transferred
}
