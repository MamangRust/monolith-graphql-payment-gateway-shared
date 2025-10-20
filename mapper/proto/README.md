# 📦 Package `protomapper`

**Source Path:** `shared/mapper/proto`

## 🧩 Types

### `AuthProtoMapper`

AuthProtoMapper defines a set of methods to map authentication-related responses
to their corresponding gRPC protobuf message formats.

```go
type AuthProtoMapper interface {
	ToProtoResponseVerifyCode func(status string, message string) (*pb.ApiResponseVerifyCode)
	ToProtoResponseForgotPassword func(status string, message string) (*pb.ApiResponseForgotPassword)
	ToProtoResponseResetPassword func(status string, message string) (*pb.ApiResponseResetPassword)
	ToProtoResponseLogin func(status string, message string, response *response.TokenResponse) (*pb.ApiResponseLogin)
	ToProtoResponseRegister func(status string, message string, response *response.UserResponse) (*pb.ApiResponseRegister)
	ToProtoResponseRefreshToken func(status string, message string, response *response.TokenResponse) (*pb.ApiResponseRefreshToken)
	ToProtoResponseGetMe func(status string, message string, response *response.UserResponse) (*pb.ApiResponseGetMe)
}
```

### `CardProtoMapper`

CardProtoMapper defines methods for mapping card-related domain responses
to gRPC protobuf response messages.

```go
type CardProtoMapper interface {
	ToProtoResponseCard func(status string, message string, card *response.CardResponse) (*pb.ApiResponseCard)
	ToProtoResponsePaginationCard func(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponse) (*pb.ApiResponsePaginationCard)
	ToProtoResponseCardDeleteAt func(status string, message string) (*pb.ApiResponseCardDelete)
	ToProtoResponseCardAll func(status string, message string) (*pb.ApiResponseCardAll)
	ToProtoResponsePaginationCardDeletedAt func(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponseDeleteAt) (*pb.ApiResponsePaginationCardDeleteAt)
	ToProtoResponseDashboardCard func(status string, message string, dash *response.DashboardCard) (*pb.ApiResponseDashboardCard)
	ToProtoResponseDashboardCardCardNumber func(status string, message string, dash *response.DashboardCardCardNumber) (*pb.ApiResponseDashboardCardNumber)
	ToProtoResponseMonthlyBalances func(status string, message string, cards []*response.CardResponseMonthBalance) (*pb.ApiResponseMonthlyBalance)
	ToProtoResponseYearlyBalances func(status string, message string, cards []*response.CardResponseYearlyBalance) (*pb.ApiResponseYearlyBalance)
	ToProtoResponseMonthlyAmounts func(status string, message string, cards []*response.CardResponseMonthAmount) (*pb.ApiResponseMonthlyAmount)
	ToProtoResponseYearlyAmounts func(status string, message string, cards []*response.CardResponseYearAmount) (*pb.ApiResponseYearlyAmount)
}
```

### `MerchantDocumentProtoMapper`

MerchantDocumentProtoMapper defines methods for mapping merchant document-related responses
into gRPC protobuf response messages.

```go
type MerchantDocumentProtoMapper interface {
	ToProtoResponseMerchantDocument func(status string, message string, doc *response.MerchantDocumentResponse) (*pb.ApiResponseMerchantDocument)
	ToProtoResponsesMerchantDocument func(status string, message string, docs []*response.MerchantDocumentResponse) (*pb.ApiResponsesMerchantDocument)
	ToProtoResponsePaginationMerchantDocument func(pagination *pb.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponse) (*pb.ApiResponsePaginationMerchantDocument)
	ToProtoResponsePaginationMerchantDocumentDeleteAt func(pagination *pb.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponseDeleteAt) (*pb.ApiResponsePaginationMerchantDocumentAt)
	ToProtoResponseMerchantDocumentDelete func(status string, message string) (*pb.ApiResponseMerchantDocumentDelete)
	ToProtoResponseMerchantDocumentAll func(status string, message string) (*pb.ApiResponseMerchantDocumentAll)
}
```

### `MerchantProtoMapper`

MerchantProtoMapper defines methods for mapping merchant-related domain responses
to gRPC protobuf response messages.

```go
type MerchantProtoMapper interface {
	ToProtoResponsePaginationMerchant func(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponse) (*pb.ApiResponsePaginationMerchant)
	ToProtoResponseMerchants func(status string, message string, res []*response.MerchantResponse) (*pb.ApiResponsesMerchant)
	ToProtoResponseMerchant func(status string, message string, res *response.MerchantResponse) (*pb.ApiResponseMerchant)
	ToProtoResponseMerchantAll func(status string, message string) (*pb.ApiResponseMerchantAll)
	ToProtoResponseMerchantDelete func(status string, message string) (*pb.ApiResponseMerchantDelete)
	ToProtoResponsePaginationMerchantDeleteAt func(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponseDeleteAt) (*pb.ApiResponsePaginationMerchantDeleteAt)
	ToProtoResponsePaginationMerchantTransaction func(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantTransactionResponse) (*pb.ApiResponsePaginationMerchantTransaction)
	ToProtoResponseMonthlyPaymentMethods func(status string, message string, ms []*response.MerchantResponseMonthlyPaymentMethod) (*pb.ApiResponseMerchantMonthlyPaymentMethod)
	ToProtoResponseYearlyPaymentMethods func(status string, message string, ms []*response.MerchantResponseYearlyPaymentMethod) (*pb.ApiResponseMerchantYearlyPaymentMethod)
	ToProtoResponseMonthlyAmounts func(status string, message string, ms []*response.MerchantResponseMonthlyAmount) (*pb.ApiResponseMerchantMonthlyAmount)
	ToProtoResponseYearlyAmounts func(status string, message string, ms []*response.MerchantResponseYearlyAmount) (*pb.ApiResponseMerchantYearlyAmount)
	ToProtoResponseMonthlyTotalAmounts func(status string, message string, ms []*response.MerchantResponseMonthlyTotalAmount) (*pb.ApiResponseMerchantMonthlyTotalAmount)
	ToProtoResponseYearlyTotalAmounts func(status string, message string, ms []*response.MerchantResponseYearlyTotalAmount) (*pb.ApiResponseMerchantYearlyTotalAmount)
}
```

### `ProtoMapper`

```go
type ProtoMapper struct {
	AuthProtoMapper AuthProtoMapper
	RoleProtoMapper RoleProtoMapper
	CardProtoMapper CardProtoMapper
	MerchantProtoMapper MerchantProtoMapper
	MerchantDocumentProtoMapper MerchantDocumentProtoMapper
	SaldoProtoMapper SaldoProtoMapper
	TopupProtoMapper TopupProtoMapper
	TransactionProtoMapper TransactionProtoMapper
	TransferProtoMapper TransferProtoMapper
	UserProtoMapper UserProtoMapper
	WithdrawalProtoMapper WithdrawalProtoMapper
}
```

### `RoleProtoMapper`

RoleProtoMapper defines methods to map role-related data structures
into gRPC protobuf response messages.

```go
type RoleProtoMapper interface {
	ToProtoResponseRoleAll func(status string, message string) (*pb.ApiResponseRoleAll)
	ToProtoResponseRoleDelete func(status string, message string) (*pb.ApiResponseRoleDelete)
	ToProtoResponseRole func(status string, message string, pbResponse *response.RoleResponse) (*pb.ApiResponseRole)
	ToProtoResponsesRole func(status string, message string, pbResponse []*response.RoleResponse) (*pb.ApiResponsesRole)
	ToProtoResponsePaginationRole func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponse) (*pb.ApiResponsePaginationRole)
	ToProtoResponsePaginationRoleDeleteAt func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponseDeleteAt) (*pb.ApiResponsePaginationRoleDeleteAt)
}
```

### `SaldoProtoMapper`

SaldoProtoMapper defines methods for mapping saldo (balance)-related responses
into gRPC protobuf response messages.

```go
type SaldoProtoMapper interface {
	ToProtoResponseSaldo func(status string, message string, pbResponse *response.SaldoResponse) (*pb.ApiResponseSaldo)
	ToProtoResponsesSaldo func(status string, message string, pbResponse []*response.SaldoResponse) (*pb.ApiResponsesSaldo)
	ToProtoResponseSaldoDelete func(status string, message string) (*pb.ApiResponseSaldoDelete)
	ToProtoResponseSaldoAll func(status string, message string) (*pb.ApiResponseSaldoAll)
	ToProtoResponseMonthTotalSaldo func(status string, message string, pbResponse []*response.SaldoMonthTotalBalanceResponse) (*pb.ApiResponseMonthTotalSaldo)
	ToProtoResponseYearTotalSaldo func(status string, message string, pbResponse []*response.SaldoYearTotalBalanceResponse) (*pb.ApiResponseYearTotalSaldo)
	ToProtoResponseMonthSaldoBalances func(status string, message string, pbResponse []*response.SaldoMonthBalanceResponse) (*pb.ApiResponseMonthSaldoBalances)
	ToProtoResponseYearSaldoBalances func(status string, message string, pbResponse []*response.SaldoYearBalanceResponse) (*pb.ApiResponseYearSaldoBalances)
	ToProtoResponsePaginationSaldo func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponse) (*pb.ApiResponsePaginationSaldo)
	ToProtoResponsePaginationSaldoDeleteAt func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponseDeleteAt) (*pb.ApiResponsePaginationSaldoDeleteAt)
}
```

### `TopupProtoMapper`

TopupProtoMapper defines methods for mapping top-up-related domain responses
into gRPC protobuf response messages.

```go
type TopupProtoMapper interface {
	ToProtoResponseTopup func(status string, message string, s *response.TopupResponse) (*pb.ApiResponseTopup)
	ToProtoResponseTopupDeletAt func(status string, message string, s *response.TopupResponseDeleteAt) (*pb.ApiResponseTopupDeleteAt)
	ToProtoResponseTopupDelete func(status string, message string) (*pb.ApiResponseTopupDelete)
	ToProtoResponseTopupAll func(status string, message string) (*pb.ApiResponseTopupAll)
	ToProtoResponsePaginationTopup func(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponse) (*pb.ApiResponsePaginationTopup)
	ToProtoResponsePaginationTopupDeleteAt func(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponseDeleteAt) (*pb.ApiResponsePaginationTopupDeleteAt)
	ToProtoResponseTopupMonthStatusSuccess func(status string, message string, s []*response.TopupResponseMonthStatusSuccess) (*pb.ApiResponseTopupMonthStatusSuccess)
	ToProtoResponseTopupYearStatusSuccess func(status string, message string, s []*response.TopupResponseYearStatusSuccess) (*pb.ApiResponseTopupYearStatusSuccess)
	ToProtoResponseTopupMonthStatusFailed func(status string, message string, s []*response.TopupResponseMonthStatusFailed) (*pb.ApiResponseTopupMonthStatusFailed)
	ToProtoResponseTopupYearStatusFailed func(status string, message string, s []*response.TopupResponseYearStatusFailed) (*pb.ApiResponseTopupYearStatusFailed)
	ToProtoResponseTopupMonthMethod func(status string, message string, s []*response.TopupMonthMethodResponse) (*pb.ApiResponseTopupMonthMethod)
	ToProtoResponseTopupYearMethod func(status string, message string, s []*response.TopupYearlyMethodResponse) (*pb.ApiResponseTopupYearMethod)
	ToProtoResponseTopupMonthAmount func(status string, message string, s []*response.TopupMonthAmountResponse) (*pb.ApiResponseTopupMonthAmount)
	ToProtoResponseTopupYearAmount func(status string, message string, s []*response.TopupYearlyAmountResponse) (*pb.ApiResponseTopupYearAmount)
}
```

### `TransactionProtoMapper`

TransactionProtoMapper defines methods for mapping transaction-related domain responses
into gRPC protobuf response messages.

```go
type TransactionProtoMapper interface {
	ToProtoResponseTransactionMonthStatusSuccess func(status string, message string, pbResponse []*response.TransactionResponseMonthStatusSuccess) (*pb.ApiResponseTransactionMonthStatusSuccess)
	ToProtoResponseTransactionYearStatusSuccess func(status string, message string, pbResponse []*response.TransactionResponseYearStatusSuccess) (*pb.ApiResponseTransactionYearStatusSuccess)
	ToProtoResponseTransactionMonthStatusFailed func(status string, message string, pbResponse []*response.TransactionResponseMonthStatusFailed) (*pb.ApiResponseTransactionMonthStatusFailed)
	ToProtoResponseTransactionYearStatusFailed func(status string, message string, pbResponse []*response.TransactionResponseYearStatusFailed) (*pb.ApiResponseTransactionYearStatusFailed)
	ToProtoResponseTransactionMonthMethod func(status string, message string, pbResponse []*response.TransactionMonthMethodResponse) (*pb.ApiResponseTransactionMonthMethod)
	ToProtoResponseTransactionYearMethod func(status string, message string, pbResponse []*response.TransactionYearMethodResponse) (*pb.ApiResponseTransactionYearMethod)
	ToProtoResponseTransactionMonthAmount func(status string, message string, pbResponse []*response.TransactionMonthAmountResponse) (*pb.ApiResponseTransactionMonthAmount)
	ToProtoResponseTransactionYearAmount func(status string, message string, pbResponse []*response.TransactionYearlyAmountResponse) (*pb.ApiResponseTransactionYearAmount)
	ToProtoResponseTransaction func(status string, message string, pbResponse *response.TransactionResponse) (*pb.ApiResponseTransaction)
	ToProtoResponseTransactions func(status string, message string, pbResponse []*response.TransactionResponse) (*pb.ApiResponseTransactions)
	ToProtoResponseTransactionDelete func(status string, message string) (*pb.ApiResponseTransactionDelete)
	ToProtoResponseTransactionAll func(status string, message string) (*pb.ApiResponseTransactionAll)
	ToProtoResponsePaginationTransaction func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) (*pb.ApiResponsePaginationTransaction)
	ToProtoResponsePaginationTransactionDeleteAt func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) (*pb.ApiResponsePaginationTransactionDeleteAt)
}
```

### `TransferProtoMapper`

```go
type TransferProtoMapper interface {
	ToProtoResponseTransferMonthStatusSuccess func(status string, message string, pbResponse []*response.TransferResponseMonthStatusSuccess) (*pb.ApiResponseTransferMonthStatusSuccess)
	ToProtoResponseTransferYearStatusSuccess func(status string, message string, pbResponse []*response.TransferResponseYearStatusSuccess) (*pb.ApiResponseTransferYearStatusSuccess)
	ToProtoResponseTransferMonthStatusFailed func(status string, message string, pbResponse []*response.TransferResponseMonthStatusFailed) (*pb.ApiResponseTransferMonthStatusFailed)
	ToProtoResponseTransferYearStatusFailed func(status string, message string, pbResponse []*response.TransferResponseYearStatusFailed) (*pb.ApiResponseTransferYearStatusFailed)
	ToProtoResponseTransferMonthAmount func(status string, message string, pbResponse []*response.TransferMonthAmountResponse) (*pb.ApiResponseTransferMonthAmount)
	ToProtoResponseTransferYearAmount func(status string, message string, pbResponse []*response.TransferYearAmountResponse) (*pb.ApiResponseTransferYearAmount)
	ToProtoResponseTransfer func(status string, message string, pbResponse *response.TransferResponse) (*pb.ApiResponseTransfer)
	ToProtoResponseTransfers func(status string, message string, pbResponse []*response.TransferResponse) (*pb.ApiResponseTransfers)
	ToProtoResponseTransferDelete func(status string, message string) (*pb.ApiResponseTransferDelete)
	ToProtoResponseTransferAll func(status string, message string) (*pb.ApiResponseTransferAll)
	ToProtoResponsePaginationTransfer func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponse) (*pb.ApiResponsePaginationTransfer)
	ToProtoResponsePaginationTransferDeleteAt func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponseDeleteAt) (*pb.ApiResponsePaginationTransferDeleteAt)
}
```

### `UserProtoMapper`

UserProtoMapper defines methods for converting user-related data structures
into gRPC protobuf response messages.

```go
type UserProtoMapper interface {
	ToProtoResponsesUser func(status string, message string, pbResponse []*response.UserResponse) (*pb.ApiResponsesUser)
	ToProtoResponseUser func(status string, message string, pbResponse *response.UserResponse) (*pb.ApiResponseUser)
	ToProtoResponseUserDeleteAt func(status string, message string, pbResponse *response.UserResponseDeleteAt) (*pb.ApiResponseUserDeleteAt)
	ToProtoResponseUserDelete func(status string, message string) (*pb.ApiResponseUserDelete)
	ToProtoResponseUserAll func(status string, message string) (*pb.ApiResponseUserAll)
	ToProtoResponsePaginationUserDeleteAt func(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) (*pb.ApiResponsePaginationUserDeleteAt)
	ToProtoResponsePaginationUser func(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponse) (*pb.ApiResponsePaginationUser)
}
```

### `WithdrawalProtoMapper`

```go
type WithdrawalProtoMapper interface {
	ToProtoResponseWithdraw func(status string, message string, withdraw *response.WithdrawResponse) (*pb.ApiResponseWithdraw)
	ToProtoResponsesWithdraw func(status string, message string, pbResponse []*response.WithdrawResponse) (*pb.ApiResponsesWithdraw)
	ToProtoResponseWithdrawDelete func(status string, message string) (*pb.ApiResponseWithdrawDelete)
	ToProtoResponseWithdrawAll func(status string, message string) (*pb.ApiResponseWithdrawAll)
	ToProtoResponsePaginationWithdraw func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponse) (*pb.ApiResponsePaginationWithdraw)
	ToProtoResponsePaginationWithdrawDeleteAt func(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponseDeleteAt) (*pb.ApiResponsePaginationWithdrawDeleteAt)
	ToProtoResponseWithdrawMonthStatusSuccess func(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusSuccess) (*pb.ApiResponseWithdrawMonthStatusSuccess)
	ToProtoResponseWithdrawYearStatusSuccess func(status string, message string, pbResponse []*response.WithdrawResponseYearStatusSuccess) (*pb.ApiResponseWithdrawYearStatusSuccess)
	ToProtoResponseWithdrawMonthStatusFailed func(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusFailed) (*pb.ApiResponseWithdrawMonthStatusFailed)
	ToProtoResponseWithdrawYearStatusFailed func(status string, message string, pbResponse []*response.WithdrawResponseYearStatusFailed) (*pb.ApiResponseWithdrawYearStatusFailed)
	ToProtoResponseWithdrawMonthAmount func(status string, message string, pbResponse []*response.WithdrawMonthlyAmountResponse) (*pb.ApiResponseWithdrawMonthAmount)
	ToProtoResponseWithdrawYearAmount func(status string, message string, pbResponse []*response.WithdrawYearlyAmountResponse) (*pb.ApiResponseWithdrawYearAmount)
}
```

### `authProtoMapper`

```go
type authProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponseForgotPassword`

ToProtoResponseForgotPassword maps the response to the proto type.

Args:

	status: The status of the response.
	message: The message of the response.

Returns:

	A pointer to the proto type.

```go
func (s *authProtoMapper) ToProtoResponseForgotPassword(status string, message string) *pb.ApiResponseForgotPassword
```

##### `ToProtoResponseGetMe`

ToProtoResponseGetMe maps a user profile response to the protobuf format.

Args:

	status: The status of the response.
	message: The message associated with the response.
	response: The UserResponse containing the user's details.

Returns:

	A pointer to the ApiResponseGetMe protobuf type containing the status, message, and user data.

```go
func (s *authProtoMapper) ToProtoResponseGetMe(status string, message string, response *response.UserResponse) *pb.ApiResponseGetMe
```

##### `ToProtoResponseLogin`

ToProtoResponseLogin maps the response to the proto type.

Args:

	status: The status of the response.
	message: The message of the response.
	response: The response containing the access and refresh tokens.

Returns:

	A pointer to the proto type.

```go
func (s *authProtoMapper) ToProtoResponseLogin(status string, message string, response *response.TokenResponse) *pb.ApiResponseLogin
```

##### `ToProtoResponseRefreshToken`

```go
func (s *authProtoMapper) ToProtoResponseRefreshToken(status string, message string, response *response.TokenResponse) *pb.ApiResponseRefreshToken
```

##### `ToProtoResponseRegister`

```go
func (s *authProtoMapper) ToProtoResponseRegister(status string, message string, response *response.UserResponse) *pb.ApiResponseRegister
```

##### `ToProtoResponseResetPassword`

ToProtoResponseResetPassword maps the response to the proto type.

Args:

	status: The status of the response.
	message: The message of the response.

Returns:

	A pointer to the proto type.

```go
func (s *authProtoMapper) ToProtoResponseResetPassword(status string, message string) *pb.ApiResponseResetPassword
```

##### `ToProtoResponseVerifyCode`

ToProtoResponseVerifyCode maps the response to the proto type.

Args:

	status: The status of the response.
	message: The message of the response.

Returns:

	A pointer to the proto type.

```go
func (s *authProtoMapper) ToProtoResponseVerifyCode(status string, message string) *pb.ApiResponseVerifyCode
```

### `cardProtoMapper`

```go
type cardProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponseCard`

ToProtoResponseCard maps a *response.CardResponse to a *pb.ApiResponseCard proto response.

It is used to generate the response for the CardService.GetCard rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseCard(status string, message string, card *response.CardResponse) *pb.ApiResponseCard
```

##### `ToProtoResponseCardAll`

ToProtoResponseCardAll maps a status and message to a *pb.ApiResponseCardAll proto response.

It is used to generate the response for the CardService.GetAllCard rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseCardAll(status string, message string) *pb.ApiResponseCardAll
```

##### `ToProtoResponseCardDeleteAt`

ToProtoResponseCardDeleteAt maps a status and message to a *pb.ApiResponseCardDelete proto response.

It is used to generate the response for the CardService.DeleteCard rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseCardDeleteAt(status string, message string) *pb.ApiResponseCardDelete
```

##### `ToProtoResponseDashboardCard`

ToProtoResponseDashboardCard maps a status, message and a *response.DashboardCard to a *pb.ApiResponseDashboardCard proto response.

It is used to generate the response for the CardService.GetDashboardCard rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseDashboardCard(status string, message string, dash *response.DashboardCard) *pb.ApiResponseDashboardCard
```

##### `ToProtoResponseDashboardCardCardNumber`

ToProtoResponseDashboardCardCardNumber maps a status, message, and a *response.DashboardCardCardNumber
to a *pb.ApiResponseDashboardCardNumber proto response.

It is used to generate the response for the CardService.GetDashboardCardCardNumber rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseDashboardCardCardNumber(status string, message string, dash *response.DashboardCardCardNumber) *pb.ApiResponseDashboardCardNumber
```

##### `ToProtoResponseMonthlyAmounts`

ToProtoResponseMonthlyAmounts maps a status, message and a list of *response.CardResponseMonthAmount
to a *pb.ApiResponseMonthlyAmount proto response.

It is used to generate the response for the CardService.GetMonthlyAmount rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseMonthlyAmounts(status string, message string, cards []*response.CardResponseMonthAmount) *pb.ApiResponseMonthlyAmount
```

##### `ToProtoResponseMonthlyBalances`

ToProtoResponseMonthlyBalances maps a status, message and a list of *response.CardResponseMonthBalance
to a *pb.ApiResponseMonthlyBalance proto response.

It is used to generate the response for the CardService.GetMonthlyBalance rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseMonthlyBalances(status string, message string, cards []*response.CardResponseMonthBalance) *pb.ApiResponseMonthlyBalance
```

##### `ToProtoResponsePaginationCard`

ToProtoResponsePaginationCard maps a pagination meta, status, message and a list of *response.CardResponse
to a *pb.ApiResponsePaginationCard proto response.

It is used to generate the response for the CardService.ListCard rpc method.

```go
func (s *cardProtoMapper) ToProtoResponsePaginationCard(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponse) *pb.ApiResponsePaginationCard
```

##### `ToProtoResponsePaginationCardDeletedAt`

ToProtoResponsePaginationCardDeletedAt maps a pagination meta, status, message and a list of *response.CardResponseDeleteAt
to a *pb.ApiResponsePaginationCardDeleteAt proto response.

It is used to generate the response for the CardService.ListDeletedCard rpc method.

```go
func (s *cardProtoMapper) ToProtoResponsePaginationCardDeletedAt(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponseDeleteAt) *pb.ApiResponsePaginationCardDeleteAt
```

##### `ToProtoResponseYearlyAmounts`

ToProtoResponseYearlyAmounts maps a status, message and a list of *response.CardResponseYearAmount
to a *pb.ApiResponseYearlyAmount proto response.

It is used to generate the response for the CardService.GetYearlyAmount rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseYearlyAmounts(status string, message string, cards []*response.CardResponseYearAmount) *pb.ApiResponseYearlyAmount
```

##### `ToProtoResponseYearlyBalances`

ToProtoResponseYearlyBalances maps a status, message and a list of *response.CardResponseYearlyBalance
to a *pb.ApiResponseYearlyBalance proto response.

It is used to generate the response for the CardService.GetYearlyBalance rpc method.

```go
func (s *cardProtoMapper) ToProtoResponseYearlyBalances(status string, message string, cards []*response.CardResponseYearlyBalance) *pb.ApiResponseYearlyBalance
```

##### `mapCardResponse`

mapCardResponse maps a *response.CardResponse to a *pb.CardResponse proto response.

It is used to generate the response for the CardService.GetCard rpc method.

```go
func (s *cardProtoMapper) mapCardResponse(card *response.CardResponse) *pb.CardResponse
```

##### `mapCardResponseDeleteAt`

mapCardResponseDeleteAt maps a *response.CardResponseDeleteAt to a *pb.CardResponseDeleteAt proto response.

This function converts a CardResponseDeleteAt, which includes deletion information, into its protobuf
representation. It handles the conversion of all fields including a nullable DeletedAt field.

```go
func (s *cardProtoMapper) mapCardResponseDeleteAt(card *response.CardResponseDeleteAt) *pb.CardResponseDeleteAt
```

##### `mapCardResponses`

mapCardResponses maps a list of *response.CardResponse to a list of *pb.CardResponse
proto response.

It is used to generate the response for the CardService.ListCard rpc method.

```go
func (s *cardProtoMapper) mapCardResponses(roles []*response.CardResponse) []*pb.CardResponse
```

##### `mapCardResponsesDeleteAt`

mapCardResponsesDeleteAt maps a list of *response.CardResponseDeleteAt to a list of
*pb.CardResponseDeleteAt proto response.

It is used to generate the response for the CardService.ListDeletedCard rpc method.

```go
func (s *cardProtoMapper) mapCardResponsesDeleteAt(roles []*response.CardResponseDeleteAt) []*pb.CardResponseDeleteAt
```

##### `mapDashboardCard`

mapDashboardCard maps a *response.DashboardCard to a *pb.CardResponseDashboard proto response.

This function converts a DashboardCard from the response domain model to its protobuf
representation, handling the mapping of all financial metrics fields.

```go
func (s *cardProtoMapper) mapDashboardCard(dash *response.DashboardCard) *pb.CardResponseDashboard
```

##### `mapDashboardCardCardNumber`

mapDashboardCardCardNumber maps a *response.DashboardCardCardNumber to a *pb.CardResponseDashboardCardNumber proto response.

This function converts a DashboardCardCardNumber from the response domain model to its protobuf
representation, handling the mapping of all financial metrics fields for a specific card.

```go
func (s *cardProtoMapper) mapDashboardCardCardNumber(dash *response.DashboardCardCardNumber) *pb.CardResponseDashboardCardNumber
```

##### `mapMonthlyAmount`

mapMonthlyAmount maps a *response.CardResponseMonthAmount to a *pb.CardResponseMonthlyAmount proto response.

It is used to generate the response for the CardService.GetMonthlyAmount rpc method.

```go
func (s *cardProtoMapper) mapMonthlyAmount(card *response.CardResponseMonthAmount) *pb.CardResponseMonthlyAmount
```

##### `mapMonthlyAmounts`

mapMonthlyAmounts maps a list of *response.CardResponseMonthAmount to a list of
*pb.CardResponseMonthlyAmount proto responses.

It iterates over each CardResponseMonthAmount in the input slice, converting
them to their protobuf equivalent using the mapMonthlyAmount function. This
function is used to generate the response data for monthly amount RPC methods.

```go
func (s *cardProtoMapper) mapMonthlyAmounts(roles []*response.CardResponseMonthAmount) []*pb.CardResponseMonthlyAmount
```

##### `mapMonthlyBalance`

mapMonthlyBalance maps a *response.CardResponseMonthBalance to a *pb.CardResponseMonthlyBalance
proto response.

It is used to generate the response for the CardService.GetMonthlyBalance rpc method.

```go
func (s *cardProtoMapper) mapMonthlyBalance(card *response.CardResponseMonthBalance) *pb.CardResponseMonthlyBalance
```

##### `mapMonthlyBalances`

mapMonthlyBalances maps a list of *response.CardResponseMonthBalance to a list of
*pb.CardResponseMonthlyBalance proto responses.

It iterates over each CardResponseMonthBalance in the input slice, converting
them to their protobuf equivalent using the mapMonthlyBalance function. This
function is used to generate the response data for monthly balance RPC methods.

```go
func (s *cardProtoMapper) mapMonthlyBalances(roles []*response.CardResponseMonthBalance) []*pb.CardResponseMonthlyBalance
```

##### `mapYearlyAmount`

mapYearlyAmount maps a *response.CardResponseYearAmount to a *pb.CardResponseYearlyAmount proto response.

This function takes a CardResponseYearAmount from the response domain model and converts it
into its protobuf representation, ensuring that the year and total amount fields are accurately mapped.

```go
func (s *cardProtoMapper) mapYearlyAmount(card *response.CardResponseYearAmount) *pb.CardResponseYearlyAmount
```

##### `mapYearlyAmounts`

mapYearlyAmounts maps a list of *response.CardResponseYearAmount to a list of
*pb.CardResponseYearlyAmount proto responses.

It iterates over each CardResponseYearAmount in the input slice, converting
them to their protobuf equivalent using the mapYearlyAmount function. This
function is used to generate the response data for yearly amount RPC methods.

```go
func (s *cardProtoMapper) mapYearlyAmounts(roles []*response.CardResponseYearAmount) []*pb.CardResponseYearlyAmount
```

##### `mapYearlyBalance`

mapYearlyBalance maps a *response.CardResponseYearlyBalance to a *pb.CardResponseYearlyBalance proto response.

This function takes a CardResponseYearlyBalance from the response domain model and converts it
into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.

```go
func (s *cardProtoMapper) mapYearlyBalance(card *response.CardResponseYearlyBalance) *pb.CardResponseYearlyBalance
```

##### `mapYearlyBalances`

mapYearlyBalances maps a list of *response.CardResponseYearlyBalance to a list of
*pb.CardResponseYearlyBalance proto responses.

It iterates over each CardResponseYearlyBalance in the input slice, converting
them to their protobuf equivalent using the mapYearlyBalance function. This
function is used to generate the response data for yearly balance RPC methods.

```go
func (s *cardProtoMapper) mapYearlyBalances(roles []*response.CardResponseYearlyBalance) []*pb.CardResponseYearlyBalance
```

### `merchantDocumentProtoMapper`

```go
type merchantDocumentProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponseMerchantDocument`

ToProtoResponseMerchantDocument converts a MerchantDocumentResponse to its protobuf representation.
It includes the status, message, and mapped document data.
Parameters:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - doc: The MerchantDocumentResponse that needs to be converted.

Returns:

	A pointer to ApiResponseMerchantDocument containing the status, message, and document data.

```go
func (m *merchantDocumentProtoMapper) ToProtoResponseMerchantDocument(status string, message string, doc *response.MerchantDocumentResponse) *pb.ApiResponseMerchantDocument
```

##### `ToProtoResponseMerchantDocumentAll`

ToProtoResponseMerchantDocumentAll converts a status and message to its protobuf representation
specifically for bulk merchant document operations.
It includes the status and message of the API response.
Parameters:
  - status: The status of the API response.
  - message: A descriptive message of the API response.

Returns:

	A pointer to ApiResponseMerchantDocumentAll containing the status and message.

```go
func (m *merchantDocumentProtoMapper) ToProtoResponseMerchantDocumentAll(status string, message string) *pb.ApiResponseMerchantDocumentAll
```

##### `ToProtoResponseMerchantDocumentDelete`

ToProtoResponseMerchantDocumentDelete converts a response message to its protobuf representation
specifically for merchant document deletion operations.
It includes the status and message of the API response.
Parameters:
  - status: The status of the API response.
  - message: A descriptive message of the API response.

Returns:

	A pointer to ApiResponseMerchantDocumentDelete containing the status and message.

```go
func (m *merchantDocumentProtoMapper) ToProtoResponseMerchantDocumentDelete(status string, message string) *pb.ApiResponseMerchantDocumentDelete
```

##### `ToProtoResponsePaginationMerchantDocument`

ToProtoResponsePaginationMerchantDocument converts a list of MerchantDocumentResponse to its protobuf representation
with pagination data.
It includes the status, message, mapped document data, and pagination data.
Parameters:
  - pagination: The pagination data for the paginated response.
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - docs: The list of MerchantDocumentResponse that needs to be converted.

Returns:

	A pointer to ApiResponsePaginationMerchantDocument containing the status, message, document data, and pagination data.

```go
func (m *merchantDocumentProtoMapper) ToProtoResponsePaginationMerchantDocument(pagination *pb.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponse) *pb.ApiResponsePaginationMerchantDocument
```

##### `ToProtoResponsePaginationMerchantDocumentDeleteAt`

ToProtoResponsePaginationMerchantDocumentDeleteAt converts a list of MerchantDocumentResponseDeleteAt to its protobuf representation
with pagination data.
It includes the status, message, mapped document data, and pagination data.
Parameters:
  - pagination: The pagination data for the paginated response.
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - docs: The list of MerchantDocumentResponseDeleteAt that needs to be converted.

Returns:

	A pointer to ApiResponsePaginationMerchantDocumentAt containing the status, message, document data, and pagination data.

```go
func (m *merchantDocumentProtoMapper) ToProtoResponsePaginationMerchantDocumentDeleteAt(pagination *pb.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponseDeleteAt) *pb.ApiResponsePaginationMerchantDocumentAt
```

##### `ToProtoResponsesMerchantDocument`

ToProtoResponsesMerchantDocument converts a list of MerchantDocumentResponse to its protobuf representation.
It includes the status, message, and mapped document data.
Parameters:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - docs: The list of MerchantDocumentResponse that needs to be converted.

Returns:

	A pointer to ApiResponsesMerchantDocument containing the status, message, and document data.

```go
func (m *merchantDocumentProtoMapper) ToProtoResponsesMerchantDocument(status string, message string, docs []*response.MerchantDocumentResponse) *pb.ApiResponsesMerchantDocument
```

##### `mapMerchantDocument`

mapMerchantDocument maps a MerchantDocumentResponse to its protobuf representation.
Parameters:
  - doc: The MerchantDocumentResponse that needs to be converted.

Returns:

	A pointer to MerchantDocument containing the document data.

```go
func (m *merchantDocumentProtoMapper) mapMerchantDocument(doc *response.MerchantDocumentResponse) *pb.MerchantDocument
```

##### `mapMerchantDocumentDeleteAt`

mapMerchantDocumentDeleteAt maps a MerchantDocumentResponseDeleteAt to its protobuf representation.
It handles the conversion of document metadata including deletion information.
Parameters:
  - doc: The MerchantDocumentResponseDeleteAt that needs to be converted.

Returns:

	A pointer to MerchantDocumentDeleteAt containing the converted document data.

```go
func (m *merchantDocumentProtoMapper) mapMerchantDocumentDeleteAt(doc *response.MerchantDocumentResponseDeleteAt) *pb.MerchantDocumentDeleteAt
```

##### `mapMerchantDocuments`

mapMerchantDocuments maps a list of MerchantDocumentResponse to its protobuf representation.
Parameters:
  - docs: The list of MerchantDocumentResponse that needs to be converted.

Returns:

	A slice of pointers to MerchantDocument containing the document data.

```go
func (m *merchantDocumentProtoMapper) mapMerchantDocuments(docs []*response.MerchantDocumentResponse) []*pb.MerchantDocument
```

##### `mapMerchantDocumentsDeleteAt`

mapMerchantDocumentsDeleteAt maps a list of MerchantDocumentResponseDeleteAt to their protobuf representation.
It processes each document in the input slice and converts it to a MerchantDocumentDeleteAt,
including all necessary metadata and deletion information.
Parameters:
  - docs: A slice of MerchantDocumentResponseDeleteAt to be converted.

Returns:

	A slice of pointers to MerchantDocumentDeleteAt containing the converted document data.

```go
func (m *merchantDocumentProtoMapper) mapMerchantDocumentsDeleteAt(docs []*response.MerchantDocumentResponseDeleteAt) []*pb.MerchantDocumentDeleteAt
```

### `merchantProtoMapper`

```go
type merchantProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponseMerchant`

ToProtoResponseMerchant maps a status, message and a *response.MerchantResponse
to a *pb.ApiResponseMerchant proto response.

It is used to generate the response for the MerchantService.GetMerchant rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseMerchant(status string, message string, res *response.MerchantResponse) *pb.ApiResponseMerchant
```

##### `ToProtoResponseMerchantAll`

ToProtoResponseMerchantAll maps a status and message to its protobuf representation
specifically for bulk merchant operations.
It includes the status and message of the API response.
Parameters:
  - status: The status of the API response.
  - message: A descriptive message for the API response.

Returns:

	A pointer to ApiResponseMerchantAll containing the status and message.

```go
func (m *merchantProtoMapper) ToProtoResponseMerchantAll(status string, message string) *pb.ApiResponseMerchantAll
```

##### `ToProtoResponseMerchantDelete`

ToProtoResponseMerchantDelete converts a status and message to its protobuf representation
specifically for merchant deletion operations.
It includes the status and message of the API response.
Parameters:
  - status: The status of the API response.
  - message: A descriptive message of the API response.

Returns:

	A pointer to ApiResponseMerchantDelete containing the status and message.

```go
func (m *merchantProtoMapper) ToProtoResponseMerchantDelete(status string, message string) *pb.ApiResponseMerchantDelete
```

##### `ToProtoResponseMerchants`

ToProtoResponseMerchants maps a status, message and a list of *response.MerchantResponse
to a *pb.ApiResponsesMerchant proto response.

It is used to generate the response for the MerchantService.ListMerchants rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseMerchants(status string, message string, res []*response.MerchantResponse) *pb.ApiResponsesMerchant
```

##### `ToProtoResponseMonthlyAmounts`

ToProtoResponseMonthlyAmounts maps a status, message and a list of *response.MerchantResponseMonthlyAmount
to a *pb.ApiResponseMerchantMonthlyAmount proto response.

It is used to generate the response for the MerchantService.GetMonthlyAmount rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseMonthlyAmounts(status string, message string, ms []*response.MerchantResponseMonthlyAmount) *pb.ApiResponseMerchantMonthlyAmount
```

##### `ToProtoResponseMonthlyPaymentMethods`

ToProtoResponseMonthlyPaymentMethods maps a status, message and a list of *response.MerchantResponseMonthlyPaymentMethod
to a *pb.ApiResponseMerchantMonthlyPaymentMethod proto response.

It is used to generate the response for the MerchantService.GetMonthlyPaymentMethod rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseMonthlyPaymentMethods(status string, message string, ms []*response.MerchantResponseMonthlyPaymentMethod) *pb.ApiResponseMerchantMonthlyPaymentMethod
```

##### `ToProtoResponseMonthlyTotalAmounts`

ToProtoResponseMonthlyTotalAmounts maps a status, message and a list of *response.MerchantResponseMonthlyTotalAmount
to a *pb.ApiResponseMerchantMonthlyTotalAmount proto response.

It is used to generate the response for the MerchantService.GetMonthlyTotalAmount rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseMonthlyTotalAmounts(status string, message string, ms []*response.MerchantResponseMonthlyTotalAmount) *pb.ApiResponseMerchantMonthlyTotalAmount
```

##### `ToProtoResponsePaginationMerchant`

ToProtoResponsePaginationMerchant converts a list of MerchantResponse and pagination metadata
into a protobuf ApiResponsePaginationMerchant. It includes the status and message for the API response.
Parameters:
  - pagination: The pagination metadata for the response.
  - status: The status of the API response.
  - message: A descriptive message for the API response.
  - merchants: A list of MerchantResponse objects to be included in the response.

Returns:

	A pointer to ApiResponsePaginationMerchant containing the status, message, merchant data, and pagination data.

```go
func (m *merchantProtoMapper) ToProtoResponsePaginationMerchant(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponse) *pb.ApiResponsePaginationMerchant
```

##### `ToProtoResponsePaginationMerchantDeleteAt`

ToProtoResponsePaginationMerchantDeleteAt maps a pagination meta, status, message and a list of *response.MerchantResponseDeleteAt
to a *pb.ApiResponsePaginationMerchantDeleteAt proto response.

It is used to generate the response for the MerchantService.ListDeletedMerchant rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponsePaginationMerchantDeleteAt(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponseDeleteAt) *pb.ApiResponsePaginationMerchantDeleteAt
```

##### `ToProtoResponsePaginationMerchantTransaction`

ToProtoResponsePaginationMerchantTransaction maps a pagination meta, status, message and a list of *response.MerchantTransactionResponse
to a *pb.ApiResponsePaginationMerchantTransaction proto response.

It is used to generate the response for the MerchantService.ListMerchantTransaction rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponsePaginationMerchantTransaction(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantTransactionResponse) *pb.ApiResponsePaginationMerchantTransaction
```

##### `ToProtoResponseYearlyAmounts`

ToProtoResponseYearlyAmounts maps a status, message and a list of *response.MerchantResponseYearlyAmount
to a *pb.ApiResponseMerchantYearlyAmount proto response.

It is used to generate the response for the MerchantService.GetYearlyAmount rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseYearlyAmounts(status string, message string, ms []*response.MerchantResponseYearlyAmount) *pb.ApiResponseMerchantYearlyAmount
```

##### `ToProtoResponseYearlyPaymentMethods`

ToProtoResponseYearlyPaymentMethods maps a status, message and a list of *response.MerchantResponseYearlyPaymentMethod
to a *pb.ApiResponseMerchantYearlyPaymentMethod proto response.

It is used to generate the response for the MerchantService.GetYearlyPaymentMethod rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseYearlyPaymentMethods(status string, message string, ms []*response.MerchantResponseYearlyPaymentMethod) *pb.ApiResponseMerchantYearlyPaymentMethod
```

##### `ToProtoResponseYearlyTotalAmounts`

ToProtoResponseYearlyTotalAmounts maps a status, message, and a list of *response.MerchantResponseYearlyTotalAmount
to a *pb.ApiResponseMerchantYearlyTotalAmount proto response.

It is used to generate the response for the MerchantService.GetYearlyTotalAmount rpc method.

```go
func (m *merchantProtoMapper) ToProtoResponseYearlyTotalAmounts(status string, message string, ms []*response.MerchantResponseYearlyTotalAmount) *pb.ApiResponseMerchantYearlyTotalAmount
```

##### `mapMerchantResponse`

mapMerchantResponse maps a *response.MerchantResponse to a *pb.MerchantResponse proto message.

It is used to generate the response for the MerchantService.GetMerchant rpc method.

```go
func (m *merchantProtoMapper) mapMerchantResponse(merchant *response.MerchantResponse) *pb.MerchantResponse
```

##### `mapMerchantResponseDeleteAt`

mapMerchantResponseDeleteAt maps a *response.MerchantResponseDeleteAt to a *pb.MerchantResponseDeleteAt proto message.

It is used to generate the response for the MerchantService.ListDeletedMerchant rpc method.

```go
func (m *merchantProtoMapper) mapMerchantResponseDeleteAt(merchant *response.MerchantResponseDeleteAt) *pb.MerchantResponseDeleteAt
```

##### `mapMerchantResponses`

mapMerchantResponses maps a list of *response.MerchantResponse to a list of
*pb.MerchantResponse proto responses.

It iterates over each MerchantResponse in the input slice, converting
them to their protobuf equivalent using the mapMerchantResponse function.
This function is used to generate the response data for merchant-related RPC methods.

```go
func (s *merchantProtoMapper) mapMerchantResponses(roles []*response.MerchantResponse) []*pb.MerchantResponse
```

##### `mapMerchantResponsesDeleteAt`

mapMerchantResponsesDeleteAt maps a list of *response.MerchantResponseDeleteAt to a list of
*pb.MerchantResponseDeleteAt proto responses.

It iterates over each MerchantResponseDeleteAt in the input slice, converting
them to their protobuf equivalent using the mapMerchantResponseDeleteAt function.
This function is used to generate the response data for the MerchantService.ListDeletedMerchant rpc method.

```go
func (s *merchantProtoMapper) mapMerchantResponsesDeleteAt(roles []*response.MerchantResponseDeleteAt) []*pb.MerchantResponseDeleteAt
```

##### `mapMerchantTransactionResponse`

mapMerchantTransactionResponse maps a *response.MerchantTransactionResponse to a *pb.MerchantTransactionResponse proto message.

It is used to generate the response for the MerchantService.ListMerchantTransaction rpc method.

```go
func (m *merchantProtoMapper) mapMerchantTransactionResponse(merchant *response.MerchantTransactionResponse) *pb.MerchantTransactionResponse
```

##### `mapMerchantTransactionResponses`

mapMerchantTransactionResponses maps a list of *response.MerchantTransactionResponse to a list of
*pb.MerchantTransactionResponse proto responses.

It iterates over each MerchantTransactionResponse in the input slice, converting
them to their protobuf equivalent using the mapMerchantTransactionResponse function.
This function is used to generate the response data for merchant transaction-related RPC methods.

```go
func (s *merchantProtoMapper) mapMerchantTransactionResponses(roles []*response.MerchantTransactionResponse) []*pb.MerchantTransactionResponse
```

##### `mapResponseMonthlyAmount`

mapResponseMonthlyAmount maps a *response.MerchantResponseMonthlyAmount to a *pb.MerchantResponseMonthlyAmount proto message.

It is used to generate the response for the MerchantService.GetMonthlyAmount rpc method.

```go
func (m *merchantProtoMapper) mapResponseMonthlyAmount(ms *response.MerchantResponseMonthlyAmount) *pb.MerchantResponseMonthlyAmount
```

##### `mapResponseMonthlyPaymentMethod`

mapResponseMonthlyPaymentMethod maps a *response.MerchantResponseMonthlyPaymentMethod to a *pb.MerchantResponseMonthlyPaymentMethod proto message.

It is used to generate the response for the MerchantService.GetMonthlyPaymentMethod rpc method.

```go
func (m *merchantProtoMapper) mapResponseMonthlyPaymentMethod(ms *response.MerchantResponseMonthlyPaymentMethod) *pb.MerchantResponseMonthlyPaymentMethod
```

##### `mapResponseMonthlyTotalAmount`

mapResponseMonthlyTotalAmount maps a *response.MerchantResponseMonthlyTotalAmount to a *pb.MerchantResponseMonthlyTotalAmount proto message.

It is used to generate the response for the MerchantService.GetMonthlyTotalAmount rpc method.

```go
func (m *merchantProtoMapper) mapResponseMonthlyTotalAmount(ms *response.MerchantResponseMonthlyTotalAmount) *pb.MerchantResponseMonthlyTotalAmount
```

##### `mapResponseYearlyAmount`

mapResponseYearlyAmount maps a *response.MerchantResponseYearlyAmount to a *pb.MerchantResponseYearlyAmount proto message.

It is used to generate the response for the MerchantService.GetYearlyAmount rpc method.

```go
func (m *merchantProtoMapper) mapResponseYearlyAmount(ms *response.MerchantResponseYearlyAmount) *pb.MerchantResponseYearlyAmount
```

##### `mapResponseYearlyPaymentMethod`

mapResponseYearlyPaymentMethod maps a *response.MerchantResponseYearlyPaymentMethod to a
*pb.MerchantResponseYearlyPaymentMethod proto message.

It is used to generate the response for the MerchantService.GetYearlyPaymentMethod rpc method.

```go
func (m *merchantProtoMapper) mapResponseYearlyPaymentMethod(ms *response.MerchantResponseYearlyPaymentMethod) *pb.MerchantResponseYearlyPaymentMethod
```

##### `mapResponseYearlyTotalAmount`

mapResponseYearlyTotalAmount maps a *response.MerchantResponseYearlyTotalAmount to a
*pb.MerchantResponseYearlyTotalAmount proto message.

It is used to generate the response for the MerchantService.GetYearlyTotalAmount rpc method.

```go
func (m *merchantProtoMapper) mapResponseYearlyTotalAmount(ms *response.MerchantResponseYearlyTotalAmount) *pb.MerchantResponseYearlyTotalAmount
```

##### `mapResponsesMonthlyAmount`

mapResponsesMonthlyAmount maps a list of *response.MerchantResponseMonthlyAmount to a list of
*pb.MerchantResponseMonthlyAmount proto responses.

It iterates over each MerchantResponseMonthlyAmount in the input slice, converting
them to their protobuf equivalent using the mapResponseMonthlyAmount function. This
function is used to generate the response data for monthly amount RPC methods.

```go
func (s *merchantProtoMapper) mapResponsesMonthlyAmount(roles []*response.MerchantResponseMonthlyAmount) []*pb.MerchantResponseMonthlyAmount
```

##### `mapResponsesMonthlyPaymentMethod`

mapResponsesMonthlyPaymentMethod maps a list of *response.MerchantResponseMonthlyPaymentMethod to a list of
*pb.MerchantResponseMonthlyPaymentMethod proto responses.

It iterates over each MerchantResponseMonthlyPaymentMethod in the input slice, converting
them to their protobuf equivalent using the mapResponseMonthlyPaymentMethod function. This
function is used to generate the response data for monthly payment method RPC methods.

```go
func (s *merchantProtoMapper) mapResponsesMonthlyPaymentMethod(roles []*response.MerchantResponseMonthlyPaymentMethod) []*pb.MerchantResponseMonthlyPaymentMethod
```

##### `mapResponsesMonthlyTotalAmount`

mapResponsesMonthlyTotalAmount maps a list of *response.MerchantResponseMonthlyTotalAmount to a list of
*pb.MerchantResponseMonthlyTotalAmount proto responses.

It iterates over each MerchantResponseMonthlyTotalAmount in the input slice, converting
them to their protobuf equivalent using the mapResponseMonthlyTotalAmount function.
This function is used to generate the response data for monthly total amount RPC methods.

```go
func (s *merchantProtoMapper) mapResponsesMonthlyTotalAmount(roles []*response.MerchantResponseMonthlyTotalAmount) []*pb.MerchantResponseMonthlyTotalAmount
```

##### `mapResponsesYearlyAmount`

mapResponsesYearlyAmount maps a list of *response.MerchantResponseYearlyAmount to a list of
*pb.MerchantResponseYearlyAmount proto responses.

It iterates over each MerchantResponseYearlyAmount in the input slice, converting
them to their protobuf equivalent using the mapResponseYearlyAmount function.
This function is used to generate the response data for yearly amount RPC methods.

```go
func (s *merchantProtoMapper) mapResponsesYearlyAmount(roles []*response.MerchantResponseYearlyAmount) []*pb.MerchantResponseYearlyAmount
```

##### `mapResponsesYearlyPaymentMethod`

mapResponsesYearlyPaymentMethod maps a list of *response.MerchantResponseYearlyPaymentMethod to a list of
*pb.MerchantResponseYearlyPaymentMethod proto responses.

It iterates over each MerchantResponseYearlyPaymentMethod in the input slice, converting
them to their protobuf equivalent using the mapResponseYearlyPaymentMethod function. This
function is used to generate the response data for yearly payment method RPC methods.

```go
func (s *merchantProtoMapper) mapResponsesYearlyPaymentMethod(roles []*response.MerchantResponseYearlyPaymentMethod) []*pb.MerchantResponseYearlyPaymentMethod
```

##### `mapResponsesYearlyTotalAmount`

mapResponsesYearlyTotalAmount maps a list of *response.MerchantResponseYearlyTotalAmount to a list of
*pb.MerchantResponseYearlyTotalAmount proto responses.

It iterates over each MerchantResponseYearlyTotalAmount in the input slice, converting
them to their protobuf equivalent using the mapResponseYearlyTotalAmount function.
This function is used to generate the response data for yearly total amount RPC methods.

```go
func (s *merchantProtoMapper) mapResponsesYearlyTotalAmount(roles []*response.MerchantResponseYearlyTotalAmount) []*pb.MerchantResponseYearlyTotalAmount
```

### `roleProtoMapper`

```go
type roleProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponsePaginationRole`

ToProtoResponsePaginationRole maps paginated role responses to a protobuf message.

Args:
  - pagination: A pointer to the pagination metadata for the response.
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.
  - pbResponse: A slice of pointers to RoleResponse to be mapped.

Returns:
  - A pointer to a pb.ApiResponsePaginationRole containing the mapped data.

```go
func (s *roleProtoMapper) ToProtoResponsePaginationRole(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsePaginationRole
```

##### `ToProtoResponsePaginationRoleDeleteAt`

ToProtoResponsePaginationRoleDeleteAt maps paginated soft-deleted role responses to a protobuf message.

Args:
  - pagination: A pointer to the pagination metadata for the response.
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.
  - pbResponse: A slice of pointers to RoleResponseDeleteAt to be mapped.

Returns:
  - A pointer to a pb.ApiResponsePaginationRoleDeleteAt containing the mapped data.

```go
func (s *roleProtoMapper) ToProtoResponsePaginationRoleDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponseDeleteAt) *pb.ApiResponsePaginationRoleDeleteAt
```

##### `ToProtoResponseRole`

ToProtoResponseRole converts a single role response to a protobuf message.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.
  - pbResponse: A pointer to the role response to be mapped.

Returns:
  - A pointer to a pb.ApiResponseRole containing the mapped data.

```go
func (s *roleProtoMapper) ToProtoResponseRole(status string, message string, pbResponse *response.RoleResponse) *pb.ApiResponseRole
```

##### `ToProtoResponseRoleAll`

ToProtoResponseRoleAll returns a protobuf message for all roles without pagination.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.

Returns:
  - A pointer to a pb.ApiResponseRoleAll containing the mapped data.

```go
func (s *roleProtoMapper) ToProtoResponseRoleAll(status string, message string) *pb.ApiResponseRoleAll
```

##### `ToProtoResponseRoleDelete`

ToProtoResponseRoleDelete returns a protobuf message indicating a role has been deleted.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.

Returns:
  - A pointer to a pb.ApiResponseRoleDelete containing the status and message.

```go
func (s *roleProtoMapper) ToProtoResponseRoleDelete(status string, message string) *pb.ApiResponseRoleDelete
```

##### `ToProtoResponsesRole`

ToProtoResponsesRole converts a list of role responses to a protobuf message.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.
  - pbResponse: A slice of pointers to RoleResponse to be mapped.

Returns:
  - A pointer to a pb.ApiResponsesRole containing the mapped data.

```go
func (s *roleProtoMapper) ToProtoResponsesRole(status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsesRole
```

##### `mapResponseRole`

mapResponseRole maps a single role response to a protobuf message.

Args:
  - role: A pointer to the role response to be mapped.

Returns:
  - A pointer to a pb.RoleResponse containing the mapped data.

```go
func (s *roleProtoMapper) mapResponseRole(role *response.RoleResponse) *pb.RoleResponse
```

##### `mapResponseRoleDeleteAt`

mapResponseRoleDeleteAt maps a RoleResponseDeleteAt to its protobuf representation.

Args:
  - role: A pointer to a RoleResponseDeleteAt containing the details of the role.

Returns:
  - A pointer to a pb.RoleResponseDeleteAt containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.

```go
func (s *roleProtoMapper) mapResponseRoleDeleteAt(role *response.RoleResponseDeleteAt) *pb.RoleResponseDeleteAt
```

##### `mapResponsesRole`

mapResponsesRole maps a list of *response.RoleResponse to a list of
*pb.RoleResponse proto responses.

It iterates over each RoleResponse in the input slice, converting
them to their protobuf equivalent using the mapResponseRole function.
This function is used to generate the response data for role RPC methods.

```go
func (s *roleProtoMapper) mapResponsesRole(roles []*response.RoleResponse) []*pb.RoleResponse
```

##### `mapResponsesRoleDeleteAt`

mapResponsesRoleDeleteAt maps a list of *response.RoleResponseDeleteAt to a list of
*pb.RoleResponseDeleteAt proto responses.

It iterates over each RoleResponseDeleteAt in the input slice, converting
them to their protobuf equivalent using the mapResponseRoleDeleteAt function.
This function is used to generate the response data for the RoleService.ListDeletedRole rpc method.

```go
func (s *roleProtoMapper) mapResponsesRoleDeleteAt(roles []*response.RoleResponseDeleteAt) []*pb.RoleResponseDeleteAt
```

### `saldoProtoMapper`

```go
type saldoProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponseMonthSaldoBalances`

ToProtoResponseMonthSaldoBalances maps a status, message and a list of *response.SaldoMonthBalanceResponse
to a *pb.ApiResponseMonthSaldoBalances proto response.

It is used to generate the response for the SaldoService.GetMonthlySaldoBalances rpc method.

```go
func (s *saldoProtoMapper) ToProtoResponseMonthSaldoBalances(status string, message string, pbResponse []*response.SaldoMonthBalanceResponse) *pb.ApiResponseMonthSaldoBalances
```

##### `ToProtoResponseMonthTotalSaldo`

ToProtoResponseMonthTotalSaldo maps a status, message and a list of *response.SaldoMonthTotalBalanceResponse
to a *pb.ApiResponseMonthTotalSaldo proto response.

It is used to generate the response for the SaldoService.GetMonthlyTotalSaldoBalance rpc method.

```go
func (s *saldoProtoMapper) ToProtoResponseMonthTotalSaldo(status string, message string, pbResponse []*response.SaldoMonthTotalBalanceResponse) *pb.ApiResponseMonthTotalSaldo
```

##### `ToProtoResponsePaginationSaldo`

ToProtoResponsePaginationSaldo maps a status, message, pagination meta, and a list of *response.SaldoResponse
to a *pb.ApiResponsePaginationSaldo proto response.

It is used to generate the response for the SaldoService.ListSaldo rpc method.

Args:

	pagination: The pagination data for the paginated response.
	status: The status of the response.
	message: The message accompanying the response.
	pbResponse: A slice of SaldoResponse to be converted.

Returns:

	A pointer to an ApiResponsePaginationSaldo containing the status, message, data, and pagination data.

```go
func (s *saldoProtoMapper) ToProtoResponsePaginationSaldo(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsePaginationSaldo
```

##### `ToProtoResponsePaginationSaldoDeleteAt`

ToProtoResponsePaginationSaldoDeleteAt creates a new instance of ApiResponsePaginationSaldoDeleteAt.

It maps the provided pagination metadata, status, message, and a list of SaldoResponseDeleteAt
to its protobuf representation.

Args:

	pagination: The pagination data for the paginated response.
	status: The status of the API response.
	message: A descriptive message of the API response.
	pbResponse: The list of SaldoResponseDeleteAt that needs to be converted.

Returns:

	A pointer to ApiResponsePaginationSaldoDeleteAt containing the status, message, saldo data, and pagination data.

```go
func (s *saldoProtoMapper) ToProtoResponsePaginationSaldoDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponseDeleteAt) *pb.ApiResponsePaginationSaldoDeleteAt
```

##### `ToProtoResponseSaldo`

ToProtoResponseSaldo maps a SaldoResponse to an ApiResponseSaldo proto message.

Args:

	status: The status of the response.
	message: The message accompanying the response.
	pbResponse: The SaldoResponse to be converted.

Returns:

	A pointer to an ApiResponseSaldo containing the mapped data.

```go
func (s *saldoProtoMapper) ToProtoResponseSaldo(status string, message string, pbResponse *response.SaldoResponse) *pb.ApiResponseSaldo
```

##### `ToProtoResponseSaldoAll`

ToProtoResponseSaldoAll maps a status and message to an ApiResponseSaldoAll proto response.

Args:

	status: The status of the response.
	message: The message accompanying the response.

Returns:

	A pointer to an ApiResponseSaldoAll containing the mapped data.

```go
func (s *saldoProtoMapper) ToProtoResponseSaldoAll(status string, message string) *pb.ApiResponseSaldoAll
```

##### `ToProtoResponseSaldoDelete`

ToProtoResponseSaldoDelete maps a status and message to an ApiResponseSaldoDelete proto response.

It is used to generate the response for the SaldoService.DeleteSaldo rpc method.

```go
func (s *saldoProtoMapper) ToProtoResponseSaldoDelete(status string, message string) *pb.ApiResponseSaldoDelete
```

##### `ToProtoResponseYearSaldoBalances`

ToProtoResponseYearSaldoBalances maps a status, message, and a list of
*response.SaldoYearBalanceResponse to a *pb.ApiResponseYearSaldoBalances proto response.

It is used to generate the response for the SaldoService.GetYearlySaldoBalances rpc method.

Args:

	status: The status of the response.
	message: The message accompanying the response.
	pbResponse: A slice of SaldoYearBalanceResponse to be converted.

Returns:

	A pointer to an ApiResponseYearSaldoBalances containing the mapped data.

```go
func (s *saldoProtoMapper) ToProtoResponseYearSaldoBalances(status string, message string, pbResponse []*response.SaldoYearBalanceResponse) *pb.ApiResponseYearSaldoBalances
```

##### `ToProtoResponseYearTotalSaldo`

ToProtoResponseYearTotalSaldo maps a status, message and a list of *response.SaldoYearTotalBalanceResponse
to a *pb.ApiResponseYearTotalSaldo proto response.

It is used to generate the response for the SaldoService.GetYearlyTotalSaldoBalance rpc method.

```go
func (s *saldoProtoMapper) ToProtoResponseYearTotalSaldo(status string, message string, pbResponse []*response.SaldoYearTotalBalanceResponse) *pb.ApiResponseYearTotalSaldo
```

##### `ToProtoResponsesSaldo`

ToProtoResponsesSaldo maps a list of SaldoResponse to an ApiResponsesSaldo proto message.

Args:

	status: The status of the response.
	message: The message accompanying the response.
	pbResponse: A slice of SaldoResponse to be converted.

Returns:

	A pointer to an ApiResponsesSaldo containing the mapped data.

```go
func (s *saldoProtoMapper) ToProtoResponsesSaldo(status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsesSaldo
```

##### `mapResponseSaldo`

mapResponseSaldo maps a SaldoResponse to its protobuf representation.

Args:

	saldo: The SaldoResponse to be converted.

Returns:

	A pointer to SaldoResponse containing the mapped data.

```go
func (s *saldoProtoMapper) mapResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse
```

##### `mapResponseSaldoDeleteAt`

mapResponseSaldoDeleteAt maps a single response.SaldoResponseDeleteAt to a single *pb.SaldoResponseDeleteAt proto response.

Args:

	saldo: The SaldoResponseDeleteAt to be converted.

Returns:

	A pointer to a SaldoResponseDeleteAt containing the mapped data.

```go
func (s *saldoProtoMapper) mapResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *pb.SaldoResponseDeleteAt
```

##### `mapResponsesSaldo`

mapResponsesSaldo maps a list of SaldoResponse to a list of *pb.SaldoResponse proto responses.

It iterates over each SaldoResponse in the input slice, converting
them to their protobuf equivalent using the mapResponseSaldo function.
This function is used to generate the response data for ListSaldo rpc methods.

```go
func (s *saldoProtoMapper) mapResponsesSaldo(saldos []*response.SaldoResponse) []*pb.SaldoResponse
```

##### `mapResponsesSaldoDeleteAt`

mapResponsesSaldoDeleteAt maps a list of *response.SaldoResponseDeleteAt to a list of
*pb.SaldoResponseDeleteAt proto responses.

It iterates over each SaldoResponseDeleteAt in the input slice, converting
them to their protobuf equivalent using the mapResponseSaldoDeleteAt function.
This function is used to generate the response data for Saldo-related RPC methods
involving deleted saldo records.

```go
func (s *saldoProtoMapper) mapResponsesSaldoDeleteAt(saldos []*response.SaldoResponseDeleteAt) []*pb.SaldoResponseDeleteAt
```

##### `mapSaldoMonthBalanceResponse`

mapSaldoMonthBalanceResponse maps a *response.SaldoMonthBalanceResponse to a *pb.SaldoMonthBalanceResponse proto response.

It takes a SaldoMonthBalanceResponse from the response domain model and converts it
into its protobuf representation, ensuring that the month and total balance fields are accurately mapped.

Args:

	ss: The SaldoMonthBalanceResponse to be converted.

Returns:

	A pointer to a SaldoMonthBalanceResponse containing the mapped data.

```go
func (s *saldoProtoMapper) mapSaldoMonthBalanceResponse(ss *response.SaldoMonthBalanceResponse) *pb.SaldoMonthBalanceResponse
```

##### `mapSaldoMonthBalanceResponses`

mapSaldoMonthBalanceResponses maps a list of *response.SaldoMonthBalanceResponse
to a list of *pb.SaldoMonthBalanceResponse proto responses.

It iterates over each SaldoMonthBalanceResponse in the input slice, converting
them to their protobuf equivalent using the mapSaldoMonthBalanceResponse function.
This function is used to generate the response data for monthly balance RPC methods.

```go
func (s *saldoProtoMapper) mapSaldoMonthBalanceResponses(ss []*response.SaldoMonthBalanceResponse) []*pb.SaldoMonthBalanceResponse
```

##### `mapSaldoMonthTotalBalanceResponse`

mapSaldoMonthTotalBalanceResponse maps a *response.SaldoMonthTotalBalanceResponse to a *pb.SaldoMonthTotalBalanceResponse.

It maps the fields of the input response to the corresponding fields of the output response.
The TotalBalance field is converted to int32.

Args:

	ss: The SaldoMonthTotalBalanceResponse to be converted.

Returns:

	A pointer to a SaldoMonthTotalBalanceResponse containing the mapped data.

```go
func (s *saldoProtoMapper) mapSaldoMonthTotalBalanceResponse(ss *response.SaldoMonthTotalBalanceResponse) *pb.SaldoMonthTotalBalanceResponse
```

##### `mapSaldoMonthTotalBalanceResponses`

mapSaldoMonthTotalBalanceResponses maps a list of *response.SaldoMonthTotalBalanceResponse
to a list of *pb.SaldoMonthTotalBalanceResponse proto responses.

It iterates over each SaldoMonthTotalBalanceResponse in the input slice, converting
them to their protobuf equivalent using the mapSaldoMonthTotalBalanceResponse function.
This function is used to generate the response data for monthly total saldo balance RPC methods.

Args:

	ss: A slice of SaldoMonthTotalBalanceResponse to be converted.

Returns:

	A slice of SaldoMonthTotalBalanceResponse containing the mapped data.

```go
func (s *saldoProtoMapper) mapSaldoMonthTotalBalanceResponses(ss []*response.SaldoMonthTotalBalanceResponse) []*pb.SaldoMonthTotalBalanceResponse
```

##### `mapSaldoYearBalanceResponse`

mapSaldoYearBalanceResponse maps a *response.SaldoYearBalanceResponse to a *pb.SaldoYearBalanceResponse proto response.

It takes a SaldoYearBalanceResponse from the response domain model and converts it
into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.

Args:

	ss: The SaldoYearBalanceResponse to be converted.

Returns:

	A pointer to SaldoYearBalanceResponse containing the mapped data.

```go
func (s *saldoProtoMapper) mapSaldoYearBalanceResponse(ss *response.SaldoYearBalanceResponse) *pb.SaldoYearBalanceResponse
```

##### `mapSaldoYearBalanceResponses`

mapSaldoYearBalanceResponses maps a list of *response.SaldoYearBalanceResponse
to a list of *pb.SaldoYearBalanceResponse proto responses.

It iterates over each SaldoYearBalanceResponse in the input slice, converting
them to their protobuf equivalent using the mapSaldoYearBalanceResponse function.
This function is used to generate the response data for yearly balance RPC methods.

```go
func (s *saldoProtoMapper) mapSaldoYearBalanceResponses(ss []*response.SaldoYearBalanceResponse) []*pb.SaldoYearBalanceResponse
```

##### `mapSaldoYearTotalBalanceResponse`

mapSaldoYearTotalBalanceResponse maps a *response.SaldoYearTotalBalanceResponse to a *pb.SaldoYearTotalBalanceResponse proto response.

It takes a SaldoYearTotalBalanceResponse from the response domain model and converts it
into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.

Args:

	ss: The SaldoYearTotalBalanceResponse to be converted.

Returns:

	A pointer to a SaldoYearTotalBalanceResponse containing the mapped data.

```go
func (s *saldoProtoMapper) mapSaldoYearTotalBalanceResponse(ss *response.SaldoYearTotalBalanceResponse) *pb.SaldoYearTotalBalanceResponse
```

##### `mapSaldoYearTotalBalanceResponses`

mapSaldoYearTotalBalanceResponses maps a list of *response.SaldoYearTotalBalanceResponse
to a list of *pb.SaldoYearTotalBalanceResponse proto responses.

It iterates over each SaldoYearTotalBalanceResponse in the input slice, converting
them to their protobuf equivalent using the mapSaldoYearTotalBalanceResponse function.
This function is used to generate the response data for yearly total saldo balance RPC methods.

```go
func (s *saldoProtoMapper) mapSaldoYearTotalBalanceResponses(ss []*response.SaldoYearTotalBalanceResponse) []*pb.SaldoYearTotalBalanceResponse
```

### `topupProtoMapper`

topupProtoMapper is responsible for mapping Topup responses to their corresponding protobuf representations.

```go
type topupProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponsePaginationTopup`

ToProtoResponsePaginationTopup maps paginated top-up records to a protobuf response.

Args:
  - pagination: The pagination metadata for the response.
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupResponse that needs to be converted.

Returns:

	A pointer to ApiResponsePaginationTopup containing the status, message, top-up data, and pagination data.

```go
func (t *topupProtoMapper) ToProtoResponsePaginationTopup(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponse) *pb.ApiResponsePaginationTopup
```

##### `ToProtoResponsePaginationTopupDeleteAt`

ToProtoResponsePaginationTopupDeleteAt maps paginated soft-deleted top-up records to a protobuf response.

Args:
  - pagination: The pagination metadata for the response.
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupResponseDeleteAt that needs to be converted.

Returns:

	A pointer to ApiResponsePaginationTopupDeleteAt containing the status, message, top-up data, and pagination data.

```go
func (t *topupProtoMapper) ToProtoResponsePaginationTopupDeleteAt(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponseDeleteAt) *pb.ApiResponsePaginationTopupDeleteAt
```

##### `ToProtoResponseTopup`

ToProtoResponseTopup maps a single top-up record to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A pointer to TopupResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTopup containing the status, message, and top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopup(status string, message string, s *response.TopupResponse) *pb.ApiResponseTopup
```

##### `ToProtoResponseTopupAll`

ToProtoResponseTopupAll returns all top-up records in a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.

Returns:
  - A pointer to ApiResponseTopupAll containing the status and message.

```go
func (t topupProtoMapper) ToProtoResponseTopupAll(status string, message string) *pb.ApiResponseTopupAll
```

##### `ToProtoResponseTopupDeletAt`

ToProtoResponseTopupDeletAt maps a soft-deleted top-up record to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A pointer to TopupResponseDeleteAt that needs to be converted.

Returns:
  - A pointer to ApiResponseTopupDeleteAt containing the status, message, and mapped top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupDeletAt(status string, message string, s *response.TopupResponseDeleteAt) *pb.ApiResponseTopupDeleteAt
```

##### `ToProtoResponseTopupDelete`

ToProtoResponseTopupDelete returns a response indicating a top-up record has been deleted.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.

Returns:

	A pointer to ApiResponseTopupDelete containing the status and message.

```go
func (t topupProtoMapper) ToProtoResponseTopupDelete(status string, message string) *pb.ApiResponseTopupDelete
```

##### `ToProtoResponseTopupMonthAmount`

ToProtoResponseTopupMonthAmount maps monthly top-up amounts to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupMonthAmountResponse that needs to be converted.

Returns:
  - A pointer to ApiResponseTopupMonthAmount containing the status, message, and mapped top-up amount data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupMonthAmount(status string, message string, s []*response.TopupMonthAmountResponse) *pb.ApiResponseTopupMonthAmount
```

##### `ToProtoResponseTopupMonthMethod`

ToProtoResponseTopupMonthMethod maps monthly top-up methods to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupMonthMethodResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTopupMonthMethod containing the status, message, and top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupMonthMethod(status string, message string, s []*response.TopupMonthMethodResponse) *pb.ApiResponseTopupMonthMethod
```

##### `ToProtoResponseTopupMonthStatusFailed`

ToProtoResponseTopupMonthStatusFailed maps monthly failed top-ups to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupResponseMonthStatusFailed that needs to be converted.

Returns:

	A pointer to ApiResponseTopupMonthStatusFailed containing the status, message, and top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupMonthStatusFailed(status string, message string, s []*response.TopupResponseMonthStatusFailed) *pb.ApiResponseTopupMonthStatusFailed
```

##### `ToProtoResponseTopupMonthStatusSuccess`

ToProtoResponseTopupMonthStatusSuccess maps monthly successful top-ups to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupResponseMonthStatusSuccess that needs to be converted.

Returns:

	A pointer to ApiResponseTopupMonthStatusSuccess containing the status, message, and top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupMonthStatusSuccess(status string, message string, s []*response.TopupResponseMonthStatusSuccess) *pb.ApiResponseTopupMonthStatusSuccess
```

##### `ToProtoResponseTopupYearAmount`

ToProtoResponseTopupYearAmount maps yearly top-up amounts to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupYearlyAmountResponse that needs to be converted.

Returns:
  - A pointer to ApiResponseTopupYearAmount containing the status, message, and mapped top-up amount data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupYearAmount(status string, message string, s []*response.TopupYearlyAmountResponse) *pb.ApiResponseTopupYearAmount
```

##### `ToProtoResponseTopupYearMethod`

ToProtoResponseTopupYearMethod maps yearly top-up methods to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupYearlyMethodResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTopupYearMethod containing the status, message, and top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupYearMethod(status string, message string, s []*response.TopupYearlyMethodResponse) *pb.ApiResponseTopupYearMethod
```

##### `ToProtoResponseTopupYearStatusFailed`

ToProtoResponseTopupYearStatusFailed maps yearly failed top-ups to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupResponseYearStatusFailed that needs to be converted.

Returns:

	A pointer to ApiResponseTopupYearStatusFailed containing the status, message, and top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupYearStatusFailed(status string, message string, s []*response.TopupResponseYearStatusFailed) *pb.ApiResponseTopupYearStatusFailed
```

##### `ToProtoResponseTopupYearStatusSuccess`

ToProtoResponseTopupYearStatusSuccess maps yearly successful top-ups to a protobuf response.

Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - s: A slice of TopupResponseYearStatusSuccess that needs to be converted.

Returns:

	A pointer to ApiResponseTopupYearStatusSuccess containing the status, message, and top-up data.

```go
func (t *topupProtoMapper) ToProtoResponseTopupYearStatusSuccess(status string, message string, s []*response.TopupResponseYearStatusSuccess) *pb.ApiResponseTopupYearStatusSuccess
```

##### `mapResponseTopup`

mapResponseTopup maps a single TopupResponse to its corresponding protobuf response.

Args:
  - topup: A pointer to TopupResponse that needs to be converted.

Returns:
  - A pointer to TopupResponse containing the mapped top-up data.

```go
func (t *topupProtoMapper) mapResponseTopup(topup *response.TopupResponse) *pb.TopupResponse
```

##### `mapResponseTopupDeleteAt`

mapResponseTopupDeleteAt maps a single TopupResponseDeleteAt to its corresponding protobuf response.

Args:
  - topup: A pointer to TopupResponseDeleteAt that needs to be converted.

Returns:
  - A pointer to TopupResponseDeleteAt containing the mapped top-up data.

```go
func (t *topupProtoMapper) mapResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *pb.TopupResponseDeleteAt
```

##### `mapResponseTopupMonthStatusFailed`

mapResponseTopupMonthStatusFailed maps a TopupResponseMonthStatusFailed to a TopupMonthStatusFailedResponse protobuf message.

Args:
  - s: A pointer to TopupResponseMonthStatusFailed that needs to be converted.

Returns:
  - A pointer to TopupMonthStatusFailedResponse containing the mapped monthly top-up failure data.

```go
func (t *topupProtoMapper) mapResponseTopupMonthStatusFailed(s *response.TopupResponseMonthStatusFailed) *pb.TopupMonthStatusFailedResponse
```

##### `mapResponseTopupMonthStatusSuccess`

mapResponseTopupMonthStatusSuccess maps a TopupResponseMonthStatusSuccess to a TopupMonthStatusSuccessResponse.

Args:
  - s: A pointer to TopupResponseMonthStatusSuccess that needs to be converted.

Returns:
  - A pointer to TopupMonthStatusSuccessResponse containing the mapped top-up data.

```go
func (t *topupProtoMapper) mapResponseTopupMonthStatusSuccess(s *response.TopupResponseMonthStatusSuccess) *pb.TopupMonthStatusSuccessResponse
```

##### `mapResponseTopupMonthlyAmount`

mapResponseTopupMonthlyAmount maps a TopupMonthAmountResponse to a TopupMonthAmountResponse protobuf message.

Args:
  - s: A pointer to TopupMonthAmountResponse that needs to be converted.

Returns:
  - A pointer to TopupMonthAmountResponse containing the mapped monthly top-up amount data.

```go
func (t *topupProtoMapper) mapResponseTopupMonthlyAmount(s *response.TopupMonthAmountResponse) *pb.TopupMonthAmountResponse
```

##### `mapResponseTopupMonthlyAmounts`

mapResponseTopupMonthlyAmounts maps a slice of TopupMonthAmountResponse to a slice of
TopupMonthAmountResponse protobuf messages.

Args:
  - s: A slice of TopupMonthAmountResponse that needs to be converted.

Returns:
  - A slice of TopupMonthAmountResponse containing the mapped monthly top-up amount data.

```go
func (t *topupProtoMapper) mapResponseTopupMonthlyAmounts(s []*response.TopupMonthAmountResponse) []*pb.TopupMonthAmountResponse
```

##### `mapResponseTopupMonthlyMethod`

mapResponseTopupMonthlyMethod maps a TopupMonthMethodResponse to a TopupMonthMethodResponse protobuf message.

Args:
  - s: A pointer to TopupMonthMethodResponse that needs to be converted.

Returns:
  - A pointer to TopupMonthMethodResponse containing the mapped monthly top-up method data.

```go
func (t *topupProtoMapper) mapResponseTopupMonthlyMethod(s *response.TopupMonthMethodResponse) *pb.TopupMonthMethodResponse
```

##### `mapResponseTopupMonthlyMethods`

mapResponseTopupMonthlyMethods maps a slice of TopupMonthMethodResponse to a slice of
TopupMonthMethodResponse protobuf messages.

Args:
  - s: A slice of TopupMonthMethodResponse that needs to be converted.

Returns:
  - A slice of TopupMonthMethodResponse containing the mapped monthly top-up method data.

```go
func (t *topupProtoMapper) mapResponseTopupMonthlyMethods(s []*response.TopupMonthMethodResponse) []*pb.TopupMonthMethodResponse
```

##### `mapResponseTopupYearlyAmount`

mapResponseTopupYearlyAmount maps a TopupYearlyAmountResponse to a TopupYearlyAmountResponse protobuf message.

Args:
  - s: A pointer to TopupYearlyAmountResponse that needs to be converted.

Returns:
  - A pointer to TopupYearlyAmountResponse containing the mapped yearly top-up amount data.

```go
func (t *topupProtoMapper) mapResponseTopupYearlyAmount(s *response.TopupYearlyAmountResponse) *pb.TopupYearlyAmountResponse
```

##### `mapResponseTopupYearlyAmounts`

mapResponseTopupYearlyAmounts maps a slice of TopupYearlyAmountResponse to a slice of
TopupYearlyAmountResponse protobuf messages.

Args:
  - s: A slice of TopupYearlyAmountResponse that needs to be converted.

Returns:
  - A slice of TopupYearlyAmountResponse containing the mapped yearly top-up amount data.

```go
func (t *topupProtoMapper) mapResponseTopupYearlyAmounts(s []*response.TopupYearlyAmountResponse) []*pb.TopupYearlyAmountResponse
```

##### `mapResponseTopupYearlyMethod`

mapResponseTopupYearlyMethod maps a TopupYearlyMethodResponse to a TopupYearlyMethodResponse protobuf message.

Args:
  - s: A pointer to TopupYearlyMethodResponse that needs to be converted.

Returns:
  - A pointer to TopupYearlyMethodResponse containing the mapped yearly top-up method data.

```go
func (t *topupProtoMapper) mapResponseTopupYearlyMethod(s *response.TopupYearlyMethodResponse) *pb.TopupYearlyMethodResponse
```

##### `mapResponseTopupYearlyMethods`

mapResponseTopupYearlyMethods maps a slice of TopupYearlyMethodResponse to a slice of
TopupYearlyMethodResponse protobuf messages.

Args:
  - s: A slice of TopupYearlyMethodResponse that needs to be converted.

Returns:
  - A slice of TopupYearlyMethodResponse containing the mapped yearly top-up method data.

```go
func (t *topupProtoMapper) mapResponseTopupYearlyMethods(s []*response.TopupYearlyMethodResponse) []*pb.TopupYearlyMethodResponse
```

##### `mapResponsesTopup`

mapResponsesTopup maps a slice of TopupResponse to a slice of TopupResponse.

Args:
  - topups: A slice of TopupResponse that needs to be converted.

Returns:
  - A slice of TopupResponse containing the mapped top-up data.

```go
func (t *topupProtoMapper) mapResponsesTopup(topups []*response.TopupResponse) []*pb.TopupResponse
```

##### `mapResponsesTopupDeleteAt`

mapResponsesTopupDeleteAt maps a slice of TopupResponseDeleteAt to a slice of TopupResponseDeleteAt.

Args:
  - topups: A slice of TopupResponseDeleteAt that needs to be converted.

Returns:
  - A slice of TopupResponseDeleteAt containing the mapped top-up data.

```go
func (t *topupProtoMapper) mapResponsesTopupDeleteAt(topups []*response.TopupResponseDeleteAt) []*pb.TopupResponseDeleteAt
```

##### `mapResponsesTopupMonthStatusFailed`

mapResponsesTopupMonthStatusFailed converts a slice of TopupResponseMonthStatusFailed
into a slice of TopupMonthStatusFailedResponse protobuf messages.

Args:
  - topups: A slice of TopupResponseMonthStatusFailed to be converted.

Returns:
  - A slice of TopupMonthStatusFailedResponse containing the mapped data
    for each monthly top-up failure entry.

```go
func (t *topupProtoMapper) mapResponsesTopupMonthStatusFailed(topups []*response.TopupResponseMonthStatusFailed) []*pb.TopupMonthStatusFailedResponse
```

##### `mapResponsesTopupMonthStatusSuccess`

mapResponsesTopupMonthStatusSuccess maps a slice of TopupResponseMonthStatusSuccess to a slice of
TopupMonthStatusSuccessResponse protobuf messages.

Args:
  - topups: A slice of TopupResponseMonthStatusSuccess that needs to be converted.

Returns:
  - A slice of TopupMonthStatusSuccessResponse containing the mapped monthly top-up success data.

```go
func (t *topupProtoMapper) mapResponsesTopupMonthStatusSuccess(topups []*response.TopupResponseMonthStatusSuccess) []*pb.TopupMonthStatusSuccessResponse
```

##### `mapTopupResponseYearStatusFailed`

mapTopupResponseYearStatusFailed maps a TopupResponseYearStatusFailed to a TopupYearStatusFailedResponse.

Args:
  - s: A pointer to TopupResponseYearStatusFailed that needs to be converted.

Returns:
  - A pointer to TopupYearStatusFailedResponse containing the mapped yearly top-up failure data.

```go
func (t *topupProtoMapper) mapTopupResponseYearStatusFailed(s *response.TopupResponseYearStatusFailed) *pb.TopupYearStatusFailedResponse
```

##### `mapTopupResponseYearStatusSuccess`

mapTopupResponseYearStatusSuccess maps a TopupResponseYearStatusSuccess to a TopupYearStatusSuccessResponse.

Args:
  - s: A pointer to TopupResponseYearStatusSuccess that needs to be converted.

Returns:
  - A pointer to TopupYearStatusSuccessResponse containing the mapped yearly top-up success data.

```go
func (t *topupProtoMapper) mapTopupResponseYearStatusSuccess(s *response.TopupResponseYearStatusSuccess) *pb.TopupYearStatusSuccessResponse
```

##### `mapTopupResponsesYearStatusFailed`

mapTopupResponsesYearStatusFailed maps a slice of TopupResponseYearStatusFailed
into a slice of TopupYearStatusFailedResponse protobuf messages.

Args:
  - topups: A slice of TopupResponseYearStatusFailed to be converted.

Returns:
  - A slice of TopupYearStatusFailedResponse containing the mapped yearly top-up failure data
    for each yearly top-up failure entry.

```go
func (t *topupProtoMapper) mapTopupResponsesYearStatusFailed(topups []*response.TopupResponseYearStatusFailed) []*pb.TopupYearStatusFailedResponse
```

##### `mapTopupResponsesYearStatusSuccess`

mapTopupResponsesYearStatusSuccess maps a slice of TopupResponseYearStatusSuccess to a slice of
TopupYearStatusSuccessResponse protobuf messages.

Args:
  - topups: A slice of TopupResponseYearStatusSuccess that needs to be converted.

Returns:
  - A slice of TopupYearStatusSuccessResponse containing the mapped yearly top-up success data.

```go
func (t *topupProtoMapper) mapTopupResponsesYearStatusSuccess(topups []*response.TopupResponseYearStatusSuccess) []*pb.TopupYearStatusSuccessResponse
```

### `transactionProtoMapper`

transactionProtoMapper is responsible for mapping Transaction responses to their corresponding protobuf representations.

```go
type transactionProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponsePaginationTransaction`

ToProtoResponsePaginationTransaction maps a pagination meta, status, message and a list of *response.TransactionResponse
to a *pb.ApiResponsePaginationTransaction proto response.

It is used to generate the response for the TransactionService.ListTransaction rpc method.

```go
func (m *transactionProtoMapper) ToProtoResponsePaginationTransaction(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction
```

##### `ToProtoResponsePaginationTransactionDeleteAt`

ToProtoResponsePaginationTransactionDeleteAt maps a pagination meta, status, message and a list of *response.TransactionResponseDeleteAt
to a *pb.ApiResponsePaginationTransactionDeleteAt proto response.

It is used to generate the response for the TransactionService.ListTransactionDeleteAt rpc method.

```go
func (m *transactionProtoMapper) ToProtoResponsePaginationTransactionDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt
```

##### `ToProtoResponseTransaction`

ToProtoResponseTransaction converts a status, message, and a TransactionResponse
to its protobuf representation.

It includes the status, message, and mapped transaction data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The TransactionResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTransaction containing the status, message, and transaction data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction
```

##### `ToProtoResponseTransactionAll`

ToProtoResponseTransactionAll generates a protobuf API response for all transactions.

It includes the status and message for the API response.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.

Returns:

	A pointer to ApiResponseTransactionAll containing the status and message.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll
```

##### `ToProtoResponseTransactionDelete`

ToProtoResponseTransactionDelete generates a protobuf API response for a transaction delete operation.

It includes the status and message for the API response.
Args:
  - status: The status of the API response.
  - message: A descriptive message for the API response.

Returns:

	A pointer to ApiResponseTransactionDelete containing the status and message.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete
```

##### `ToProtoResponseTransactionMonthAmount`

ToProtoResponseTransactionMonthAmount converts a status, message, and a list of TransactionMonthAmountResponse
to its protobuf representation.

It includes the status, message, and mapped transaction month amount data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionMonthAmountResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionMonthAmount containing the status, message, and transaction month amount data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthAmount(status string, message string, pbResponse []*response.TransactionMonthAmountResponse) *pb.ApiResponseTransactionMonthAmount
```

##### `ToProtoResponseTransactionMonthMethod`

ToProtoResponseTransactionMonthMethod converts a status, message, and a list of TransactionMonthMethodResponse
to its protobuf representation.

It includes the status, message, and mapped transaction method data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionMonthMethodResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionMonthMethod containing the status, message, and transaction method data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthMethod(status string, message string, pbResponse []*response.TransactionMonthMethodResponse) *pb.ApiResponseTransactionMonthMethod
```

##### `ToProtoResponseTransactionMonthStatusFailed`

ToProtoResponseTransactionMonthStatusFailed converts a status, message, and a list of TransactionResponseMonthStatusFailed
to its protobuf representation.

It includes the status, message, and mapped transaction data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionResponseMonthStatusFailed that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionMonthStatusFailed containing the status, message, and transaction data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthStatusFailed(status string, message string, pbResponse []*response.TransactionResponseMonthStatusFailed) *pb.ApiResponseTransactionMonthStatusFailed
```

##### `ToProtoResponseTransactionMonthStatusSuccess`

ToProtoResponseTransactionMonthStatusSuccess converts a status, message, and a list of TransactionResponseMonthStatusSuccess
to its protobuf representation.

It includes the status, message, and mapped transaction data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionResponseMonthStatusSuccess that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionMonthStatusSuccess containing the status, message, and transaction data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseMonthStatusSuccess) *pb.ApiResponseTransactionMonthStatusSuccess
```

##### `ToProtoResponseTransactionYearAmount`

ToProtoResponseTransactionYearAmount converts a status, message, and a list of TransactionYearlyAmountResponse
to its protobuf representation.

It includes the status, message, and mapped transaction year amount data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionYearlyAmountResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionYearAmount containing the status, message, and transaction year amount data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionYearAmount(status string, message string, pbResponse []*response.TransactionYearlyAmountResponse) *pb.ApiResponseTransactionYearAmount
```

##### `ToProtoResponseTransactionYearMethod`

ToProtoResponseTransactionYearMethod converts a status, message, and a list of TransactionYearMethodResponse
to its protobuf representation.

It includes the status, message, and mapped transaction method data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionYearMethodResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionYearMethod containing the status, message, and transaction method data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionYearMethod(status string, message string, pbResponse []*response.TransactionYearMethodResponse) *pb.ApiResponseTransactionYearMethod
```

##### `ToProtoResponseTransactionYearStatusFailed`

ToProtoResponseTransactionYearStatusFailed converts a status, message, and a list of TransactionResponseYearStatusFailed
to its protobuf representation.

It includes the status, message, and mapped transaction data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionResponseYearStatusFailed that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionYearStatusFailed containing the status, message, and transaction data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionYearStatusFailed(status string, message string, pbResponse []*response.TransactionResponseYearStatusFailed) *pb.ApiResponseTransactionYearStatusFailed
```

##### `ToProtoResponseTransactionYearStatusSuccess`

ToProtoResponseTransactionYearStatusSuccess converts a status, message, and a list of TransactionResponseYearStatusSuccess
to its protobuf representation.

It includes the status, message, and mapped transaction data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionResponseYearStatusSuccess that needs to be converted.

Returns:

	A pointer to ApiResponseTransactionYearStatusSuccess containing the status, message, and transaction data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactionYearStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseYearStatusSuccess) *pb.ApiResponseTransactionYearStatusSuccess
```

##### `ToProtoResponseTransactions`

ToProtoResponseTransactions converts a status, message, and a list of TransactionResponse
to its protobuf representation.

It includes the status, message, and mapped transaction data.
Args:
  - status: The status of the API response.
  - message: A descriptive message of the API response.
  - pbResponse: The list of TransactionResponse that needs to be converted.

Returns:

	A pointer to ApiResponseTransactions containing the status, message, and transaction data.

```go
func (m *transactionProtoMapper) ToProtoResponseTransactions(status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponseTransactions
```

##### `mapResponseTransaction`

mapResponseTransaction maps a single response.TransactionResponse to its protobuf representation.

It takes a response.TransactionResponse as input and returns a pointer to the converted TransactionResponse proto object.

```go
func (m *transactionProtoMapper) mapResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse
```

##### `mapResponseTransactionDeleteAt`

mapResponseTransactionDeleteAt maps a single response.TransactionResponseDeleteAt to its protobuf representation.

It takes a response.TransactionResponseDeleteAt as input and returns a pointer to the converted TransactionResponseDeleteAt proto object.
The mapping includes fields like ID, TransactionNo, CardNumber, Amount, PaymentMethod, TransactionTime, MerchantId, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *transactionProtoMapper) mapResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt
```

##### `mapResponseTransactionMonthAmount`

mapResponseTransactionMonthAmount maps a *response.TransactionMonthAmountResponse to a *pb.TransactionMonthAmountResponse proto message.

It takes a *response.TransactionMonthAmountResponse as input and returns a pointer to the converted TransactionMonthAmountResponse proto object.
The mapping includes fields like Month and TotalAmount.

```go
func (m *transactionProtoMapper) mapResponseTransactionMonthAmount(s *response.TransactionMonthAmountResponse) *pb.TransactionMonthAmountResponse
```

##### `mapResponseTransactionMonthAmounts`

mapResponseTransactionMonthAmounts maps a slice of TransactionMonthAmountResponse
to a slice of protobuf TransactionMonthAmountResponse.

It iterates through each TransactionMonthAmountResponse, converting it to its
protobuf equivalent using the mapResponseTransactionMonthAmount function.

Args:
  - s: A slice of TransactionMonthAmountResponse to be converted.

Returns:
  - A slice of pointers to TransactionMonthAmountResponse containing the mapped data.

```go
func (m *transactionProtoMapper) mapResponseTransactionMonthAmounts(s []*response.TransactionMonthAmountResponse) []*pb.TransactionMonthAmountResponse
```

##### `mapResponseTransactionMonthMethod`

mapResponseTransactionMonthMethod maps a *response.TransactionMonthMethodResponse to a *pb.TransactionMonthMethodResponse proto message.

It takes a *response.TransactionMonthMethodResponse as input and returns a pointer to the converted TransactionMonthMethodResponse proto object.
The mapping includes fields like Month, PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (m *transactionProtoMapper) mapResponseTransactionMonthMethod(s *response.TransactionMonthMethodResponse) *pb.TransactionMonthMethodResponse
```

##### `mapResponseTransactionMonthMethods`

mapResponseTransactionMonthMethods maps a slice of TransactionMonthMethodResponse to a slice of protobuf TransactionMonthMethodResponse.

It takes a slice of TransactionMonthMethodResponse as input and returns a slice of corresponding protobuf TransactionMonthMethodResponse.
The mapping includes fields like Month, PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (m *transactionProtoMapper) mapResponseTransactionMonthMethods(s []*response.TransactionMonthMethodResponse) []*pb.TransactionMonthMethodResponse
```

##### `mapResponseTransactionMonthStatusFailed`

mapResponseTransactionMonthStatusFailed maps a response.TransactionResponseMonthStatusFailed to its protobuf representation.

It takes a response.TransactionResponseMonthStatusFailed as input and returns a pointer to the converted TransactionMonthStatusFailedResponse proto object.
The mapping includes fields like Year, Month, TotalFailed, and TotalAmount.

```go
func (t *transactionProtoMapper) mapResponseTransactionMonthStatusFailed(s *response.TransactionResponseMonthStatusFailed) *pb.TransactionMonthStatusFailedResponse
```

##### `mapResponseTransactionMonthStatusSuccess`

mapResponseTransactionMonthStatusSuccess maps a response.TransactionResponseMonthStatusSuccess to its protobuf representation.

It takes a response.TransactionResponseMonthStatusSuccess as input and returns a pointer to the converted TransactionMonthStatusSuccessResponse proto object.
The mapping includes fields like Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *transactionProtoMapper) mapResponseTransactionMonthStatusSuccess(s *response.TransactionResponseMonthStatusSuccess) *pb.TransactionMonthStatusSuccessResponse
```

##### `mapResponseTransactionYearMethod`

mapResponseTransactionYearMethod maps a *response.TransactionYearMethodResponse to a *pb.TransactionYearMethodResponse proto message.

It takes a *response.TransactionYearMethodResponse as input and returns a pointer to the converted TransactionYearMethodResponse proto object.
The mapping includes fields like Year, PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (m *transactionProtoMapper) mapResponseTransactionYearMethod(s *response.TransactionYearMethodResponse) *pb.TransactionYearMethodResponse
```

##### `mapResponseTransactionYearMethods`

mapResponseTransactionYearMethods maps a slice of TransactionYearMethodResponse to a slice of protobuf TransactionYearMethodResponse.

It takes a slice of TransactionYearMethodResponse as input and returns a slice of corresponding protobuf TransactionYearMethodResponse.
The mapping includes fields like Year, PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (m *transactionProtoMapper) mapResponseTransactionYearMethods(s []*response.TransactionYearMethodResponse) []*pb.TransactionYearMethodResponse
```

##### `mapResponseTransactionYearlyAmount`

mapResponseTransactionYearlyAmount maps a *response.TransactionYearlyAmountResponse to a *pb.TransactionYearlyAmountResponse proto message.

It takes a *response.TransactionYearlyAmountResponse as input and returns a pointer to the converted TransactionYearlyAmountResponse proto object.
The mapping includes fields like Year and TotalAmount.

```go
func (m *transactionProtoMapper) mapResponseTransactionYearlyAmount(s *response.TransactionYearlyAmountResponse) *pb.TransactionYearlyAmountResponse
```

##### `mapResponseTransactionYearlyAmounts`

mapResponseTransactionYearlyAmounts maps a slice of TransactionYearlyAmountResponse
to a slice of protobuf TransactionYearlyAmountResponse.

It iterates through each TransactionYearlyAmountResponse, converting it to its
protobuf equivalent using the mapResponseTransactionYearlyAmount function.

Args:
  - s: A slice of TransactionYearlyAmountResponse to be converted.

Returns:
  - A slice of pointers to TransactionYearlyAmountResponse containing the mapped data.

```go
func (m *transactionProtoMapper) mapResponseTransactionYearlyAmounts(s []*response.TransactionYearlyAmountResponse) []*pb.TransactionYearlyAmountResponse
```

##### `mapResponsesTransaction`

mapResponsesTransaction maps a slice of response.TransactionResponse to a slice of protobuf TransactionResponse.

It takes a slice of response.TransactionResponse as input and returns a slice of corresponding protobuf TransactionResponse.
The mapping includes fields like ID, TransactionNo, CardNumber, Amount, PaymentMethod, TransactionTime, MerchantId, CreatedAt, and UpdatedAt.

```go
func (m *transactionProtoMapper) mapResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse
```

##### `mapResponsesTransactionDeleteAt`

mapResponsesTransactionDeleteAt maps a slice of response.TransactionResponseDeleteAt to a slice of protobuf TransactionResponseDeleteAt.

Args:
  - transactions: A slice of TransactionResponseDeleteAt to be converted.

Returns:

	A slice of pointers to TransactionResponseDeleteAt containing the mapped data.

```go
func (m *transactionProtoMapper) mapResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt
```

##### `mapResponsesTransactionMonthStatusFailed`

mapResponsesTransactionMonthStatusFailed converts a slice of TransactionResponseMonthStatusFailed
to a slice of protobuf TransactionMonthStatusFailedResponse.

It iterates through each TransactionResponseMonthStatusFailed, mapping it to its protobuf equivalent
using the mapResponseTransactionMonthStatusFailed function.

Args:
  - Transactions: A slice of TransactionResponseMonthStatusFailed to be converted.

Returns:

	A slice of pointers to TransactionMonthStatusFailedResponse containing the mapped data.

```go
func (t *transactionProtoMapper) mapResponsesTransactionMonthStatusFailed(Transactions []*response.TransactionResponseMonthStatusFailed) []*pb.TransactionMonthStatusFailedResponse
```

##### `mapResponsesTransactionMonthStatusSuccess`

mapResponsesTransactionMonthStatusSuccess converts a slice of TransactionResponseMonthStatusSuccess
to a slice of protobuf TransactionMonthStatusSuccessResponse.

It iterates through each TransactionResponseMonthStatusSuccess, mapping it to its protobuf equivalent
using the mapResponseTransactionMonthStatusSuccess function.

Args:
  - Transactions: A slice of TransactionResponseMonthStatusSuccess to be converted.

Returns:

	A slice of pointers to TransactionMonthStatusSuccessResponse containing the mapped data.

```go
func (t *transactionProtoMapper) mapResponsesTransactionMonthStatusSuccess(Transactions []*response.TransactionResponseMonthStatusSuccess) []*pb.TransactionMonthStatusSuccessResponse
```

##### `mapTransactionResponseYearStatusFailed`

mapTransactionResponseYearStatusFailed maps a response.TransactionResponseYearStatusFailed to its protobuf representation.

It takes a response.TransactionResponseYearStatusFailed as input and returns a pointer to the converted
TransactionYearStatusFailedResponse proto object. The mapping includes fields like Year, TotalFailed, and TotalAmount.

Args:
  - s: A pointer to a TransactionResponseYearStatusFailed to be converted.

Returns:
  - A pointer to a TransactionYearStatusFailedResponse containing the mapped data.

```go
func (t *transactionProtoMapper) mapTransactionResponseYearStatusFailed(s *response.TransactionResponseYearStatusFailed) *pb.TransactionYearStatusFailedResponse
```

##### `mapTransactionResponseYearStatusSuccess`

mapTransactionResponseYearStatusSuccess maps a response.TransactionResponseYearStatusSuccess to its protobuf representation.

It takes a response.TransactionResponseYearStatusSuccess as input and returns a pointer to the converted TransactionYearStatusSuccessResponse proto object.
The mapping includes fields like Year, TotalSuccess, and TotalAmount.

```go
func (t *transactionProtoMapper) mapTransactionResponseYearStatusSuccess(s *response.TransactionResponseYearStatusSuccess) *pb.TransactionYearStatusSuccessResponse
```

##### `mapTransactionResponsesYearStatusFailed`

mapTransactionResponsesYearStatusFailed maps a slice of TransactionResponseYearStatusFailed
to a slice of protobuf TransactionYearStatusFailedResponse.

It iterates through each TransactionResponseYearStatusFailed, mapping it to its protobuf equivalent
using the mapTransactionResponseYearStatusFailed function.

Args:
  - Transactions: A slice of TransactionResponseYearStatusFailed to be converted.

Returns:

	A slice of pointers to TransactionYearStatusFailedResponse containing the mapped data.

```go
func (t *transactionProtoMapper) mapTransactionResponsesYearStatusFailed(Transactions []*response.TransactionResponseYearStatusFailed) []*pb.TransactionYearStatusFailedResponse
```

##### `mapTransactionResponsesYearStatusSuccess`

mapTransactionResponsesYearStatusSuccess maps a slice of TransactionResponseYearStatusSuccess to a slice of protobuf TransactionYearStatusSuccessResponse.

It takes a slice of TransactionResponseYearStatusSuccess as input and returns a slice of corresponding protobuf TransactionYearStatusSuccessResponse.
The mapping includes fields like Year, TotalSuccess, and TotalAmount.

```go
func (t *transactionProtoMapper) mapTransactionResponsesYearStatusSuccess(Transactions []*response.TransactionResponseYearStatusSuccess) []*pb.TransactionYearStatusSuccessResponse
```

### `transferProtoMapper`

transferProtoMapper is responsible for mapping Transfer responses to their corresponding protobuf representations.

```go
type transferProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponsePaginationTransfer`

ToProtoResponsePaginationTransfer maps a pagination meta, status, message, and a list of TransferResponse
to an ApiResponsePaginationTransfer proto message.

Args:
  - pagination: The pagination meta that needs to be converted.
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferResponse to be converted.

Returns:
  - A pointer to an ApiResponsePaginationTransfer containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponsePaginationTransfer(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponsePaginationTransfer
```

##### `ToProtoResponsePaginationTransferDeleteAt`

ToProtoResponsePaginationTransferDeleteAt maps paginated soft-deleted transfer responses to a protobuf message.

Args:
  - pagination: The pagination metadata for the response.
  - status: The status of the API response.
  - message: A descriptive message for the API response.
  - pbResponse: A slice of TransferResponseDeleteAt objects to be included in the response.

Returns:
  - A pointer to an ApiResponsePaginationTransferDeleteAt containing the status, message, transfer data, and pagination data.

```go
func (m *transferProtoMapper) ToProtoResponsePaginationTransferDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponseDeleteAt) *pb.ApiResponsePaginationTransferDeleteAt
```

##### `ToProtoResponseTransfer`

ToProtoResponseTransfer maps a status, message, and a single TransferResponse
to an ApiResponseTransfer proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: The TransferResponse to be converted.

Returns:
  - A pointer to an ApiResponseTransfer containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransfer(status string, message string, pbResponse *response.TransferResponse) *pb.ApiResponseTransfer
```

##### `ToProtoResponseTransferAll`

ToProtoResponseTransferAll maps a status, message, and no data
to an ApiResponseTransferAll proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.

Returns:
  - A pointer to an ApiResponseTransferAll containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferAll(status string, message string) *pb.ApiResponseTransferAll
```

##### `ToProtoResponseTransferDelete`

ToProtoResponseTransferDelete maps a status, message, and no data
to an ApiResponseTransferDelete proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.

Returns:
  - A pointer to an ApiResponseTransferDelete containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferDelete(status string, message string) *pb.ApiResponseTransferDelete
```

##### `ToProtoResponseTransferMonthAmount`

ToProtoResponseTransferMonthAmount maps a status, message, and a list of TransferMonthAmountResponse
to an ApiResponseTransferMonthAmount proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferMonthAmountResponse to be converted.

Returns:
  - A pointer to an ApiResponseTransferMonthAmount containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferMonthAmount(status string, message string, pbResponse []*response.TransferMonthAmountResponse) *pb.ApiResponseTransferMonthAmount
```

##### `ToProtoResponseTransferMonthStatusFailed`

ToProtoResponseTransferMonthStatusFailed maps a status, message, and a list of TransferResponseMonthStatusFailed
to an ApiResponseTransferMonthStatusFailed proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferResponseMonthStatusFailed to be converted.

Returns:
  - A pointer to an ApiResponseTransferMonthStatusFailed containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferMonthStatusFailed(status string, message string, pbResponse []*response.TransferResponseMonthStatusFailed) *pb.ApiResponseTransferMonthStatusFailed
```

##### `ToProtoResponseTransferMonthStatusSuccess`

ToProtoResponseTransferMonthStatusSuccess maps a status, message, and a list of TransferResponseMonthStatusSuccess
to an ApiResponseTransferMonthStatusSuccess proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferResponseMonthStatusSuccess to be converted.

Returns:
  - A pointer to an ApiResponseTransferMonthStatusSuccess containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferMonthStatusSuccess(status string, message string, pbResponse []*response.TransferResponseMonthStatusSuccess) *pb.ApiResponseTransferMonthStatusSuccess
```

##### `ToProtoResponseTransferYearAmount`

ToProtoResponseTransferYearAmount maps a status, message, and a list of TransferYearAmountResponse
to an ApiResponseTransferYearAmount proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferYearAmountResponse to be converted.

Returns:
  - A pointer to an ApiResponseTransferYearAmount containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferYearAmount(status string, message string, pbResponse []*response.TransferYearAmountResponse) *pb.ApiResponseTransferYearAmount
```

##### `ToProtoResponseTransferYearStatusFailed`

ToProtoResponseTransferYearStatusFailed maps a status, message, and a list of TransferResponseYearStatusFailed
to an ApiResponseTransferYearStatusFailed proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferResponseYearStatusFailed to be converted.

Returns:
  - A pointer to an ApiResponseTransferYearStatusFailed containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferYearStatusFailed(status string, message string, pbResponse []*response.TransferResponseYearStatusFailed) *pb.ApiResponseTransferYearStatusFailed
```

##### `ToProtoResponseTransferYearStatusSuccess`

ToProtoResponseTransferYearStatusSuccess maps a status, message, and a list of TransferResponseYearStatusSuccess
to an ApiResponseTransferYearStatusSuccess proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferResponseYearStatusSuccess to be converted.

Returns:
  - A pointer to an ApiResponseTransferYearStatusSuccess containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransferYearStatusSuccess(status string, message string, pbResponse []*response.TransferResponseYearStatusSuccess) *pb.ApiResponseTransferYearStatusSuccess
```

##### `ToProtoResponseTransfers`

ToProtoResponseTransfers maps a status, message, and a list of TransferResponse
to an ApiResponseTransfers proto message.

Args:
  - status: The status of the response.
  - message: The accompanying message of the response.
  - pbResponse: A slice of TransferResponse to be converted.

Returns:
  - A pointer to an ApiResponseTransfers containing the mapped data.

```go
func (m *transferProtoMapper) ToProtoResponseTransfers(status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponseTransfers
```

##### `mapResponseTransfer`

mapResponseTransfer maps a TransferResponse to a TransferResponse proto message.

Args:
  - transfer: The TransferResponse to be converted.

Returns:
  - A pointer to a TransferResponse containing the mapped data.

```go
func (t *transferProtoMapper) mapResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse
```

##### `mapResponseTransferDeleteAt`

mapResponseTransferDeleteAt maps a TransferResponseDeleteAt to its protobuf representation.

Args:
  - transfer: The TransferResponseDeleteAt that needs to be converted.

Returns:
  - A pointer to a TransferResponseDeleteAt protobuf message containing the mapped data.

```go
func (t *transferProtoMapper) mapResponseTransferDeleteAt(transfer *response.TransferResponseDeleteAt) *pb.TransferResponseDeleteAt
```

##### `mapResponseTransferMonthAmount`

mapResponseTransferMonthAmount maps a TransferMonthAmountResponse to its protobuf representation.

Args:
  - s: The TransferMonthAmountResponse that needs to be converted.

Returns:
  - A pointer to a TransferMonthAmountResponse protobuf message containing the mapped data.

```go
func (m *transferProtoMapper) mapResponseTransferMonthAmount(s *response.TransferMonthAmountResponse) *pb.TransferMonthAmountResponse
```

##### `mapResponseTransferMonthAmounts`

mapResponseTransferMonthAmounts maps a slice of TransferMonthAmountResponse
to a slice of TransferMonthAmountResponse protobuf messages.

It iterates over the input slice, converting each TransferMonthAmountResponse
to its corresponding protobuf representation using the mapResponseTransferMonthAmount method.

Args:
  - s: A slice of TransferMonthAmountResponse to be converted.

Returns:
  - A slice of TransferMonthAmountResponse containing the mapped protobuf data.

```go
func (m *transferProtoMapper) mapResponseTransferMonthAmounts(s []*response.TransferMonthAmountResponse) []*pb.TransferMonthAmountResponse
```

##### `mapResponseTransferMonthStatusFailed`

mapResponseTransferMonthStatusFailed maps a TransferResponseMonthStatusFailed to its protobuf representation.

Args:
  - s: The TransferResponseMonthStatusFailed that needs to be converted.

Returns:
  - A pointer to a TransferMonthStatusFailedResponse protobuf message containing the mapped data.

```go
func (t *transferProtoMapper) mapResponseTransferMonthStatusFailed(s *response.TransferResponseMonthStatusFailed) *pb.TransferMonthStatusFailedResponse
```

##### `mapResponseTransferMonthStatusSuccess`

mapResponseTransferMonthStatusSuccess maps a TransferResponseMonthStatusSuccess to its protobuf representation.

Args:
  - s: The TransferResponseMonthStatusSuccess that needs to be converted.

Returns:
  - A pointer to a TransferMonthStatusSuccessResponse protobuf message containing the mapped data.

```go
func (t *transferProtoMapper) mapResponseTransferMonthStatusSuccess(s *response.TransferResponseMonthStatusSuccess) *pb.TransferMonthStatusSuccessResponse
```

##### `mapResponseTransferYearAmount`

mapResponseTransferYearAmount maps a TransferYearAmountResponse to its protobuf representation.

Args:
  - s: The TransferYearAmountResponse that needs to be converted.

Returns:
  - A pointer to a TransferYearAmountResponse protobuf message containing the mapped data.

```go
func (m *transferProtoMapper) mapResponseTransferYearAmount(s *response.TransferYearAmountResponse) *pb.TransferYearAmountResponse
```

##### `mapResponseTransferYearAmounts`

mapResponseTransferYearAmounts maps a slice of TransferYearAmountResponse
to a slice of TransferYearAmountResponse protobuf messages.

It iterates over the input slice, converting each TransferYearAmountResponse
to its corresponding protobuf representation using the mapResponseTransferYearAmount method.

Args:
  - s: A slice of TransferYearAmountResponse to be converted.

Returns:
  - A slice of TransferYearAmountResponse containing the mapped protobuf data.

```go
func (m *transferProtoMapper) mapResponseTransferYearAmounts(s []*response.TransferYearAmountResponse) []*pb.TransferYearAmountResponse
```

##### `mapResponsesTransfer`

mapResponsesTransfer maps a slice of TransferResponse to a slice of TransferResponse proto messages.

Args:
  - transfers: A slice of TransferResponse to be converted.

Returns:
  - A slice of pointers to TransferResponse containing the mapped data.

```go
func (t *transferProtoMapper) mapResponsesTransfer(transfers []*response.TransferResponse) []*pb.TransferResponse
```

##### `mapResponsesTransferDeleteAt`

mapResponsesTransferDeleteAt maps a slice of TransferResponseDeleteAt to a slice of
their protobuf representations.

Args:
  - transfers: A slice of TransferResponseDeleteAt to be converted.

Returns:
  - A slice of pointers to TransferResponseDeleteAt protobuf messages containing the mapped data.

```go
func (t *transferProtoMapper) mapResponsesTransferDeleteAt(transfers []*response.TransferResponseDeleteAt) []*pb.TransferResponseDeleteAt
```

##### `mapResponsesTransferMonthStatusFailed`

mapResponsesTransferMonthStatusFailed maps a slice of TransferResponseMonthStatusFailed
to a slice of TransferMonthStatusFailedResponse protobuf messages.

It iterates over the input slice, converting each TransferResponseMonthStatusFailed
to its corresponding protobuf representation using the mapResponseTransferMonthStatusFailed method.

Args:
  - Transfers: A slice of TransferResponseMonthStatusFailed to be converted.

Returns:
  - A slice of TransferMonthStatusFailedResponse containing the mapped protobuf data.

```go
func (t *transferProtoMapper) mapResponsesTransferMonthStatusFailed(Transfers []*response.TransferResponseMonthStatusFailed) []*pb.TransferMonthStatusFailedResponse
```

##### `mapResponsesTransferMonthStatusSuccess`

mapResponsesTransferMonthStatusSuccess maps a slice of TransferResponseMonthStatusSuccess
to a slice of TransferMonthStatusSuccessResponse protobuf messages.

It iterates over the input slice, converting each TransferResponseMonthStatusSuccess
to its corresponding protobuf representation using the mapResponseTransferMonthStatusSuccess method.

Args:
  - Transfers: A slice of TransferResponseMonthStatusSuccess to be converted.

Returns:
  - A slice of TransferMonthStatusSuccessResponse containing the mapped protobuf data.

```go
func (t *transferProtoMapper) mapResponsesTransferMonthStatusSuccess(Transfers []*response.TransferResponseMonthStatusSuccess) []*pb.TransferMonthStatusSuccessResponse
```

##### `mapTransferResponseYearStatusFailed`

mapTransferResponseYearStatusFailed maps a TransferResponseYearStatusFailed to its protobuf representation.

Args:
  - s: The TransferResponseYearStatusFailed that needs to be converted.

Returns:
  - A pointer to a TransferYearStatusFailedResponse protobuf message containing the mapped data.

```go
func (t *transferProtoMapper) mapTransferResponseYearStatusFailed(s *response.TransferResponseYearStatusFailed) *pb.TransferYearStatusFailedResponse
```

##### `mapTransferResponseYearStatusSuccess`

mapTransferResponseYearStatusSuccess maps a TransferResponseYearStatusSuccess to its protobuf representation.

Args:
  - s: The TransferResponseYearStatusSuccess that needs to be converted.

Returns:
  - A pointer to a TransferYearStatusSuccessResponse protobuf message containing the mapped data.

```go
func (t *transferProtoMapper) mapTransferResponseYearStatusSuccess(s *response.TransferResponseYearStatusSuccess) *pb.TransferYearStatusSuccessResponse
```

##### `mapTransferResponsesYearStatusFailed`

mapTransferResponsesYearStatusFailed maps a slice of TransferResponseYearStatusFailed
to a slice of TransferYearStatusFailedResponse protobuf messages.

It iterates over the input slice, converting each TransferResponseYearStatusFailed
to its corresponding protobuf representation using the mapTransferResponseYearStatusFailed method.

Args:
  - Transfers: A slice of TransferResponseYearStatusFailed to be converted.

Returns:
  - A slice of TransferYearStatusFailedResponse containing the mapped protobuf data.

```go
func (t *transferProtoMapper) mapTransferResponsesYearStatusFailed(Transfers []*response.TransferResponseYearStatusFailed) []*pb.TransferYearStatusFailedResponse
```

##### `mapTransferResponsesYearStatusSuccess`

mapTransferResponsesYearStatusSuccess maps a slice of TransferResponseYearStatusSuccess
to a slice of TransferYearStatusSuccessResponse protobuf messages.

It iterates over the input slice, converting each TransferResponseYearStatusSuccess
to its corresponding protobuf representation using the mapTransferResponseYearStatusSuccess method.

Args:
  - Transfers: A slice of TransferResponseYearStatusSuccess to be converted.

Returns:
  - A slice of TransferYearStatusSuccessResponse containing the mapped protobuf data.

```go
func (t *transferProtoMapper) mapTransferResponsesYearStatusSuccess(Transfers []*response.TransferResponseYearStatusSuccess) []*pb.TransferYearStatusSuccessResponse
```

### `userProtoMapper`

UserProtoMapper provides methods for mapping domain models to pb.UserResponse

```go
type userProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponsePaginationUser`

ToProtoResponsePaginationUser maps a list of UserResponse and pagination metadata
to a protobuf ApiResponsePaginationUser. It includes the status and message for the API response.

Args:
  - pagination: The pagination metadata for the response.
  - status: The status of the API response.
  - message: A descriptive message for the API response.
  - users: A list of UserResponse objects to be included in the response.

Returns:

	A pointer to ApiResponsePaginationUser containing the status, message, user data, and pagination data.

```go
func (u *userProtoMapper) ToProtoResponsePaginationUser(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponse) *pb.ApiResponsePaginationUser
```

##### `ToProtoResponsePaginationUserDeleteAt`

ToProtoResponsePaginationUserDeleteAt maps a list of UserResponseDeleteAt and pagination metadata
to a protobuf ApiResponsePaginationUserDeleteAt. It includes the status and message for the API response.

Args:
  - pagination: The pagination metadata for the response.
  - status: The status of the API response.
  - message: A descriptive message for the API response.
  - users: A list of UserResponseDeleteAt objects to be included in the response.

Returns:

	A pointer to ApiResponsePaginationUserDeleteAt containing the status, message, user data, and pagination data.

```go
func (u *userProtoMapper) ToProtoResponsePaginationUserDeleteAt(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) *pb.ApiResponsePaginationUserDeleteAt
```

##### `ToProtoResponseUser`

ToProtoResponseUser maps a UserResponse to a pb.UserResponse

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.
  - pbResponse: A pointer to the UserResponse to be mapped.

Returns:
  - A pointer to a pb.UserResponse containing the mapped data.
  - If pbResponse is nil, returns nil.

```go
func (u *userProtoMapper) ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser
```

##### `ToProtoResponseUserAll`

ToProtoResponseUserAll returns a protobuf message for all users without pagination.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.

Returns:
  - A pointer to a pb.ApiResponseUserAll containing the mapped data.

```go
func (u *userProtoMapper) ToProtoResponseUserAll(status string, message string) *pb.ApiResponseUserAll
```

##### `ToProtoResponseUserDelete`

ToProtoResponseUserDelete maps a UserResponseDelete to a pb.ApiResponseUserDelete proto message.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.

Returns:
  - A pointer to a pb.ApiResponseUserDelete containing the mapped data.

```go
func (u *userProtoMapper) ToProtoResponseUserDelete(status string, message string) *pb.ApiResponseUserDelete
```

##### `ToProtoResponseUserDeleteAt`

ToProtoResponseUserDeleteAt maps a UserResponseDeleteAt to a pb.ApiResponseUserDeleteAt proto message.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.
  - pbResponse: A pointer to the UserResponseDeleteAt to be mapped.

Returns:
  - A pointer to a pb.ApiResponseUserDeleteAt containing the mapped data.
  - If pbResponse is nil, returns nil.

```go
func (u *userProtoMapper) ToProtoResponseUserDeleteAt(status string, message string, pbResponse *response.UserResponseDeleteAt) *pb.ApiResponseUserDeleteAt
```

##### `ToProtoResponsesUser`

ToProtoResponsesUser maps a list of UserResponse to a pb.ApiResponsesUser proto message.

Args:
  - status: A string representing the status of the response.
  - message: A string representing the message associated with the response.
  - pbResponse: A slice of pointers to UserResponse to be mapped.

Returns:
  - A pointer to a pb.ApiResponsesUser containing the mapped data.
  - If pbResponse is nil, returns nil.

```go
func (u *userProtoMapper) ToProtoResponsesUser(status string, message string, pbResponse []*response.UserResponse) *pb.ApiResponsesUser
```

##### `mapResponseUser`

mapResponseUser maps a UserResponse to a protobuf UserResponse.

Args:
  - user: A pointer to a UserResponse to be mapped.

Returns:
  - A pointer to a protobuf UserResponse containing the mapped data.

```go
func (u *userProtoMapper) mapResponseUser(user *response.UserResponse) *pb.UserResponse
```

##### `mapResponseUserDelete`

mapResponseUserDelete maps a UserResponseDeleteAt to a protobuf UserResponseDeleteAt.

Args:
  - user: A pointer to a UserResponseDeleteAt containing the user's details.

Returns:
  - A pointer to a protobuf UserResponseDeleteAt containing the mapped data, including
    ID, Firstname, Lastname, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.

```go
func (u *userProtoMapper) mapResponseUserDelete(user *response.UserResponseDeleteAt) *pb.UserResponseDeleteAt
```

##### `mapResponsesUser`

mapResponsesUser maps a slice of UserResponse to a slice of protobuf UserResponse.

Args:
  - users: A slice of pointers to UserResponse to be mapped.

Returns:
  - A slice of pointers to protobuf UserResponse containing the mapped data for each user.

```go
func (u *userProtoMapper) mapResponsesUser(users []*response.UserResponse) []*pb.UserResponse
```

##### `mapResponsesUserDeleteAt`

mapResponsesUserDeleteAt maps a slice of UserResponseDeleteAt to a slice of protobuf UserResponseDeleteAt.

Args:
  - users: A slice of pointers to UserResponseDeleteAt to be mapped.

Returns:
  - A slice of pointers to protobuf UserResponseDeleteAt containing the mapped data for each user, including
    ID, Firstname, Lastname, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.

```go
func (u *userProtoMapper) mapResponsesUserDeleteAt(users []*response.UserResponseDeleteAt) []*pb.UserResponseDeleteAt
```

### `withdrawProtoMapper`

```go
type withdrawProtoMapper struct {
}
```

#### Methods

##### `ToProtoResponsePaginationWithdraw`

ToProtoResponsePaginationWithdraw maps a status, message, a pagination meta, and a list of WithdrawResponse
to an ApiResponsePaginationWithdraw proto message.

Args:
  - pagination: The pagination meta that needs to be converted.
  - status: The status of the response.
  - message: The message accompanying the response.
  - withdraws: The list of WithdrawResponse that needs to be converted.

Returns:

	A pointer to an ApiResponsePaginationWithdraw containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponsePaginationWithdraw(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsePaginationWithdraw
```

##### `ToProtoResponsePaginationWithdrawDeleteAt`

ToProtoResponsePaginationWithdrawDeleteAt maps a pagination meta, status, message, and a list of WithdrawResponseDeleteAt
to an ApiResponsePaginationWithdrawDeleteAt proto message.

Args:
  - pagination: The pagination meta that needs to be converted.
  - status: The status of the response.
  - message: The message accompanying the response.
  - withdraws: The list of WithdrawResponseDeleteAt that needs to be converted.

Returns:

	A pointer to an ApiResponsePaginationWithdrawDeleteAt containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponsePaginationWithdrawDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponseDeleteAt) *pb.ApiResponsePaginationWithdrawDeleteAt
```

##### `ToProtoResponseWithdraw`

ToProtoResponseWithdraw maps a status, message, and a WithdrawResponse
to an ApiResponseWithdraw proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - withdraw: The WithdrawResponse that needs to be converted.

Returns:

	A pointer to an ApiResponseWithdraw containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdraw(status string, message string, withdraw *response.WithdrawResponse) *pb.ApiResponseWithdraw
```

##### `ToProtoResponseWithdrawAll`

ToProtoResponseWithdrawAll maps a status and message to an ApiResponseWithdrawAll proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.

Returns:

	A pointer to an ApiResponseWithdrawAll containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawAll(status string, message string) *pb.ApiResponseWithdrawAll
```

##### `ToProtoResponseWithdrawDelete`

ToProtoResponseWithdrawDelete maps a status and message to an ApiResponseWithdrawDelete proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.

Returns:

	A pointer to an ApiResponseWithdrawDelete containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawDelete(status string, message string) *pb.ApiResponseWithdrawDelete
```

##### `ToProtoResponseWithdrawMonthAmount`

ToProtoResponseWithdrawMonthAmount maps a status, message, and a list of WithdrawMonthlyAmountResponse
to an ApiResponseWithdrawMonthAmount proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - pbResponse: The list of WithdrawMonthlyAmountResponse that needs to be converted.

Returns:

	A pointer to an ApiResponseWithdrawMonthAmount containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawMonthAmount(status string, message string, pbResponse []*response.WithdrawMonthlyAmountResponse) *pb.ApiResponseWithdrawMonthAmount
```

##### `ToProtoResponseWithdrawMonthStatusFailed`

ToProtoResponseWithdrawMonthStatusFailed maps a status, message, and a list of WithdrawResponseMonthStatusFailed
to an ApiResponseWithdrawMonthStatusFailed proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - pbResponse: The list of WithdrawResponseMonthStatusFailed that needs to be converted.

Returns:

	A pointer to an ApiResponseWithdrawMonthStatusFailed containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawMonthStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusFailed) *pb.ApiResponseWithdrawMonthStatusFailed
```

##### `ToProtoResponseWithdrawMonthStatusSuccess`

ToProtoResponseWithdrawMonthStatusSuccess maps a status, message, and a list of WithdrawResponseMonthStatusSuccess
to an ApiResponseWithdrawMonthStatusSuccess proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - pbResponse: The list of WithdrawResponseMonthStatusSuccess that needs to be converted.

Returns:

	A pointer to an ApiResponseWithdrawMonthStatusSuccess containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawMonthStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusSuccess) *pb.ApiResponseWithdrawMonthStatusSuccess
```

##### `ToProtoResponseWithdrawYearAmount`

ToProtoResponseWithdrawYearAmount maps a status, message, and a list of WithdrawYearlyAmountResponse
to an ApiResponseWithdrawYearAmount proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - pbResponse: The list of WithdrawYearlyAmountResponse that needs to be converted.

Returns:

	A pointer to an ApiResponseWithdrawYearAmount containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawYearAmount(status string, message string, pbResponse []*response.WithdrawYearlyAmountResponse) *pb.ApiResponseWithdrawYearAmount
```

##### `ToProtoResponseWithdrawYearStatusFailed`

ToProtoResponseWithdrawYearStatusFailed maps a status, message, and a list of WithdrawResponseYearStatusFailed
to an ApiResponseWithdrawYearStatusFailed proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - pbResponse: The list of WithdrawResponseYearStatusFailed that needs to be converted.

Returns:

	A pointer to an ApiResponseWithdrawYearStatusFailed containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawYearStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseYearStatusFailed) *pb.ApiResponseWithdrawYearStatusFailed
```

##### `ToProtoResponseWithdrawYearStatusSuccess`

ToProtoResponseWithdrawYearStatusSuccess maps a status, message, and a list of WithdrawResponseYearStatusSuccess
to an ApiResponseWithdrawYearStatusSuccess proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - pbResponse: The list of WithdrawResponseYearStatusSuccess that needs to be converted.

Returns:

	A pointer to an ApiResponseWithdrawYearStatusSuccess containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponseWithdrawYearStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseYearStatusSuccess) *pb.ApiResponseWithdrawYearStatusSuccess
```

##### `ToProtoResponsesWithdraw`

ToProtoResponsesWithdraw maps a status, message, and a list of WithdrawResponse
to an ApiResponsesWithdraw proto message.

Args:
  - status: The status of the response.
  - message: The message accompanying the response.
  - withdraws: The list of WithdrawResponse that needs to be converted.

Returns:

	A pointer to an ApiResponsesWithdraw containing the mapped data.

```go
func (m *withdrawProtoMapper) ToProtoResponsesWithdraw(status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsesWithdraw
```

##### `mapResponseWithdrawMonthStatusFailed`

mapResponseWithdrawMonthStatusFailed maps a single WithdrawResponseMonthStatusFailed to its protobuf representation.

Args:
  - s: The WithdrawResponseMonthStatusFailed that needs to be converted.

Returns:
  - A pointer to a WithdrawMonthStatusFailedResponse containing the mapped data, with fields Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawProtoMapper) mapResponseWithdrawMonthStatusFailed(s *response.WithdrawResponseMonthStatusFailed) *pb.WithdrawMonthStatusFailedResponse
```

##### `mapResponseWithdrawMonthStatusSuccess`

mapResponseWithdrawMonthStatusSuccess maps a single WithdrawResponseMonthStatusSuccess to its protobuf representation.

Args:
  - s: The WithdrawResponseMonthStatusSuccess that needs to be converted.

Returns:
  - A pointer to a WithdrawMonthStatusSuccessResponse protobuf message containing the mapped data.

```go
func (t *withdrawProtoMapper) mapResponseWithdrawMonthStatusSuccess(s *response.WithdrawResponseMonthStatusSuccess) *pb.WithdrawMonthStatusSuccessResponse
```

##### `mapResponseWithdrawMonthlyAmount`

mapResponseWithdrawMonthlyAmount maps a WithdrawMonthlyAmountResponse to its protobuf representation.

Args:
  - s: The WithdrawMonthlyAmountResponse that needs to be converted.

Returns:
  - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data, with fields Month and TotalAmount.

```go
func (m *withdrawProtoMapper) mapResponseWithdrawMonthlyAmount(s *response.WithdrawMonthlyAmountResponse) *pb.WithdrawMonthlyAmountResponse
```

##### `mapResponseWithdrawMonthlyAmounts`

mapResponseWithdrawMonthlyAmounts maps a slice of WithdrawMonthlyAmountResponse to a slice of protobuf WithdrawMonthlyAmountResponse.

It takes a slice of WithdrawMonthlyAmountResponse as input and returns a slice of corresponding protobuf WithdrawMonthlyAmountResponse.
The mapping includes fields like Month and TotalAmount.

```go
func (m *withdrawProtoMapper) mapResponseWithdrawMonthlyAmounts(s []*response.WithdrawMonthlyAmountResponse) []*pb.WithdrawMonthlyAmountResponse
```

##### `mapResponseWithdrawYearlyAmount`

mapResponseWithdrawYearlyAmount maps a WithdrawYearlyAmountResponse to its protobuf representation.

Args:
  - s: The WithdrawYearlyAmountResponse that needs to be converted.

Returns:
  - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, with fields Year and TotalAmount.

```go
func (m *withdrawProtoMapper) mapResponseWithdrawYearlyAmount(s *response.WithdrawYearlyAmountResponse) *pb.WithdrawYearlyAmountResponse
```

##### `mapResponseWithdrawYearlyAmounts`

mapResponseWithdrawYearlyAmounts maps a slice of WithdrawYearlyAmountResponse to a slice of protobuf WithdrawYearlyAmountResponse.

It takes a slice of WithdrawYearlyAmountResponse as input and returns a slice of corresponding protobuf WithdrawYearlyAmountResponse.
The mapping includes fields like Year and TotalAmount.

```go
func (m *withdrawProtoMapper) mapResponseWithdrawYearlyAmounts(s []*response.WithdrawYearlyAmountResponse) []*pb.WithdrawYearlyAmountResponse
```

##### `mapResponseWithdrawal`

mapResponseWithdrawal maps a single WithdrawResponse to its protobuf representation.

Args:
  - withdraw: The WithdrawResponse that needs to be converted.

Returns:
  - A pointer to a WithdrawResponse protobuf message containing the mapped data.

```go
func (w *withdrawProtoMapper) mapResponseWithdrawal(withdraw *response.WithdrawResponse) *pb.WithdrawResponse
```

##### `mapResponseWithdrawalDeleteAt`

mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to its protobuf representation.

Args:
  - withdraw: The WithdrawResponseDeleteAt that needs to be converted.

Returns:
  - A pointer to a WithdrawResponseDeleteAt protobuf message containing the mapped data.

```go
func (w *withdrawProtoMapper) mapResponseWithdrawalDeleteAt(withdraw *response.WithdrawResponseDeleteAt) *pb.WithdrawResponseDeleteAt
```

##### `mapResponsesWithdrawMonthStatusFailed`

mapResponsesWithdrawMonthStatusFailed maps a slice of WithdrawResponseMonthStatusFailed to a slice of protobuf WithdrawMonthStatusFailedResponse.

It takes a slice of WithdrawResponseMonthStatusFailed as input and returns a slice of corresponding protobuf WithdrawMonthStatusFailedResponse.
The mapping includes fields like Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawProtoMapper) mapResponsesWithdrawMonthStatusFailed(Withdraws []*response.WithdrawResponseMonthStatusFailed) []*pb.WithdrawMonthStatusFailedResponse
```

##### `mapResponsesWithdrawMonthStatusSuccess`

mapResponsesWithdrawMonthStatusSuccess maps a slice of WithdrawResponseMonthStatusSuccess to a slice of protobuf WithdrawMonthStatusSuccessResponse.

It takes a slice of WithdrawResponseMonthStatusSuccess as input and returns a slice of corresponding protobuf WithdrawMonthStatusSuccessResponse.
The mapping includes fields like Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *withdrawProtoMapper) mapResponsesWithdrawMonthStatusSuccess(Withdraws []*response.WithdrawResponseMonthStatusSuccess) []*pb.WithdrawMonthStatusSuccessResponse
```

##### `mapResponsesWithdrawal`

mapResponsesWithdrawal maps a slice of WithdrawResponse to a slice of protobuf WithdrawResponse.

It takes a slice of WithdrawResponse as input and returns a slice of corresponding protobuf WithdrawResponse.
The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.

```go
func (w *withdrawProtoMapper) mapResponsesWithdrawal(withdraws []*response.WithdrawResponse) []*pb.WithdrawResponse
```

##### `mapResponsesWithdrawalDeleteAt`

mapResponsesWithdrawalDeleteAt maps a slice of WithdrawResponseDeleteAt to a slice of protobuf WithdrawResponseDeleteAt.

It takes a slice of WithdrawResponseDeleteAt as input and returns a slice of corresponding protobuf WithdrawResponseDeleteAt.
The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (w *withdrawProtoMapper) mapResponsesWithdrawalDeleteAt(withdraws []*response.WithdrawResponseDeleteAt) []*pb.WithdrawResponseDeleteAt
```

##### `mapWithdrawResponseYearStatusFailed`

mapWithdrawResponseYearStatusFailed maps a WithdrawResponseYearStatusFailed to a WithdrawYearStatusFailedResponse proto message.

Args:
  - s: The WithdrawResponseYearStatusFailed that needs to be converted.

Returns:
  - A pointer to a WithdrawYearStatusFailedResponse containing the mapped data, with fields Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawProtoMapper) mapWithdrawResponseYearStatusFailed(s *response.WithdrawResponseYearStatusFailed) *pb.WithdrawYearStatusFailedResponse
```

##### `mapWithdrawResponseYearStatusSuccess`

mapWithdrawResponseYearStatusSuccess maps a single WithdrawResponseYearStatusSuccess to a WithdrawYearStatusSuccessResponse proto message.

Args:
  - s: The WithdrawResponseYearStatusSuccess that needs to be converted.

Returns:
  - A pointer to a WithdrawYearStatusSuccessResponse containing the mapped data, with fields Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawProtoMapper) mapWithdrawResponseYearStatusSuccess(s *response.WithdrawResponseYearStatusSuccess) *pb.WithdrawYearStatusSuccessResponse
```

##### `mapWithdrawResponsesYearStatusFailed`

mapWithdrawResponsesYearStatusFailed maps a slice of WithdrawResponseYearStatusFailed to a slice of protobuf WithdrawYearStatusFailedResponse.

It takes a slice of WithdrawResponseYearStatusFailed as input and returns a slice of corresponding protobuf WithdrawYearStatusFailedResponse.
The mapping includes fields like Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawProtoMapper) mapWithdrawResponsesYearStatusFailed(Withdraws []*response.WithdrawResponseYearStatusFailed) []*pb.WithdrawYearStatusFailedResponse
```

##### `mapWithdrawResponsesYearStatusSuccess`

mapWithdrawResponsesYearStatusSuccess maps a slice of WithdrawResponseYearStatusSuccess to a slice of protobuf WithdrawYearStatusSuccessResponse.

It takes a slice of WithdrawResponseYearStatusSuccess as input and returns a slice of corresponding protobuf WithdrawYearStatusSuccessResponse.
The mapping includes fields like Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawProtoMapper) mapWithdrawResponsesYearStatusSuccess(Withdraws []*response.WithdrawResponseYearStatusSuccess) []*pb.WithdrawYearStatusSuccessResponse
```

## 🚀 Functions

### `mapPaginationMeta`

mapPaginationMeta maps a sharedPagination.PaginationMeta to a protobuf PaginationMeta.

Args:
  - s: A pointer to a sharedPagination.PaginationMeta to be mapped.

Returns:
  - A pointer to a protobuf PaginationMeta containing the mapped data.

```go
func mapPaginationMeta(s *pb.PaginationMeta) *pb.PaginationMeta
```