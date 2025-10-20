# 📦 Package `record`

**Source Path:** `shared/domain/record`

## 🧩 Types

### `CardEmailRecord`

CardEmailRecord represents a card record with associated email information.
This is typically used for transactions or notifications where email is required.

```go
type CardEmailRecord struct {
	ID int `json:"id"`
	Email string `json:"email"`
	UserID int `json:"user_id"`
	CardNumber string `json:"card_number"`
	CardType string `json:"card_type"`
	ExpireDate string `json:"expire_date"`
	CVV string `json:"cvv"`
	CardProvider string `json:"card_provider"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `CardMonthAmount`

CardMonthAmount represents the total transaction amount for a card for a specific month.
Used for monthly spending analysis and reporting.

```go
type CardMonthAmount struct {
	Month string `json:"month"`
	TotalAmount int64 `json:"total_amount"`
}
```

### `CardMonthBalance`

CardMonthBalance represents the total balance for a card for a specific month.
Used for monthly balance tracking and reporting.

```go
type CardMonthBalance struct {
	Month string `json:"month"`
	TotalBalance int64 `json:"total_balance"`
}
```

### `CardRecord`

CardRecord represents a credit/debit card record in the system.
It contains all the essential card information along with timestamps for tracking.

```go
type CardRecord struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	CardNumber string `json:"card_number"`
	CardType string `json:"card_type"`
	ExpireDate string `json:"expire_date"`
	CVV string `json:"cvv"`
	CardProvider string `json:"card_provider"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `CardYearAmount`

CardYearAmount represents the total transaction amount for a card for a specific year.
Used for annual spending analysis and reporting.

```go
type CardYearAmount struct {
	Year string `json:"year"`
	TotalAmount int64 `json:"total_amount"`
}
```

### `CardYearlyBalance`

CardYearlyBalance represents the total balance for a card for a specific year.
Used for annual balance tracking and reporting.

```go
type CardYearlyBalance struct {
	Year string `json:"year"`
	TotalBalance int64 `json:"total_balance"`
}
```

### `MerchantDocumentRecord`

MerchantDocumentRecord represents a document associated with a merchant.
This is typically used for verification, compliance, or contractual documents.

```go
type MerchantDocumentRecord struct {
	ID int `json:"id"`
	MerchantID int `json:"merchant_id"`
	DocumentType string `json:"document_type"`
	DocumentURL string `json:"document_url"`
	Status string `json:"status"`
	Note string `json:"note"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `MerchantMonthlyAmount`

MerchantMonthlyAmount represents the total transaction amount for a merchant in a specific month.

```go
type MerchantMonthlyAmount struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantMonthlyPaymentMethod`

MerchantMonthlyPaymentMethod represents monthly transaction totals by payment method for a merchant.
Used for financial reporting and analytics.

```go
type MerchantMonthlyPaymentMethod struct {
	Month string `json:"month"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantMonthlyTotalAmount`

MerchantMonthlyTotalAmount represents the total transaction amount for a merchant in a specific month of a year.
This is similar to MerchantMonthlyAmount but includes the year explicitly.

```go
type MerchantMonthlyTotalAmount struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantRecord`

MerchantRecord represents a merchant entity in the system.
Merchants are businesses or service providers that can process transactions.

```go
type MerchantRecord struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
	UserID int `json:"user_id"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `MerchantTransactionsRecord`

MerchantTransactionsRecord represents a transaction processed by a merchant.

```go
type MerchantTransactionsRecord struct {
	TransactionID int32 `json:"transaction_id"`
	CardNumber string `json:"card_number"`
	Amount int32 `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	MerchantID int32 `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	TransactionTime time.Time `json:"transaction_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `MerchantYearlyAmount`

MerchantYearlyAmount represents the total transaction amount for a merchant in a specific year.

```go
type MerchantYearlyAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantYearlyPaymentMethod`

MerchantYearlyPaymentMethod represents yearly transaction totals by payment method for a merchant.
Used for financial reporting and analytics.

```go
type MerchantYearlyPaymentMethod struct {
	Year string `json:"year"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantYearlyTotalAmount`

MerchantYearlyTotalAmount represents the total transaction amount for a merchant in a specific year.
This is an alternative to MerchantYearlyAmount with potentially different aggregation.

```go
type MerchantYearlyTotalAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `RefreshTokenRecord`

RefreshTokenRecord represents a refresh token used for maintaining user sessions.
Refresh tokens allow users to obtain new access tokens without re-authenticating.

```go
type RefreshTokenRecord struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Token string `json:"token"`
	ExpiredAt string `json:"expired_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `ResetTokenRecord`

ResetTokenRecord represents a password reset token used for user authentication recovery.
These tokens are typically short-lived and sent to users via email for password reset flows.

```go
type ResetTokenRecord struct {
	ID int64 `json:"id"`
	Token string `json:"token"`
	UserID int64 `json:"user_id"`
	ExpiredAt string `json:"expired_at"`
}
```

### `RoleRecord`

RoleRecord represents a role in the system that can be assigned to users.
Roles are used to define permissions and access levels for different user types.

```go
type RoleRecord struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `SaldoMonthSaldoBalance`

SaldoMonthSaldoBalance represents monthly balance information.
Similar to SaldoMonthTotalBalance but with different aggregation context.

```go
type SaldoMonthSaldoBalance struct {
	Month string `json:"month"`
	TotalBalance int `json:"total_balance"`
}
```

### `SaldoMonthTotalBalance`

SaldoMonthTotalBalance represents the total balance aggregated by month and year.
Used for monthly balance reporting and analytics.

```go
type SaldoMonthTotalBalance struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalBalance int `json:"total_balance"`
}
```

### `SaldoRecord`

SaldoRecord represents a balance/withdrawal record for a card.
Tracks the current balance and withdrawal transactions for a specific card.

```go
type SaldoRecord struct {
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	TotalBalance int `json:"total_balance"`
	WithdrawAmount int `json:"withdraw_amount"`
	WithdrawTime string `json:"withdraw_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `SaldoYearSaldoBalance`

SaldoYearSaldoBalance represents yearly balance information.
Similar to SaldoYearTotalBalance but with different aggregation context.

```go
type SaldoYearSaldoBalance struct {
	Year string `json:"year"`
	TotalBalance int `json:"total_balance"`
}
```

### `SaldoYearTotalBalance`

SaldoYearTotalBalance represents the total balance aggregated by year.
Used for annual balance reporting and analytics.

```go
type SaldoYearTotalBalance struct {
	Year string `json:"year"`
	TotalBalance int `json:"total_balance"`
}
```

### `TopupMonthAmount`

TopupMonthAmount represents total top-up amounts by month.
Used for monthly financial reporting and analysis.

```go
type TopupMonthAmount struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupMonthMethod`

TopupMonthMethod represents top-ups aggregated by month and payment method.
Used for analyzing payment method trends and preferences.

```go
type TopupMonthMethod struct {
	Month string `json:"month"`
	TopupMethod string `json:"topup_method"`
	TotalTopups int `json:"total_topups"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupRecord`

TopupRecord represents a single top-up transaction in the system.
Records the details of funds being added to a card/account.

```go
type TopupRecord struct {
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	TopupNo string `json:"topup_no"`
	TopupAmount int `json:"topup_amount"`
	TopupMethod string `json:"topup_method"`
	TopupTime string `json:"topup_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `TopupRecordMonthStatusFailed`

TopupRecordMonthStatusFailed represents aggregated failed top-ups by month.
Used for monthly failure rate reporting and analytics.

```go
type TopupRecordMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupRecordMonthStatusSuccess`

TopupRecordMonthStatusSuccess represents aggregated successful top-ups by month.
Used for monthly success rate reporting and analytics.

```go
type TopupRecordMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupRecordYearStatusFailed`

TopupRecordYearStatusFailed represents aggregated failed top-ups by year.
Used for annual failure rate reporting and analytics.

```go
type TopupRecordYearStatusFailed struct {
	Year string `json:"year"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupRecordYearStatusSuccess`

TopupRecordYearStatusSuccess represents aggregated successful top-ups by year.
Used for annual success rate reporting and analytics.

```go
type TopupRecordYearStatusSuccess struct {
	Year string `json:"year"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupYearlyAmount`

TopupYearlyAmount represents total top-up amounts by year.
Used for annual financial reporting and analysis.

```go
type TopupYearlyAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupYearlyMethod`

TopupYearlyMethod represents top-ups aggregated by year and payment method.
Used for annual analysis of payment method trends.

```go
type TopupYearlyMethod struct {
	Year string `json:"year"`
	TopupMethod string `json:"topup_method"`
	TotalTopups int `json:"total_topups"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionMonthAmount`

TransactionMonthAmount represents total transaction amounts by month.
Used for monthly financial reporting.

```go
type TransactionMonthAmount struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionMonthMethod`

TransactionMonthMethod represents transactions aggregated by month and payment method.
Used for analyzing payment method trends by month.

```go
type TransactionMonthMethod struct {
	Month string `json:"month"`
	PaymentMethod string `json:"payment_method"`
	TotalTransactions int `json:"total_transactions"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionRecord`

TransactionRecord represents a financial transaction in the system.
Records details of payments made from a card to a merchant.

```go
type TransactionRecord struct {
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	TransactionNo string `json:"transaction_no"`
	Amount int `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	MerchantID int `json:"merchant_id"`
	TransactionTime string `json:"transaction_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `TransactionRecordMonthStatusFailed`

TransactionRecordMonthStatusFailed represents failed transactions aggregated by month.
Used for monthly failure analysis and reporting.

```go
type TransactionRecordMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionRecordMonthStatusSuccess`

TransactionRecordMonthStatusSuccess represents successful transactions aggregated by month.
Used for monthly success rate analysis and reporting.

```go
type TransactionRecordMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionRecordYearStatusFailed`

TransactionRecordYearStatusFailed represents failed transactions aggregated by year.
Used for annual failure analysis and reporting.

```go
type TransactionRecordYearStatusFailed struct {
	Year string `json:"year"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionRecordYearStatusSuccess`

TransactionRecordYearStatusSuccess represents successful transactions aggregated by year.
Used for annual success rate analysis and reporting.

```go
type TransactionRecordYearStatusSuccess struct {
	Year string `json:"year"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionYearMethod`

TransactionYearMethod represents transactions aggregated by year and payment method.
Used for analyzing annual payment method trends.

```go
type TransactionYearMethod struct {
	Year string `json:"year"`
	PaymentMethod string `json:"payment_method"`
	TotalTransactions int `json:"total_transactions"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionYearlyAmount`

TransactionYearlyAmount represents total transaction amounts by year.
Used for annual financial reporting.

```go
type TransactionYearlyAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferMonthAmount`

TransferMonthAmount represents total transfer amounts by month.
Used for monthly financial reporting and analysis.

```go
type TransferMonthAmount struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferRecord`

TransferRecord represents a funds transfer transaction between accounts.
Records the movement of money from one account/card to another.

```go
type TransferRecord struct {
	ID int `json:"id"`
	TransferNo string `json:"transfer_no"`
	TransferFrom string `json:"transfer_from"`
	TransferTo string `json:"transfer_to"`
	TransferAmount int `json:"transfer_amount"`
	TransferTime string `json:"transfer_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `TransferRecordMonthStatusFailed`

TransferRecordMonthStatusFailed represents failed transfers aggregated by month.
Used for monthly failure analysis and reporting.

```go
type TransferRecordMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferRecordMonthStatusSuccess`

TransferRecordMonthStatusSuccess represents successful transfers aggregated by month.
Used for monthly success rate analysis and reporting.

```go
type TransferRecordMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferRecordYearStatusFailed`

TransferRecordYearStatusFailed represents failed transfers aggregated by year.
Used for annual failure analysis and reporting.

```go
type TransferRecordYearStatusFailed struct {
	Year string `json:"year"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferRecordYearStatusSuccess`

TransferRecordYearStatusSuccess represents successful transfers aggregated by year.
Used for annual success rate analysis and reporting.

```go
type TransferRecordYearStatusSuccess struct {
	Year string `json:"year"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferYearAmount`

TransferYearAmount represents total transfer amounts by year.
Used for annual financial reporting and analysis.

```go
type TransferYearAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `UserRecord`

UserRecord represents a user account in the system.
Contains core user information and authentication details.

```go
type UserRecord struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	VerifiedCode string `json:"verified_code"`
	IsVerified bool `json:"is_verified"`
	ConfirmPassword string `json:"confirm_password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `UserRoleRecord`

UserRoleRecord represents the association between a user and a role.
This is a join table that defines which roles are assigned to which users.

```go
type UserRoleRecord struct {
	UserRoleID int32 `json:"user_role_id"`
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
	RoleName string `json:"role_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
```

### `WithdrawMonthlyAmount`

WithdrawMonthlyAmount represents total withdrawal amounts by month.
Used for monthly cash flow analysis and liquidity reporting.

```go
type WithdrawMonthlyAmount struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `WithdrawRecord`

WithdrawRecord represents a cash withdrawal transaction from a card/account.
Records the details of funds being withdrawn from the system.

```go
type WithdrawRecord struct {
	ID int `json:"id"`
	WithdrawNo string `json:"withdraw_no"`
	CardNumber string `json:"card_number"`
	WithdrawAmount int `json:"withdraw_amount"`
	WithdrawTime string `json:"withdraw_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `WithdrawRecordMonthStatusFailed`

WithdrawRecordMonthStatusFailed represents failed withdrawal attempts aggregated by month.
Used for monthly failure analysis and system health monitoring.

```go
type WithdrawRecordMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `WithdrawRecordMonthStatusSuccess`

WithdrawRecordMonthStatusSuccess represents successful withdrawals aggregated by month.
Used for monthly success rate reporting and cash flow analysis.

```go
type WithdrawRecordMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `WithdrawRecordYearStatusFailed`

WithdrawRecordYearStatusFailed represents failed withdrawal attempts aggregated by year.
Used for annual failure analysis and system reliability reporting.

```go
type WithdrawRecordYearStatusFailed struct {
	Year string `json:"year"`
	TotalFailed int `json:"total_failed"`
	TotalAmount int `json:"total_amount"`
}
```

### `WithdrawRecordYearStatusSuccess`

WithdrawRecordYearStatusSuccess represents successful withdrawals aggregated by year.
Used for annual success rate reporting and financial analysis.

```go
type WithdrawRecordYearStatusSuccess struct {
	Year string `json:"year"`
	TotalSuccess int `json:"total_success"`
	TotalAmount int `json:"total_amount"`
}
```

### `WithdrawYearlyAmount`

WithdrawYearlyAmount represents total withdrawal amounts by year.
Used for annual financial reporting and cash flow analysis.

```go
type WithdrawYearlyAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```