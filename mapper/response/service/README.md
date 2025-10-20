# 📦 Package `responseservice`

**Source Path:** `shared/mapper/response/service`

## 🧩 Types

### `CardResponseMapper`

CardResponseMapper defines methods for converting internal Card records
into HTTP/REST API response formats.

```go
type CardResponseMapper interface {
	ToCardResponse func(card *record.CardRecord) (*response.CardResponse)
	ToCardsResponse func(cards []*record.CardRecord) ([]*response.CardResponse)
	ToCardResponseDeleteAt func(card *record.CardRecord) (*response.CardResponseDeleteAt)
	ToCardsResponseDeleteAt func(cards []*record.CardRecord) ([]*response.CardResponseDeleteAt)
	ToGetMonthlyBalance func(card *record.CardMonthBalance) (*response.CardResponseMonthBalance)
	ToGetMonthlyBalances func(cards []*record.CardMonthBalance) ([]*response.CardResponseMonthBalance)
	ToGetYearlyBalance func(card *record.CardYearlyBalance) (*response.CardResponseYearlyBalance)
	ToGetYearlyBalances func(cards []*record.CardYearlyBalance) ([]*response.CardResponseYearlyBalance)
	ToGetMonthlyAmount func(card *record.CardMonthAmount) (*response.CardResponseMonthAmount)
	ToGetMonthlyAmounts func(cards []*record.CardMonthAmount) ([]*response.CardResponseMonthAmount)
	ToGetYearlyAmount func(card *record.CardYearAmount) (*response.CardResponseYearAmount)
	ToGetYearlyAmounts func(cards []*record.CardYearAmount) ([]*response.CardResponseYearAmount)
}
```

### `MerchantDocumentResponseMapper`

MerchantDocumentResponseMapper defines methods for converting merchant document
records into API-compatible response types.

```go
type MerchantDocumentResponseMapper interface {
	ToMerchantDocumentResponse func(doc *record.MerchantDocumentRecord) (*response.MerchantDocumentResponse)
	ToMerchantDocumentsResponse func(docs []*record.MerchantDocumentRecord) ([]*response.MerchantDocumentResponse)
	ToMerchantDocumentResponseDeleteAt func(doc *record.MerchantDocumentRecord) (*response.MerchantDocumentResponseDeleteAt)
	ToMerchantDocumentsResponseDeleteAt func(docs []*record.MerchantDocumentRecord) ([]*response.MerchantDocumentResponseDeleteAt)
}
```

### `MerchantResponseMapper`

MerchantResponseMapper defines methods for converting merchant-related database
records into API-compatible response types.

```go
type MerchantResponseMapper interface {
	ToMerchantResponse func(merchant *record.MerchantRecord) (*response.MerchantResponse)
	ToMerchantsResponse func(merchants []*record.MerchantRecord) ([]*response.MerchantResponse)
	ToMerchantMonthlyTotalAmount func(ms *record.MerchantMonthlyTotalAmount) (*response.MerchantResponseMonthlyTotalAmount)
	ToMerchantMonthlyTotalAmounts func(ms []*record.MerchantMonthlyTotalAmount) ([]*response.MerchantResponseMonthlyTotalAmount)
	ToMerchantYearlyTotalAmount func(ms *record.MerchantYearlyTotalAmount) (*response.MerchantResponseYearlyTotalAmount)
	ToMerchantYearlyTotalAmounts func(ms []*record.MerchantYearlyTotalAmount) ([]*response.MerchantResponseYearlyTotalAmount)
	ToMerchantTransactionResponse func(merchant *record.MerchantTransactionsRecord) (*response.MerchantTransactionResponse)
	ToMerchantsTransactionResponse func(merchants []*record.MerchantTransactionsRecord) ([]*response.MerchantTransactionResponse)
	ToMerchantMonthlyPaymentMethod func(ms *record.MerchantMonthlyPaymentMethod) (*response.MerchantResponseMonthlyPaymentMethod)
	ToMerchantMonthlyPaymentMethods func(ms []*record.MerchantMonthlyPaymentMethod) ([]*response.MerchantResponseMonthlyPaymentMethod)
	ToMerchantYearlyPaymentMethod func(ms *record.MerchantYearlyPaymentMethod) (*response.MerchantResponseYearlyPaymentMethod)
	ToMerchantYearlyPaymentMethods func(ms []*record.MerchantYearlyPaymentMethod) ([]*response.MerchantResponseYearlyPaymentMethod)
	ToMerchantMonthlyAmount func(ms *record.MerchantMonthlyAmount) (*response.MerchantResponseMonthlyAmount)
	ToMerchantMonthlyAmounts func(ms []*record.MerchantMonthlyAmount) ([]*response.MerchantResponseMonthlyAmount)
	ToMerchantYearlyAmount func(ms *record.MerchantYearlyAmount) (*response.MerchantResponseYearlyAmount)
	ToMerchantYearlyAmounts func(ms []*record.MerchantYearlyAmount) ([]*response.MerchantResponseYearlyAmount)
	ToMerchantResponseDeleteAt func(merchant *record.MerchantRecord) (*response.MerchantResponseDeleteAt)
	ToMerchantsResponseDeleteAt func(merchants []*record.MerchantRecord) ([]*response.MerchantResponseDeleteAt)
}
```

### `RefreshTokenResponseMapper`

RefreshTokenResponseMapper defines methods to map RefreshTokenRecord
into their corresponding API response structures.

```go
type RefreshTokenResponseMapper interface {
	ToRefreshTokenResponse func(refresh *record.RefreshTokenRecord) (*response.RefreshTokenResponse)
	ToRefreshTokenResponses func(refreshs []*record.RefreshTokenRecord) ([]*response.RefreshTokenResponse)
}
```

### `ResponseServiceMapper`

```go
type ResponseServiceMapper struct {
	CardResponseMapper CardResponseMapper
	RoleResponseMapper RoleResponseMapper
	RefreshTokenResponseMapper RefreshTokenResponseMapper
	SaldoResponseMapper SaldoResponseMapper
	TransactionResponseMapper TransactionResponseMapper
	TransferResponseMapper TransferResponseMapper
	TopupResponseMapper TopupResponseMapper
	WithdrawResponseMapper WithdrawResponseMapper
	UserResponseMapper UserResponseMapper
	MerchantResponseMapper MerchantResponseMapper
	MerchantDocumentResponseMapper MerchantDocumentResponseMapper
}
```

### `RoleResponseMapper`

RoleResponseMapper defines methods to map RoleRecord domain models
into structured API response representations.

```go
type RoleResponseMapper interface {
	ToRoleResponse func(role *record.RoleRecord) (*response.RoleResponse)
	ToRolesResponse func(roles []*record.RoleRecord) ([]*response.RoleResponse)
	ToRoleResponseDeleteAt func(role *record.RoleRecord) (*response.RoleResponseDeleteAt)
	ToRolesResponseDeleteAt func(roles []*record.RoleRecord) ([]*response.RoleResponseDeleteAt)
}
```

### `SaldoResponseMapper`

SaldoResponseMapper defines methods to map Saldo-related records
into structured API response formats.

```go
type SaldoResponseMapper interface {
	ToSaldoResponse func(saldo *record.SaldoRecord) (*response.SaldoResponse)
	ToSaldoResponses func(saldos []*record.SaldoRecord) ([]*response.SaldoResponse)
	ToSaldoMonthTotalBalanceResponse func(ss *record.SaldoMonthTotalBalance) (*response.SaldoMonthTotalBalanceResponse)
	ToSaldoMonthTotalBalanceResponses func(ss []*record.SaldoMonthTotalBalance) ([]*response.SaldoMonthTotalBalanceResponse)
	ToSaldoYearTotalBalanceResponse func(ss *record.SaldoYearTotalBalance) (*response.SaldoYearTotalBalanceResponse)
	ToSaldoYearTotalBalanceResponses func(ss []*record.SaldoYearTotalBalance) ([]*response.SaldoYearTotalBalanceResponse)
	ToSaldoMonthBalanceResponse func(ss *record.SaldoMonthSaldoBalance) (*response.SaldoMonthBalanceResponse)
	ToSaldoMonthBalanceResponses func(ss []*record.SaldoMonthSaldoBalance) ([]*response.SaldoMonthBalanceResponse)
	ToSaldoYearBalanceResponse func(ss *record.SaldoYearSaldoBalance) (*response.SaldoYearBalanceResponse)
	ToSaldoYearBalanceResponses func(ss []*record.SaldoYearSaldoBalance) ([]*response.SaldoYearBalanceResponse)
	ToSaldoResponseDeleteAt func(saldo *record.SaldoRecord) (*response.SaldoResponseDeleteAt)
	ToSaldoResponsesDeleteAt func(saldos []*record.SaldoRecord) ([]*response.SaldoResponseDeleteAt)
}
```

### `TopupResponseMapper`

TopupResponseMapper defines methods for mapping top-up records into
structured API response objects used in the application.

```go
type TopupResponseMapper interface {
	ToTopupResponse func(topup *record.TopupRecord) (*response.TopupResponse)
	ToTopupResponses func(topups []*record.TopupRecord) ([]*response.TopupResponse)
	ToTopupResponseMonthStatusSuccess func(s *record.TopupRecordMonthStatusSuccess) (*response.TopupResponseMonthStatusSuccess)
	ToTopupResponsesMonthStatusSuccess func(topups []*record.TopupRecordMonthStatusSuccess) ([]*response.TopupResponseMonthStatusSuccess)
	ToTopupResponseYearStatusSuccess func(s *record.TopupRecordYearStatusSuccess) (*response.TopupResponseYearStatusSuccess)
	ToTopupResponsesYearStatusSuccess func(topups []*record.TopupRecordYearStatusSuccess) ([]*response.TopupResponseYearStatusSuccess)
	ToTopupResponseMonthStatusFailed func(s *record.TopupRecordMonthStatusFailed) (*response.TopupResponseMonthStatusFailed)
	ToTopupResponsesMonthStatusFailed func(topups []*record.TopupRecordMonthStatusFailed) ([]*response.TopupResponseMonthStatusFailed)
	ToTopupResponseYearStatusFailed func(s *record.TopupRecordYearStatusFailed) (*response.TopupResponseYearStatusFailed)
	ToTopupResponsesYearStatusFailed func(topups []*record.TopupRecordYearStatusFailed) ([]*response.TopupResponseYearStatusFailed)
	ToTopupMonthlyMethodResponse func(s *record.TopupMonthMethod) (*response.TopupMonthMethodResponse)
	ToTopupMonthlyMethodResponses func(s []*record.TopupMonthMethod) ([]*response.TopupMonthMethodResponse)
	ToTopupYearlyMethodResponse func(s *record.TopupYearlyMethod) (*response.TopupYearlyMethodResponse)
	ToTopupYearlyMethodResponses func(s []*record.TopupYearlyMethod) ([]*response.TopupYearlyMethodResponse)
	ToTopupMonthlyAmountResponse func(s *record.TopupMonthAmount) (*response.TopupMonthAmountResponse)
	ToTopupMonthlyAmountResponses func(s []*record.TopupMonthAmount) ([]*response.TopupMonthAmountResponse)
	ToTopupYearlyAmountResponse func(s *record.TopupYearlyAmount) (*response.TopupYearlyAmountResponse)
	ToTopupYearlyAmountResponses func(s []*record.TopupYearlyAmount) ([]*response.TopupYearlyAmountResponse)
	ToTopupResponseDeleteAt func(topup *record.TopupRecord) (*response.TopupResponseDeleteAt)
	ToTopupResponsesDeleteAt func(topups []*record.TopupRecord) ([]*response.TopupResponseDeleteAt)
}
```

### `TransactionResponseMapper`

TransactionResponseMapper defines a set of methods for converting transaction records
from the data layer into structured API response objects for use in handlers.

```go
type TransactionResponseMapper interface {
	ToTransactionResponse func(transaction *record.TransactionRecord) (*response.TransactionResponse)
	ToTransactionsResponse func(transactions []*record.TransactionRecord) ([]*response.TransactionResponse)
	ToTransactionResponseMonthStatusSuccess func(s *record.TransactionRecordMonthStatusSuccess) (*response.TransactionResponseMonthStatusSuccess)
	ToTransactionResponsesMonthStatusSuccess func(Transactions []*record.TransactionRecordMonthStatusSuccess) ([]*response.TransactionResponseMonthStatusSuccess)
	ToTransactionResponseYearStatusSuccess func(s *record.TransactionRecordYearStatusSuccess) (*response.TransactionResponseYearStatusSuccess)
	ToTransactionResponsesYearStatusSuccess func(Transactions []*record.TransactionRecordYearStatusSuccess) ([]*response.TransactionResponseYearStatusSuccess)
	ToTransactionResponseMonthStatusFailed func(s *record.TransactionRecordMonthStatusFailed) (*response.TransactionResponseMonthStatusFailed)
	ToTransactionResponsesMonthStatusFailed func(Transactions []*record.TransactionRecordMonthStatusFailed) ([]*response.TransactionResponseMonthStatusFailed)
	ToTransactionResponseYearStatusFailed func(s *record.TransactionRecordYearStatusFailed) (*response.TransactionResponseYearStatusFailed)
	ToTransactionResponsesYearStatusFailed func(Transactions []*record.TransactionRecordYearStatusFailed) ([]*response.TransactionResponseYearStatusFailed)
	ToTransactionMonthlyMethodResponse func(s *record.TransactionMonthMethod) (*response.TransactionMonthMethodResponse)
	ToTransactionMonthlyMethodResponses func(s []*record.TransactionMonthMethod) ([]*response.TransactionMonthMethodResponse)
	ToTransactionYearlyMethodResponse func(s *record.TransactionYearMethod) (*response.TransactionYearMethodResponse)
	ToTransactionYearlyMethodResponses func(s []*record.TransactionYearMethod) ([]*response.TransactionYearMethodResponse)
	ToTransactionMonthlyAmountResponse func(s *record.TransactionMonthAmount) (*response.TransactionMonthAmountResponse)
	ToTransactionMonthlyAmountResponses func(s []*record.TransactionMonthAmount) ([]*response.TransactionMonthAmountResponse)
	ToTransactionYearlyAmountResponse func(s *record.TransactionYearlyAmount) (*response.TransactionYearlyAmountResponse)
	ToTransactionYearlyAmountResponses func(s []*record.TransactionYearlyAmount) ([]*response.TransactionYearlyAmountResponse)
	ToTransactionResponseDeleteAt func(transaction *record.TransactionRecord) (*response.TransactionResponseDeleteAt)
	ToTransactionsResponseDeleteAt func(transactions []*record.TransactionRecord) ([]*response.TransactionResponseDeleteAt)
}
```

### `TransferResponseMapper`

TransferResponseMapper defines a set of methods for converting transfer-related records
from the database layer into structured API responses used by the application layer.

```go
type TransferResponseMapper interface {
	ToTransferResponse func(transfer *record.TransferRecord) (*response.TransferResponse)
	ToTransfersResponse func(transfers []*record.TransferRecord) ([]*response.TransferResponse)
	ToTransferResponseMonthStatusSuccess func(s *record.TransferRecordMonthStatusSuccess) (*response.TransferResponseMonthStatusSuccess)
	ToTransferResponsesMonthStatusSuccess func(Transfers []*record.TransferRecordMonthStatusSuccess) ([]*response.TransferResponseMonthStatusSuccess)
	ToTransferResponseYearStatusSuccess func(s *record.TransferRecordYearStatusSuccess) (*response.TransferResponseYearStatusSuccess)
	ToTransferResponsesYearStatusSuccess func(Transfers []*record.TransferRecordYearStatusSuccess) ([]*response.TransferResponseYearStatusSuccess)
	ToTransferResponseMonthStatusFailed func(s *record.TransferRecordMonthStatusFailed) (*response.TransferResponseMonthStatusFailed)
	ToTransferResponsesMonthStatusFailed func(Transfers []*record.TransferRecordMonthStatusFailed) ([]*response.TransferResponseMonthStatusFailed)
	ToTransferResponseYearStatusFailed func(s *record.TransferRecordYearStatusFailed) (*response.TransferResponseYearStatusFailed)
	ToTransferResponsesYearStatusFailed func(Transfers []*record.TransferRecordYearStatusFailed) ([]*response.TransferResponseYearStatusFailed)
	ToTransferResponseMonthAmount func(s *record.TransferMonthAmount) (*response.TransferMonthAmountResponse)
	ToTransferResponsesMonthAmount func(s []*record.TransferMonthAmount) ([]*response.TransferMonthAmountResponse)
	ToTransferResponseYearAmount func(s *record.TransferYearAmount) (*response.TransferYearAmountResponse)
	ToTransferResponsesYearAmount func(s []*record.TransferYearAmount) ([]*response.TransferYearAmountResponse)
	ToTransferResponseDeleteAt func(transfer *record.TransferRecord) (*response.TransferResponseDeleteAt)
	ToTransfersResponseDeleteAt func(transfers []*record.TransferRecord) ([]*response.TransferResponseDeleteAt)
}
```

### `UserResponseMapper`

UserResponseMapper defines methods for converting internal User records
into HTTP/REST API response formats.

```go
type UserResponseMapper interface {
	ToUserResponse func(user *record.UserRecord) (*response.UserResponse)
	ToUsersResponse func(users []*record.UserRecord) ([]*response.UserResponse)
	ToUserResponseDeleteAt func(user *record.UserRecord) (*response.UserResponseDeleteAt)
	ToUsersResponseDeleteAt func(users []*record.UserRecord) ([]*response.UserResponseDeleteAt)
}
```

### `WithdrawResponseMapper`

WithdrawResponseMapper defines methods to map withdrawal-related database records
to structured API response objects used in handlers or services.

```go
type WithdrawResponseMapper interface {
	ToWithdrawResponse func(withdraw *record.WithdrawRecord) (*response.WithdrawResponse)
	ToWithdrawsResponse func(withdraws []*record.WithdrawRecord) ([]*response.WithdrawResponse)
	ToWithdrawResponseMonthStatusSuccess func(s *record.WithdrawRecordMonthStatusSuccess) (*response.WithdrawResponseMonthStatusSuccess)
	ToWithdrawResponsesMonthStatusSuccess func(Withdraws []*record.WithdrawRecordMonthStatusSuccess) ([]*response.WithdrawResponseMonthStatusSuccess)
	ToWithdrawResponseYearStatusSuccess func(s *record.WithdrawRecordYearStatusSuccess) (*response.WithdrawResponseYearStatusSuccess)
	ToWithdrawResponsesYearStatusSuccess func(Withdraws []*record.WithdrawRecordYearStatusSuccess) ([]*response.WithdrawResponseYearStatusSuccess)
	ToWithdrawResponseMonthStatusFailed func(s *record.WithdrawRecordMonthStatusFailed) (*response.WithdrawResponseMonthStatusFailed)
	ToWithdrawResponsesMonthStatusFailed func(Withdraws []*record.WithdrawRecordMonthStatusFailed) ([]*response.WithdrawResponseMonthStatusFailed)
	ToWithdrawResponseYearStatusFailed func(s *record.WithdrawRecordYearStatusFailed) (*response.WithdrawResponseYearStatusFailed)
	ToWithdrawResponsesYearStatusFailed func(Withdraws []*record.WithdrawRecordYearStatusFailed) ([]*response.WithdrawResponseYearStatusFailed)
	ToWithdrawAmountMonthlyResponse func(s *record.WithdrawMonthlyAmount) (*response.WithdrawMonthlyAmountResponse)
	ToWithdrawsAmountMonthlyResponses func(s []*record.WithdrawMonthlyAmount) ([]*response.WithdrawMonthlyAmountResponse)
	ToWithdrawAmountYearlyResponse func(s *record.WithdrawYearlyAmount) (*response.WithdrawYearlyAmountResponse)
	ToWithdrawsAmountYearlyResponses func(s []*record.WithdrawYearlyAmount) ([]*response.WithdrawYearlyAmountResponse)
	ToWithdrawResponseDeleteAt func(withdraw *record.WithdrawRecord) (*response.WithdrawResponseDeleteAt)
	ToWithdrawsResponseDeleteAt func(withdraws []*record.WithdrawRecord) ([]*response.WithdrawResponseDeleteAt)
}
```

### `cardResponseMapper`

```go
type cardResponseMapper struct {
}
```

#### Methods

##### `ToCardResponse`

ToCardResponse converts a single card record into a CardResponse.

Args:
  - card: A pointer to a CardRecord representing the card record.

Returns:
  - A pointer to a CardResponse containing the mapped data, including ID, UserID, CardNumber,
    CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.

```go
func (s *cardResponseMapper) ToCardResponse(card *record.CardRecord) *response.CardResponse
```

##### `ToCardResponseDeleteAt`

ToCardResponseDeleteAt converts a CardRecord into a CardResponseDeleteAt.

Args:
  - card: A pointer to a CardRecord representing the card record.

Returns:
  - A pointer to a CardResponseDeleteAt containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardResponseMapper) ToCardResponseDeleteAt(card *record.CardRecord) *response.CardResponseDeleteAt
```

##### `ToCardsResponse`

ToCardsResponse converts a list of card records into a list of CardResponse.

Args:
  - cards: A pointer to a slice of CardRecord representing the card records.

Returns:
  - A pointer to a slice of CardResponse containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.

```go
func (s *cardResponseMapper) ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse
```

##### `ToCardsResponseDeleteAt`

```go
func (s *cardResponseMapper) ToCardsResponseDeleteAt(cards []*record.CardRecord) []*response.CardResponseDeleteAt
```

##### `ToGetMonthlyAmount`

```go
func (s *cardResponseMapper) ToGetMonthlyAmount(card *record.CardMonthAmount) *response.CardResponseMonthAmount
```

##### `ToGetMonthlyAmounts`

```go
func (s *cardResponseMapper) ToGetMonthlyAmounts(cards []*record.CardMonthAmount) []*response.CardResponseMonthAmount
```

##### `ToGetMonthlyBalance`

```go
func (s *cardResponseMapper) ToGetMonthlyBalance(card *record.CardMonthBalance) *response.CardResponseMonthBalance
```

##### `ToGetMonthlyBalances`

```go
func (s *cardResponseMapper) ToGetMonthlyBalances(cards []*record.CardMonthBalance) []*response.CardResponseMonthBalance
```

##### `ToGetYearlyAmount`

```go
func (s *cardResponseMapper) ToGetYearlyAmount(card *record.CardYearAmount) *response.CardResponseYearAmount
```

##### `ToGetYearlyAmounts`

```go
func (s *cardResponseMapper) ToGetYearlyAmounts(cards []*record.CardYearAmount) []*response.CardResponseYearAmount
```

##### `ToGetYearlyBalance`

```go
func (s *cardResponseMapper) ToGetYearlyBalance(card *record.CardYearlyBalance) *response.CardResponseYearlyBalance
```

##### `ToGetYearlyBalances`

```go
func (s *cardResponseMapper) ToGetYearlyBalances(cards []*record.CardYearlyBalance) []*response.CardResponseYearlyBalance
```

### `merchantDocumentResponseMapper`

merchantDocumentResponseMapper provides methods to map MerchantDocumentRecord domain models to MerchantDocumentResponse API-compatible response types.

```go
type merchantDocumentResponseMapper struct {
}
```

#### Methods

##### `ToMerchantDocumentResponse`

ToMerchantDocumentResponse maps a single MerchantDocumentRecord to a MerchantDocumentResponse API-compatible response.
Args:
  - doc: A pointer to a MerchantDocumentRecord containing the data to be mapped.

Returns:
  - A pointer to a MerchantDocumentResponse containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *merchantDocumentResponseMapper) ToMerchantDocumentResponse(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponse
```

##### `ToMerchantDocumentResponseDeleteAt`

ToMerchantDocumentResponseDeleteAt maps a soft-deleted MerchantDocumentRecord to its corresponding response.
Args:
  - doc: A pointer to a MerchantDocumentRecord containing the data to be mapped.

Returns:
  - A pointer to a MerchantDocumentResponseDeleteAt containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *merchantDocumentResponseMapper) ToMerchantDocumentResponseDeleteAt(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponseDeleteAt
```

##### `ToMerchantDocumentsResponse`

ToMerchantDocumentsResponse maps multiple MerchantDocumentRecords to a slice of MerchantDocumentResponse API-compatible responses.
It constructs a slice of MerchantDocumentResponse by mapping each MerchantDocumentRecord to a MerchantDocumentResponse using
the ToMerchantDocumentResponse method.

Args:
  - docs: A slice of pointers to MerchantDocumentRecord containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantDocumentResponse containing the mapped data.

```go
func (s *merchantDocumentResponseMapper) ToMerchantDocumentsResponse(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponse
```

##### `ToMerchantDocumentsResponseDeleteAt`

ToMerchantDocumentsResponseDeleteAt maps multiple soft-deleted MerchantDocumentRecords to a slice of MerchantDocumentResponseDeleteAt API-compatible responses.
It constructs a slice of MerchantDocumentResponseDeleteAt by mapping each MerchantDocumentRecord to a MerchantDocumentResponseDeleteAt using
the ToMerchantDocumentResponseDeleteAt method.

Args:
  - docs: A slice of pointers to MerchantDocumentRecord containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantDocumentResponseDeleteAt containing the mapped data.

```go
func (s *merchantDocumentResponseMapper) ToMerchantDocumentsResponseDeleteAt(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponseDeleteAt
```

### `merchantResponseMapper`

merchantResponseMapper provides methods to map MerchantRecord domain models to MerchantResponse API-compatible response types.

```go
type merchantResponseMapper struct {
}
```

#### Methods

##### `ToMerchantMonthlyAmount`

ToMerchantMonthlyAmount maps a single monthly amount record to a
MerchantResponseMonthlyAmount API-compatible response.

Args:
  - ms: A pointer to a MerchantMonthlyAmount containing the data to be mapped.

Returns:
  - A pointer to a MerchantResponseMonthlyAmount containing the mapped data, including
    Month and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantMonthlyAmount(ms *record.MerchantMonthlyAmount) *response.MerchantResponseMonthlyAmount
```

##### `ToMerchantMonthlyAmounts`

ToMerchantMonthlyAmounts maps multiple monthly amount records into a slice of MerchantResponseMonthlyAmount.

Args:
  - ms: A slice of pointers to MerchantMonthlyAmount records containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantResponseMonthlyAmount containing the mapped data, including
    Month and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantMonthlyAmounts(ms []*record.MerchantMonthlyAmount) []*response.MerchantResponseMonthlyAmount
```

##### `ToMerchantMonthlyPaymentMethod`

```go
func (s *merchantResponseMapper) ToMerchantMonthlyPaymentMethod(ms *record.MerchantMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod
```

##### `ToMerchantMonthlyPaymentMethods`

ToMerchantMonthlyPaymentMethods maps multiple monthly payment method records into a slice of MerchantResponseMonthlyPaymentMethod.

Args:
  - ms: A slice of pointers to MerchantMonthlyPaymentMethod records containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantResponseMonthlyPaymentMethod containing the mapped data, including
    Month, PaymentMethod, and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantMonthlyPaymentMethods(ms []*record.MerchantMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod
```

##### `ToMerchantMonthlyTotalAmount`

ToMerchantMonthlyTotalAmount maps a single MerchantMonthlyTotalAmount record to a
MerchantResponseMonthlyTotalAmount API-compatible response.

Args:
  - ms: A pointer to a MerchantMonthlyTotalAmount containing the data to be mapped.

Returns:
  - A pointer to a MerchantResponseMonthlyTotalAmount containing the mapped data, including
    Month, Year, and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantMonthlyTotalAmount(ms *record.MerchantMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount
```

##### `ToMerchantMonthlyTotalAmounts`

ToMerchantMonthlyTotalAmounts maps multiple monthly total amount records into a slice of
MerchantResponseMonthlyTotalAmount.

Args:
  - ms: A slice of pointers to MerchantMonthlyTotalAmount records containing the data
    to be mapped.

Returns:
  - A slice of pointers to MerchantResponseMonthlyTotalAmount containing the mapped
    data, including Month, Year, and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantMonthlyTotalAmounts(ms []*record.MerchantMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount
```

##### `ToMerchantResponse`

ToMerchantResponse maps a single MerchantRecord to a MerchantResponse API-compatible response.
Args:
  - merchant: A pointer to a MerchantRecord containing the data to be mapped.

Returns:
  - A pointer to a MerchantResponse containing the mapped data, including
    ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.

```go
func (s *merchantResponseMapper) ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse
```

##### `ToMerchantResponseDeleteAt`

ToMerchantResponseDeleteAt maps a MerchantRecord to a MerchantResponseDeleteAt,
which includes additional deletion data.
This function is useful for handling soft-deleted merchants where the deletion
timestamp is relevant.

Args:
  - merchant: A pointer to a MerchantRecord containing the data to be mapped.

Returns:
  - A pointer to a MerchantResponseDeleteAt containing the mapped data,
    including ID, Name, UserID, Status, ApiKey, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *merchantResponseMapper) ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt
```

##### `ToMerchantTransactionResponse`

ToMerchantTransactionResponse maps a single MerchantTransactionsRecord to a MerchantTransactionResponse API-compatible response.

Args:
  - merchant: A pointer to a MerchantTransactionsRecord containing the data to be mapped.

Returns:
  - A pointer to a MerchantTransactionResponse containing the mapped data, including
    ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt.

```go
func (m *merchantResponseMapper) ToMerchantTransactionResponse(merchant *record.MerchantTransactionsRecord) *response.MerchantTransactionResponse
```

##### `ToMerchantYearlyAmount`

ToMerchantYearlyAmount maps a single yearly amount record to a
MerchantResponseYearlyAmount API-compatible response.

Args:
  - ms: A pointer to a MerchantYearlyAmount containing the data to be mapped.

Returns:
  - A pointer to a MerchantResponseYearlyAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantYearlyAmount(ms *record.MerchantYearlyAmount) *response.MerchantResponseYearlyAmount
```

##### `ToMerchantYearlyAmounts`

ToMerchantYearlyAmounts maps multiple yearly amount records into a slice of
MerchantResponseYearlyAmount.

Args:
  - ms: A slice of pointers to MerchantYearlyAmount records containing the data
    to be mapped.

Returns:
  - A slice of pointers to MerchantResponseYearlyAmount containing the mapped
    data, including Year and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantYearlyAmounts(ms []*record.MerchantYearlyAmount) []*response.MerchantResponseYearlyAmount
```

##### `ToMerchantYearlyPaymentMethod`

ToMerchantYearlyPaymentMethod maps a single yearly payment method record to a
MerchantResponseYearlyPaymentMethod API-compatible response.

Args:
  - ms: A pointer to a MerchantYearlyPaymentMethod containing the data to be mapped.

Returns:
  - A pointer to a MerchantResponseYearlyPaymentMethod containing the mapped data, including
    Year, PaymentMethod, and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantYearlyPaymentMethod(ms *record.MerchantYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod
```

##### `ToMerchantYearlyPaymentMethods`

ToMerchantYearlyPaymentMethods maps multiple yearly payment method records into a slice of MerchantResponseYearlyPaymentMethod.

Args:
  - ms: A slice of pointers to MerchantYearlyPaymentMethod records containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantResponseYearlyPaymentMethod containing the mapped data, including
    Year, PaymentMethod, and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantYearlyPaymentMethods(ms []*record.MerchantYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod
```

##### `ToMerchantYearlyTotalAmount`

ToMerchantYearlyTotalAmount maps a single yearly total amount record to a
MerchantResponseYearlyTotalAmount API-compatible response.

Args:
  - ms: A pointer to a MerchantYearlyTotalAmount containing the data to be mapped.

Returns:
  - A pointer to a MerchantResponseYearlyTotalAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (s *merchantResponseMapper) ToMerchantYearlyTotalAmount(ms *record.MerchantYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount
```

##### `ToMerchantYearlyTotalAmounts`

```go
func (s *merchantResponseMapper) ToMerchantYearlyTotalAmounts(ms []*record.MerchantYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount
```

##### `ToMerchantsResponse`

ToMerchantsResponse maps multiple MerchantRecords to a slice of MerchantResponse API-compatible responses.
Args:
  - merchants: A slice of pointers to MerchantRecord containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantResponse containing the mapped data, including
    ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.

```go
func (s *merchantResponseMapper) ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse
```

##### `ToMerchantsResponseDeleteAt`

ToMerchantsResponseDeleteAt maps multiple soft-deleted MerchantRecords to their corresponding responses.
Args:
  - merchants: A slice of pointers to MerchantRecord containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantResponseDeleteAt containing the mapped data, including
    ID, Name, UserID, Status, ApiKey, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *merchantResponseMapper) ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt
```

##### `ToMerchantsTransactionResponse`

ToMerchantsTransactionResponse maps multiple MerchantTransactionsRecords to a slice of MerchantTransactionResponse API-compatible responses.
Args:
  - merchants: A slice of pointers to MerchantTransactionsRecord containing the data to be mapped.

Returns:
  - A slice of pointers to MerchantTransactionResponse containing the mapped data, including
    ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt.

```go
func (m *merchantResponseMapper) ToMerchantsTransactionResponse(merchants []*record.MerchantTransactionsRecord) []*response.MerchantTransactionResponse
```

### `refreshTokenResponseMapper`

refreshTokenResponseMapper provides methods to map RefreshTokenRecord domain models to RefreshTokenResponse API-compatible response types.

```go
type refreshTokenResponseMapper struct {
}
```

#### Methods

##### `ToRefreshTokenResponse`

ToRefreshTokenResponse maps a RefreshTokenRecord domain model to a RefreshTokenResponse API-compatible response type.
Args:
  - refresh: A pointer to a RefreshTokenRecord representing the domain model.

Returns:
  - A pointer to a RefreshTokenResponse representing the API-compatible response type.

```go
func (r *refreshTokenResponseMapper) ToRefreshTokenResponse(refresh *record.RefreshTokenRecord) *response.RefreshTokenResponse
```

##### `ToRefreshTokenResponses`

ToRefreshTokenResponses maps a slice of RefreshTokenRecord domain models to a slice of RefreshTokenResponse API-compatible response types.
Args:
  - refreshs: A slice of pointers to RefreshTokenRecord representing the domain models.

Returns:
  - A slice of pointers to RefreshTokenResponse representing the API-compatible response types.

```go
func (r *refreshTokenResponseMapper) ToRefreshTokenResponses(refreshs []*record.RefreshTokenRecord) []*response.RefreshTokenResponse
```

### `roleResponseMapper`

roleResponseMapper provides methods to map RoleRecord domain models to RoleResponse API-compatible response types.

```go
type roleResponseMapper struct {
}
```

#### Methods

##### `ToRoleResponse`

ToRoleResponse converts a single RoleRecord into RoleResponse.

Args:
  - role: A pointer to the RoleRecord to be mapped.

Returns:
  - A pointer to a RoleResponse containing the mapped data, with fields ID, Name,
    CreatedAt, and UpdatedAt extracted from the RoleRecord.

```go
func (s *roleResponseMapper) ToRoleResponse(role *record.RoleRecord) *response.RoleResponse
```

##### `ToRoleResponseDeleteAt`

ToRoleResponseDeleteAt converts a RoleRecord with deletion information
into a RoleResponseDeleteAt structure.

Args:
  - role: A pointer to the RoleRecord to be mapped.

Returns:
  - A pointer to a RoleResponseDeleteAt containing the mapped data, with fields ID, Name,
    CreatedAt, UpdatedAt, and DeletedAt extracted from the RoleRecord.

```go
func (s *roleResponseMapper) ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt
```

##### `ToRolesResponse`

ToRolesResponse maps a slice of RoleRecord domain models to a slice of RoleResponse API-compatible response types.

Args:
  - roles: A slice of pointers to RoleRecord domain models to be mapped.

Returns:
  - A slice of pointers to RoleResponse, each containing the mapped data, with fields ID, Name,
    CreatedAt, and UpdatedAt extracted from the RoleRecord.

```go
func (s *roleResponseMapper) ToRolesResponse(roles []*record.RoleRecord) []*response.RoleResponse
```

##### `ToRolesResponseDeleteAt`

ToRolesResponseDeleteAt maps a slice of RoleRecord domain models with deletion information
to a slice of RoleResponseDeleteAt API-compatible response types.

Args:
  - roles: A slice of pointers to RoleRecord domain models to be mapped,
    which must have deletion information (DeletedAt != nil).

Returns:
  - A slice of pointers to RoleResponseDeleteAt, each containing the mapped data, with fields ID, Name,
    CreatedAt, UpdatedAt, and DeletedAt extracted from the RoleRecord.

```go
func (s *roleResponseMapper) ToRolesResponseDeleteAt(roles []*record.RoleRecord) []*response.RoleResponseDeleteAt
```

### `saldoResponseMapper`

saldoResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.

```go
type saldoResponseMapper struct {
}
```

#### Methods

##### `ToSaldoMonthBalanceResponse`

ToSaldoMonthBalanceResponse maps a SaldoMonthSaldoBalance record to a SaldoMonthBalanceResponse.

Args:
  - ss: A pointer to a SaldoMonthSaldoBalance containing the data to be mapped.

Returns:
  - A pointer to a SaldoMonthBalanceResponse containing the mapped data, including Month, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoMonthBalanceResponse(ss *record.SaldoMonthSaldoBalance) *response.SaldoMonthBalanceResponse
```

##### `ToSaldoMonthBalanceResponses`

ToSaldoMonthBalanceResponses maps a list of SaldoMonthSaldoBalance records to a list of SaldoMonthBalanceResponse.

Args:

	ss: A slice of pointers to SaldoMonthSaldoBalance containing the data to be mapped.

Returns:

	A slice of pointers to SaldoMonthBalanceResponse containing the mapped data, including Month, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoMonthBalanceResponses(ss []*record.SaldoMonthSaldoBalance) []*response.SaldoMonthBalanceResponse
```

##### `ToSaldoMonthTotalBalanceResponse`

ToSaldoMonthTotalBalanceResponse maps a SaldoMonthTotalBalance record to a SaldoMonthTotalBalanceResponse.

Args:
  - ss: A pointer to a SaldoMonthTotalBalance containing the data to be mapped.

Returns:
  - A pointer to a SaldoMonthTotalBalanceResponse containing the mapped data, including Month, Year, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoMonthTotalBalanceResponse(ss *record.SaldoMonthTotalBalance) *response.SaldoMonthTotalBalanceResponse
```

##### `ToSaldoMonthTotalBalanceResponses`

ToSaldoMonthTotalBalanceResponses maps a list of SaldoMonthTotalBalance records to a list of SaldoMonthTotalBalanceResponse.

Args:

	ss: A slice of pointers to SaldoMonthTotalBalance containing the data to be mapped.

Returns:

	A slice of pointers to SaldoMonthTotalBalanceResponse containing the mapped data, including Month, Year, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoMonthTotalBalanceResponses(ss []*record.SaldoMonthTotalBalance) []*response.SaldoMonthTotalBalanceResponse
```

##### `ToSaldoResponse`

ToSaldoResponse maps a single SaldoRecord to a SaldoResponse.

Args:
  - saldo: A pointer to a SaldoRecord containing the data to be mapped.

Returns:
  - A pointer to a SaldoResponse containing the mapped data.

```go
func (s *saldoResponseMapper) ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse
```

##### `ToSaldoResponseDeleteAt`

ToSaldoResponseDeleteAt maps a single SaldoRecord to a SaldoResponseDeleteAt.

Args:
  - saldo: A pointer to a SaldoRecord containing the data to be mapped.

Returns:
  - A pointer to a SaldoResponseDeleteAt containing the mapped data.

```go
func (s *saldoResponseMapper) ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt
```

##### `ToSaldoResponses`

ToSaldoResponses maps a list of SaldoRecord into a list of SaldoResponse.

Args:

	saldos: A slice of pointers to SaldoRecord containing the data to be mapped.

Returns:

	A slice of pointers to SaldoResponse containing the mapped data.

```go
func (s *saldoResponseMapper) ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse
```

##### `ToSaldoResponsesDeleteAt`

ToSaldoResponsesDeleteAt maps a list of SaldoRecord into a list of SaldoResponseDeleteAt.

Args:

	saldos: A slice of pointers to SaldoRecord containing the data to be mapped.

Returns:

	A slice of pointers to SaldoResponseDeleteAt containing the mapped data.

```go
func (s *saldoResponseMapper) ToSaldoResponsesDeleteAt(saldos []*record.SaldoRecord) []*response.SaldoResponseDeleteAt
```

##### `ToSaldoYearBalanceResponse`

ToSaldoYearBalanceResponse maps a SaldoYearSaldoBalance record to a SaldoYearBalanceResponse.

Args:
  - ss: A pointer to a SaldoYearSaldoBalance containing the data to be mapped.

Returns:
  - A pointer to a SaldoYearBalanceResponse containing the mapped data, including Year, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoYearBalanceResponse(ss *record.SaldoYearSaldoBalance) *response.SaldoYearBalanceResponse
```

##### `ToSaldoYearBalanceResponses`

ToSaldoYearBalanceResponses maps a list of SaldoYearSaldoBalance records to a list of SaldoYearBalanceResponse.

Args:

	ss: A slice of pointers to SaldoYearSaldoBalance containing the data to be mapped.

Returns:

	A slice of pointers to SaldoYearBalanceResponse containing the mapped data, including Year, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoYearBalanceResponses(ss []*record.SaldoYearSaldoBalance) []*response.SaldoYearBalanceResponse
```

##### `ToSaldoYearTotalBalanceResponse`

ToSaldoYearTotalBalanceResponse maps a SaldoYearTotalBalance record to a SaldoYearTotalBalanceResponse.

Args:
  - ss: A pointer to a SaldoYearTotalBalance containing the data to be mapped.

Returns:
  - A pointer to a SaldoYearTotalBalanceResponse containing the mapped data, including Year, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoYearTotalBalanceResponse(ss *record.SaldoYearTotalBalance) *response.SaldoYearTotalBalanceResponse
```

##### `ToSaldoYearTotalBalanceResponses`

ToSaldoYearTotalBalanceResponses maps a list of SaldoYearTotalBalance records to a list of SaldoYearTotalBalanceResponse.

Args:
  - ss: A slice of pointers to SaldoYearTotalBalance containing the data to be mapped.

Returns:
  - A slice of pointers to SaldoYearTotalBalanceResponse containing the mapped data, including Year, and TotalBalance.

```go
func (s *saldoResponseMapper) ToSaldoYearTotalBalanceResponses(ss []*record.SaldoYearTotalBalance) []*response.SaldoYearTotalBalanceResponse
```

### `topupResponseMapper`

topupResponseMapper provides methods to map TopupRecord domain models to TopupResponse API-compatible response types.

```go
type topupResponseMapper struct {
}
```

#### Methods

##### `ToTopupMonthlyAmountResponse`

ToTopupMonthlyAmountResponse converts a TopupMonthAmount domain model
into a TopupMonthAmountResponse API-compatible response type.

Args:
  - s: A pointer to a TopupMonthAmount containing the data to be mapped.

Returns:
  - A pointer to TopupMonthAmountResponse containing the mapped data, including
    Month, TotalAmount.

```go
func (t *topupResponseMapper) ToTopupMonthlyAmountResponse(s *record.TopupMonthAmount) *response.TopupMonthAmountResponse
```

##### `ToTopupMonthlyAmountResponses`

ToTopupMonthlyAmountResponses converts a slice of TopupMonthAmount domain models
into a slice of TopupMonthAmountResponse API-compatible response types.

Args:
  - s: A slice of TopupMonthAmount containing the data to be mapped.

Returns:
  - A slice of TopupMonthAmountResponse containing the mapped data, including
    Month and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupMonthlyAmountResponses(s []*record.TopupMonthAmount) []*response.TopupMonthAmountResponse
```

##### `ToTopupMonthlyMethodResponse`

ToTopupMonthlyMethodResponse converts a TopupMonthMethod domain model
into a TopupMonthMethodResponse API-compatible response type.

Args:
  - s: A pointer to a TopupMonthMethod containing the data to be mapped.

Returns:
  - A pointer to TopupMonthMethodResponse containing the mapped data, including
    Month, TopupMethod, TotalTopups, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupMonthlyMethodResponse(s *record.TopupMonthMethod) *response.TopupMonthMethodResponse
```

##### `ToTopupMonthlyMethodResponses`

ToTopupMonthlyMethodResponses maps a slice of TopupMonthMethod domain models
to a slice of TopupMonthMethodResponse API-compatible response types.

Args:
  - s: A slice of TopupMonthMethod containing the data to be mapped.

Returns:
  - A slice of TopupMonthMethodResponse containing the mapped data, including
    Month, TopupMethod, TotalTopups, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupMonthlyMethodResponses(s []*record.TopupMonthMethod) []*response.TopupMonthMethodResponse
```

##### `ToTopupResponse`

ToTopupResponse maps a single TopupRecord domain model to a TopupResponse API-compatible response type.
Args:
  - topup: A pointer to a TopupRecord containing the data to be mapped.

Returns:
  - A pointer to a TopupResponse containing the mapped data, including
    ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *topupResponseMapper) ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse
```

##### `ToTopupResponseDeleteAt`

ToTopupResponseDeleteAt maps a single TopupRecord domain model to a TopupResponseDeleteAt API-compatible response type.
It includes soft delete information by mapping the DeletedAt field.

Args:
  - topup: A pointer to a TopupRecord containing the data to be mapped.

Returns:
  - A pointer to a TopupResponseDeleteAt containing the mapped data, including
    ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *topupResponseMapper) ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt
```

##### `ToTopupResponseMonthStatusFailed`

ToTopupResponseMonthStatusFailed converts a TopupRecordMonthStatusFailed domain model
into a TopupResponseMonthStatusFailed API-compatible response type.

Args:
  - s: A pointer to a TopupRecordMonthStatusFailed containing the data to be mapped.

Returns:
  - A pointer to TopupResponseMonthStatusFailed containing the mapped data, including
    Year, Month, TotalFailed, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponseMonthStatusFailed(s *record.TopupRecordMonthStatusFailed) *response.TopupResponseMonthStatusFailed
```

##### `ToTopupResponseMonthStatusSuccess`

ToTopupResponseMonthStatusSuccess converts a TopupRecordMonthStatusSuccess domain model
into a TopupResponseMonthStatusSuccess API-compatible response type.

Args:
  - s: A pointer to a TopupRecordMonthStatusSuccess containing the data to be mapped.

Returns:
  - A pointer to TopupResponseMonthStatusSuccess containing the mapped data, including
    Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponseMonthStatusSuccess(s *record.TopupRecordMonthStatusSuccess) *response.TopupResponseMonthStatusSuccess
```

##### `ToTopupResponseYearStatusFailed`

ToTopupResponseYearStatusFailed maps a TopupRecordYearStatusFailed domain model
to a TopupResponseYearStatusFailed API-compatible response type.

Args:
  - s: A pointer to a TopupRecordYearStatusFailed containing the data to be mapped.

Returns:
  - A pointer to TopupResponseYearStatusFailed containing the mapped data, including
    Year, TotalFailed, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponseYearStatusFailed(s *record.TopupRecordYearStatusFailed) *response.TopupResponseYearStatusFailed
```

##### `ToTopupResponseYearStatusSuccess`

ToTopupResponseYearStatusSuccess converts a TopupRecordYearStatusSuccess domain model
into a TopupResponseYearStatusSuccess API-compatible response type.

Args:
  - s: A pointer to a TopupRecordYearStatusSuccess containing the data to be mapped.

Returns:
  - A pointer to TopupResponseYearStatusSuccess containing the mapped data, including
    Year, TotalSuccess, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponseYearStatusSuccess(s *record.TopupRecordYearStatusSuccess) *response.TopupResponseYearStatusSuccess
```

##### `ToTopupResponses`

ToTopupResponses maps a slice of TopupRecord to a slice of TopupResponse API-compatible response types.

Args:
  - topups: A slice of TopupRecord containing the data to be mapped.

Returns:
  - A slice of TopupResponse containing the mapped data, including
    ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *topupResponseMapper) ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse
```

##### `ToTopupResponsesDeleteAt`

ToTopupResponsesDeleteAt maps a slice of TopupRecord domain models to a slice of TopupResponseDeleteAt API-compatible response types.
It includes soft delete information by mapping the DeletedAt field.

Args:
  - topups: A slice of TopupRecord containing the data to be mapped.

Returns:
  - A slice of TopupResponseDeleteAt containing the mapped data, including
    ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *topupResponseMapper) ToTopupResponsesDeleteAt(topups []*record.TopupRecord) []*response.TopupResponseDeleteAt
```

##### `ToTopupResponsesMonthStatusFailed`

ToTopupResponsesMonthStatusFailed maps a slice of TopupRecordMonthStatusFailed domain models
to a slice of TopupResponseMonthStatusFailed API-compatible response types.

Args:
  - topups: A slice of TopupRecordMonthStatusFailed containing the data to be mapped.

Returns:
  - A slice of TopupResponseMonthStatusFailed containing the mapped data, including
    Year, Month, TotalFailed, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponsesMonthStatusFailed(topups []*record.TopupRecordMonthStatusFailed) []*response.TopupResponseMonthStatusFailed
```

##### `ToTopupResponsesMonthStatusSuccess`

ToTopupResponsesMonthStatusSuccess maps a slice of TopupRecordMonthStatusSuccess domain models
to a slice of TopupResponseMonthStatusSuccess API-compatible response types.

Args:
  - topups: A slice of TopupRecordMonthStatusSuccess containing the data to be mapped.

Returns:
  - A slice of TopupResponseMonthStatusSuccess containing the mapped data, including
    Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponsesMonthStatusSuccess(topups []*record.TopupRecordMonthStatusSuccess) []*response.TopupResponseMonthStatusSuccess
```

##### `ToTopupResponsesYearStatusFailed`

ToTopupResponsesYearStatusFailed maps a slice of TopupRecordYearStatusFailed domain models
to a slice of TopupResponseYearStatusFailed API-compatible response types.

Args:
  - topups: A slice of TopupRecordYearStatusFailed containing the data to be mapped.

Returns:
  - A slice of TopupResponseYearStatusFailed containing the mapped data, including
    Year, TotalFailed, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponsesYearStatusFailed(topups []*record.TopupRecordYearStatusFailed) []*response.TopupResponseYearStatusFailed
```

##### `ToTopupResponsesYearStatusSuccess`

ToTopupResponsesYearStatusSuccess maps a slice of TopupRecordYearStatusSuccess domain models
to a slice of TopupResponseYearStatusSuccess API-compatible response types.

Args:
  - topups: A slice of TopupRecordYearStatusSuccess containing the data to be mapped.

Returns:
  - A slice of TopupResponseYearStatusSuccess containing the mapped data, including
    Year, TotalSuccess, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupResponsesYearStatusSuccess(topups []*record.TopupRecordYearStatusSuccess) []*response.TopupResponseYearStatusSuccess
```

##### `ToTopupYearlyAmountResponse`

ToTopupYearlyAmountResponse converts a TopupYearlyAmount domain model
into a TopupYearlyAmountResponse API-compatible response type.

Args:
  - s: A pointer to a TopupYearlyAmount containing the data to be mapped.

Returns:
  - A pointer to TopupYearlyAmountResponse containing the mapped data, including
    Year, TotalAmount.

```go
func (t *topupResponseMapper) ToTopupYearlyAmountResponse(s *record.TopupYearlyAmount) *response.TopupYearlyAmountResponse
```

##### `ToTopupYearlyAmountResponses`

ToTopupYearlyAmountResponses maps a slice of TopupYearlyAmount domain models
to a slice of TopupYearlyAmountResponse API-compatible response types.

Args:
  - s: A slice of TopupYearlyAmount containing the data to be mapped.

Returns:
  - A slice of TopupYearlyAmountResponse containing the mapped data, including
    Year, TotalAmount.

```go
func (t *topupResponseMapper) ToTopupYearlyAmountResponses(s []*record.TopupYearlyAmount) []*response.TopupYearlyAmountResponse
```

##### `ToTopupYearlyMethodResponse`

ToTopupYearlyMethodResponse converts a TopupYearlyMethod domain model
into a TopupYearlyMethodResponse API-compatible response type.

Args:
  - s: A pointer to a TopupYearlyMethod containing the data to be mapped.

Returns:
  - A pointer to TopupYearlyMethodResponse containing the mapped data, including
    Year, TopupMethod, TotalTopups, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupYearlyMethodResponse(s *record.TopupYearlyMethod) *response.TopupYearlyMethodResponse
```

##### `ToTopupYearlyMethodResponses`

ToTopupYearlyMethodResponses maps a slice of TopupYearlyMethod domain models
to a slice of TopupYearlyMethodResponse API-compatible response types.

Args:
  - s: A slice of TopupYearlyMethod containing the data to be mapped.

Returns:
  - A slice of TopupYearlyMethodResponse containing the mapped data, including
    Year, TopupMethod, TotalTopups, and TotalAmount.

```go
func (t *topupResponseMapper) ToTopupYearlyMethodResponses(s []*record.TopupYearlyMethod) []*response.TopupYearlyMethodResponse
```

### `transactionResponseMapper`

transactionResponseMapper provides methods to map TransactionRecord domain models to TransactionResponse API-compatible response types.

```go
type transactionResponseMapper struct {
}
```

#### Methods

##### `ToTransactionMonthlyAmountResponse`

ToTransactionMonthlyAmountResponse converts a monthly amount record into a TransactionMonthAmountResponse.

Args:
  - s: A pointer to TransactionMonthAmount containing the data to be mapped.

Returns:
  - A pointer to TransactionMonthAmountResponse containing the mapped data, including Month and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionMonthlyAmountResponse(s *record.TransactionMonthAmount) *response.TransactionMonthAmountResponse
```

##### `ToTransactionMonthlyAmountResponses`

ToTransactionMonthlyAmountResponses maps a slice of TransactionMonthAmount domain models
into a slice of TransactionMonthAmountResponse API-compatible response types.

Args:
  - s: A slice of pointers to TransactionMonthAmount containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionMonthAmountResponse containing the mapped data, including
    Month and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionMonthlyAmountResponses(s []*record.TransactionMonthAmount) []*response.TransactionMonthAmountResponse
```

##### `ToTransactionMonthlyMethodResponse`

ToTransactionMonthlyMethodResponse maps a record of monthly transaction methods into a TransactionMonthMethodResponse.

Args:
  - s: A pointer to TransactionMonthMethod containing the data to be mapped.

Returns:
  - A pointer to TransactionMonthMethodResponse containing the mapped data, including Month,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionMonthlyMethodResponse(s *record.TransactionMonthMethod) *response.TransactionMonthMethodResponse
```

##### `ToTransactionMonthlyMethodResponses`

ToTransactionMonthlyMethodResponses maps a slice of TransactionMonthMethod domain models
into a slice of TransactionMonthMethodResponse API-compatible response types.

Args:
  - s: A slice of pointers to TransactionMonthMethod containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionMonthMethodResponse containing the mapped data, including
    Month, PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionMonthlyMethodResponses(s []*record.TransactionMonthMethod) []*response.TransactionMonthMethodResponse
```

##### `ToTransactionResponse`

ToTransactionResponse converts a single transaction record into a TransactionResponse.

Args:
  - transaction: A pointer to a TransactionRecord containing the data to be mapped.

Returns:
  - A pointer to a TransactionResponse containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, and UpdatedAt.

```go
func (s *transactionResponseMapper) ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse
```

##### `ToTransactionResponseDeleteAt`

ToTransactionResponseDeleteAt maps a soft-deleted transaction record into a TransactionResponseDeleteAt.

Args:
  - transaction: A pointer to a TransactionRecord containing the data to be mapped.

Returns:
  - A pointer to a TransactionResponseDeleteAt containing the mapped data,
    including ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID,
    TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionResponseMapper) ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt
```

##### `ToTransactionResponseMonthStatusFailed`

ToTransactionResponseMonthStatusFailed maps a record of monthly transaction failed status into a TransactionResponseMonthStatusFailed.

Args:
  - s: A pointer to TransactionRecordMonthStatusFailed containing the data to be mapped.

Returns:
  - A pointer to TransactionResponseMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponseMonthStatusFailed(s *record.TransactionRecordMonthStatusFailed) *response.TransactionResponseMonthStatusFailed
```

##### `ToTransactionResponseMonthStatusSuccess`

ToTransactionResponseMonthStatusSuccess maps a record of monthly transaction success status
into a TransactionResponseMonthStatusSuccess.

Args:
  - s: A pointer to TransactionRecordMonthStatusSuccess containing the data to be mapped.

Returns:
  - A pointer to TransactionResponseMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponseMonthStatusSuccess(s *record.TransactionRecordMonthStatusSuccess) *response.TransactionResponseMonthStatusSuccess
```

##### `ToTransactionResponseYearStatusFailed`

ToTransactionResponseYearStatusFailed maps a yearly transaction failed status record
into a TransactionResponseYearStatusFailed.

Args:
  - s: A pointer to TransactionRecordYearStatusFailed containing the data to be mapped.

Returns:
  - A pointer to TransactionResponseYearStatusFailed containing the mapped data, including
    Year, TotalFailed, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponseYearStatusFailed(s *record.TransactionRecordYearStatusFailed) *response.TransactionResponseYearStatusFailed
```

##### `ToTransactionResponseYearStatusSuccess`

ToTransactionResponseYearStatusSuccess maps a record of yearly transaction success status
into a TransactionResponseYearStatusSuccess.

Args:
  - s: A pointer to TransactionRecordYearStatusSuccess containing the data to be mapped.

Returns:
  - A pointer to TransactionResponseYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponseYearStatusSuccess(s *record.TransactionRecordYearStatusSuccess) *response.TransactionResponseYearStatusSuccess
```

##### `ToTransactionResponsesMonthStatusFailed`

ToTransactionResponsesMonthStatusFailed converts multiple records of monthly transaction failed status into a slice of
TransactionResponseMonthStatusFailed.

Args:
  - Transactions: A slice of pointers to TransactionRecordMonthStatusFailed containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionResponseMonthStatusFailed containing the mapped data, including Year, Month, TotalFailed, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponsesMonthStatusFailed(Transactions []*record.TransactionRecordMonthStatusFailed) []*response.TransactionResponseMonthStatusFailed
```

##### `ToTransactionResponsesMonthStatusSuccess`

ToTransactionResponsesMonthStatusSuccess converts multiple records of monthly transaction success status into a slice of
TransactionResponseMonthStatusSuccess.

Args:
  - Transactions: A slice of pointers to TransactionRecordMonthStatusSuccess containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionResponseMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponsesMonthStatusSuccess(Transactions []*record.TransactionRecordMonthStatusSuccess) []*response.TransactionResponseMonthStatusSuccess
```

##### `ToTransactionResponsesYearStatusFailed`

ToTransactionResponsesYearStatusFailed converts multiple yearly transaction failed status records
into a slice of TransactionResponseYearStatusFailed.

Args:
  - Transactions: A slice of pointers to TransactionRecordYearStatusFailed containing the records to be mapped.

Returns:
  - A slice of pointers to TransactionResponseYearStatusFailed containing the mapped data, including
    Year, TotalFailed, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponsesYearStatusFailed(Transactions []*record.TransactionRecordYearStatusFailed) []*response.TransactionResponseYearStatusFailed
```

##### `ToTransactionResponsesYearStatusSuccess`

ToTransactionResponsesYearStatusSuccess converts multiple records of yearly transaction success status into a slice of
TransactionResponseYearStatusSuccess.

Args:
  - Transactions: A slice of pointers to TransactionRecordYearStatusSuccess containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionResponseYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionResponsesYearStatusSuccess(Transactions []*record.TransactionRecordYearStatusSuccess) []*response.TransactionResponseYearStatusSuccess
```

##### `ToTransactionYearlyAmountResponse`

ToTransactionYearlyAmountResponse maps a yearly amount record into a TransactionYearlyAmountResponse.

Args:
  - s: A pointer to TransactionYearlyAmount containing the data to be mapped.

Returns:
  - A pointer to TransactionYearlyAmountResponse containing the mapped data, including Year and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionYearlyAmountResponse(s *record.TransactionYearlyAmount) *response.TransactionYearlyAmountResponse
```

##### `ToTransactionYearlyAmountResponses`

ToTransactionYearlyAmountResponses maps a slice of TransactionYearlyAmount domain models
into a slice of TransactionYearlyAmountResponse API-compatible response types.

Args:
  - s: A slice of pointers to TransactionYearlyAmount containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionYearlyAmountResponse containing the mapped data, including
    Year and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionYearlyAmountResponses(s []*record.TransactionYearlyAmount) []*response.TransactionYearlyAmountResponse
```

##### `ToTransactionYearlyMethodResponse`

ToTransactionYearlyMethodResponse maps a single yearly transaction method record into a TransactionYearMethodResponse.

Args:
  - s: A pointer to TransactionYearMethod containing the data to be mapped.

Returns:
  - A pointer to TransactionYearMethodResponse containing the mapped data, including Year,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionYearlyMethodResponse(s *record.TransactionYearMethod) *response.TransactionYearMethodResponse
```

##### `ToTransactionYearlyMethodResponses`

ToTransactionYearlyMethodResponses maps a slice of TransactionYearMethod domain models
into a slice of TransactionYearMethodResponse API-compatible response types.

Args:
  - s: A slice of pointers to TransactionYearMethod containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionYearMethodResponse containing the mapped data, including
    Year, PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (t *transactionResponseMapper) ToTransactionYearlyMethodResponses(s []*record.TransactionYearMethod) []*response.TransactionYearMethodResponse
```

##### `ToTransactionsResponse`

ToTransactionsResponse converts multiple transaction records into a slice of TransactionResponse.

Args:
  - transactions: A slice of pointers to TransactionRecord containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionResponse containing the mapped data, including ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, and UpdatedAt.

```go
func (s *transactionResponseMapper) ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse
```

##### `ToTransactionsResponseDeleteAt`

ToTransactionsResponseDeleteAt converts multiple soft-deleted transaction records into a slice of TransactionResponseDeleteAt.

Args:
  - transactions: A slice of pointers to TransactionRecord containing the data to be mapped.

Returns:
  - A slice of pointers to TransactionResponseDeleteAt containing the mapped data, including
    ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt,
    UpdatedAt, and DeletedAt.

```go
func (s *transactionResponseMapper) ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt
```

### `transferResponseMapper`

transferResponseMapper provides methods to map TransferRecord domain models to TransferResponse API-compatible response types.

```go
type transferResponseMapper struct {
}
```

#### Methods

##### `ToTransferResponse`

ToTransferResponse converts a single transfer record into a TransferResponse.

Args:
  - transfer: The transfer record to be converted.

Returns:
  - A pointer to a response.TransferResponse containing the mapped data.

```go
func (s *transferResponseMapper) ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse
```

##### `ToTransferResponseDeleteAt`

ToTransferResponseDeleteAt converts a single transfer record into a TransferResponseDeleteAt.

Args:
  - transfer: The transfer record to be converted.

Returns:
  - A pointer to a response.TransferResponseDeleteAt containing the mapped data.

```go
func (s *transferResponseMapper) ToTransferResponseDeleteAt(transfer *record.TransferRecord) *response.TransferResponseDeleteAt
```

##### `ToTransferResponseMonthAmount`

ToTransferResponseMonthAmount maps a TransferMonthAmount record to a TransferMonthAmountResponse.

Args:
  - s: A pointer to a TransferMonthAmount containing the data to be mapped.

Returns:
  - A pointer to a TransferMonthAmountResponse containing the mapped data.

```go
func (t *transferResponseMapper) ToTransferResponseMonthAmount(s *record.TransferMonthAmount) *response.TransferMonthAmountResponse
```

##### `ToTransferResponseMonthStatusFailed`

ToTransferResponseMonthStatusFailed converts a monthly failed status transfer record
into a TransferResponseMonthStatusFailed.

Args:
  - s: A pointer to a TransferRecordMonthStatusFailed containing the data to be mapped.

Returns:
  - A pointer to a TransferResponseMonthStatusFailed containing the mapped data, including
    Year, Month, TotalFailed, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponseMonthStatusFailed(s *record.TransferRecordMonthStatusFailed) *response.TransferResponseMonthStatusFailed
```

##### `ToTransferResponseMonthStatusSuccess`

ToTransferResponseMonthStatusSuccess converts a monthly success status transfer record
into a TransferResponseMonthStatusSuccess.

Args:
  - s: A pointer to a TransferRecordMonthStatusSuccess containing the data to be mapped.

Returns:
  - A pointer to a TransferResponseMonthStatusSuccess containing the mapped data, including
    Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponseMonthStatusSuccess(s *record.TransferRecordMonthStatusSuccess) *response.TransferResponseMonthStatusSuccess
```

##### `ToTransferResponseYearAmount`

ToTransferResponseYearAmount converts a yearly transfer amount record into a TransferYearAmountResponse.

Args:
  - s: A pointer to a TransferYearAmount containing the data to be mapped.

Returns:
  - A pointer to a TransferYearAmountResponse containing the mapped data, including Year and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponseYearAmount(s *record.TransferYearAmount) *response.TransferYearAmountResponse
```

##### `ToTransferResponseYearStatusFailed`

ToTransferResponseYearStatusFailed converts a yearly failed status transfer record
into a TransferResponseYearStatusFailed.

Args:
  - s: A pointer to a TransferRecordYearStatusFailed containing the data to be mapped.

Returns:
  - A pointer to a TransferResponseYearStatusFailed containing the mapped data, including
    Year, TotalFailed, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponseYearStatusFailed(s *record.TransferRecordYearStatusFailed) *response.TransferResponseYearStatusFailed
```

##### `ToTransferResponseYearStatusSuccess`

ToTransferResponseYearStatusSuccess converts a yearly success status transfer record
into a TransferResponseYearStatusSuccess.

Args:
  - s: A pointer to a TransferRecordYearStatusSuccess containing the data to be mapped.

Returns:
  - A pointer to a TransferResponseYearStatusSuccess containing the mapped data, including
    Year, TotalSuccess, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponseYearStatusSuccess(s *record.TransferRecordYearStatusSuccess) *response.TransferResponseYearStatusSuccess
```

##### `ToTransferResponsesMonthAmount`

ToTransferResponsesMonthAmount converts a slice of TransferMonthAmount records
into a slice of TransferMonthAmountResponse API-compatible response types.

Args:
  - s: A slice of pointers to TransferMonthAmount containing the data to be mapped.

Returns:
  - A slice of pointers to TransferMonthAmountResponse containing the mapped data.

```go
func (t *transferResponseMapper) ToTransferResponsesMonthAmount(s []*record.TransferMonthAmount) []*response.TransferMonthAmountResponse
```

##### `ToTransferResponsesMonthStatusFailed`

ToTransferResponsesMonthStatusFailed converts multiple monthly failed status transfer records
into a slice of TransferResponseMonthStatusFailed.

Args:
  - Transfers: A slice of pointers to TransferRecordMonthStatusFailed containing the data to be mapped.

Returns:
  - A slice of pointers to TransferResponseMonthStatusFailed containing the mapped data, including
    Year, Month, TotalFailed, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponsesMonthStatusFailed(Transfers []*record.TransferRecordMonthStatusFailed) []*response.TransferResponseMonthStatusFailed
```

##### `ToTransferResponsesMonthStatusSuccess`

ToTransferResponsesMonthStatusSuccess converts multiple monthly success status transfer records
into a slice of TransferResponseMonthStatusSuccess.

Args:
  - Transfers: A slice of pointers to TransferRecordMonthStatusSuccess containing the data to be mapped.

Returns:
  - A slice of pointers to TransferResponseMonthStatusSuccess containing the mapped data, including
    Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponsesMonthStatusSuccess(Transfers []*record.TransferRecordMonthStatusSuccess) []*response.TransferResponseMonthStatusSuccess
```

##### `ToTransferResponsesYearAmount`

ToTransferResponsesYearAmount converts multiple yearly transfer amount records into a slice of TransferYearAmountResponse API-compatible response types.

Args:
  - s: A slice of pointers to TransferYearAmount containing the data to be mapped.

Returns:
  - A slice of pointers to TransferYearAmountResponse containing the mapped data.

```go
func (t *transferResponseMapper) ToTransferResponsesYearAmount(s []*record.TransferYearAmount) []*response.TransferYearAmountResponse
```

##### `ToTransferResponsesYearStatusFailed`

ToTransferResponsesYearStatusFailed converts multiple yearly failed status transfer records
into a slice of TransferResponseYearStatusFailed.

Args:
  - Transfers: A slice of pointers to TransferRecordYearStatusFailed containing the data to be mapped.

Returns:
  - A slice of pointers to TransferResponseYearStatusFailed containing the mapped data, including
    Year, TotalFailed, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponsesYearStatusFailed(Transfers []*record.TransferRecordYearStatusFailed) []*response.TransferResponseYearStatusFailed
```

##### `ToTransferResponsesYearStatusSuccess`

ToTransferResponsesYearStatusSuccess converts multiple yearly success status transfer records
into a slice of TransferResponseYearStatusSuccess.

Args:
  - Transfers: A slice of pointers to TransferRecordYearStatusSuccess containing the data to be mapped.

Returns:
  - A slice of pointers to TransferResponseYearStatusSuccess containing the mapped data, including
    Year, TotalSuccess, and TotalAmount.

```go
func (t *transferResponseMapper) ToTransferResponsesYearStatusSuccess(Transfers []*record.TransferRecordYearStatusSuccess) []*response.TransferResponseYearStatusSuccess
```

##### `ToTransfersResponse`

ToTransfersResponse converts multiple transfer records into a slice of TransferResponse.

Args:
  - transfers: A slice of transfer records to be converted.

Returns:
  - A slice of pointers to response.TransferResponse containing the mapped data.

```go
func (s *transferResponseMapper) ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse
```

##### `ToTransfersResponseDeleteAt`

ToTransfersResponseDeleteAt converts multiple soft-deleted transfer records into a slice of TransferResponseDeleteAt.

Args:
  - transfers: A slice of pointers to TransferRecord containing the data to be mapped.

Returns:
  - A slice of pointers to TransferResponseDeleteAt containing the mapped data, including
    ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime, CreatedAt,
    UpdatedAt, and DeletedAt.

```go
func (s *transferResponseMapper) ToTransfersResponseDeleteAt(transfers []*record.TransferRecord) []*response.TransferResponseDeleteAt
```

### `userResponseMapper`

userResponseMapper provides methods to map UserRecord domain models to UserResponse API-compatible response types.

```go
type userResponseMapper struct {
}
```

#### Methods

##### `ToUserResponse`

ToUserResponse maps a UserRecord domain model to a UserResponse API-compatible format.

Args:
  - user: A pointer to a UserRecord containing the user data to be mapped.

Returns:
  - A pointer to a UserResponse with the mapped user data, including ID, FirstName, LastName,
    Email, IsVerified, CreatedAt, and UpdatedAt.

```go
func (s *userResponseMapper) ToUserResponse(user *record.UserRecord) *response.UserResponse
```

##### `ToUserResponseDeleteAt`

ToUserResponseDeleteAt converts a UserRecord domain model into a UserResponseDeleteAt API-compatible format.

Args:
  - user: A pointer to a UserRecord containing the user data to be mapped.

Returns:
  - A pointer to a UserResponseDeleteAt with the mapped user data, including ID, FirstName, LastName,
    Email, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userResponseMapper) ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt
```

##### `ToUsersResponse`

ToUsersResponse converts a slice of UserRecord domain models into a slice of UserResponse API-compatible formats.

Args:
  - users: A slice of pointers to UserRecord containing the user data to be mapped.

Returns:
  - A slice of pointers to UserResponse with the mapped user data for each user.

```go
func (s *userResponseMapper) ToUsersResponse(users []*record.UserRecord) []*response.UserResponse
```

##### `ToUsersResponseDeleteAt`

ToUsersResponseDeleteAt converts a slice of UserRecord domain models into a slice of UserResponseDeleteAt API-compatible formats.

Args:
  - users: A slice of pointers to UserRecord containing the user data to be mapped.

Returns:
  - A slice of pointers to UserResponseDeleteAt with the mapped user data for each user.

```go
func (s *userResponseMapper) ToUsersResponseDeleteAt(users []*record.UserRecord) []*response.UserResponseDeleteAt
```

### `withdrawResponseMapper`

withdrawResponseMapper provides methods to map WithdrawRecord domain models to WithdrawResponse API-compatible response types.

```go
type withdrawResponseMapper struct {
}
```

#### Methods

##### `ToWithdrawAmountMonthlyResponse`

ToWithdrawAmountMonthlyResponse converts a single WithdrawMonthlyAmount record
into a WithdrawMonthlyAmountResponse.

Args:
  - s: The WithdrawMonthlyAmount that needs to be converted.

Returns:
  - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data,
    including Month and TotalAmount.

```go
func (w *withdrawResponseMapper) ToWithdrawAmountMonthlyResponse(s *record.WithdrawMonthlyAmount) *response.WithdrawMonthlyAmountResponse
```

##### `ToWithdrawAmountYearlyResponse`

ToWithdrawAmountYearlyResponse converts a WithdrawYearlyAmount record
into a WithdrawYearlyAmountResponse.

Args:
  - s: The WithdrawYearlyAmount that needs to be converted.

Returns:
  - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, including
    Year and TotalAmount.

```go
func (w *withdrawResponseMapper) ToWithdrawAmountYearlyResponse(s *record.WithdrawYearlyAmount) *response.WithdrawYearlyAmountResponse
```

##### `ToWithdrawResponse`

ToWithdrawResponse converts a single WithdrawRecord into a WithdrawResponse.

Args:
  - withdraw: The WithdrawRecord that needs to be converted.

Returns:
  - A pointer to a WithdrawResponse containing the mapped data, including ID, WithdrawNo,
    CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.

```go
func (s *withdrawResponseMapper) ToWithdrawResponse(withdraw *record.WithdrawRecord) *response.WithdrawResponse
```

##### `ToWithdrawResponseDeleteAt`

ToWithdrawResponseDeleteAt converts a single WithdrawRecord into a WithdrawResponseDeleteAt.

Args:
  - withdraw: The WithdrawRecord that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including ID, WithdrawNo,
    CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *withdrawResponseMapper) ToWithdrawResponseDeleteAt(withdraw *record.WithdrawRecord) *response.WithdrawResponseDeleteAt
```

##### `ToWithdrawResponseMonthStatusFailed`

ToWithdrawResponseMonthStatusFailed converts a single WithdrawRecordMonthStatusFailed into a WithdrawResponseMonthStatusFailed.

Args:
  - s: The WithdrawRecordMonthStatusFailed that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseMonthStatusFailed containing the mapped data, including Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponseMonthStatusFailed(s *record.WithdrawRecordMonthStatusFailed) *response.WithdrawResponseMonthStatusFailed
```

##### `ToWithdrawResponseMonthStatusSuccess`

ToWithdrawResponseMonthStatusSuccess converts a single WithdrawRecordMonthStatusSuccess into a WithdrawResponseMonthStatusSuccess.

Args:
  - s: The WithdrawRecordMonthStatusSuccess that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseMonthStatusSuccess containing the mapped data, including Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponseMonthStatusSuccess(s *record.WithdrawRecordMonthStatusSuccess) *response.WithdrawResponseMonthStatusSuccess
```

##### `ToWithdrawResponseYearStatusFailed`

ToWithdrawResponseYearStatusFailed converts a single WithdrawRecordYearStatusFailed into
a WithdrawResponseYearStatusFailed.

Args:
  - s: The WithdrawRecordYearStatusFailed that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseYearStatusFailed containing the mapped data, including
    Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponseYearStatusFailed(s *record.WithdrawRecordYearStatusFailed) *response.WithdrawResponseYearStatusFailed
```

##### `ToWithdrawResponseYearStatusSuccess`

ToWithdrawResponseYearStatusSuccess converts a single WithdrawRecordYearStatusSuccess into a WithdrawResponseYearStatusSuccess.

Args:
  - s: The WithdrawRecordYearStatusSuccess that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseYearStatusSuccess containing the mapped data, including Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponseYearStatusSuccess(s *record.WithdrawRecordYearStatusSuccess) *response.WithdrawResponseYearStatusSuccess
```

##### `ToWithdrawResponsesMonthStatusFailed`

ToWithdrawResponsesMonthStatusFailed converts multiple WithdrawRecordMonthStatusFailed into a slice of WithdrawResponseMonthStatusFailed.

Args:
  - Withdraws: A slice of pointers to WithdrawRecordMonthStatusFailed that need to be converted.

Returns:
  - A slice of pointers to WithdrawResponseMonthStatusFailed, each containing the mapped data
    including Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponsesMonthStatusFailed(Withdraws []*record.WithdrawRecordMonthStatusFailed) []*response.WithdrawResponseMonthStatusFailed
```

##### `ToWithdrawResponsesMonthStatusSuccess`

ToWithdrawResponsesMonthStatusSuccess converts multiple WithdrawRecordMonthStatusSuccess into a slice of WithdrawResponseMonthStatusSuccess.

Args:
  - Withdraws: The WithdrawRecordMonthStatusSuccess that needs to be converted.

Returns:
  - A slice of WithdrawResponseMonthStatusSuccess containing the mapped data, including Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponsesMonthStatusSuccess(Withdraws []*record.WithdrawRecordMonthStatusSuccess) []*response.WithdrawResponseMonthStatusSuccess
```

##### `ToWithdrawResponsesYearStatusFailed`

ToWithdrawResponsesYearStatusFailed converts multiple yearly failed-status WithdrawRecords
into a slice of WithdrawResponseYearStatusFailed.

Args:
  - Withdraws: A slice of pointers to WithdrawRecordYearStatusFailed that need to be converted.

Returns:
  - A slice of pointers to WithdrawResponseYearStatusFailed, each containing the mapped data
    including Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponsesYearStatusFailed(Withdraws []*record.WithdrawRecordYearStatusFailed) []*response.WithdrawResponseYearStatusFailed
```

##### `ToWithdrawResponsesYearStatusSuccess`

ToWithdrawResponsesYearStatusSuccess converts multiple yearly success-status WithdrawRecords
into a slice of WithdrawResponseYearStatusSuccess.

Args:
  - Withdraws: A slice of pointers to WithdrawRecordYearStatusSuccess that need to be converted.

Returns:
  - A slice of pointers to WithdrawResponseYearStatusSuccess, each containing the mapped data
    including Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawResponseMapper) ToWithdrawResponsesYearStatusSuccess(Withdraws []*record.WithdrawRecordYearStatusSuccess) []*response.WithdrawResponseYearStatusSuccess
```

##### `ToWithdrawsAmountMonthlyResponses`

ToWithdrawsAmountMonthlyResponses converts a slice of WithdrawMonthlyAmount
records into a slice of WithdrawMonthlyAmountResponse objects.

Args:
  - s: A slice of pointers to WithdrawMonthlyAmount records that need to be converted.

Returns:
  - A slice of pointers to WithdrawMonthlyAmountResponse, each containing the mapped data,
    including Month and TotalAmount.

```go
func (w *withdrawResponseMapper) ToWithdrawsAmountMonthlyResponses(s []*record.WithdrawMonthlyAmount) []*response.WithdrawMonthlyAmountResponse
```

##### `ToWithdrawsAmountYearlyResponses`

ToWithdrawsAmountYearlyResponses converts a slice of WithdrawYearlyAmount records
into a slice of WithdrawYearlyAmountResponse.

Args:
  - s: A slice of pointers to WithdrawYearlyAmount records that need to be converted.

Returns:
  - A slice of pointers to WithdrawYearlyAmountResponse, each containing the mapped data,
    including Year and TotalAmount.

```go
func (w *withdrawResponseMapper) ToWithdrawsAmountYearlyResponses(s []*record.WithdrawYearlyAmount) []*response.WithdrawYearlyAmountResponse
```

##### `ToWithdrawsResponse`

ToWithdrawsResponse converts multiple WithdrawRecords into a slice of WithdrawResponse.

Args:
  - withdraws: The WithdrawRecords that needs to be converted.

Returns:
  - A slice of WithdrawResponse containing the mapped data, including ID, WithdrawNo,
    CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.

```go
func (s *withdrawResponseMapper) ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse
```

##### `ToWithdrawsResponseDeleteAt`

```go
func (s *withdrawResponseMapper) ToWithdrawsResponseDeleteAt(withdraws []*record.WithdrawRecord) []*response.WithdrawResponseDeleteAt
```