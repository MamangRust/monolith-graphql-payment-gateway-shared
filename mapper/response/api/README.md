# 📦 Package `apimapper`

**Source Path:** `shared/mapper/response/api`

## 🧩 Types

### `AuthResponseMapper`

AuthResponseMapper defines methods to map gRPC Auth responses to HTTP API responses.

```go
type AuthResponseMapper interface {
	ToResponseVerifyCode func(res *pb.ApiResponseVerifyCode) (*response.ApiResponseVerifyCode)
	ToResponseForgotPassword func(res *pb.ApiResponseForgotPassword) (*response.ApiResponseForgotPassword)
	ToResponseResetPassword func(res *pb.ApiResponseResetPassword) (*response.ApiResponseResetPassword)
	ToResponseLogin func(res *pb.ApiResponseLogin) (*response.ApiResponseLogin)
	ToResponseRegister func(res *pb.ApiResponseRegister) (*response.ApiResponseRegister)
	ToResponseRefreshToken func(res *pb.ApiResponseRefreshToken) (*response.ApiResponseRefreshToken)
	ToResponseGetMe func(res *pb.ApiResponseGetMe) (*response.ApiResponseGetMe)
}
```

### `CardResponseMapper`

CardResponseMapper defines methods for converting gRPC card-related responses
(protobuf format) into HTTP-compatible API response structures used in the presentation layer.

```go
type CardResponseMapper interface {
	ToApiResponseCard func(card *pb.ApiResponseCard) (*response.ApiResponseCard)
	ToApiResponsesCard func(cards *pb.ApiResponsePaginationCard) (*response.ApiResponsePaginationCard)
	ToApiResponseCardAll func(card *pb.ApiResponseCardAll) (*response.ApiResponseCardAll)
	ToApiResponseCardDeleteAt func(card *pb.ApiResponseCardDelete) (*response.ApiResponseCardDelete)
	ToApiResponsesCardDeletedAt func(cards *pb.ApiResponsePaginationCardDeleteAt) (*response.ApiResponsePaginationCardDeleteAt)
	ToApiResponseDashboardCard func(dash *pb.ApiResponseDashboardCard) (*response.ApiResponseDashboardCard)
	ToApiResponseDashboardCardCardNumber func(dash *pb.ApiResponseDashboardCardNumber) (*response.ApiResponseDashboardCardNumber)
	ToApiResponseMonthlyBalances func(cards *pb.ApiResponseMonthlyBalance) (*response.ApiResponseMonthlyBalance)
	ToApiResponseYearlyBalances func(cards *pb.ApiResponseYearlyBalance) (*response.ApiResponseYearlyBalance)
	ToApiResponseMonthlyAmounts func(cards *pb.ApiResponseMonthlyAmount) (*response.ApiResponseMonthlyAmount)
	ToApiResponseYearlyAmounts func(cards *pb.ApiResponseYearlyAmount) (*response.ApiResponseYearlyAmount)
}
```

### `MerchantDocumentResponseMapper`

MerchantDocumentResponseMapper defines methods to map gRPC merchant document responses
into HTTP API response structures used in the presentation layer.

```go
type MerchantDocumentResponseMapper interface {
	ToApiResponseMerchantDocument func(doc *pb.ApiResponseMerchantDocument) (*response.ApiResponseMerchantDocument)
	ToApiResponsesMerchantDocument func(docs *pb.ApiResponsesMerchantDocument) (*response.ApiResponsesMerchantDocument)
	ToApiResponsePaginationMerchantDocument func(docs *pb.ApiResponsePaginationMerchantDocument) (*response.ApiResponsePaginationMerchantDocument)
	ToApiResponsePaginationMerchantDocumentDeleteAt func(docs *pb.ApiResponsePaginationMerchantDocumentAt) (*response.ApiResponsePaginationMerchantDocumentDeleteAt)
	ToApiResponseMerchantDocumentAll func(resp *pb.ApiResponseMerchantDocumentAll) (*response.ApiResponseMerchantDocumentAll)
	ToApiResponseMerchantDocumentDeleteAt func(resp *pb.ApiResponseMerchantDocumentDelete) (*response.ApiResponseMerchantDocumentDelete)
}
```

### `MerchantResponseMapper`

MerchantResponseMapper defines methods for converting gRPC merchant-related responses
to HTTP-compatible API response formats used in the REST layer.

```go
type MerchantResponseMapper interface {
	ToApiResponseMerchant func(merchants *pb.ApiResponseMerchant) (*response.ApiResponseMerchant)
	ToApiResponseMerchants func(merchants *pb.ApiResponsesMerchant) (*response.ApiResponsesMerchant)
	ToApiResponsesMerchant func(merchants *pb.ApiResponsePaginationMerchant) (*response.ApiResponsePaginationMerchant)
	ToApiResponsesMerchantDeleteAt func(merchants *pb.ApiResponsePaginationMerchantDeleteAt) (*response.ApiResponsePaginationMerchantDeleteAt)
	ToApiResponseMerchantAll func(card *pb.ApiResponseMerchantAll) (*response.ApiResponseMerchantAll)
	ToApiResponseMerchantDeleteAt func(card *pb.ApiResponseMerchantDelete) (*response.ApiResponseMerchantDelete)
	ToApiResponseMerchantsTransactionResponse func(merchants *pb.ApiResponsePaginationMerchantTransaction) (*response.ApiResponsePaginationMerchantTransaction)
	ToApiResponseMonthlyPaymentMethods func(ms *pb.ApiResponseMerchantMonthlyPaymentMethod) (*response.ApiResponseMerchantMonthlyPaymentMethod)
	ToApiResponseYearlyPaymentMethods func(ms *pb.ApiResponseMerchantYearlyPaymentMethod) (*response.ApiResponseMerchantYearlyPaymentMethod)
	ToApiResponseMonthlyAmounts func(ms *pb.ApiResponseMerchantMonthlyAmount) (*response.ApiResponseMerchantMonthlyAmount)
	ToApiResponseYearlyAmounts func(ms *pb.ApiResponseMerchantYearlyAmount) (*response.ApiResponseMerchantYearlyAmount)
	ToApiResponseMonthlyTotalAmounts func(ms *pb.ApiResponseMerchantMonthlyTotalAmount) (*response.ApiResponseMerchantMonthlyTotalAmount)
	ToApiResponseYearlyTotalAmounts func(ms *pb.ApiResponseMerchantYearlyTotalAmount) (*response.ApiResponseMerchantYearlyTotalAmount)
}
```

### `ResponseApiMapper`

```go
type ResponseApiMapper struct {
	AuthResponseMapper AuthResponseMapper
	CardResponseMapper CardResponseMapper
	RoleResponseMapper RoleResponseMapper
	SaldoResponseMapper SaldoResponseMapper
	TransactionResponseMapper TransactionResponseMapper
	TransferResponseMapper TransferResponseMapper
	TopupResponseMapper TopupResponseMapper
	WithdrawResponseMapper WithdrawResponseMapper
	UserResponseMapper UserResponseMapper
	MerchantResponseMapper MerchantResponseMapper
	MerchantDocumentProMapper MerchantDocumentResponseMapper
}
```

### `RoleResponseMapper`

RoleResponseMapper defines a set of methods to map gRPC Role API responses

```go
type RoleResponseMapper interface {
	ToApiResponseRoleAll func(pbResponse *pb.ApiResponseRoleAll) (*response.ApiResponseRoleAll)
	ToApiResponseRoleDelete func(pbResponse *pb.ApiResponseRoleDelete) (*response.ApiResponseRoleDelete)
	ToApiResponseRole func(pbResponse *pb.ApiResponseRole) (*response.ApiResponseRole)
	ToApiResponsesRole func(pbResponse *pb.ApiResponsesRole) (*response.ApiResponsesRole)
	ToApiResponsePaginationRole func(pbResponse *pb.ApiResponsePaginationRole) (*response.ApiResponsePaginationRole)
	ToApiResponsePaginationRoleDeleteAt func(pbResponse *pb.ApiResponsePaginationRoleDeleteAt) (*response.ApiResponsePaginationRoleDeleteAt)
}
```

### `SaldoResponseMapper`

SaldoResponseMapper defines methods for transforming gRPC saldo-related responses
into HTTP-friendly API response formats used in the RESTful layer.

```go
type SaldoResponseMapper interface {
	ToApiResponseSaldo func(pbResponse *pb.ApiResponseSaldo) (*response.ApiResponseSaldo)
	ToApiResponsesSaldo func(pbResponse *pb.ApiResponsesSaldo) (*response.ApiResponsesSaldo)
	ToApiResponseSaldoDelete func(pbResponse *pb.ApiResponseSaldoDelete) (*response.ApiResponseSaldoDelete)
	ToApiResponseSaldoAll func(pbResponse *pb.ApiResponseSaldoAll) (*response.ApiResponseSaldoAll)
	ToApiResponseMonthTotalSaldo func(pbResponse *pb.ApiResponseMonthTotalSaldo) (*response.ApiResponseMonthTotalSaldo)
	ToApiResponseYearTotalSaldo func(pbResponse *pb.ApiResponseYearTotalSaldo) (*response.ApiResponseYearTotalSaldo)
	ToApiResponseMonthSaldoBalances func(pbResponse *pb.ApiResponseMonthSaldoBalances) (*response.ApiResponseMonthSaldoBalances)
	ToApiResponseYearSaldoBalances func(pbResponse *pb.ApiResponseYearSaldoBalances) (*response.ApiResponseYearSaldoBalances)
	ToApiResponsePaginationSaldo func(pbResponse *pb.ApiResponsePaginationSaldo) (*response.ApiResponsePaginationSaldo)
	ToApiResponsePaginationSaldoDeleteAt func(pbResponse *pb.ApiResponsePaginationSaldoDeleteAt) (*response.ApiResponsePaginationSaldoDeleteAt)
}
```

### `TopupResponseMapper`

TopupResponseMapper defines methods to convert gRPC top-up responses into API responses
for use in the HTTP layer of the application.

```go
type TopupResponseMapper interface {
	ToApiResponseTopup func(s *pb.ApiResponseTopup) (*response.ApiResponseTopup)
	ToApiResponseTopupDeleteAt func(s *pb.ApiResponseTopupDeleteAt) (*response.ApiResponseTopupDeleteAt)
	ToApiResponseTopupAll func(s *pb.ApiResponseTopupAll) (*response.ApiResponseTopupAll)
	ToApiResponseTopupDelete func(s *pb.ApiResponseTopupDelete) (*response.ApiResponseTopupDelete)
	ToApiResponsePaginationTopup func(s *pb.ApiResponsePaginationTopup) (*response.ApiResponsePaginationTopup)
	ToApiResponsePaginationTopupDeleteAt func(s *pb.ApiResponsePaginationTopupDeleteAt) (*response.ApiResponsePaginationTopupDeleteAt)
	ToApiResponseTopupMonthStatusSuccess func(s *pb.ApiResponseTopupMonthStatusSuccess) (*response.ApiResponseTopupMonthStatusSuccess)
	ToApiResponseTopupYearStatusSuccess func(s *pb.ApiResponseTopupYearStatusSuccess) (*response.ApiResponseTopupYearStatusSuccess)
	ToApiResponseTopupMonthStatusFailed func(s *pb.ApiResponseTopupMonthStatusFailed) (*response.ApiResponseTopupMonthStatusFailed)
	ToApiResponseTopupYearStatusFailed func(s *pb.ApiResponseTopupYearStatusFailed) (*response.ApiResponseTopupYearStatusFailed)
	ToApiResponseTopupMonthMethod func(s *pb.ApiResponseTopupMonthMethod) (*response.ApiResponseTopupMonthMethod)
	ToApiResponseTopupYearMethod func(s *pb.ApiResponseTopupYearMethod) (*response.ApiResponseTopupYearMethod)
	ToApiResponseTopupMonthAmount func(s *pb.ApiResponseTopupMonthAmount) (*response.ApiResponseTopupMonthAmount)
	ToApiResponseTopupYearAmount func(s *pb.ApiResponseTopupYearAmount) (*response.ApiResponseTopupYearAmount)
}
```

### `TransactionResponseMapper`

TransactionResponseMapper defines methods for converting gRPC transaction-related responses
into API responses suitable for the HTTP/REST layer.

```go
type TransactionResponseMapper interface {
	ToApiResponseTransactionMonthStatusSuccess func(pbResponse *pb.ApiResponseTransactionMonthStatusSuccess) (*response.ApiResponseTransactionMonthStatusSuccess)
	ToApiResponseTransactionYearStatusSuccess func(pbResponse *pb.ApiResponseTransactionYearStatusSuccess) (*response.ApiResponseTransactionYearStatusSuccess)
	ToApiResponseTransactionMonthStatusFailed func(pbResponse *pb.ApiResponseTransactionMonthStatusFailed) (*response.ApiResponseTransactionMonthStatusFailed)
	ToApiResponseTransactionYearStatusFailed func(pbResponse *pb.ApiResponseTransactionYearStatusFailed) (*response.ApiResponseTransactionYearStatusFailed)
	ToApiResponseTransactionMonthMethod func(pbResponse *pb.ApiResponseTransactionMonthMethod) (*response.ApiResponseTransactionMonthMethod)
	ToApiResponseTransactionYearMethod func(pbResponse *pb.ApiResponseTransactionYearMethod) (*response.ApiResponseTransactionYearMethod)
	ToApiResponseTransactionMonthAmount func(pbResponse *pb.ApiResponseTransactionMonthAmount) (*response.ApiResponseTransactionMonthAmount)
	ToApiResponseTransactionYearAmount func(pbResponse *pb.ApiResponseTransactionYearAmount) (*response.ApiResponseTransactionYearAmount)
	ToApiResponseTransaction func(pbResponse *pb.ApiResponseTransaction) (*response.ApiResponseTransaction)
	ToApiResponseTransactions func(pbResponse *pb.ApiResponseTransactions) (*response.ApiResponseTransactions)
	ToApiResponseTransactionDelete func(pbResponse *pb.ApiResponseTransactionDelete) (*response.ApiResponseTransactionDelete)
	ToApiResponseTransactionAll func(pbResponse *pb.ApiResponseTransactionAll) (*response.ApiResponseTransactionAll)
	ToApiResponsePaginationTransaction func(pbResponse *pb.ApiResponsePaginationTransaction) (*response.ApiResponsePaginationTransaction)
	ToApiResponsePaginationTransactionDeleteAt func(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) (*response.ApiResponsePaginationTransactionDeleteAt)
}
```

### `TransferResponseMapper`

TransferResponseMapper defines methods for mapping gRPC transfer-related responses
into HTTP/REST API responses.

```go
type TransferResponseMapper interface {
	ToApiResponseTransferMonthStatusSuccess func(pbResponse *pb.ApiResponseTransferMonthStatusSuccess) (*response.ApiResponseTransferMonthStatusSuccess)
	ToApiResponseTransferYearStatusSuccess func(pbResponse *pb.ApiResponseTransferYearStatusSuccess) (*response.ApiResponseTransferYearStatusSuccess)
	ToApiResponseTransferMonthStatusFailed func(pbResponse *pb.ApiResponseTransferMonthStatusFailed) (*response.ApiResponseTransferMonthStatusFailed)
	ToApiResponseTransferYearStatusFailed func(pbResponse *pb.ApiResponseTransferYearStatusFailed) (*response.ApiResponseTransferYearStatusFailed)
	ToApiResponseTransferMonthAmount func(pbResponse *pb.ApiResponseTransferMonthAmount) (*response.ApiResponseTransferMonthAmount)
	ToApiResponseTransferYearAmount func(pbResponse *pb.ApiResponseTransferYearAmount) (*response.ApiResponseTransferYearAmount)
	ToApiResponseTransfer func(pbResponse *pb.ApiResponseTransfer) (*response.ApiResponseTransfer)
	ToApiResponseTransfers func(pbResponse *pb.ApiResponseTransfers) (*response.ApiResponseTransfers)
	ToApiResponseTransferDelete func(pbResponse *pb.ApiResponseTransferDelete) (*response.ApiResponseTransferDelete)
	ToApiResponseTransferAll func(pbResponse *pb.ApiResponseTransferAll) (*response.ApiResponseTransferAll)
	ToApiResponsePaginationTransfer func(pbResponse *pb.ApiResponsePaginationTransfer) (*response.ApiResponsePaginationTransfer)
	ToApiResponsePaginationTransferDeleteAt func(pbResponse *pb.ApiResponsePaginationTransferDeleteAt) (*response.ApiResponsePaginationTransferDeleteAt)
}
```

### `UserResponseMapper`

UserResponseMapper defines methods for mapping gRPC user-related responses
into HTTP/REST API responses.

```go
type UserResponseMapper interface {
	ToApiResponseUser func(pbResponse *pb.ApiResponseUser) (*response.ApiResponseUser)
	ToApiResponseUserDeleteAt func(pbResponse *pb.ApiResponseUserDeleteAt) (*response.ApiResponseUserDeleteAt)
	ToApiResponsesUser func(pbResponse *pb.ApiResponsesUser) (*response.ApiResponsesUser)
	ToApiResponseUserDelete func(pbResponse *pb.ApiResponseUserDelete) (*response.ApiResponseUserDelete)
	ToApiResponseUserAll func(pbResponse *pb.ApiResponseUserAll) (*response.ApiResponseUserAll)
	ToApiResponsePaginationUserDeleteAt func(pbResponse *pb.ApiResponsePaginationUserDeleteAt) (*response.ApiResponsePaginationUserDeleteAt)
	ToApiResponsePaginationUser func(pbResponse *pb.ApiResponsePaginationUser) (*response.ApiResponsePaginationUser)
}
```

### `WithdrawResponseMapper`

WithdrawResponseMapper defines methods for converting gRPC withdraw-related responses
into HTTP/REST API responses.

```go
type WithdrawResponseMapper interface {
	ToApiResponseWithdraw func(pbResponse *pb.ApiResponseWithdraw) (*response.ApiResponseWithdraw)
	ToApiResponsesWithdraw func(pbResponse *pb.ApiResponsesWithdraw) (*response.ApiResponsesWithdraw)
	ToApiResponseWithdrawDelete func(pbResponse *pb.ApiResponseWithdrawDelete) (*response.ApiResponseWithdrawDelete)
	ToApiResponseWithdrawAll func(pbResponse *pb.ApiResponseWithdrawAll) (*response.ApiResponseWithdrawAll)
	ToApiResponsePaginationWithdraw func(pbResponse *pb.ApiResponsePaginationWithdraw) (*response.ApiResponsePaginationWithdraw)
	ToApiResponsePaginationWithdrawDeleteAt func(pbResponse *pb.ApiResponsePaginationWithdrawDeleteAt) (*response.ApiResponsePaginationWithdrawDeleteAt)
	ToApiResponseWithdrawMonthStatusSuccess func(pbResponse *pb.ApiResponseWithdrawMonthStatusSuccess) (*response.ApiResponseWithdrawMonthStatusSuccess)
	ToApiResponseWithdrawYearStatusSuccess func(pbResponse *pb.ApiResponseWithdrawYearStatusSuccess) (*response.ApiResponseWithdrawYearStatusSuccess)
	ToApiResponseWithdrawMonthStatusFailed func(pbResponse *pb.ApiResponseWithdrawMonthStatusFailed) (*response.ApiResponseWithdrawMonthStatusFailed)
	ToApiResponseWithdrawYearStatusFailed func(pbResponse *pb.ApiResponseWithdrawYearStatusFailed) (*response.ApiResponseWithdrawYearStatusFailed)
	ToApiResponseWithdrawMonthAmount func(pbResponse *pb.ApiResponseWithdrawMonthAmount) (*response.ApiResponseWithdrawMonthAmount)
	ToApiResponseWithdrawYearAmount func(pbResponse *pb.ApiResponseWithdrawYearAmount) (*response.ApiResponseWithdrawYearAmount)
}
```

### `authResponseMapper`

authResponseMapper provides methods to map the auth response from the domain
to the api response.

```go
type authResponseMapper struct {
}
```

#### Methods

##### `ToResponseForgotPassword`

ToResponseForgotPassword maps the ApiResponseForgotPassword from the domain to the
ApiResponseForgotPassword of the api.

```go
func (s *authResponseMapper) ToResponseForgotPassword(res *pb.ApiResponseForgotPassword) *response.ApiResponseForgotPassword
```

##### `ToResponseGetMe`

ToResponseGetMe maps the ApiResponseGetMe from the domain to the ApiResponseGetMe of the api.

Args:
  - res: A pointer to a pb.ApiResponseGetMe representing the ApiResponseGetMe from the domain.

Returns:
  - A pointer to a response.ApiResponseGetMe containing the mapped data, including status, message,
    and user data such as ID, FirstName, LastName, Email, CreatedAt, and UpdatedAt.

```go
func (s *authResponseMapper) ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe
```

##### `ToResponseLogin`

ToResponseLogin maps the ApiResponseLogin from the domain to the ApiResponseLogin of the api.

```go
func (s *authResponseMapper) ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin
```

##### `ToResponseRefreshToken`

ToResponseRefreshToken maps the ApiResponseRefreshToken from the domain to the
ApiResponseRefreshToken of the api.

This function takes a pointer to a pb.ApiResponseRefreshToken and returns a
pointer to a response.ApiResponseRefreshToken, which includes status, message,
and token data such as access and refresh tokens.

```go
func (s *authResponseMapper) ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken
```

##### `ToResponseRegister`

ToResponseRegister maps the ApiResponseRegister from the domain to the
ApiResponseRegister of the api.

```go
func (s *authResponseMapper) ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister
```

##### `ToResponseResetPassword`

ToResponseResetPassword maps the ApiResponseResetPassword from the domain to the
ApiResponseResetPassword of the api.

```go
func (s *authResponseMapper) ToResponseResetPassword(res *pb.ApiResponseResetPassword) *response.ApiResponseResetPassword
```

##### `ToResponseVerifyCode`

ToResponseVerifyCode maps the ApiResponseVerifyCode from the domain to the
ApiResponseVerifyCode of the api.

```go
func (s *authResponseMapper) ToResponseVerifyCode(res *pb.ApiResponseVerifyCode) *response.ApiResponseVerifyCode
```

### `cardResponseMapper`

cardResponseMapper maps between card related response messages and domain response objects.

```go
type cardResponseMapper struct {
}
```

#### Methods

##### `ToApiResponseCard`

ToApiResponseCard maps the ApiResponseCard from the domain to the ApiResponseCard of the api.

Args:
  - card: A pointer to a pb.ApiResponseCard representing the ApiResponseCard from the domain.

Returns:
  - A pointer to a response.ApiResponseCard containing the mapped data, including status, message,
    and data.

```go
func (s *cardResponseMapper) ToApiResponseCard(card *pb.ApiResponseCard) *response.ApiResponseCard
```

##### `ToApiResponseCardAll`

ToApiResponseCardAll maps the ApiResponseCardAll from the domain to the ApiResponseCardAll of the api.

Args:
  - card: A pointer to a pb.ApiResponseCardAll representing the ApiResponseCardAll from the domain.

Returns:
  - A pointer to a response.ApiResponseCardAll containing the mapped data, including status and message.

```go
func (s *cardResponseMapper) ToApiResponseCardAll(card *pb.ApiResponseCardAll) *response.ApiResponseCardAll
```

##### `ToApiResponseCardDeleteAt`

ToApiResponseCardDeleteAt maps the ApiResponseCardDelete from the domain to the ApiResponseCardDelete of the api.

Args:
  - card: A pointer to a pb.ApiResponseCardDelete representing the ApiResponseCardDelete from the domain.

Returns:
  - A pointer to a response.ApiResponseCardDelete containing the mapped data, including status and message.

```go
func (s *cardResponseMapper) ToApiResponseCardDeleteAt(card *pb.ApiResponseCardDelete) *response.ApiResponseCardDelete
```

##### `ToApiResponseDashboardCard`

ToApiResponseDashboardCard maps the ApiResponseDashboardCard from the domain to the ApiResponseDashboardCard of the api.

Args:
  - dash: A pointer to a pb.ApiResponseDashboardCard representing the ApiResponseDashboardCard from the domain.

Returns:
  - A pointer to a response.ApiResponseDashboardCard containing the mapped data, including status, message, and data.

```go
func (s *cardResponseMapper) ToApiResponseDashboardCard(dash *pb.ApiResponseDashboardCard) *response.ApiResponseDashboardCard
```

##### `ToApiResponseDashboardCardCardNumber`

ToApiResponseDashboardCardCardNumber maps the ApiResponseDashboardCardNumber from the domain to the ApiResponseDashboardCardNumber of the api.

Args:
  - dash: A pointer to a pb.ApiResponseDashboardCardNumber representing the ApiResponseDashboardCardNumber from the domain.

Returns:
  - A pointer to a response.ApiResponseDashboardCardNumber containing the mapped data, including status, message, and data.

```go
func (s *cardResponseMapper) ToApiResponseDashboardCardCardNumber(dash *pb.ApiResponseDashboardCardNumber) *response.ApiResponseDashboardCardNumber
```

##### `ToApiResponseMonthlyAmounts`

ToApiResponseMonthlyAmounts maps the ApiResponseMonthlyAmount from the domain to the ApiResponseMonthlyAmount of the api.

Args:
  - cards: A pointer to a pb.ApiResponseMonthlyAmount representing the ApiResponseMonthlyAmount from the domain.

Returns:
  - A pointer to a response.ApiResponseMonthlyAmount containing the mapped data, including status, message, and data.

```go
func (s *cardResponseMapper) ToApiResponseMonthlyAmounts(cards *pb.ApiResponseMonthlyAmount) *response.ApiResponseMonthlyAmount
```

##### `ToApiResponseMonthlyBalances`

ToApiResponseMonthlyBalances maps the ApiResponseMonthlyBalance from the domain to the ApiResponseMonthlyBalance of the api.

Args:
  - cards: A pointer to a pb.ApiResponseMonthlyBalance representing the ApiResponseMonthlyBalance from the domain.

Returns:
  - A pointer to a response.ApiResponseMonthlyBalance containing the mapped data, including status, message, and data.

```go
func (s *cardResponseMapper) ToApiResponseMonthlyBalances(cards *pb.ApiResponseMonthlyBalance) *response.ApiResponseMonthlyBalance
```

##### `ToApiResponseYearlyAmounts`

ToApiResponseYearlyAmounts maps the ApiResponseYearlyAmount from the domain to the ApiResponseYearlyAmount of the api.

Args:
  - cards: A pointer to a pb.ApiResponseYearlyAmount representing the ApiResponseYearlyAmount from the domain.

Returns:
  - A pointer to a response.ApiResponseYearlyAmount containing the mapped data, including status, message, and data.

```go
func (s *cardResponseMapper) ToApiResponseYearlyAmounts(cards *pb.ApiResponseYearlyAmount) *response.ApiResponseYearlyAmount
```

##### `ToApiResponseYearlyBalances`

ToApiResponseYearlyBalances maps the ApiResponseYearlyBalance from the domain to the ApiResponseYearlyBalance of the api.

Args:
  - cards: A pointer to a pb.ApiResponseYearlyBalance representing the ApiResponseYearlyBalance from the domain.

Returns:
  - A pointer to a response.ApiResponseYearlyBalance containing the mapped data, including status, message, and data.

```go
func (s *cardResponseMapper) ToApiResponseYearlyBalances(cards *pb.ApiResponseYearlyBalance) *response.ApiResponseYearlyBalance
```

##### `ToApiResponsesCard`

ToApiResponsesCard maps the ApiResponsePaginationCard from the domain to the ApiResponsePaginationCard of the api.

Args:
  - cards: A pointer to a pb.ApiResponsePaginationCard representing the ApiResponsePaginationCard from the domain.

Returns:
  - A pointer to a response.ApiResponsePaginationCard containing the mapped data, including status, message, data, and pagination details.

```go
func (s *cardResponseMapper) ToApiResponsesCard(cards *pb.ApiResponsePaginationCard) *response.ApiResponsePaginationCard
```

##### `ToApiResponsesCardDeletedAt`

ToApiResponsesCardDeletedAt maps the ApiResponsePaginationCardDeleteAt from the domain to the
ApiResponsePaginationCardDeleteAt of the api.

Args:
  - cards: A pointer to a pb.ApiResponsePaginationCardDeleteAt representing the ApiResponsePaginationCardDeleteAt from the domain.

Returns:
  - A pointer to a response.ApiResponsePaginationCardDeleteAt containing the mapped data, including status, message, data, and pagination details.

```go
func (s *cardResponseMapper) ToApiResponsesCardDeletedAt(cards *pb.ApiResponsePaginationCardDeleteAt) *response.ApiResponsePaginationCardDeleteAt
```

##### `mapCardResponse`

mapCardResponse maps a CardResponse from the domain representation to the API response representation.

Args:
  - card: A pointer to a pb.CardResponse representing the domain CardResponse object.

Returns:
  - A pointer to a response.CardResponse containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.

```go
func (s *cardResponseMapper) mapCardResponse(card *pb.CardResponse) *response.CardResponse
```

##### `mapCardResponseDeleteAt`

mapCardResponseDeleteAt maps a CardResponseDeleteAt from the domain representation to the API response representation.

Args:
  - card: A pointer to a pb.CardResponseDeleteAt representing the domain CardResponseDeleteAt object.

Returns:
  - A pointer to a response.CardResponseDeleteAt containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardResponseMapper) mapCardResponseDeleteAt(card *pb.CardResponseDeleteAt) *response.CardResponseDeleteAt
```

##### `mapCardResponses`

mapCardResponses maps a slice of CardResponse from the domain representation to the API response representation.

Args:
  - cards: A pointer to a slice of pb.CardResponse representing the domain CardResponse objects.

Returns:
  - A pointer to a slice of response.CardResponse containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.

```go
func (s *cardResponseMapper) mapCardResponses(cards []*pb.CardResponse) []*response.CardResponse
```

##### `mapCardResponsesDeleteAt`

mapCardResponsesDeleteAt maps a slice of CardResponseDeleteAt from the domain representation to the API response representation.

Args:
  - cards: A pointer to a slice of pb.CardResponseDeleteAt representing the domain CardResponseDeleteAt objects.

Returns:
  - A pointer to a slice of response.CardResponseDeleteAt containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardResponseMapper) mapCardResponsesDeleteAt(cards []*pb.CardResponseDeleteAt) []*response.CardResponseDeleteAt
```

##### `mapDashboardCard`

mapDashboardCard maps a CardResponseDashboard from the domain representation to the API response representation.

Args:
  - dash: A pointer to a pb.CardResponseDashboard representing the domain CardResponseDashboard object.

Returns:
  - A pointer to a response.DashboardCard containing the mapped data, including total balance,
    total withdraw, total topup, total transfer, and total transaction.

```go
func (s *cardResponseMapper) mapDashboardCard(dash *pb.CardResponseDashboard) *response.DashboardCard
```

##### `mapDashboardCardCardNumber`

mapDashboardCardCardNumber maps a CardResponseDashboardCardNumber from the domain representation to the API response representation.

Args:
  - dash: A pointer to a pb.CardResponseDashboardCardNumber representing the domain CardResponseDashboardCardNumber object.

Returns:
  - A pointer to a response.DashboardCardCardNumber containing the mapped data, including total balance,
    total withdraw, total topup, total transfer send, total transfer receiver, and total transaction.

```go
func (s *cardResponseMapper) mapDashboardCardCardNumber(dash *pb.CardResponseDashboardCardNumber) *response.DashboardCardCardNumber
```

##### `mapMonthlyAmount`

mapMonthlyAmount maps a CardResponseMonthlyAmount from the domain representation to the API response representation.

Args:
  - cards: A pointer to a pb.CardResponseMonthlyAmount representing the domain CardResponseMonthlyAmount object.

Returns:
  - A pointer to a response.CardResponseMonthAmount containing the mapped data, including month and total amount.

```go
func (s *cardResponseMapper) mapMonthlyAmount(cards *pb.CardResponseMonthlyAmount) *response.CardResponseMonthAmount
```

##### `mapMonthlyAmounts`

mapMonthlyAmounts maps a slice of CardResponseMonthlyAmount from the domain representation to the API response representation.

Args:
  - cards: A slice of pointers to pb.CardResponseMonthlyAmount representing the domain CardResponseMonthlyAmount objects.

Returns:
  - A slice of pointers to response.CardResponseMonthAmount containing the mapped data, including month and total amount.

```go
func (s *cardResponseMapper) mapMonthlyAmounts(cards []*pb.CardResponseMonthlyAmount) []*response.CardResponseMonthAmount
```

##### `mapMonthlyBalance`

mapMonthlyBalance maps a CardResponseMonthlyBalance from the domain representation to the API response representation.

Args:
  - cards: A pointer to a pb.CardResponseMonthlyBalance representing the domain CardResponseMonthlyBalance object.

Returns:
  - A pointer to a response.CardResponseMonthBalance containing the mapped data, including month and total balance.

```go
func (s *cardResponseMapper) mapMonthlyBalance(cards *pb.CardResponseMonthlyBalance) *response.CardResponseMonthBalance
```

##### `mapMonthlyBalances`

mapMonthlyBalances maps a slice of CardResponseMonthlyBalance from the domain representation to the API response representation.

Args:
  - cards: A slice of pointers to pb.CardResponseMonthlyBalance representing the domain CardResponseMonthlyBalance objects.

Returns:
  - A slice of pointers to response.CardResponseMonthBalance containing the mapped data, including month and total balance.

```go
func (s *cardResponseMapper) mapMonthlyBalances(cards []*pb.CardResponseMonthlyBalance) []*response.CardResponseMonthBalance
```

##### `mapYearlyAmount`

mapYearlyAmount maps a CardResponseYearlyAmount from the domain representation to the API response representation.

Args:
  - cards: A pointer to a pb.CardResponseYearlyAmount representing the domain CardResponseYearlyAmount object.

Returns:
  - A pointer to a response.CardResponseYearAmount containing the mapped data, including year and total amount.

```go
func (s *cardResponseMapper) mapYearlyAmount(cards *pb.CardResponseYearlyAmount) *response.CardResponseYearAmount
```

##### `mapYearlyAmounts`

mapYearlyAmounts maps a slice of CardResponseYearlyAmount from the domain representation to the API response representation.

Args:
  - cards: A slice of pointers to pb.CardResponseYearlyAmount representing the domain CardResponseYearlyAmount objects.

Returns:
  - A slice of pointers to response.CardResponseYearAmount containing the mapped data, including year and total amount.

```go
func (s *cardResponseMapper) mapYearlyAmounts(cards []*pb.CardResponseYearlyAmount) []*response.CardResponseYearAmount
```

##### `mapYearlyBalance`

mapYearlyBalance maps a CardResponseYearlyBalance from the domain representation to the API response representation.

Args:
  - cards: A pointer to a pb.CardResponseYearlyBalance representing the domain CardResponseYearlyBalance object.

Returns:
  - A pointer to a response.CardResponseYearlyBalance containing the mapped data, including year and total balance.

```go
func (s *cardResponseMapper) mapYearlyBalance(cards *pb.CardResponseYearlyBalance) *response.CardResponseYearlyBalance
```

##### `mapYearlyBalances`

mapYearlyBalances maps a slice of CardResponseYearlyBalance from the domain representation to the API response representation.

Args:
  - cards: A slice of pointers to pb.CardResponseYearlyBalance representing the domain CardResponseYearlyBalance objects.

Returns:
  - A slice of pointers to response.CardResponseYearlyBalance containing the mapped data, including year and total balance.

```go
func (s *cardResponseMapper) mapYearlyBalances(cards []*pb.CardResponseYearlyBalance) []*response.CardResponseYearlyBalance
```

### `merchantDocumentResponse`

merchantDocumentResponse provides methods to map gRPC merchant document responses to HTTP API responses.

```go
type merchantDocumentResponse struct {
}
```

#### Methods

##### `ToApiResponseMerchantDocument`

ToApiResponseMerchantDocument maps a gRPC merchant document response to an HTTP API response.
It constructs an ApiResponseMerchantDocument by copying the status and message fields
and mapping the document data to a MerchantDocumentResponse.

Args:

	doc: A pointer to a pb.ApiResponseMerchantDocument containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantDocument with mapped data.

```go
func (m *merchantDocumentResponse) ToApiResponseMerchantDocument(doc *pb.ApiResponseMerchantDocument) *response.ApiResponseMerchantDocument
```

##### `ToApiResponseMerchantDocumentAll`

ToApiResponseMerchantDocumentAll maps a gRPC response containing all merchant documents
to an HTTP API response. It constructs an ApiResponseMerchantDocumentAll by copying
the status and message fields from the gRPC response.

Args:

	resp: A pointer to a pb.ApiResponseMerchantDocumentAll containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantDocumentAll with the status and message set.

```go
func (m *merchantDocumentResponse) ToApiResponseMerchantDocumentAll(resp *pb.ApiResponseMerchantDocumentAll) *response.ApiResponseMerchantDocumentAll
```

##### `ToApiResponseMerchantDocumentDeleteAt`

ToApiResponseMerchantDocumentDeleteAt maps a soft-deleted gRPC merchant document response
to an HTTP API response. It constructs an ApiResponseMerchantDocumentDelete by copying
the status and message fields from the gRPC response.

Args:

	resp: A pointer to a pb.ApiResponseMerchantDocumentDelete containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantDocumentDelete with the status and message set.

```go
func (m *merchantDocumentResponse) ToApiResponseMerchantDocumentDeleteAt(resp *pb.ApiResponseMerchantDocumentDelete) *response.ApiResponseMerchantDocumentDelete
```

##### `ToApiResponsePaginationMerchantDocument`

ToApiResponsePaginationMerchantDocument maps a paginated gRPC merchant document response
to an HTTP API response. It constructs an ApiResponsePaginationMerchantDocument by copying
the status and message fields, mapping the document data slice to a slice of MerchantDocumentResponse,
and including pagination metadata.

Args:

	docs: A pointer to a pb.ApiResponsePaginationMerchantDocument containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationMerchantDocument with mapped data and pagination info.

```go
func (m *merchantDocumentResponse) ToApiResponsePaginationMerchantDocument(docs *pb.ApiResponsePaginationMerchantDocument) *response.ApiResponsePaginationMerchantDocument
```

##### `ToApiResponsePaginationMerchantDocumentDeleteAt`

ToApiResponsePaginationMerchantDocumentDeleteAt maps a paginated gRPC response of
soft-deleted merchant documents to an HTTP API response. It constructs an
ApiResponsePaginationMerchantDocumentDeleteAt by copying the status and message fields,
mapping the document data slice to a slice of MerchantDocumentResponseDeleteAt, and including
pagination metadata.

Args:

	docs: A pointer to a pb.ApiResponsePaginationMerchantDocumentAt containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationMerchantDocumentDeleteAt with mapped data and pagination info.

```go
func (m *merchantDocumentResponse) ToApiResponsePaginationMerchantDocumentDeleteAt(docs *pb.ApiResponsePaginationMerchantDocumentAt) *response.ApiResponsePaginationMerchantDocumentDeleteAt
```

##### `ToApiResponsesMerchantDocument`

ToApiResponsesMerchantDocument maps a gRPC merchant document response slice to an HTTP API response.
It constructs an ApiResponsesMerchantDocument by copying the status and message fields
and mapping the document data slice to a slice of MerchantDocumentResponse.

Args:

	docs: A pointer to a pb.ApiResponsesMerchantDocument containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsesMerchantDocument with mapped data.

```go
func (m *merchantDocumentResponse) ToApiResponsesMerchantDocument(docs *pb.ApiResponsesMerchantDocument) *response.ApiResponsesMerchantDocument
```

##### `mapMerchantDocument`

mapMerchantDocument maps a gRPC merchant document response to an HTTP API response.
It constructs a MerchantDocumentResponse by copying the fields from the gRPC response.

Args:

	doc: A pointer to a pb.MerchantDocument containing the gRPC response data.

Returns:

	A pointer to a response.MerchantDocumentResponse with the mapped data.

```go
func (m *merchantDocumentResponse) mapMerchantDocument(doc *pb.MerchantDocument) *response.MerchantDocumentResponse
```

##### `mapMerchantDocumentDeletedAt`

mapMerchantDocumentDeletedAt maps a gRPC merchant document response for soft-deleted documents
to an HTTP API response. It constructs a MerchantDocumentResponseDeleteAt by copying
the fields from the gRPC response and handling the optional DeletedAt field.

Args:

	doc: A pointer to a pb.MerchantDocumentDeleteAt containing the gRPC response data.

Returns:

	A pointer to a response.MerchantDocumentResponseDeleteAt with the mapped data.

```go
func (m *merchantDocumentResponse) mapMerchantDocumentDeletedAt(doc *pb.MerchantDocumentDeleteAt) *response.MerchantDocumentResponseDeleteAt
```

##### `mapMerchantDocuments`

mapMerchantDocuments maps a slice of gRPC merchant document responses to a slice of HTTP API responses.
It constructs a slice of MerchantDocumentResponse by copying the fields from the gRPC responses.

Args:

	docs: A slice of pointers to pb.MerchantDocument containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantDocumentResponse with the mapped data.

```go
func (m *merchantDocumentResponse) mapMerchantDocuments(docs []*pb.MerchantDocument) []*response.MerchantDocumentResponse
```

##### `mapMerchantDocumentsDeletedAt`

mapMerchantDocumentsDeletedAt maps a slice of gRPC merchant document responses for soft-deleted documents
to a slice of HTTP API responses. It constructs a slice of MerchantDocumentResponseDeleteAt by iterating
over each MerchantDocumentDeleteAt in the input slice, mapping them using mapMerchantDocumentDeletedAt,
and handling the optional DeletedAt field.

Args:

	docs: A slice of pointers to pb.MerchantDocumentDeleteAt containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantDocumentResponseDeleteAt with the mapped data.

```go
func (m *merchantDocumentResponse) mapMerchantDocumentsDeletedAt(docs []*pb.MerchantDocumentDeleteAt) []*response.MerchantDocumentResponseDeleteAt
```

### `merchantResponse`

merchantResponse provides methods to map gRPC merchant responses to HTTP API responses.

```go
type merchantResponse struct {
}
```

#### Methods

##### `ToApiResponseMerchant`

ToApiResponseMerchant maps a gRPC merchant response to an HTTP API response. It constructs
an ApiResponseMerchant by copying the status and message fields and mapping the merchant
data to a MerchantResponse.

Args:

	merchants: A pointer to a pb.ApiResponseMerchant containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchant with mapped data.

```go
func (m *merchantResponse) ToApiResponseMerchant(merchants *pb.ApiResponseMerchant) *response.ApiResponseMerchant
```

##### `ToApiResponseMerchantAll`

ToApiResponseMerchantAll maps a gRPC response containing all merchants to an HTTP API response.
It constructs an ApiResponseMerchantAll by copying the status and message fields from the
gRPC response.

Args:

	card: A pointer to a pb.ApiResponseMerchantAll containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantAll with the status and message set.

```go
func (s *merchantResponse) ToApiResponseMerchantAll(card *pb.ApiResponseMerchantAll) *response.ApiResponseMerchantAll
```

##### `ToApiResponseMerchantDeleteAt`

ToApiResponseMerchantDeleteAt maps a gRPC response for a deleted merchant to an HTTP API response. It
constructs an ApiResponseMerchantDelete by copying the status and message fields.

Args:

	card: A pointer to a pb.ApiResponseMerchantDelete containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantDelete with mapped data.

```go
func (s *merchantResponse) ToApiResponseMerchantDeleteAt(card *pb.ApiResponseMerchantDelete) *response.ApiResponseMerchantDelete
```

##### `ToApiResponseMerchants`

ToApiResponseMerchants maps a gRPC merchant response slice to an HTTP API response. It constructs
an ApiResponsesMerchant by copying the status and message fields and mapping the merchant
data to a slice of MerchantResponse.

```go
func (m *merchantResponse) ToApiResponseMerchants(merchants *pb.ApiResponsesMerchant) *response.ApiResponsesMerchant
```

##### `ToApiResponseMerchantsTransactionResponse`

ToApiResponseMerchantsTransactionResponse maps a paginated gRPC response of merchant transactions to an HTTP API response.
It constructs an ApiResponsePaginationMerchantTransaction by copying the status and message fields,
mapping the transaction data slice to a slice of MerchantTransactionResponse, and including pagination metadata.

Args:

	merchants: A pointer to a pb.ApiResponsePaginationMerchantTransaction containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationMerchantTransaction with mapped data and pagination info.

```go
func (m *merchantResponse) ToApiResponseMerchantsTransactionResponse(merchants *pb.ApiResponsePaginationMerchantTransaction) *response.ApiResponsePaginationMerchantTransaction
```

##### `ToApiResponseMonthlyAmounts`

ToApiResponseMonthlyAmounts converts monthly financial amounts data from a gRPC response
into an HTTP API response format. It constructs an ApiResponseMerchantMonthlyAmount by
copying the status and message fields and mapping the data to a slice of
MerchantResponseMonthlyAmount.

Args:

	ms: A pointer to a pb.ApiResponseMerchantMonthlyAmount containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantMonthlyAmount with mapped data.

```go
func (m *merchantResponse) ToApiResponseMonthlyAmounts(ms *pb.ApiResponseMerchantMonthlyAmount) *response.ApiResponseMerchantMonthlyAmount
```

##### `ToApiResponseMonthlyPaymentMethods`

ToApiResponseMonthlyPaymentMethods converts monthly payment method statistics from a gRPC
response to an HTTP API response. It constructs an ApiResponseMerchantMonthlyPaymentMethod
by copying the status and message fields and mapping the data to a slice of
MerchantResponseMonthlyPaymentMethod.

Args:

	ms: A pointer to a pb.ApiResponseMerchantMonthlyPaymentMethod containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantMonthlyPaymentMethod with mapped data.

```go
func (m *merchantResponse) ToApiResponseMonthlyPaymentMethods(ms *pb.ApiResponseMerchantMonthlyPaymentMethod) *response.ApiResponseMerchantMonthlyPaymentMethod
```

##### `ToApiResponseMonthlyTotalAmounts`

ToApiResponseMonthlyTotalAmounts converts monthly total amount data from a gRPC response
into an HTTP API response format. It constructs an ApiResponseMerchantMonthlyTotalAmount by
copying the status and message fields and mapping the data to a slice of
MerchantResponseMonthlyTotalAmount.

Args:

	ms: A pointer to a pb.ApiResponseMerchantMonthlyTotalAmount containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantMonthlyTotalAmount with mapped data.

```go
func (m *merchantResponse) ToApiResponseMonthlyTotalAmounts(ms *pb.ApiResponseMerchantMonthlyTotalAmount) *response.ApiResponseMerchantMonthlyTotalAmount
```

##### `ToApiResponseYearlyAmounts`

ToApiResponseYearlyAmounts converts yearly financial amounts data from a gRPC response
into an HTTP API response format. It constructs an ApiResponseMerchantYearlyAmount by
copying the status and message fields and mapping the data to a slice of
MerchantResponseYearlyAmount.

Args:

	ms: A pointer to a pb.ApiResponseMerchantYearlyAmount containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantYearlyAmount with mapped data.

```go
func (m *merchantResponse) ToApiResponseYearlyAmounts(ms *pb.ApiResponseMerchantYearlyAmount) *response.ApiResponseMerchantYearlyAmount
```

##### `ToApiResponseYearlyPaymentMethods`

ToApiResponseYearlyPaymentMethods converts yearly payment method statistics from
a gRPC response to an HTTP API response. It constructs an ApiResponseMerchantYearlyPaymentMethod
by copying the status and message fields and mapping the data to a slice of
MerchantResponseYearlyPaymentMethod.

Args:

	ms: A pointer to a pb.ApiResponseMerchantYearlyPaymentMethod containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantYearlyPaymentMethod with mapped data.

```go
func (m *merchantResponse) ToApiResponseYearlyPaymentMethods(ms *pb.ApiResponseMerchantYearlyPaymentMethod) *response.ApiResponseMerchantYearlyPaymentMethod
```

##### `ToApiResponseYearlyTotalAmounts`

ToApiResponseYearlyTotalAmounts converts yearly total amount data from a gRPC response
into an HTTP API response format. It constructs an ApiResponseMerchantYearlyTotalAmount by
copying the status and message fields and mapping the data to a slice of
MerchantResponseYearlyTotalAmount.

Args:

	ms: A pointer to a pb.ApiResponseMerchantYearlyTotalAmount containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMerchantYearlyTotalAmount with mapped data.

```go
func (m *merchantResponse) ToApiResponseYearlyTotalAmounts(ms *pb.ApiResponseMerchantYearlyTotalAmount) *response.ApiResponseMerchantYearlyTotalAmount
```

##### `ToApiResponsesMerchant`

ToApiResponsesMerchant maps a paginated list of merchants to a paginated HTTP API response.
It constructs an ApiResponsePaginationMerchant by copying the status and message fields
and mapping the merchant data slice to a slice of MerchantResponse, and including pagination
metadata.

Args:

	merchants: A pointer to a pb.ApiResponsePaginationMerchant containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationMerchant with mapped data and pagination info.

```go
func (m *merchantResponse) ToApiResponsesMerchant(merchants *pb.ApiResponsePaginationMerchant) *response.ApiResponsePaginationMerchant
```

##### `ToApiResponsesMerchantDeleteAt`

ToApiResponsesMerchantDeleteAt maps a paginated list of soft-deleted merchants to an HTTP API response.
It constructs an ApiResponsePaginationMerchantDeleteAt by copying the status and message fields,
mapping the merchant data slice to a slice of MerchantResponseDeleteAt, and including pagination metadata.

Args:

	merchants: A pointer to a pb.ApiResponsePaginationMerchantDeleteAt containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationMerchantDeleteAt with mapped data and pagination info.

```go
func (m *merchantResponse) ToApiResponsesMerchantDeleteAt(merchants *pb.ApiResponsePaginationMerchantDeleteAt) *response.ApiResponsePaginationMerchantDeleteAt
```

##### `mapMerchantResponse`

mapMerchantResponse maps a gRPC MerchantResponse to an HTTP API response. It
constructs an ApiResponseMerchant by copying the status, message, ID, name,
status, API key, user ID, created at, and updated at fields from the gRPC
response.

Args:

	merchant: A pointer to a pb.MerchantResponse containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponse with the mapped data.

```go
func (m *merchantResponse) mapMerchantResponse(merchant *pb.MerchantResponse) *response.MerchantResponse
```

##### `mapMerchantResponseDeleteAt`

mapMerchantResponseDeleteAt maps a gRPC MerchantResponseDeleteAt to an HTTP API response. It
constructs an ApiResponseMerchantDeleteAt by copying the status, message, ID, name,
status, API key, user ID, created at, updated at, and deleted at fields from the gRPC
response.

Args:

	merchant: A pointer to a pb.MerchantResponseDeleteAt containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponseDeleteAt with the mapped data.

```go
func (m *merchantResponse) mapMerchantResponseDeleteAt(merchant *pb.MerchantResponseDeleteAt) *response.MerchantResponseDeleteAt
```

##### `mapMerchantResponses`

mapMerchantResponses maps a slice of gRPC MerchantResponse to a slice of HTTP API responses. It
constructs a slice of ApiResponseMerchant by copying the status, message, ID, name,
status, API key, user ID, created at, and updated at fields from the gRPC
response.

Args:

	r: A slice of pointers to pb.MerchantResponse containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponse with the mapped data.

```go
func (m *merchantResponse) mapMerchantResponses(r []*pb.MerchantResponse) []*response.MerchantResponse
```

##### `mapMerchantResponsesDeleteAt`

mapMerchantResponsesDeleteAt maps a slice of gRPC MerchantResponseDeleteAt to a slice of HTTP API responses.
It constructs a slice of ApiResponseMerchantDeleteAt by iterating over each MerchantResponseDeleteAt
in the input slice and mapping them using mapMerchantResponseDeleteAt. It handles the optional DeletedAt field.

Args:

	r: A slice of pointers to pb.MerchantResponseDeleteAt containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponseDeleteAt with the mapped data.

```go
func (m *merchantResponse) mapMerchantResponsesDeleteAt(r []*pb.MerchantResponseDeleteAt) []*response.MerchantResponseDeleteAt
```

##### `mapMerchantTransactionResponse`

mapMerchantTransactionResponse maps a gRPC MerchantTransactionResponse to an HTTP API response.
It constructs a MerchantTransactionResponse by copying fields such as ID, CardNumber,
Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt
from the gRPC response.

Args:

	merchant: A pointer to a pb.MerchantTransactionResponse containing the gRPC response data.

Returns:

	A pointer to a response.MerchantTransactionResponse with the mapped data.

```go
func (m *merchantResponse) mapMerchantTransactionResponse(merchant *pb.MerchantTransactionResponse) *response.MerchantTransactionResponse
```

##### `mapMerchantTransactionResponses`

mapMerchantTransactionResponses maps a slice of gRPC MerchantTransactionResponses to a slice of HTTP API responses.
It constructs a slice of MerchantTransactionResponse by copying fields such as ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
TransactionTime, CreatedAt, and UpdatedAt from the gRPC response.

Args:

	r: A slice of pointers to pb.MerchantTransactionResponse containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantTransactionResponse with the mapped data.

```go
func (m *merchantResponse) mapMerchantTransactionResponses(r []*pb.MerchantTransactionResponse) []*response.MerchantTransactionResponse
```

##### `mapResponseMonthlyAmount`

mapResponseMonthlyAmount maps a single gRPC MerchantResponseMonthlyAmount to an HTTP API response.
It constructs a MerchantResponseMonthlyAmount by copying the month and total amount from the gRPC response.

Args:

	ms: A pointer to a pb.MerchantResponseMonthlyAmount containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponseMonthlyAmount with the mapped data.

```go
func (m *merchantResponse) mapResponseMonthlyAmount(ms *pb.MerchantResponseMonthlyAmount) *response.MerchantResponseMonthlyAmount
```

##### `mapResponseMonthlyPaymentMethod`

mapResponseMonthlyPaymentMethod maps a single gRPC MerchantResponseMonthlyPaymentMethod to an HTTP API response.
It constructs a MerchantResponseMonthlyPaymentMethod by copying the month, payment method, and total amount
from the gRPC response.

Args:

	ms: A pointer to a pb.MerchantResponseMonthlyPaymentMethod containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponseMonthlyPaymentMethod with the mapped data.

```go
func (m *merchantResponse) mapResponseMonthlyPaymentMethod(ms *pb.MerchantResponseMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod
```

##### `mapResponseMonthlyTotalAmount`

mapResponseMonthlyTotalAmount maps a single gRPC MerchantResponseMonthlyTotalAmount to an HTTP API response.
It constructs a MerchantResponseMonthlyTotalAmount by copying the month, year, and total amount from the gRPC response.

Args:

	ms: A pointer to a pb.MerchantResponseMonthlyTotalAmount containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponseMonthlyTotalAmount with the mapped data.

```go
func (m *merchantResponse) mapResponseMonthlyTotalAmount(ms *pb.MerchantResponseMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount
```

##### `mapResponseYearlyAmount`

mapResponseYearlyAmount maps a single gRPC MerchantResponseYearlyAmount to an HTTP API response.
It constructs a MerchantResponseYearlyAmount by copying the year and total amount from the gRPC response.

Args:

	ms: A pointer to a pb.MerchantResponseYearlyAmount containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponseYearlyAmount with the mapped data.

```go
func (m *merchantResponse) mapResponseYearlyAmount(ms *pb.MerchantResponseYearlyAmount) *response.MerchantResponseYearlyAmount
```

##### `mapResponseYearlyPaymentMethod`

mapResponseYearlyPaymentMethod maps a single gRPC MerchantResponseYearlyPaymentMethod to an HTTP API response.
It constructs a MerchantResponseYearlyPaymentMethod by copying the year, payment method, and total amount
from the gRPC response.

Args:

	ms: A pointer to a pb.MerchantResponseYearlyPaymentMethod containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponseYearlyPaymentMethod with the mapped data.

```go
func (m *merchantResponse) mapResponseYearlyPaymentMethod(ms *pb.MerchantResponseYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod
```

##### `mapResponseYearlyTotalAmount`

mapResponseYearlyTotalAmount maps a single gRPC MerchantResponseYearlyTotalAmount to an HTTP API response.
It constructs a MerchantResponseYearlyTotalAmount by copying the year and total amount from the gRPC response.

Args:

	ms: A pointer to a pb.MerchantResponseYearlyTotalAmount containing the gRPC response data.

Returns:

	A pointer to a response.MerchantResponseYearlyTotalAmount with the mapped data.

```go
func (m *merchantResponse) mapResponseYearlyTotalAmount(ms *pb.MerchantResponseYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount
```

##### `mapResponsesMonthlyAmount`

mapResponsesMonthlyAmount maps a slice of gRPC MerchantResponseMonthlyAmount
to a slice of HTTP API responses. It constructs a slice of MerchantResponseMonthlyAmount
by copying the month and total amount from each gRPC response.

Args:

	r: A slice of pointers to pb.MerchantResponseMonthlyAmount containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponseMonthlyAmount with the mapped data.

```go
func (m *merchantResponse) mapResponsesMonthlyAmount(r []*pb.MerchantResponseMonthlyAmount) []*response.MerchantResponseMonthlyAmount
```

##### `mapResponsesMonthlyPaymentMethod`

mapResponsesMonthlyPaymentMethod maps a slice of gRPC MerchantResponseMonthlyPaymentMethods to a slice of HTTP API responses.
It constructs a slice of MerchantResponseMonthlyPaymentMethod by copying the month, payment method, and total amount
from the gRPC response.

Args:

	r: A slice of pointers to pb.MerchantResponseMonthlyPaymentMethod containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponseMonthlyPaymentMethod with the mapped data.

```go
func (m *merchantResponse) mapResponsesMonthlyPaymentMethod(r []*pb.MerchantResponseMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod
```

##### `mapResponsesMonthlyTotalAmount`

mapResponsesMonthlyTotalAmount maps a slice of gRPC MerchantResponseMonthlyTotalAmount
to a slice of HTTP API responses. It constructs a slice of MerchantResponseMonthlyTotalAmount
by iterating over each gRPC response and mapping them using mapResponseMonthlyTotalAmount.

Args:

	r: A slice of pointers to pb.MerchantResponseMonthlyTotalAmount containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponseMonthlyTotalAmount with the mapped data.

```go
func (m *merchantResponse) mapResponsesMonthlyTotalAmount(r []*pb.MerchantResponseMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount
```

##### `mapResponsesYearlyAmount`

mapResponsesYearlyAmount maps a slice of gRPC MerchantResponseYearlyAmount
to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyAmount
by copying the year and total amount from each gRPC response.

Args:

	r: A slice of pointers to pb.MerchantResponseYearlyAmount containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponseYearlyAmount with the mapped data.

```go
func (m *merchantResponse) mapResponsesYearlyAmount(r []*pb.MerchantResponseYearlyAmount) []*response.MerchantResponseYearlyAmount
```

##### `mapResponsesYearlyPaymentMethod`

mapResponsesYearlyPaymentMethod maps a slice of gRPC MerchantResponseYearlyPaymentMethod
to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyPaymentMethod
by copying the year, payment method, and total amount from each gRPC response.

Args:

	r: A slice of pointers to pb.MerchantResponseYearlyPaymentMethod containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponseYearlyPaymentMethod with the mapped data.

```go
func (m *merchantResponse) mapResponsesYearlyPaymentMethod(r []*pb.MerchantResponseYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod
```

##### `mapResponsesYearlyTotalAmount`

mapResponsesYearlyTotalAmount maps a slice of gRPC MerchantResponseYearlyTotalAmount
to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyTotalAmount
by copying the year and total amount from each gRPC response.

Args:

	r: A slice of pointers to pb.MerchantResponseYearlyTotalAmount containing the gRPC response data.

Returns:

	A slice of pointers to response.MerchantResponseYearlyTotalAmount with the mapped data.

```go
func (m *merchantResponse) mapResponsesYearlyTotalAmount(r []*pb.MerchantResponseYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount
```

### `roleResponseMapper`

roleResponseMapper maps between gRPC role-related responses and HTTP-compatible API response formats

```go
type roleResponseMapper struct {
}
```

#### Methods

##### `ToApiResponsePaginationRole`

ToApiResponsePaginationRole maps a paginated gRPC response of roles
into a paginated HTTP API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationRole containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationRole containing the mapped data, including status,
    message, data, and pagination metadata.

```go
func (s *roleResponseMapper) ToApiResponsePaginationRole(pbResponse *pb.ApiResponsePaginationRole) *response.ApiResponsePaginationRole
```

##### `ToApiResponsePaginationRoleDeleteAt`

ToApiResponsePaginationRoleDeleteAt maps a paginated gRPC response of soft-deleted roles
into a paginated HTTP API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationRoleDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationRoleDeleteAt containing the mapped data, including status,
    message, data, and pagination metadata.

```go
func (s *roleResponseMapper) ToApiResponsePaginationRoleDeleteAt(pbResponse *pb.ApiResponsePaginationRoleDeleteAt) *response.ApiResponsePaginationRoleDeleteAt
```

##### `ToApiResponseRole`

ToApiResponseRole maps a single gRPC role response into an HTTP API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponseRole containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseRole containing the mapped status, message, and data.

```go
func (s *roleResponseMapper) ToApiResponseRole(pbResponse *pb.ApiResponseRole) *response.ApiResponseRole
```

##### `ToApiResponseRoleAll`

ToApiResponseRoleAll maps a gRPC response containing all roles to an HTTP API response format.
Args:
  - pbResponse: A pointer to a pb.ApiResponseRoleAll containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseRoleAll containing the mapped data, including status and message.

```go
func (s *roleResponseMapper) ToApiResponseRoleAll(pbResponse *pb.ApiResponseRoleAll) *response.ApiResponseRoleAll
```

##### `ToApiResponseRoleDelete`

ToApiResponseRoleDelete maps a gRPC delete role response to an HTTP API response format.
Args:
  - pbResponse: A pointer to a pb.ApiResponseRoleDelete containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseRoleDelete containing the mapped status and message.

```go
func (s *roleResponseMapper) ToApiResponseRoleDelete(pbResponse *pb.ApiResponseRoleDelete) *response.ApiResponseRoleDelete
```

##### `ToApiResponsesRole`

ToApiResponsesRole maps a gRPC response containing multiple roles
into a list HTTP API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponsesRole containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsesRole containing the mapped data, including status, message,
    and data.

```go
func (s *roleResponseMapper) ToApiResponsesRole(pbResponse *pb.ApiResponsesRole) *response.ApiResponsesRole
```

##### `mapResponseRole`

mapResponseRole maps a gRPC RoleResponse to an HTTP-compatible RoleResponse.

Args:
  - role: A pointer to a pb.RoleResponse containing the gRPC role response data.

Returns:
  - A pointer to a response.RoleResponse containing the mapped data, with fields ID, Name,
    CreatedAt, and UpdatedAt extracted from the gRPC response.

```go
func (s *roleResponseMapper) mapResponseRole(role *pb.RoleResponse) *response.RoleResponse
```

##### `mapResponseRoleDeleteAt`

mapResponseRoleDeleteAt maps a gRPC RoleResponseDeleteAt to an HTTP-compatible RoleResponseDeleteAt.

Args:
  - role: A pointer to a pb.RoleResponseDeleteAt containing the gRPC role response data including deletion info.

Returns:
  - A pointer to a response.RoleResponseDeleteAt containing the mapped data, with fields ID, Name,
    CreatedAt, UpdatedAt, and DeletedAt extracted from the gRPC response. If DeletedAt is present,
    it is mapped to a string; otherwise, it is set to nil.

```go
func (s *roleResponseMapper) mapResponseRoleDeleteAt(role *pb.RoleResponseDeleteAt) *response.RoleResponseDeleteAt
```

##### `mapResponsesRole`

mapResponsesRole maps a slice of gRPC RoleResponse to a slice of HTTP-compatible RoleResponse.

Args:
  - roles: A slice of pointers to pb.RoleResponse containing the gRPC role response data.

Returns:
  - A slice of pointers to response.RoleResponse containing the mapped data, with fields ID, Name,
    CreatedAt, and UpdatedAt extracted from the gRPC response.

```go
func (s *roleResponseMapper) mapResponsesRole(roles []*pb.RoleResponse) []*response.RoleResponse
```

##### `mapResponsesRoleDeleteAt`

mapResponsesRoleDeleteAt maps a slice of gRPC RoleResponseDeleteAt to a slice of HTTP-compatible
RoleResponseDeleteAt.

Args:
  - roles: A slice of pointers to pb.RoleResponseDeleteAt containing the gRPC role response data
    including deletion info.

Returns:
  - A slice of pointers to response.RoleResponseDeleteAt containing the mapped data, with fields ID,
    Name, CreatedAt, UpdatedAt, and DeletedAt extracted from the gRPC response. If DeletedAt is
    present, it is mapped to a string; otherwise, it is set to nil.

```go
func (s *roleResponseMapper) mapResponsesRoleDeleteAt(roles []*pb.RoleResponseDeleteAt) []*response.RoleResponseDeleteAt
```

### `saldoResponse`

saldoResponse provides methods to map gRPC saldo responses to HTTP API responses.

```go
type saldoResponse struct {
}
```

#### Methods

##### `ToApiResponseMonthSaldoBalances`

ToApiResponseMonthSaldoBalances maps a gRPC response containing monthly saldo
balance statistics to an HTTP API response. It constructs an
ApiResponseMonthSaldoBalances by copying the status and message fields and mapping
the saldo statistics to a SaldoMonthBalanceResponse slice.

Args:

	pbResponse: A pointer to a pb.ApiResponseMonthSaldoBalances containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMonthSaldoBalances with mapped data.

```go
func (s *saldoResponse) ToApiResponseMonthSaldoBalances(pbResponse *pb.ApiResponseMonthSaldoBalances) *response.ApiResponseMonthSaldoBalances
```

##### `ToApiResponseMonthTotalSaldo`

ToApiResponseMonthTotalSaldo maps a gRPC response containing monthly total saldo
statistics to an HTTP API response. It constructs an ApiResponseMonthTotalSaldo by
copying the status and message fields and mapping the saldo statistics to a
SaldoMonthTotalBalanceResponse slice.

Args:

	pbResponse: A pointer to a pb.ApiResponseMonthTotalSaldo containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseMonthTotalSaldo with mapped data.

```go
func (s *saldoResponse) ToApiResponseMonthTotalSaldo(pbResponse *pb.ApiResponseMonthTotalSaldo) *response.ApiResponseMonthTotalSaldo
```

##### `ToApiResponsePaginationSaldo`

ToApiResponsePaginationSaldo maps a paginated gRPC response of saldo records to an HTTP API response.
It constructs an ApiResponsePaginationSaldo by copying the status and message fields,
mapping the saldo data slice to a slice of SaldoResponse, and including pagination metadata.

Args:

	pbResponse: A pointer to a pb.ApiResponsePaginationSaldo containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationSaldo with mapped data and pagination info.

```go
func (s *saldoResponse) ToApiResponsePaginationSaldo(pbResponse *pb.ApiResponsePaginationSaldo) *response.ApiResponsePaginationSaldo
```

##### `ToApiResponsePaginationSaldoDeleteAt`

ToApiResponsePaginationSaldoDeleteAt maps a paginated gRPC response of soft-deleted saldo records
to an HTTP API response. It constructs an ApiResponsePaginationSaldoDeleteAt by copying the
status and message fields, mapping the saldo data slice to a slice of SaldoResponseDeleteAt,
and including pagination metadata.

Args:

	pbResponse: A pointer to a pb.ApiResponsePaginationSaldoDeleteAt containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationSaldoDeleteAt with mapped data and pagination info.

```go
func (s *saldoResponse) ToApiResponsePaginationSaldoDeleteAt(pbResponse *pb.ApiResponsePaginationSaldoDeleteAt) *response.ApiResponsePaginationSaldoDeleteAt
```

##### `ToApiResponseSaldo`

ToApiResponseSaldo maps a gRPC saldo response to an HTTP API response. It constructs an ApiResponseSaldo by copying the status and message fields and mapping the saldo data to a SaldoResponse.

Args:

	pbResponse: A pointer to a pb.ApiResponseSaldo containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseSaldo with mapped data.

```go
func (s *saldoResponse) ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo
```

##### `ToApiResponseSaldoAll`

ToApiResponseSaldoAll maps a gRPC response containing all saldo records to an HTTP API response. It constructs an ApiResponseSaldoAll by copying the status and message fields from the gRPC response.

Args:

	pbResponse: A pointer to a pb.ApiResponseSaldoAll containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseSaldoAll with mapped data.

```go
func (s *saldoResponse) ToApiResponseSaldoAll(pbResponse *pb.ApiResponseSaldoAll) *response.ApiResponseSaldoAll
```

##### `ToApiResponseSaldoDelete`

ToApiResponseSaldoDelete maps a gRPC response containing a deleted saldo record to an HTTP API response.
It constructs an ApiResponseSaldoDelete by copying the status and message fields from the gRPC response.

Args:

	pbResponse: A pointer to a pb.ApiResponseSaldoDelete containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseSaldoDelete with mapped data.

```go
func (s *saldoResponse) ToApiResponseSaldoDelete(pbResponse *pb.ApiResponseSaldoDelete) *response.ApiResponseSaldoDelete
```

##### `ToApiResponseYearSaldoBalances`

ToApiResponseYearSaldoBalances maps a gRPC response containing yearly saldo
balance statistics to an HTTP API response. It constructs an
ApiResponseYearSaldoBalances by copying the status and message fields and
mapping the saldo statistics to a SaldoYearBalanceResponse slice.

Args:

	pbResponse: A pointer to a pb.ApiResponseYearSaldoBalances containing the
	gRPC response data.

Returns:

	A pointer to a response.ApiResponseYearSaldoBalances with mapped data.

```go
func (s *saldoResponse) ToApiResponseYearSaldoBalances(pbResponse *pb.ApiResponseYearSaldoBalances) *response.ApiResponseYearSaldoBalances
```

##### `ToApiResponseYearTotalSaldo`

ToApiResponseYearTotalSaldo maps a gRPC response containing yearly total saldo
statistics to an HTTP API response. It constructs an ApiResponseYearTotalSaldo by
copying the status and message fields and mapping the saldo statistics to a
SaldoYearTotalBalanceResponse slice.

Args:

	pbResponse: A pointer to a pb.ApiResponseYearTotalSaldo containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponseYearTotalSaldo with mapped data.

```go
func (s *saldoResponse) ToApiResponseYearTotalSaldo(pbResponse *pb.ApiResponseYearTotalSaldo) *response.ApiResponseYearTotalSaldo
```

##### `ToApiResponsesSaldo`

ToApiResponsesSaldo maps a gRPC response containing multiple saldo records to an HTTP API response. It constructs an ApiResponseSaldo by copying the status and message fields and mapping the saldo data slice to a slice of SaldoResponse.

Args:

	pbResponse: A pointer to a pb.ApiResponsesSaldo containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsesSaldo with mapped data.

```go
func (s *saldoResponse) ToApiResponsesSaldo(pbResponse *pb.ApiResponsesSaldo) *response.ApiResponsesSaldo
```

##### `mapResponseSaldo`

mapResponseSaldo maps a gRPC SaldoResponse to an HTTP API SaldoResponse.
It constructs a SaldoResponse by converting and copying relevant fields
such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
CreatedAt, and UpdatedAt.

Args:

	saldo: A pointer to a pb.SaldoResponse containing the gRPC response data.

Returns:

	A pointer to a response.SaldoResponse with mapped data.

```go
func (s *saldoResponse) mapResponseSaldo(saldo *pb.SaldoResponse) *response.SaldoResponse
```

##### `mapResponseSaldoDeleteAt`

mapResponseSaldoDeleteAt maps a gRPC SaldoResponseDeleteAt to an HTTP API SaldoResponseDeleteAt.
It constructs a SaldoResponseDeleteAt by converting and copying relevant fields
such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
CreatedAt, UpdatedAt, and DeletedAt.

Args:

	saldo: A pointer to a pb.SaldoResponseDeleteAt containing the gRPC response data.

Returns:

	A pointer to a response.SaldoResponseDeleteAt with mapped data.

```go
func (s *saldoResponse) mapResponseSaldoDeleteAt(saldo *pb.SaldoResponseDeleteAt) *response.SaldoResponseDeleteAt
```

##### `mapResponsesSaldo`

mapResponsesSaldo maps a slice of gRPC SaldoResponses to a slice of HTTP API SaldoResponses.
It constructs a slice of SaldoResponse by converting and copying relevant fields
such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
CreatedAt, and UpdatedAt from each gRPC SaldoResponse.

Args:

	saldos: A slice of pointers to pb.SaldoResponse containing the gRPC response data.

Returns:

	A slice of pointers to response.SaldoResponse with mapped data.

```go
func (s *saldoResponse) mapResponsesSaldo(saldos []*pb.SaldoResponse) []*response.SaldoResponse
```

##### `mapResponsesSaldoDeleteAt`

mapResponsesSaldoDeleteAt maps a slice of gRPC SaldoResponseDeleteAt to a slice of HTTP API SaldoResponseDeleteAt.
It constructs a slice of SaldoResponseDeleteAt by converting and copying relevant fields
such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
CreatedAt, UpdatedAt, and DeletedAt from each gRPC SaldoResponseDeleteAt.

Args:

	saldos: A slice of pointers to pb.SaldoResponseDeleteAt containing the gRPC response data.

Returns:

	A slice of pointers to response.SaldoResponseDeleteAt with mapped data.

```go
func (s *saldoResponse) mapResponsesSaldoDeleteAt(saldos []*pb.SaldoResponseDeleteAt) []*response.SaldoResponseDeleteAt
```

##### `mapSaldoMonthBalanceResponse`

mapSaldoMonthBalanceResponse maps a gRPC SaldoMonthBalanceResponse to an HTTP API SaldoMonthBalanceResponse.
It constructs a SaldoMonthBalanceResponse by copying the Month field and mapping the TotalBalance field to an integer.

Args:

	ss: A pointer to a pb.SaldoMonthBalanceResponse containing the gRPC response data.

Returns:

	A pointer to a response.SaldoMonthBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoMonthBalanceResponse(ss *pb.SaldoMonthBalanceResponse) *response.SaldoMonthBalanceResponse
```

##### `mapSaldoMonthBalanceResponses`

mapSaldoMonthBalanceResponses maps a slice of gRPC SaldoMonthBalanceResponses to a slice of HTTP API SaldoMonthBalanceResponses.
It constructs a slice of SaldoMonthBalanceResponse by copying the Month field and mapping the TotalBalance field to an integer from each gRPC SaldoMonthBalanceResponse.

Args:

	ss: A slice of pointers to pb.SaldoMonthBalanceResponse containing the gRPC response data.

Returns:

	A slice of pointers to response.SaldoMonthBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoMonthBalanceResponses(ss []*pb.SaldoMonthBalanceResponse) []*response.SaldoMonthBalanceResponse
```

##### `mapSaldoMonthTotalBalanceResponse`

mapSaldoMonthTotalBalanceResponse maps a gRPC SaldoMonthTotalBalanceResponse to an HTTP API SaldoMonthTotalBalanceResponse.
It constructs a SaldoMonthTotalBalanceResponse by copying the Month and Year fields and mapping the TotalBalance field to an integer.

Args:

	ss: A pointer to a pb.SaldoMonthTotalBalanceResponse containing the gRPC response data.

Returns:

	A pointer to a response.SaldoMonthTotalBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoMonthTotalBalanceResponse(ss *pb.SaldoMonthTotalBalanceResponse) *response.SaldoMonthTotalBalanceResponse
```

##### `mapSaldoMonthTotalBalanceResponses`

mapSaldoMonthTotalBalanceResponses maps a slice of gRPC SaldoMonthTotalBalanceResponses to a slice of HTTP API SaldoMonthTotalBalanceResponses.
It constructs a slice of SaldoMonthTotalBalanceResponse by copying the Month and Year fields and mapping the TotalBalance field to an integer from each gRPC SaldoMonthTotalBalanceResponse.

Args:

	ss: A slice of pointers to pb.SaldoMonthTotalBalanceResponse containing the gRPC response data.

Returns:

	A slice of pointers to response.SaldoMonthTotalBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoMonthTotalBalanceResponses(ss []*pb.SaldoMonthTotalBalanceResponse) []*response.SaldoMonthTotalBalanceResponse
```

##### `mapSaldoYearBalanceResponse`

mapSaldoYearBalanceResponse maps a gRPC SaldoYearBalanceResponse to an HTTP API SaldoYearBalanceResponse.
It constructs a SaldoYearBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer.

Args:

	ss: A pointer to a pb.SaldoYearBalanceResponse containing the gRPC response data.

Returns:

	A pointer to a response.SaldoYearBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoYearBalanceResponse(ss *pb.SaldoYearBalanceResponse) *response.SaldoYearBalanceResponse
```

##### `mapSaldoYearBalanceResponses`

mapSaldoYearBalanceResponses maps a slice of gRPC SaldoYearBalanceResponses to a slice of HTTP API SaldoYearBalanceResponses.
It constructs a slice of SaldoYearBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer from each gRPC SaldoYearBalanceResponse.

Args:

	ss: A slice of pointers to pb.SaldoYearBalanceResponse containing the gRPC response data.

Returns:

	A slice of pointers to response.SaldoYearBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoYearBalanceResponses(ss []*pb.SaldoYearBalanceResponse) []*response.SaldoYearBalanceResponse
```

##### `mapSaldoYearTotalBalanceResponse`

mapSaldoYearTotalBalanceResponse maps a gRPC SaldoYearTotalBalanceResponse to an HTTP API SaldoYearTotalBalanceResponse.
It constructs a SaldoYearTotalBalanceResponse by converting and copying the Year
field and mapping the TotalBalance field to an integer.

Args:

	ss: A pointer to a pb.SaldoYearTotalBalanceResponse containing the gRPC response data.

Returns:

	A pointer to a response.SaldoYearTotalBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoYearTotalBalanceResponse(ss *pb.SaldoYearTotalBalanceResponse) *response.SaldoYearTotalBalanceResponse
```

##### `mapSaldoYearTotalBalanceResponses`

mapSaldoYearTotalBalanceResponses maps a slice of gRPC SaldoYearTotalBalanceResponses to a slice of HTTP API SaldoYearTotalBalanceResponses.
It constructs a slice of SaldoYearTotalBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer from each gRPC SaldoYearTotalBalanceResponse.

Args:

	ss: A slice of pointers to pb.SaldoYearTotalBalanceResponse containing the gRPC response data.

Returns:

	A slice of pointers to response.SaldoYearTotalBalanceResponse with mapped data.

```go
func (s *saldoResponse) mapSaldoYearTotalBalanceResponses(ss []*pb.SaldoYearTotalBalanceResponse) []*response.SaldoYearTotalBalanceResponse
```

### `topupResponseMapper`

topupResponseMapper provides methods to map gRPC top-up responses to HTTP-compatible API responses.

```go
type topupResponseMapper struct {
}
```

#### Methods

##### `ToApiResponsePaginationTopup`

ToApiResponsePaginationTopup maps a paginated gRPC top-up response to an HTTP API response. It constructs an
ApiResponsePaginationTopup by copying the status and message fields, mapping the top-up data slice to a slice
of TopupResponse, and including pagination metadata.

Args:
  - s: A pointer to a pb.ApiResponsePaginationTopup containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationTopup with mapped data and pagination info.

```go
func (t *topupResponseMapper) ToApiResponsePaginationTopup(s *pb.ApiResponsePaginationTopup) *response.ApiResponsePaginationTopup
```

##### `ToApiResponsePaginationTopupDeleteAt`

ToApiResponsePaginationTopupDeleteAt maps a paginated gRPC response of soft-deleted top-ups
to an HTTP API response. It constructs an ApiResponsePaginationTopupDeleteAt by copying
the status and message fields, mapping the top-up data slice to a slice of TopupDeleteAtResponse,
and including pagination metadata.

Args:
  - s: A pointer to a pb.ApiResponsePaginationTopupDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationTopupDeleteAt with mapped data and pagination info.

```go
func (t *topupResponseMapper) ToApiResponsePaginationTopupDeleteAt(s *pb.ApiResponsePaginationTopupDeleteAt) *response.ApiResponsePaginationTopupDeleteAt
```

##### `ToApiResponseTopup`

ToApiResponseTopup maps a single gRPC top-up response to an HTTP API response.

Args:
  - s: A pointer to a pb.ApiResponseTopup containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopup containing the mapped data, including status, message, and a single
    mapped top-up response.

```go
func (t *topupResponseMapper) ToApiResponseTopup(s *pb.ApiResponseTopup) *response.ApiResponseTopup
```

##### `ToApiResponseTopupAll`

ToApiResponseTopupAll maps a gRPC response containing all top-up records to an HTTP API response.

Args:
  - s: A pointer to a pb.ApiResponseTopupAll containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupAll containing the mapped data, including status and message.

```go
func (t *topupResponseMapper) ToApiResponseTopupAll(s *pb.ApiResponseTopupAll) *response.ApiResponseTopupAll
```

##### `ToApiResponseTopupDelete`

ToApiResponseTopupDelete maps a single gRPC top-up delete response to an HTTP API response.

Args:
  - s: A pointer to a pb.ApiResponseTopupDelete containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupDelete containing the mapped data, including status and message.

```go
func (t *topupResponseMapper) ToApiResponseTopupDelete(s *pb.ApiResponseTopupDelete) *response.ApiResponseTopupDelete
```

##### `ToApiResponseTopupDeleteAt`

ToApiResponseTopupDeleteAt maps a single gRPC soft-deleted top-up response to an HTTP API response.

Args:
  - s: A pointer to a pb.ApiResponseTopupDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupDeleteAt containing the mapped data, including status, message, and a single
    mapped soft-deleted top-up response.

```go
func (t *topupResponseMapper) ToApiResponseTopupDeleteAt(s *pb.ApiResponseTopupDeleteAt) *response.ApiResponseTopupDeleteAt
```

##### `ToApiResponseTopupMonthAmount`

ToApiResponseTopupMonthAmount maps a gRPC response containing a month's worth of top-up amounts
to an HTTP API response. It constructs an ApiResponseTopupMonthAmount by copying the status and message fields,
mapping the TopupMonthAmount data slice to a slice of TopupMonthAmountResponse, and assigning it to the response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupMonthAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupMonthAmount with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupMonthAmount(s *pb.ApiResponseTopupMonthAmount) *response.ApiResponseTopupMonthAmount
```

##### `ToApiResponseTopupMonthMethod`

ToApiResponseTopupMonthMethod maps a gRPC response containing a month's worth of top-up statistics by payment method
to an HTTP API response. It constructs an ApiResponseTopupMonthMethod by copying the status and message fields,
mapping the TopupMonthMethod data slice to a slice of TopupMonthMethodResponse, and assigning it to the response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupMonthMethod containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupMonthMethod with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupMonthMethod(s *pb.ApiResponseTopupMonthMethod) *response.ApiResponseTopupMonthMethod
```

##### `ToApiResponseTopupMonthStatusFailed`

ToApiResponseTopupMonthStatusFailed maps a gRPC response containing a month's worth of failed top-up statistics
to an HTTP API response. It constructs an ApiResponseTopupMonthStatusFailed by copying the status and message fields,
mapping the TopupMonthStatusFailed data slice to a slice of TopupMonthStatusFailedResponse, and assigning it to the
response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupMonthStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupMonthStatusFailed with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupMonthStatusFailed(s *pb.ApiResponseTopupMonthStatusFailed) *response.ApiResponseTopupMonthStatusFailed
```

##### `ToApiResponseTopupMonthStatusSuccess`

ToApiResponseTopupMonthStatusSuccess maps a gRPC response containing a month's worth of successful top-up statistics
to an HTTP API response. It constructs an ApiResponseTopupMonthStatusSuccess by copying the status and message fields,
mapping the TopupMonthStatusSuccess data slice to a slice of TopupMonthStatusSuccessResponse, and assigning it to the
response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupMonthStatusSuccess containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupMonthStatusSuccess with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupMonthStatusSuccess(s *pb.ApiResponseTopupMonthStatusSuccess) *response.ApiResponseTopupMonthStatusSuccess
```

##### `ToApiResponseTopupYearAmount`

ToApiResponseTopupYearAmount maps a gRPC response containing a year's worth of top-up amounts
to an HTTP API response. It constructs an ApiResponseTopupYearAmount by copying the status and message fields,
mapping the TopupYearAmount data slice to a slice of TopupYearAmountResponse, and assigning it to the response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupYearAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupYearAmount with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupYearAmount(s *pb.ApiResponseTopupYearAmount) *response.ApiResponseTopupYearAmount
```

##### `ToApiResponseTopupYearMethod`

ToApiResponseTopupYearMethod maps a gRPC response containing a year's worth of top-up statistics by payment method
to an HTTP API response. It constructs an ApiResponseTopupYearMethod by copying the status and message fields,
mapping the TopupYearMethod data slice to a slice of TopupYearMethodResponse, and assigning it to the response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupYearMethod containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupYearMethod with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupYearMethod(s *pb.ApiResponseTopupYearMethod) *response.ApiResponseTopupYearMethod
```

##### `ToApiResponseTopupYearStatusFailed`

ToApiResponseTopupYearStatusFailed maps a gRPC response containing a year's worth of failed top-up statistics
to an HTTP API response. It constructs an ApiResponseTopupYearStatusFailed by copying the status and message fields,
mapping the TopupYearStatusFailed data slice to a slice of TopupYearStatusFailedResponse, and assigning it to the
response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupYearStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupYearStatusFailed with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupYearStatusFailed(s *pb.ApiResponseTopupYearStatusFailed) *response.ApiResponseTopupYearStatusFailed
```

##### `ToApiResponseTopupYearStatusSuccess`

ToApiResponseTopupYearStatusSuccess maps a gRPC response containing a year's worth of successful top-up statistics
to an HTTP API response. It constructs an ApiResponseTopupYearStatusSuccess by copying the status and message fields,
mapping the TopupYearStatusSuccess data slice to a slice of TopupYearStatusSuccessResponse, and assigning it to the
response's Data field.

Args:
  - s: A pointer to a pb.ApiResponseTopupYearStatusSuccess containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTopupYearStatusSuccess with mapped data.

```go
func (t *topupResponseMapper) ToApiResponseTopupYearStatusSuccess(s *pb.ApiResponseTopupYearStatusSuccess) *response.ApiResponseTopupYearStatusSuccess
```

##### `mapResponseTopup`

mapResponseTopup maps a single gRPC top-up response to an HTTP API response.
It constructs a TopupResponse by copying the fields from the gRPC response.

Args:
  - topup: A pointer to a pb.TopupResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TopupResponse with mapped data.

```go
func (t *topupResponseMapper) mapResponseTopup(topup *pb.TopupResponse) *response.TopupResponse
```

##### `mapResponseTopupDeleteAt`

mapResponseTopupDeleteAt maps a gRPC soft-deleted top-up response to an HTTP API response format.
It constructs a TopupResponseDeleteAt by copying the fields from the gRPC response, converting
the ID and TopupAmount fields to integers, and handling the potentially nil DeletedAt field.

Args:
  - topup: A pointer to a pb.TopupResponseDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.TopupResponseDeleteAt with the mapped data, including a non-nil DeletedAt field if present.

```go
func (t *topupResponseMapper) mapResponseTopupDeleteAt(topup *pb.TopupResponseDeleteAt) *response.TopupResponseDeleteAt
```

##### `mapResponseTopupMonthStatusFailed`

mapResponseTopupMonthStatusFailed maps a gRPC response containing a month's worth of failed top-up statistics
to an HTTP API response format. It constructs a TopupResponseMonthStatusFailed by copying the year, month,
total failed count, and total amount fields from the gRPC response to the API response.

```go
func (t *topupResponseMapper) mapResponseTopupMonthStatusFailed(s *pb.TopupMonthStatusFailedResponse) *response.TopupResponseMonthStatusFailed
```

##### `mapResponseTopupMonthStatusSuccess`

mapResponseTopupMonthStatusSuccess maps a gRPC response containing a month's worth of successful top-up statistics
to an HTTP API response format. It constructs a TopupResponseMonthStatusSuccess by copying the year, month,
total success count, and total amount fields from the gRPC response to the API response.

```go
func (t *topupResponseMapper) mapResponseTopupMonthStatusSuccess(s *pb.TopupMonthStatusSuccessResponse) *response.TopupResponseMonthStatusSuccess
```

##### `mapResponseTopupMonthlyAmount`

mapResponseTopupMonthlyAmount maps a gRPC response containing a month's worth of top-up
amount statistics to an HTTP API response format. It constructs a TopupMonthAmountResponse
by copying the month and total amount fields from the gRPC response to the API response.

```go
func (t *topupResponseMapper) mapResponseTopupMonthlyAmount(s *pb.TopupMonthAmountResponse) *response.TopupMonthAmountResponse
```

##### `mapResponseTopupMonthlyAmounts`

mapResponseTopupMonthlyAmounts maps a slice of gRPC responses containing monthly top-up amount statistics
to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
using mapResponseTopupMonthlyAmount, and returns a slice of mapped TopupMonthAmountResponse.

Args:
  - s: A slice of pointers to pb.TopupMonthAmountResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupMonthAmountResponse with the mapped data.

```go
func (t *topupResponseMapper) mapResponseTopupMonthlyAmounts(s []*pb.TopupMonthAmountResponse) []*response.TopupMonthAmountResponse
```

##### `mapResponseTopupMonthlyMethod`

mapResponseTopupMonthlyMethod maps a gRPC response containing a month's worth of top-up statistics
by payment method to an HTTP API response format. It constructs a TopupMonthMethodResponse by
copying the month, top-up method, total top-ups, and total amount fields from the gRPC response
to the API response.

```go
func (t *topupResponseMapper) mapResponseTopupMonthlyMethod(s *pb.TopupMonthMethodResponse) *response.TopupMonthMethodResponse
```

##### `mapResponseTopupMonthlyMethods`

mapResponseTopupMonthlyMethods maps a slice of gRPC responses containing monthly top-up statistics by payment method
to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
using mapResponseTopupMonthlyMethod, and returns a slice of mapped TopupMonthMethodResponse.

Args:
  - s: A slice of pointers to pb.TopupMonthMethodResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupMonthMethodResponse with the mapped data.

```go
func (t *topupResponseMapper) mapResponseTopupMonthlyMethods(s []*pb.TopupMonthMethodResponse) []*response.TopupMonthMethodResponse
```

##### `mapResponseTopupYearlyAmount`

mapResponseTopupYearlyAmount maps a gRPC response containing a year's worth of top-up amount statistics
to an HTTP API response format. It constructs a TopupYearlyAmountResponse by copying the year and
total amount fields from the gRPC response to the API response.

```go
func (t *topupResponseMapper) mapResponseTopupYearlyAmount(s *pb.TopupYearlyAmountResponse) *response.TopupYearlyAmountResponse
```

##### `mapResponseTopupYearlyAmounts`

mapResponseTopupYearlyAmounts maps a slice of gRPC responses containing yearly top-up amount statistics
to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
using mapResponseTopupYearlyAmount, and returns a slice of mapped TopupYearlyAmountResponse.

Args:
  - s: A slice of pointers to pb.TopupYearlyAmountResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupYearlyAmountResponse with the mapped data.

```go
func (t *topupResponseMapper) mapResponseTopupYearlyAmounts(s []*pb.TopupYearlyAmountResponse) []*response.TopupYearlyAmountResponse
```

##### `mapResponseTopupYearlyMethod`

mapResponseTopupYearlyMethod maps a gRPC response containing a year's worth of top-up statistics
by payment method to an HTTP API response format. It constructs a TopupYearlyMethodResponse by
copying the year, top-up method, total top-ups, and total amount fields from the gRPC response
to the API response.

Args:
  - s: A pointer to a pb.TopupYearlyMethodResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TopupYearlyMethodResponse with the mapped data.

```go
func (t *topupResponseMapper) mapResponseTopupYearlyMethod(s *pb.TopupYearlyMethodResponse) *response.TopupYearlyMethodResponse
```

##### `mapResponseTopupYearlyMethods`

mapResponseTopupYearlyMethods maps a slice of gRPC responses containing yearly top-up statistics
by payment method to a slice of HTTP API response formats. It iterates over the gRPC response slice,
mapping each individual response using mapResponseTopupYearlyMethod, and returns a slice of mapped
TopupYearlyMethodResponse.

Args:
  - s: A slice of pointers to pb.TopupYearlyMethodResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupYearlyMethodResponse with the mapped data.

```go
func (t *topupResponseMapper) mapResponseTopupYearlyMethods(s []*pb.TopupYearlyMethodResponse) []*response.TopupYearlyMethodResponse
```

##### `mapResponsesTopup`

mapResponsesTopup maps a slice of gRPC top-up responses to a slice of HTTP API responses.
It iterates over the slice of gRPC responses and maps each one to an HTTP API response
using mapResponseTopup, returning the slice of mapped responses.

Args:
  - topups: A slice of pointers to pb.TopupResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupResponse with the mapped data.

```go
func (t *topupResponseMapper) mapResponsesTopup(topups []*pb.TopupResponse) []*response.TopupResponse
```

##### `mapResponsesTopupDeleteAt`

mapResponsesTopupDeleteAt maps a slice of gRPC soft-deleted top-up responses to a slice
of HTTP API responses. It iterates over the gRPC response slice, mapping each individual
response using mapResponseTopupDeleteAt, and returns a slice of mapped TopupResponseDeleteAt.

Args:
  - topups: A slice of pointers to pb.TopupResponseDeleteAt containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupResponseDeleteAt with the mapped data.

```go
func (t *topupResponseMapper) mapResponsesTopupDeleteAt(topups []*pb.TopupResponseDeleteAt) []*response.TopupResponseDeleteAt
```

##### `mapResponsesTopupMonthStatusFailed`

mapResponsesTopupMonthStatusFailed maps a slice of gRPC responses containing a month's worth of failed
top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
each individual response using mapResponseTopupMonthStatusFailed, and returns a slice of mapped
TopupResponseMonthStatusFailed.

Args:
  - topups: A slice of pointers to pb.TopupMonthStatusFailedResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupResponseMonthStatusFailed with the mapped data.

```go
func (t *topupResponseMapper) mapResponsesTopupMonthStatusFailed(topups []*pb.TopupMonthStatusFailedResponse) []*response.TopupResponseMonthStatusFailed
```

##### `mapResponsesTopupMonthStatusSuccess`

mapResponsesTopupMonthStatusSuccess maps a slice of gRPC responses containing a month's worth of successful
top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
each individual response using mapResponseTopupMonthStatusSuccess, and returns a slice of mapped
TopupResponseMonthStatusSuccess.

Args:
  - topups: A slice of pointers to pb.TopupMonthStatusSuccessResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupResponseMonthStatusSuccess with the mapped data.

```go
func (t *topupResponseMapper) mapResponsesTopupMonthStatusSuccess(topups []*pb.TopupMonthStatusSuccessResponse) []*response.TopupResponseMonthStatusSuccess
```

##### `mapTopupResponseYearStatusFailed`

mapTopupResponseYearStatusFailed maps a gRPC response containing a year's worth
of failed top-up statistics to an HTTP API response format. It constructs a
TopupResponseYearStatusFailed by copying the year, total failed count, and
total amount fields from the gRPC response to the API response.

```go
func (t *topupResponseMapper) mapTopupResponseYearStatusFailed(s *pb.TopupYearStatusFailedResponse) *response.TopupResponseYearStatusFailed
```

##### `mapTopupResponseYearStatusSuccess`

mapTopupResponseYearStatusSuccess maps a gRPC response containing a year's worth of successful top-up statistics
to an HTTP API response format. It constructs a TopupResponseYearStatusSuccess by copying the year, total success count,
and total amount fields from the gRPC response to the API response.

```go
func (t *topupResponseMapper) mapTopupResponseYearStatusSuccess(s *pb.TopupYearStatusSuccessResponse) *response.TopupResponseYearStatusSuccess
```

##### `mapTopupResponsesYearStatusFailed`

mapTopupResponsesYearStatusFailed maps a slice of gRPC responses containing a year's worth
of failed top-up statistics to a slice of HTTP API responses. It iterates over the gRPC
response slice, mapping each individual response using mapTopupResponseYearStatusFailed,
and returns a slice of mapped TopupResponseYearStatusFailed.

Args:
  - topups: A slice of pointers to pb.TopupYearStatusFailedResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupResponseYearStatusFailed with the mapped data.

```go
func (t *topupResponseMapper) mapTopupResponsesYearStatusFailed(topups []*pb.TopupYearStatusFailedResponse) []*response.TopupResponseYearStatusFailed
```

##### `mapTopupResponsesYearStatusSuccess`

mapTopupResponsesYearStatusSuccess maps a slice of gRPC responses containing a year's worth of successful
top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
each individual response using mapTopupResponseYearStatusSuccess, and returns a slice of mapped
TopupResponseYearStatusSuccess.

Args:
  - topups: A slice of pointers to pb.TopupYearStatusSuccessResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TopupResponseYearStatusSuccess with the mapped data.

```go
func (t *topupResponseMapper) mapTopupResponsesYearStatusSuccess(topups []*pb.TopupYearStatusSuccessResponse) []*response.TopupResponseYearStatusSuccess
```

### `transactionResponseMapper`

transactionResponseMapper provides methods to map gRPC transaction responses to HTTP-compatible API responses.

```go
type transactionResponseMapper struct {
}
```

#### Methods

##### `ToApiResponsePaginationTransaction`

ToApiResponsePaginationTransaction maps a paginated gRPC response of transactions to an HTTP API response.
It constructs an ApiResponsePaginationTransaction by copying the status and message fields,
mapping the transaction data slice to a slice of TransactionResponse, and including pagination metadata.

Args:

	pbResponse: A pointer to a pb.ApiResponsePaginationTransaction containing the gRPC response data.

Returns:

	A pointer to a response.ApiResponsePaginationTransaction with mapped data and pagination info.

```go
func (m *transactionResponseMapper) ToApiResponsePaginationTransaction(pbResponse *pb.ApiResponsePaginationTransaction) *response.ApiResponsePaginationTransaction
```

##### `ToApiResponsePaginationTransactionDeleteAt`

ToApiResponsePaginationTransactionDeleteAt maps a paginated gRPC response of
soft-deleted transactions to an HTTP API response. It constructs an
ApiResponsePaginationTransactionDeleteAt by copying the status and message fields,
mapping the transaction data slice to a slice of TransactionResponseDeleteAt, and
including pagination metadata.

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationTransactionDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationTransactionDeleteAt with mapped data and pagination info.

```go
func (m *transactionResponseMapper) ToApiResponsePaginationTransactionDeleteAt(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) *response.ApiResponsePaginationTransactionDeleteAt
```

##### `ToApiResponseTransaction`

ToApiResponseTransaction maps a single gRPC transaction response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransaction containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransaction containing the mapped data, including status, message,
    and a single mapped transaction response.

```go
func (m *transactionResponseMapper) ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction
```

##### `ToApiResponseTransactionAll`

ToApiResponseTransactionAll maps a single gRPC transaction all response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionAll containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionAll containing the mapped data, including status and message.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionAll(pbResponse *pb.ApiResponseTransactionAll) *response.ApiResponseTransactionAll
```

##### `ToApiResponseTransactionDelete`

ToApiResponseTransactionDelete maps a single gRPC transaction delete response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionDelete containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionDelete containing the mapped data, including status and message.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionDelete(pbResponse *pb.ApiResponseTransactionDelete) *response.ApiResponseTransactionDelete
```

##### `ToApiResponseTransactionMonthAmount`

ToApiResponseTransactionMonthAmount maps a gRPC transaction month amount response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionMonthAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionMonthAmount containing the mapped data, including status, message,
    and a single mapped transaction month amount response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionMonthAmount(pbResponse *pb.ApiResponseTransactionMonthAmount) *response.ApiResponseTransactionMonthAmount
```

##### `ToApiResponseTransactionMonthMethod`

ToApiResponseTransactionMonthMethod maps a single gRPC transaction month method response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionMonthMethod containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionMonthMethod containing the mapped data, including status, message,
    and a single mapped transaction month method response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionMonthMethod(pbResponse *pb.ApiResponseTransactionMonthMethod) *response.ApiResponseTransactionMonthMethod
```

##### `ToApiResponseTransactionMonthStatusFailed`

ToApiResponseTransactionMonthStatusFailed maps a single gRPC transaction month status failed response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionMonthStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionMonthStatusFailed containing the mapped data, including status, message,
    and a single mapped transaction month status failed response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionMonthStatusFailed(pbResponse *pb.ApiResponseTransactionMonthStatusFailed) *response.ApiResponseTransactionMonthStatusFailed
```

##### `ToApiResponseTransactionMonthStatusSuccess`

ToApiResponseTransactionMonthStatusSuccess maps a single gRPC transaction month status response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionMonthStatusSuccess containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionMonthStatusSuccess containing the mapped data, including status, message,
    and a single mapped transaction month status response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionMonthStatusSuccess(pbResponse *pb.ApiResponseTransactionMonthStatusSuccess) *response.ApiResponseTransactionMonthStatusSuccess
```

##### `ToApiResponseTransactionYearAmount`

ToApiResponseTransactionYearAmount maps a single gRPC transaction year amount response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionYearAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionYearAmount containing the mapped data, including status, message,
    and a single mapped transaction year amount response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionYearAmount(pbResponse *pb.ApiResponseTransactionYearAmount) *response.ApiResponseTransactionYearAmount
```

##### `ToApiResponseTransactionYearMethod`

ToApiResponseTransactionYearMethod maps a single gRPC transaction year method response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionYearMethod containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionYearMethod containing the mapped data, including status, message,
    and a single mapped transaction year method response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionYearMethod(pbResponse *pb.ApiResponseTransactionYearMethod) *response.ApiResponseTransactionYearMethod
```

##### `ToApiResponseTransactionYearStatusFailed`

ToApiResponseTransactionYearStatusFailed maps a single gRPC transaction year status failed response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionYearStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionYearStatusFailed containing the mapped data, including status, message,
    and a single mapped transaction year status failed response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionYearStatusFailed(pbResponse *pb.ApiResponseTransactionYearStatusFailed) *response.ApiResponseTransactionYearStatusFailed
```

##### `ToApiResponseTransactionYearStatusSuccess`

ToApiResponseTransactionYearStatusSuccess maps a single gRPC transaction year status response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactionYearStatusSuccess containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactionYearStatusSuccess containing the mapped data, including status, message,
    and a single mapped transaction year status response.

```go
func (m *transactionResponseMapper) ToApiResponseTransactionYearStatusSuccess(pbResponse *pb.ApiResponseTransactionYearStatusSuccess) *response.ApiResponseTransactionYearStatusSuccess
```

##### `ToApiResponseTransactions`

ToApiResponseTransactions maps multiple gRPC transaction responses into an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransactions containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransactions containing the mapped data, including status, message,
    and a slice of mapped transaction responses.

```go
func (m *transactionResponseMapper) ToApiResponseTransactions(pbResponse *pb.ApiResponseTransactions) *response.ApiResponseTransactions
```

##### `ToResponsesTransactionDeleteAt`

ToResponsesTransactionDeleteAt maps a slice of gRPC transaction delete responses to a slice of HTTP API responses.

Args:
  - transactions: A slice of pointers to pb.TransactionResponseDeleteAt containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionResponseDeleteAt containing the mapped data.

```go
func (m *transactionResponseMapper) ToResponsesTransactionDeleteAt(transactions []*pb.TransactionResponseDeleteAt) []*response.TransactionResponseDeleteAt
```

##### `ToResponsesTransactionMonthStatusSuccess`

ToResponsesTransactionMonthStatusSuccess maps a slice of gRPC transaction month status success responses
to a slice of HTTP API responses.

Args:
  - transactions: A slice of pointers to pb.TransactionMonthStatusSuccessResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionResponseMonthStatusSuccess containing the mapped data.

```go
func (m *transactionResponseMapper) ToResponsesTransactionMonthStatusSuccess(transactions []*pb.TransactionMonthStatusSuccessResponse) []*response.TransactionResponseMonthStatusSuccess
```

##### `mapResponseTransaction`

mapResponseTransaction maps a single gRPC transaction response to an HTTP API response.

Args:
  - transaction: A pointer to a pb.TransactionResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionResponse containing the mapped data.

```go
func (m *transactionResponseMapper) mapResponseTransaction(transaction *pb.TransactionResponse) *response.TransactionResponse
```

##### `mapResponseTransactionDeleteAt`

mapResponseTransactionDeleteAt maps a single gRPC transaction delete response to an HTTP API response.

Args:
  - transaction: A pointer to a pb.TransactionResponseDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionResponseDeleteAt containing the mapped data.

```go
func (m *transactionResponseMapper) mapResponseTransactionDeleteAt(transaction *pb.TransactionResponseDeleteAt) *response.TransactionResponseDeleteAt
```

##### `mapResponseTransactionMonthAmount`

mapResponseTransactionMonthAmount maps a gRPC transaction month amount response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionMonthAmountResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionMonthAmountResponse containing the mapped data,
    including the month and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionMonthAmount(s *pb.TransactionMonthAmountResponse) *response.TransactionMonthAmountResponse
```

##### `mapResponseTransactionMonthAmounts`

mapResponseTransactionMonthAmounts maps a slice of gRPC transaction month amount responses
to a slice of HTTP API responses.

Args:
  - s: A slice of pointers to pb.TransactionMonthAmountResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionMonthAmountResponse containing the mapped data,
    including the month and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionMonthAmounts(s []*pb.TransactionMonthAmountResponse) []*response.TransactionMonthAmountResponse
```

##### `mapResponseTransactionMonthMethod`

mapResponseTransactionMonthMethod maps a gRPC transaction month method response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionMonthMethodResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionMonthMethodResponse containing the mapped data,
    including the month, payment method, total transactions, and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionMonthMethod(s *pb.TransactionMonthMethodResponse) *response.TransactionMonthMethodResponse
```

##### `mapResponseTransactionMonthMethods`

mapResponseTransactionMonthMethods maps a slice of gRPC transaction month method responses
to a slice of HTTP API responses.

Args:
  - s: A slice of pointers to pb.TransactionMonthMethodResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionMonthMethodResponse containing the mapped data,
    including the month, payment method, total transactions, and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionMonthMethods(s []*pb.TransactionMonthMethodResponse) []*response.TransactionMonthMethodResponse
```

##### `mapResponseTransactionMonthStatusFailed`

mapResponseTransactionMonthStatusFailed maps a gRPC transaction month status failed response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionMonthStatusFailedResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionResponseMonthStatusFailed containing the mapped data,
    including the year, month, total failed transactions, and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionMonthStatusFailed(s *pb.TransactionMonthStatusFailedResponse) *response.TransactionResponseMonthStatusFailed
```

##### `mapResponseTransactionMonthStatusSuccess`

mapResponseTransactionMonthStatusSuccess maps a gRPC transaction month status success response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionMonthStatusSuccessResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionResponseMonthStatusSuccess containing the mapped data,
    including the year, month, total successful transactions, and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionMonthStatusSuccess(s *pb.TransactionMonthStatusSuccessResponse) *response.TransactionResponseMonthStatusSuccess
```

##### `mapResponseTransactionYearMethod`

mapResponseTransactionYearMethod maps a single gRPC transaction year method response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionYearMethodResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionYearMethodResponse containing the mapped data,
    including the year, payment method, total transactions, and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionYearMethod(s *pb.TransactionYearMethodResponse) *response.TransactionYearMethodResponse
```

##### `mapResponseTransactionYearMethods`

mapResponseTransactionYearMethods maps a slice of gRPC transaction year method responses
to a slice of HTTP API responses.

Args:
  - s: A slice of pointers to pb.TransactionYearMethodResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionYearMethodResponse containing the mapped data,
    including the year, payment method, total transactions, and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionYearMethods(s []*pb.TransactionYearMethodResponse) []*response.TransactionYearMethodResponse
```

##### `mapResponseTransactionYearlyAmount`

mapResponseTransactionYearlyAmount maps a gRPC transaction yearly amount response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionYearlyAmountResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionYearlyAmountResponse containing the mapped data,
    including the year and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionYearlyAmount(s *pb.TransactionYearlyAmountResponse) *response.TransactionYearlyAmountResponse
```

##### `mapResponseTransactionYearlyAmounts`

mapResponseTransactionYearlyAmounts maps a slice of gRPC transaction yearly amount responses
to a slice of HTTP API responses.

Args:
  - s: A slice of pointers to pb.TransactionYearlyAmountResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionYearlyAmountResponse containing the mapped data,
    including the year and total amount.

```go
func (m *transactionResponseMapper) mapResponseTransactionYearlyAmounts(s []*pb.TransactionYearlyAmountResponse) []*response.TransactionYearlyAmountResponse
```

##### `mapResponsesTransaction`

mapResponsesTransaction maps a slice of gRPC transaction responses to a slice of HTTP API responses.

Args:
  - transactions: A slice of pointers to pb.TransactionResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionResponse containing the mapped data.

```go
func (m *transactionResponseMapper) mapResponsesTransaction(transactions []*pb.TransactionResponse) []*response.TransactionResponse
```

##### `mapResponsesTransactionMonthStatusFailed`

mapResponsesTransactionMonthStatusFailed maps a slice of gRPC transaction month status failed responses
to a slice of HTTP API responses.

Args:
  - transactions: A slice of pointers to pb.TransactionMonthStatusFailedResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionResponseMonthStatusFailed containing the mapped data.

```go
func (m *transactionResponseMapper) mapResponsesTransactionMonthStatusFailed(transactions []*pb.TransactionMonthStatusFailedResponse) []*response.TransactionResponseMonthStatusFailed
```

##### `mapTransactionResponseYearStatusFailed`

mapTransactionResponseYearStatusFailed maps a gRPC transaction year status failed response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionYearStatusFailedResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionResponseYearStatusFailed containing the mapped data,
    including the year, total failed transactions, and total amount.

```go
func (m *transactionResponseMapper) mapTransactionResponseYearStatusFailed(s *pb.TransactionYearStatusFailedResponse) *response.TransactionResponseYearStatusFailed
```

##### `mapTransactionResponseYearStatusSuccess`

mapTransactionResponseYearStatusSuccess maps a gRPC transaction year status success response
to an HTTP API response format.

Args:
  - s: A pointer to a pb.TransactionYearStatusSuccessResponse containing the gRPC response data.

Returns:
  - A pointer to a response.TransactionResponseYearStatusSuccess containing the mapped data,
    including the year, total successful transactions, and total amount.

```go
func (m *transactionResponseMapper) mapTransactionResponseYearStatusSuccess(s *pb.TransactionYearStatusSuccessResponse) *response.TransactionResponseYearStatusSuccess
```

##### `mapTransactionResponsesYearStatusFailed`

mapTransactionResponsesYearStatusFailed maps a slice of gRPC transaction year status failed responses
to a slice of HTTP API responses.

Args:
  - transactions: A slice of pointers to pb.TransactionYearStatusFailedResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionResponseYearStatusFailed containing the mapped data.

```go
func (m *transactionResponseMapper) mapTransactionResponsesYearStatusFailed(transactions []*pb.TransactionYearStatusFailedResponse) []*response.TransactionResponseYearStatusFailed
```

##### `mapTransactionResponsesYearStatusSuccess`

mapTransactionResponsesYearStatusSuccess maps a slice of gRPC transaction year status success responses
to a slice of HTTP API responses.

Args:
  - transactions: A slice of pointers to pb.TransactionYearStatusSuccessResponse containing the gRPC response data.

Returns:
  - A slice of pointers to response.TransactionResponseYearStatusSuccess containing the mapped data.

```go
func (m *transactionResponseMapper) mapTransactionResponsesYearStatusSuccess(transactions []*pb.TransactionYearStatusSuccessResponse) []*response.TransactionResponseYearStatusSuccess
```

### `transferResponseMapper`

```go
type transferResponseMapper struct {
}
```

#### Methods

##### `ToApiResponsePaginationTransfer`

ToApiResponsePaginationTransfer maps a pagination meta, status, message, and a list of TransferResponse
to a response.ApiResponsePaginationTransfer proto message.

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationTransfer containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationTransfer containing the mapped data.

```go
func (m *transferResponseMapper) ToApiResponsePaginationTransfer(pbResponse *pb.ApiResponsePaginationTransfer) *response.ApiResponsePaginationTransfer
```

##### `ToApiResponsePaginationTransferDeleteAt`

ToApiResponsePaginationTransferDeleteAt maps a pagination meta, status, message, and a list of TransferResponseDeleteAt
to a response.ApiResponsePaginationTransferDeleteAt proto message.

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationTransferDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationTransferDeleteAt containing the mapped data.

```go
func (m *transferResponseMapper) ToApiResponsePaginationTransferDeleteAt(pbResponse *pb.ApiResponsePaginationTransferDeleteAt) *response.ApiResponsePaginationTransferDeleteAt
```

##### `ToApiResponseTransfer`

ToApiResponseTransfer maps a single gRPC transfer response to an HTTP API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransfer containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransfer containing the mapped data, including status, message, and a single mapped transfer response.

```go
func (m *transferResponseMapper) ToApiResponseTransfer(pbResponse *pb.ApiResponseTransfer) *response.ApiResponseTransfer
```

##### `ToApiResponseTransferAll`

ToApiResponseTransferAll maps a gRPC response containing all transfer records into an HTTP API response.
It constructs an ApiResponseTransferAll by copying the status and message fields from the gRPC response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferAll containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransferAll with the mapped status and message.

```go
func (m *transferResponseMapper) ToApiResponseTransferAll(pbResponse *pb.ApiResponseTransferAll) *response.ApiResponseTransferAll
```

##### `ToApiResponseTransferDelete`

ToApiResponseTransferDelete maps a gRPC transfer delete response to an HTTP API response.
It constructs an ApiResponseTransferDelete by copying the status and message fields from
the gRPC response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferDelete containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransferDelete with the mapped status and message.

```go
func (m *transferResponseMapper) ToApiResponseTransferDelete(pbResponse *pb.ApiResponseTransferDelete) *response.ApiResponseTransferDelete
```

##### `ToApiResponseTransferMonthAmount`

ToApiResponseTransferMonthAmount maps a gRPC response containing a month's worth of transfer amounts
into an HTTP API response. It constructs an ApiResponseTransferMonthAmount by copying the status and message fields,
mapping the TransferMonthAmount data slice to a slice of TransferMonthAmountResponse, and assigning it to the response's Data field.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferMonthAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransferMonthAmount with mapped data.

```go
func (m *transferResponseMapper) ToApiResponseTransferMonthAmount(pbResponse *pb.ApiResponseTransferMonthAmount) *response.ApiResponseTransferMonthAmount
```

##### `ToApiResponseTransferMonthStatusFailed`

ToApiResponseTransferMonthStatusFailed maps a gRPC response containing a month's worth of failed transfer
statistics into an HTTP API response. It constructs an ApiResponseTransferMonthStatusFailed by copying the
status and message fields, mapping the TransferMonthStatusFailed data slice to a slice of
TransferResponseMonthStatusFailed, and assigning it to the response's Data field.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferMonthStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransferMonthStatusFailed with mapped data.

```go
func (m *transferResponseMapper) ToApiResponseTransferMonthStatusFailed(pbResponse *pb.ApiResponseTransferMonthStatusFailed) *response.ApiResponseTransferMonthStatusFailed
```

##### `ToApiResponseTransferMonthStatusSuccess`

ToApiResponseTransferMonthStatusSuccess converts a gRPC response containing a month's worth
of successful transfer statistics into an HTTP API response. It constructs an
ApiResponseTransferMonthStatusSuccess by copying the status and message fields, mapping the
TransferMonthStatusSuccess data slice to a slice of TransferResponseMonthStatusSuccess, and
assigning it to the response's Data field.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferMonthStatusSuccess containing the gRPC
    response data.

Returns:
  - A pointer to a response.ApiResponseTransferMonthStatusSuccess with mapped data.

```go
func (m *transferResponseMapper) ToApiResponseTransferMonthStatusSuccess(pbResponse *pb.ApiResponseTransferMonthStatusSuccess) *response.ApiResponseTransferMonthStatusSuccess
```

##### `ToApiResponseTransferYearAmount`

ToApiResponseTransferYearAmount maps a gRPC response containing a year's worth of transfer amounts
into an HTTP API response. It constructs an ApiResponseTransferYearAmount by copying the status and message fields,
mapping the TransferYearAmount data slice to a slice of TransferYearAmountResponse, and assigning it to the response's Data field.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferYearAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransferYearAmount with mapped data.

```go
func (m *transferResponseMapper) ToApiResponseTransferYearAmount(pbResponse *pb.ApiResponseTransferYearAmount) *response.ApiResponseTransferYearAmount
```

##### `ToApiResponseTransferYearStatusFailed`

ToApiResponseTransferYearStatusFailed maps a gRPC response containing a year's worth of failed transfer
statistics into an HTTP API response. It constructs an ApiResponseTransferYearStatusFailed by copying the
status and message fields, mapping the TransferYearStatusFailed data slice to a slice of
TransferResponseYearStatusFailed, and assigning it to the response's Data field.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferYearStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransferYearStatusFailed with mapped data.

```go
func (m *transferResponseMapper) ToApiResponseTransferYearStatusFailed(pbResponse *pb.ApiResponseTransferYearStatusFailed) *response.ApiResponseTransferYearStatusFailed
```

##### `ToApiResponseTransferYearStatusSuccess`

ToApiResponseTransferYearStatusSuccess maps a gRPC response containing a year's worth of successful transfer
statistics to an HTTP API response. It constructs an ApiResponseTransferYearStatusSuccess by copying the
status and message fields, mapping the TransferYearStatusSuccess data slice to a slice of
TransferResponseYearStatusSuccess, and assigning it to the response's Data field.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransferYearStatusSuccess containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransferYearStatusSuccess with mapped data.

```go
func (m *transferResponseMapper) ToApiResponseTransferYearStatusSuccess(pbResponse *pb.ApiResponseTransferYearStatusSuccess) *response.ApiResponseTransferYearStatusSuccess
```

##### `ToApiResponseTransfers`

ToApiResponseTransfers maps a gRPC response containing multiple transfer records into an HTTP API response.
It constructs an ApiResponseTransfers by copying the status and message fields, mapping the Transfer data slice to a slice of
TransferResponse, and assigning it to the response's Data field.

Args:
  - pbResponse: A pointer to a pb.ApiResponseTransfers containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseTransfers with mapped data.

```go
func (m *transferResponseMapper) ToApiResponseTransfers(pbResponse *pb.ApiResponseTransfers) *response.ApiResponseTransfers
```

##### `mapResponseTransfer`

mapResponseTransfer maps a TransferResponse from a protobuf message to a response.TransferResponse.

Args:
  - transfer: A pointer to a pb.TransferResponse containing the data to be mapped.

Returns:
  - A pointer to a response.TransferResponse containing the mapped data.

```go
func (t *transferResponseMapper) mapResponseTransfer(transfer *pb.TransferResponse) *response.TransferResponse
```

##### `mapResponseTransferDeleteAt`

mapResponseTransferDeleteAt maps a TransferResponseDeleteAt from a protobuf message to a response.TransferResponseDeleteAt.

Args:
  - transfer: A pointer to a pb.TransferResponseDeleteAt containing the data to be mapped.

Returns:
  - A pointer to a response.TransferResponseDeleteAt containing the mapped data.

```go
func (t *transferResponseMapper) mapResponseTransferDeleteAt(transfer *pb.TransferResponseDeleteAt) *response.TransferResponseDeleteAt
```

##### `mapResponseTransferMonthAmount`

mapResponseTransferMonthAmount maps a TransferMonthAmountResponse protobuf message to a
response.TransferMonthAmountResponse.

Args:
  - s: A pointer to a pb.TransferMonthAmountResponse containing the data to be mapped.

Returns:
  - A pointer to a response.TransferMonthAmountResponse containing the mapped data, including Month,
    and TotalAmount.

```go
func (m *transferResponseMapper) mapResponseTransferMonthAmount(s *pb.TransferMonthAmountResponse) *response.TransferMonthAmountResponse
```

##### `mapResponseTransferMonthAmounts`

mapResponseTransferMonthAmounts maps a slice of TransferMonthAmountResponse
protobuf messages to a slice of TransferMonthAmountResponse.

It iterates over the input slice, converting each TransferMonthAmountResponse
to its corresponding response representation using the mapResponseTransferMonthAmount method.

Args:
  - s: A slice of pointers to pb.TransferMonthAmountResponse containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferMonthAmountResponse containing the mapped data.

```go
func (m *transferResponseMapper) mapResponseTransferMonthAmounts(s []*pb.TransferMonthAmountResponse) []*response.TransferMonthAmountResponse
```

##### `mapResponseTransferMonthStatusFailed`

mapResponseTransferMonthStatusFailed maps a TransferMonthStatusFailedResponse protobuf message to a
response.TransferResponseMonthStatusFailed.

Args:
  - s: A pointer to a pb.TransferMonthStatusFailedResponse containing the data to be mapped.

Returns:
  - A pointer to a response.TransferResponseMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transferResponseMapper) mapResponseTransferMonthStatusFailed(s *pb.TransferMonthStatusFailedResponse) *response.TransferResponseMonthStatusFailed
```

##### `mapResponseTransferMonthStatusSuccess`

mapResponseTransferMonthStatusSuccess maps a TransferMonthStatusSuccessResponse protobuf message to a
response.TransferResponseMonthStatusSuccess.

Args:
  - s: A pointer to a pb.TransferMonthStatusSuccessResponse containing the data to be mapped.

Returns:
  - A pointer to a response.TransferResponseMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transferResponseMapper) mapResponseTransferMonthStatusSuccess(s *pb.TransferMonthStatusSuccessResponse) *response.TransferResponseMonthStatusSuccess
```

##### `mapResponseTransferYearAmount`

mapResponseTransferYearAmount maps a TransferYearAmountResponse protobuf message to a
response.TransferYearAmountResponse.

Args:
  - s: A pointer to a pb.TransferYearAmountResponse containing the data to be mapped.

Returns:
  - A pointer to a response.TransferYearAmountResponse containing the mapped data, including Year
    and TotalAmount.

```go
func (m *transferResponseMapper) mapResponseTransferYearAmount(s *pb.TransferYearAmountResponse) *response.TransferYearAmountResponse
```

##### `mapResponseTransferYearAmounts`

mapResponseTransferYearAmounts maps a slice of TransferYearAmountResponse
protobuf messages to a slice of TransferYearAmountResponse.

It iterates over the input slice, converting each TransferYearAmountResponse
to its corresponding response representation using the mapResponseTransferYearAmount method.

Args:
  - s: A slice of pointers to pb.TransferYearAmountResponse containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferYearAmountResponse containing the mapped data.

```go
func (m *transferResponseMapper) mapResponseTransferYearAmounts(s []*pb.TransferYearAmountResponse) []*response.TransferYearAmountResponse
```

##### `mapResponsesTransfer`

mapResponsesTransfer maps a slice of TransferResponse from protobuf messages to a slice of response.TransferResponse.

Args:
  - transfers: A slice of pointers to pb.TransferResponse containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferResponse containing the mapped data.

```go
func (t *transferResponseMapper) mapResponsesTransfer(transfers []*pb.TransferResponse) []*response.TransferResponse
```

##### `mapResponsesTransferDeleteAt`

mapResponsesTransferDeleteAt maps a slice of TransferResponseDeleteAt from protobuf messages to a slice of response.TransferResponseDeleteAt.

It iterates over the input slice, converting each TransferResponseDeleteAt to its corresponding response.TransferResponseDeleteAt
using the mapResponseTransferDeleteAt method.

Args:
  - transfers: A slice of pointers to pb.TransferResponseDeleteAt containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferResponseDeleteAt containing the mapped data.

```go
func (t *transferResponseMapper) mapResponsesTransferDeleteAt(transfers []*pb.TransferResponseDeleteAt) []*response.TransferResponseDeleteAt
```

##### `mapResponsesTransferMonthStatusFailed`

mapResponsesTransferMonthStatusFailed maps a slice of TransferMonthStatusFailedResponse
protobuf messages to a slice of TransferResponseMonthStatusFailed.

It iterates over the input slice, converting each TransferMonthStatusFailedResponse
to its corresponding response representation using the mapResponseTransferMonthStatusFailed method.

Args:
  - Transfers: A slice of pointers to pb.TransferMonthStatusFailedResponse containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferResponseMonthStatusFailed containing the mapped data.

```go
func (t *transferResponseMapper) mapResponsesTransferMonthStatusFailed(Transfers []*pb.TransferMonthStatusFailedResponse) []*response.TransferResponseMonthStatusFailed
```

##### `mapResponsesTransferMonthStatusSuccess`

mapResponsesTransferMonthStatusSuccess maps a slice of TransferMonthStatusSuccessResponse
protobuf messages to a slice of TransferResponseMonthStatusSuccess.

It iterates over the input slice, converting each TransferMonthStatusSuccessResponse
to its corresponding response representation using the mapResponseTransferMonthStatusSuccess method.

Args:
  - Transfers: A slice of pointers to pb.TransferMonthStatusSuccessResponse containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferResponseMonthStatusSuccess containing the mapped data.

```go
func (t *transferResponseMapper) mapResponsesTransferMonthStatusSuccess(Transfers []*pb.TransferMonthStatusSuccessResponse) []*response.TransferResponseMonthStatusSuccess
```

##### `mapTransferResponseYearStatusFailed`

mapTransferResponseYearStatusFailed maps a TransferYearStatusFailedResponse protobuf message to a
response.TransferResponseYearStatusFailed.

Args:
  - s: A pointer to a pb.TransferYearStatusFailedResponse containing the data to be mapped.

Returns:
  - A pointer to a response.TransferResponseYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transferResponseMapper) mapTransferResponseYearStatusFailed(s *pb.TransferYearStatusFailedResponse) *response.TransferResponseYearStatusFailed
```

##### `mapTransferResponseYearStatusSuccess`

mapTransferResponseYearStatusSuccess maps a TransferYearStatusSuccessResponse protobuf message to a
response.TransferResponseYearStatusSuccess.

Args:
  - s: A pointer to a pb.TransferYearStatusSuccessResponse containing the data to be mapped.

Returns:
  - A pointer to a response.TransferResponseYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transferResponseMapper) mapTransferResponseYearStatusSuccess(s *pb.TransferYearStatusSuccessResponse) *response.TransferResponseYearStatusSuccess
```

##### `mapTransferResponsesYearStatusFailed`

mapTransferResponsesYearStatusFailed maps a slice of TransferYearStatusFailedResponse
protobuf messages to a slice of TransferResponseYearStatusFailed.

It iterates over the input slice, converting each TransferYearStatusFailedResponse
to its corresponding response representation using the mapTransferResponseYearStatusFailed method.

Args:
  - Transfers: A slice of pointers to pb.TransferYearStatusFailedResponse containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferResponseYearStatusFailed containing the mapped data.

```go
func (t *transferResponseMapper) mapTransferResponsesYearStatusFailed(Transfers []*pb.TransferYearStatusFailedResponse) []*response.TransferResponseYearStatusFailed
```

##### `mapTransferResponsesYearStatusSuccess`

mapTransferResponsesYearStatusSuccess maps a slice of TransferYearStatusSuccessResponse
protobuf messages to a slice of TransferResponseYearStatusSuccess.

It iterates over the input slice, converting each TransferYearStatusSuccessResponse
to its corresponding response representation using the mapTransferResponseYearStatusSuccess method.

Args:
  - Transfers: A slice of pointers to pb.TransferYearStatusSuccessResponse containing the data to be mapped.

Returns:
  - A slice of pointers to response.TransferResponseYearStatusSuccess containing the mapped data.

```go
func (t *transferResponseMapper) mapTransferResponsesYearStatusSuccess(Transfers []*pb.TransferYearStatusSuccessResponse) []*response.TransferResponseYearStatusSuccess
```

### `userResponseMapper`

userResponseMapper provides methods to map gRPC user responses to HTTP API responses.

```go
type userResponseMapper struct {
}
```

#### Methods

##### `ToApiResponsePaginationUser`

ToApiResponsePaginationUser maps a pagination meta, status, message, and a list of UserResponse

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationUser containing the user data.

Returns:
  - A pointer to a response.ApiResponsePaginationUser containing the mapped data.

```go
func (u *userResponseMapper) ToApiResponsePaginationUser(pbResponse *pb.ApiResponsePaginationUser) *response.ApiResponsePaginationUser
```

##### `ToApiResponsePaginationUserDeleteAt`

ToApiResponsePaginationUserDeleteAt maps a pagination meta, status, message, and a list of UserResponseDeleteAt

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationUserDeleteAt containing the user data.

Returns:
  - A pointer to a response.ApiResponsePaginationUserDeleteAt containing the mapped data.

```go
func (u *userResponseMapper) ToApiResponsePaginationUserDeleteAt(pbResponse *pb.ApiResponsePaginationUserDeleteAt) *response.ApiResponsePaginationUserDeleteAt
```

##### `ToApiResponseUser`

ToApiResponseUser converts a single user response into an API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseUser containing the user data.

Returns:
  - A pointer to a response.ApiResponseUser containing the mapped user data, including
    Status, Message, and Data.

```go
func (u *userResponseMapper) ToApiResponseUser(pbResponse *pb.ApiResponseUser) *response.ApiResponseUser
```

##### `ToApiResponseUserAll`

ToApiResponseUserAll maps a pb.ApiResponseUserAll to a response.ApiResponseUserAll.

Args:
  - pbResponse: A pointer to a pb.ApiResponseUserAll containing the status and message.

Returns:
  - A pointer to a response.ApiResponseUserAll containing the mapped status and message.

```go
func (u *userResponseMapper) ToApiResponseUserAll(pbResponse *pb.ApiResponseUserAll) *response.ApiResponseUserAll
```

##### `ToApiResponseUserDelete`

ToApiResponseUserDelete maps a permanently deleted user response to an API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseUserDelete containing the user data.

Returns:
  - A pointer to a response.ApiResponseUserDelete containing the mapped status and message.

```go
func (u *userResponseMapper) ToApiResponseUserDelete(pbResponse *pb.ApiResponseUserDelete) *response.ApiResponseUserDelete
```

##### `ToApiResponseUserDeleteAt`

ToApiResponseUserDeleteAt maps a soft-deleted user response into an API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponseUserDeleteAt containing the user data.

Returns:
  - A pointer to a response.ApiResponseUserDeleteAt containing the mapped user data, including
    Status, Message, and Data.

```go
func (u *userResponseMapper) ToApiResponseUserDeleteAt(pbResponse *pb.ApiResponseUserDeleteAt) *response.ApiResponseUserDeleteAt
```

##### `ToApiResponsesUser`

ToApiResponsesUser converts multiple user responses into a grouped API response.

Args:
  - pbResponse: A pointer to a pb.ApiResponsesUser containing the user data.

Returns:
  - A pointer to a response.ApiResponsesUser containing the mapped user data, including
    Status, Message, and Data.

```go
func (u *userResponseMapper) ToApiResponsesUser(pbResponse *pb.ApiResponsesUser) *response.ApiResponsesUser
```

##### `ToResponseUser`

ToResponseUser maps a protobuf UserResponse to a domain UserResponse.

Args:
  - user: A pointer to a pb.UserResponse containing the user data.

Returns:
  - A pointer to a response.UserResponse containing the mapped user data, including
    ID, FirstName, LastName, Email, CreatedAt, and UpdatedAt.

```go
func (u *userResponseMapper) ToResponseUser(user *pb.UserResponse) *response.UserResponse
```

##### `ToResponseUserDeleteAt`

ToResponseUserDeleteAt maps a protobuf UserResponseDeleteAt to a domain UserResponseDeleteAt.

Args:
  - user: A pointer to a pb.UserResponseDeleteAt containing the user data.

Returns:
  - A pointer to a response.UserResponseDeleteAt containing the mapped user data, including
    ID, FirstName, LastName, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.

```go
func (u *userResponseMapper) ToResponseUserDeleteAt(user *pb.UserResponseDeleteAt) *response.UserResponseDeleteAt
```

##### `ToResponsesUser`

ToResponsesUser maps a slice of pb.UserResponse to a slice of response.UserResponse.

Args:
  - users: A slice of pointers to pb.UserResponse to be mapped.

Returns:
  - A slice of pointers to response.UserResponse containing the mapped user data for each user, including
    ID, FirstName, LastName, Email, CreatedAt, and UpdatedAt.

```go
func (u *userResponseMapper) ToResponsesUser(users []*pb.UserResponse) []*response.UserResponse
```

##### `ToResponsesUserDeleteAt`

ToResponsesUserDeleteAt maps a slice of protobuf UserResponseDeleteAt to a slice of domain UserResponseDeleteAt.

Args:
  - users: A slice of pointers to pb.UserResponseDeleteAt to be mapped.

Returns:
  - A slice of pointers to response.UserResponseDeleteAt containing the mapped user data for each user, including
    ID, FirstName, LastName, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.

```go
func (u *userResponseMapper) ToResponsesUserDeleteAt(users []*pb.UserResponseDeleteAt) []*response.UserResponseDeleteAt
```

### `withdrawResponseMapper`

withdrawResponseMapper provides methods to map gRPC withdraw responses to HTTP API responses

```go
type withdrawResponseMapper struct {
}
```

#### Methods

##### `ToApiResponsePaginationWithdraw`

ToApiResponsePaginationWithdraw maps a pagination meta, status, message, and a list of WithdrawResponse
to a response.ApiResponsePaginationWithdraw proto message.

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationWithdraw containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationWithdraw containing the mapped data.

```go
func (m *withdrawResponseMapper) ToApiResponsePaginationWithdraw(pbResponse *pb.ApiResponsePaginationWithdraw) *response.ApiResponsePaginationWithdraw
```

##### `ToApiResponsePaginationWithdrawDeleteAt`

ToApiResponsePaginationWithdrawDeleteAt maps a pagination meta, status, message, and a list of WithdrawResponseDeleteAt
to a response.ApiResponsePaginationWithdrawDeleteAt proto message.

Args:
  - pbResponse: A pointer to a pb.ApiResponsePaginationWithdrawDeleteAt containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponsePaginationWithdrawDeleteAt containing the mapped data.

```go
func (m *withdrawResponseMapper) ToApiResponsePaginationWithdrawDeleteAt(pbResponse *pb.ApiResponsePaginationWithdrawDeleteAt) *response.ApiResponsePaginationWithdrawDeleteAt
```

##### `ToApiResponseWithdraw`

ToApiResponseWithdraw maps a single withdraw gRPC response to an API response.

Args:
  - pbResponse: The gRPC response that needs to be converted.

Returns:
  - A pointer to an ApiResponseWithdraw containing the mapped data.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdraw(pbResponse *pb.ApiResponseWithdraw) *response.ApiResponseWithdraw
```

##### `ToApiResponseWithdrawAll`

ToApiResponseWithdrawAll maps a gRPC response containing all withdraw records to an API response.

Args:
  - pbResponse: The gRPC response that needs to be converted.

Returns:
  - A pointer to an ApiResponseWithdrawAll containing the mapped data.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawAll(pbResponse *pb.ApiResponseWithdrawAll) *response.ApiResponseWithdrawAll
```

##### `ToApiResponseWithdrawDelete`

ToApiResponseWithdrawDelete maps a gRPC response indicating a withdraw has been deleted to an API response.

Args:
  - pbResponse: The gRPC response that needs to be converted.

Returns:
  - A pointer to an ApiResponseWithdrawDelete containing the mapped data.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawDelete(pbResponse *pb.ApiResponseWithdrawDelete) *response.ApiResponseWithdrawDelete
```

##### `ToApiResponseWithdrawMonthAmount`

ToApiResponseWithdrawMonthAmount converts a gRPC response containing monthly withdraw amounts
into an API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseWithdrawMonthAmount containing the mapped data, including status, message,
    and detailed information about the monthly withdraw amounts.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawMonthAmount(pbResponse *pb.ApiResponseWithdrawMonthAmount) *response.ApiResponseWithdrawMonthAmount
```

##### `ToApiResponseWithdrawMonthStatusFailed`

ToApiResponseWithdrawMonthStatusFailed maps a gRPC response containing monthly failed withdraw statistics
into an API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseWithdrawMonthStatusFailed containing the mapped data including status, message,
    and detailed information about the monthly failed withdraw statistics.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawMonthStatusFailed(pbResponse *pb.ApiResponseWithdrawMonthStatusFailed) *response.ApiResponseWithdrawMonthStatusFailed
```

##### `ToApiResponseWithdrawMonthStatusSuccess`

ToApiResponseWithdrawMonthStatusSuccess converts a gRPC response containing monthly successful withdraw statistics
into an API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthStatusSuccess containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseWithdrawMonthStatusSuccess containing the mapped data including status, message,
    and detailed information about the monthly successful withdraw statistics.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawMonthStatusSuccess(pbResponse *pb.ApiResponseWithdrawMonthStatusSuccess) *response.ApiResponseWithdrawMonthStatusSuccess
```

##### `ToApiResponseWithdrawYearAmount`

ToApiResponseWithdrawYearAmount maps a gRPC response containing yearly withdraw amounts
into an API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponseWithdrawYearAmount containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseWithdrawYearAmount containing the mapped data, including status, message,
    and detailed information about the yearly withdraw amounts.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawYearAmount(pbResponse *pb.ApiResponseWithdrawYearAmount) *response.ApiResponseWithdrawYearAmount
```

##### `ToApiResponseWithdrawYearStatusFailed`

ToApiResponseWithdrawYearStatusFailed maps a gRPC response containing yearly failed withdraw statistics
into an API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponseWithdrawYearStatusFailed containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseWithdrawYearStatusFailed containing the mapped data including status, message,
    and detailed information about the yearly failed withdraw statistics.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawYearStatusFailed(pbResponse *pb.ApiResponseWithdrawYearStatusFailed) *response.ApiResponseWithdrawYearStatusFailed
```

##### `ToApiResponseWithdrawYearStatusSuccess`

ToApiResponseWithdrawYearStatusSuccess maps a gRPC response containing yearly successful withdraw statistics
to an HTTP API response format.

Args:
  - pbResponse: A pointer to a pb.ApiResponseWithdrawYearStatusSuccess containing the gRPC response data.

Returns:
  - A pointer to a response.ApiResponseWithdrawYearStatusSuccess containing the mapped data, including status, message,
    and detailed information about the yearly successful withdraw statistics.

```go
func (m *withdrawResponseMapper) ToApiResponseWithdrawYearStatusSuccess(pbResponse *pb.ApiResponseWithdrawYearStatusSuccess) *response.ApiResponseWithdrawYearStatusSuccess
```

##### `ToApiResponsesWithdraw`

ToApiResponsesWithdraw maps a list of withdraw gRPC responses to an API response.

Args:
  - pbResponse: The gRPC response that needs to be converted.

Returns:
  - A pointer to an ApiResponsesWithdraw containing the mapped data.

```go
func (m *withdrawResponseMapper) ToApiResponsesWithdraw(pbResponse *pb.ApiResponsesWithdraw) *response.ApiResponsesWithdraw
```

##### `mapResponseWithdrawMonthStatusFailed`

mapResponseWithdrawMonthStatusFailed maps a single WithdrawMonthStatusFailedResponse gRPC message
to an API response WithdrawResponseMonthStatusFailed object.

Args:
  - s: The WithdrawMonthStatusFailedResponse gRPC message that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseMonthStatusFailed containing the mapped data, including
    fields like Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawResponseMapper) mapResponseWithdrawMonthStatusFailed(s *pb.WithdrawMonthStatusFailedResponse) *response.WithdrawResponseMonthStatusFailed
```

##### `mapResponseWithdrawMonthStatusSuccess`

```go
func (t *withdrawResponseMapper) mapResponseWithdrawMonthStatusSuccess(s *pb.WithdrawMonthStatusSuccessResponse) *response.WithdrawResponseMonthStatusSuccess
```

##### `mapResponseWithdrawMonthlyAmount`

mapResponseWithdrawMonthlyAmount maps a single WithdrawMonthlyAmountResponse gRPC message
to an API response WithdrawMonthlyAmountResponse object.

Args:
  - s: The WithdrawMonthlyAmountResponse gRPC message that needs to be converted.

Returns:
  - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data, including
    fields like Month and TotalAmount.

```go
func (m *withdrawResponseMapper) mapResponseWithdrawMonthlyAmount(s *pb.WithdrawMonthlyAmountResponse) *response.WithdrawMonthlyAmountResponse
```

##### `mapResponseWithdrawMonthlyAmounts`

mapResponseWithdrawMonthlyAmounts maps a slice of gRPC WithdrawMonthlyAmountResponse messages
to a slice of API response WithdrawMonthlyAmountResponse objects.

Args:
  - s: A slice of pointers to pb.WithdrawMonthlyAmountResponse messages that need to be converted.

Returns:
  - A slice of pointers to response.WithdrawMonthlyAmountResponse objects containing the mapped data,
    including fields like Month and TotalAmount.

```go
func (m *withdrawResponseMapper) mapResponseWithdrawMonthlyAmounts(s []*pb.WithdrawMonthlyAmountResponse) []*response.WithdrawMonthlyAmountResponse
```

##### `mapResponseWithdrawYearlyAmount`

mapResponseWithdrawYearlyAmount maps a single WithdrawYearlyAmountResponse gRPC message
to an API response WithdrawYearlyAmountResponse object.

Args:
  - s: The WithdrawYearlyAmountResponse gRPC message that needs to be converted.

Returns:
  - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, including
    fields like Year and TotalAmount.

```go
func (m *withdrawResponseMapper) mapResponseWithdrawYearlyAmount(s *pb.WithdrawYearlyAmountResponse) *response.WithdrawYearlyAmountResponse
```

##### `mapResponseWithdrawYearlyAmounts`

mapResponseWithdrawYearlyAmounts maps a slice of gRPC WithdrawYearlyAmountResponse messages
to a slice of API response WithdrawYearlyAmountResponse objects.

Args:
  - s: A slice of pointers to pb.WithdrawYearlyAmountResponse messages that need to be converted.

Returns:
  - A slice of pointers to response.WithdrawYearlyAmountResponse objects containing the mapped data,
    including fields like Year and TotalAmount.

```go
func (m *withdrawResponseMapper) mapResponseWithdrawYearlyAmounts(s []*pb.WithdrawYearlyAmountResponse) []*response.WithdrawYearlyAmountResponse
```

##### `mapResponseWithdrawal`

mapResponseWithdrawal maps a single withdraw gRPC response to an API response.

Args:
  - withdraw: The gRPC response that needs to be converted.

Returns:
  - A pointer to a WithdrawResponse containing the mapped data.

```go
func (w *withdrawResponseMapper) mapResponseWithdrawal(withdraw *pb.WithdrawResponse) *response.WithdrawResponse
```

##### `mapResponseWithdrawalDeleteAt`

mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to an API response.

Args:
  - withdraw: The WithdrawResponseDeleteAt that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including id, withdraw_no, card_number,
    withdraw_amount, withdraw_time, created_at, updated_at, and deleted_at.

```go
func (w *withdrawResponseMapper) mapResponseWithdrawalDeleteAt(withdraw *pb.WithdrawResponseDeleteAt) *response.WithdrawResponseDeleteAt
```

##### `mapResponsesWithdrawMonthStatusFailed`

mapResponsesWithdrawMonthStatusFailed maps a slice of WithdrawMonthStatusFailedResponse gRPC messages
to a slice of API response WithdrawResponseMonthStatusFailed objects.

Args:
  - Withdraws: A slice of pointers to pb.WithdrawMonthStatusFailedResponse messages that need to be converted.

Returns:
  - A slice of pointers to response.WithdrawResponseMonthStatusFailed objects containing the mapped data,
    including fields like Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawResponseMapper) mapResponsesWithdrawMonthStatusFailed(Withdraws []*pb.WithdrawMonthStatusFailedResponse) []*response.WithdrawResponseMonthStatusFailed
```

##### `mapResponsesWithdrawMonthStatusSuccess`

mapResponsesWithdrawMonthStatusSuccess maps a slice of WithdrawMonthStatusSuccessResponse gRPC messages to a slice
of API response WithdrawResponseMonthStatusSuccess objects.

Args:
  - Withdraws: A slice of WithdrawMonthStatusSuccessResponse gRPC messages that need to be converted.

Returns:
  - A slice of WithdrawResponseMonthStatusSuccess API responses containing the mapped data.

```go
func (t *withdrawResponseMapper) mapResponsesWithdrawMonthStatusSuccess(Withdraws []*pb.WithdrawMonthStatusSuccessResponse) []*response.WithdrawResponseMonthStatusSuccess
```

##### `mapResponsesWithdrawal`

mapResponsesWithdrawal maps a slice of WithdrawResponse to a slice of WithdrawResponse.

It takes a slice of WithdrawResponse as input and returns a slice of corresponding WithdrawResponse.
The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.

```go
func (w *withdrawResponseMapper) mapResponsesWithdrawal(withdraws []*pb.WithdrawResponse) []*response.WithdrawResponse
```

##### `mapResponsesWithdrawalDeleteAt`

mapResponsesWithdrawalDeleteAt maps a slice of WithdrawResponseDeleteAt gRPC messages
to a slice of API response WithdrawResponseDeleteAt objects.

Args:
  - withdraws: A slice of WithdrawResponseDeleteAt gRPC messages that need to be converted.

Returns:
  - A slice of WithdrawResponseDeleteAt API responses containing the mapped data.

```go
func (w *withdrawResponseMapper) mapResponsesWithdrawalDeleteAt(withdraws []*pb.WithdrawResponseDeleteAt) []*response.WithdrawResponseDeleteAt
```

##### `mapWithdrawResponseYearStatusFailed`

mapWithdrawResponseYearStatusFailed maps a WithdrawYearStatusFailedResponse gRPC message
to an API response WithdrawResponseYearStatusFailed object.

Args:
  - s: The WithdrawYearStatusFailedResponse gRPC message that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseYearStatusFailed containing the mapped data, including
    fields like Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawResponseMapper) mapWithdrawResponseYearStatusFailed(s *pb.WithdrawYearStatusFailedResponse) *response.WithdrawResponseYearStatusFailed
```

##### `mapWithdrawResponseYearStatusSuccess`

mapWithdrawResponseYearStatusSuccess maps a single WithdrawYearStatusSuccessResponse gRPC message
to an API response WithdrawResponseYearStatusSuccess object.

Args:
  - s: The WithdrawYearStatusSuccessResponse gRPC message that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseYearStatusSuccess containing the mapped data, including
    fields like Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawResponseMapper) mapWithdrawResponseYearStatusSuccess(s *pb.WithdrawYearStatusSuccessResponse) *response.WithdrawResponseYearStatusSuccess
```

##### `mapWithdrawResponsesYearStatusFailed`

mapWithdrawResponsesYearStatusFailed maps a slice of gRPC WithdrawYearStatusFailedResponse messages
to a slice of API WithdrawResponseYearStatusFailed objects.

Args:
  - Withdraws: A slice of pointers to pb.WithdrawYearStatusFailedResponse messages that need to be converted.

Returns:
  - A slice of pointers to response.WithdrawResponseYearStatusFailed objects containing the mapped data,
    including fields like Year, TotalFailed, and TotalAmount.s

```go
func (t *withdrawResponseMapper) mapWithdrawResponsesYearStatusFailed(Withdraws []*pb.WithdrawYearStatusFailedResponse) []*response.WithdrawResponseYearStatusFailed
```

##### `mapWithdrawResponsesYearStatusSuccess`

mapWithdrawResponsesYearStatusSuccess maps a slice of gRPC WithdrawYearStatusSuccessResponse messages
to a slice of API WithdrawResponseYearStatusSuccess objects.

Args:
  - Withdraws: A slice of pointers to pb.WithdrawYearStatusSuccessResponse messages that need to be converted.

Returns:
  - A slice of pointers to response.WithdrawResponseYearStatusSuccess objects containing the mapped data,
    including fields like Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawResponseMapper) mapWithdrawResponsesYearStatusSuccess(Withdraws []*pb.WithdrawYearStatusSuccessResponse) []*response.WithdrawResponseYearStatusSuccess
```

## 🚀 Functions

### `mapPaginationMeta`

mapPaginationMeta maps a gRPC PaginationMeta to an HTTP-compatible API response PaginationMeta.

Args:
  - s: A pointer to a pb.PaginationMeta containing the gRPC response data.

Returns:
  - A pointer to a response.PaginationMeta containing the mapped data, including current page, page size,
    total records, and total pages.

```go
func mapPaginationMeta(s *pb.PaginationMeta) *response.PaginationMeta
```