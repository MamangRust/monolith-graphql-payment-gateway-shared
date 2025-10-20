# 📦 Package `recordmapper`

**Source Path:** `shared/mapper/record`

## 🧩 Types

### `CardRecordMapping`

CardRecordMapping provides methods to map database rows to Card domain models.

```go
type CardRecordMapping interface {
	ToCardEmailRecord func(card *db.GetUserEmailByCardNumberRow) (*record.CardEmailRecord)
	ToCardRecord func(card *db.Card) (*record.CardRecord)
	ToCardsRecord func(cards []*db.GetCardsRow) ([]*record.CardRecord)
	ToCardRecordActive func(card *db.GetActiveCardsWithCountRow) (*record.CardRecord)
	ToCardRecordsActive func(cards []*db.GetActiveCardsWithCountRow) ([]*record.CardRecord)
	ToCardRecordTrashed func(card *db.GetTrashedCardsWithCountRow) (*record.CardRecord)
	ToCardRecordsTrashed func(cards []*db.GetTrashedCardsWithCountRow) ([]*record.CardRecord)
	ToMonthlyBalance func(card *db.GetMonthlyBalancesRow) (*record.CardMonthBalance)
	ToMonthlyBalances func(cards []*db.GetMonthlyBalancesRow) ([]*record.CardMonthBalance)
	ToYearlyBalance func(card *db.GetYearlyBalancesRow) (*record.CardYearlyBalance)
	ToYearlyBalances func(cards []*db.GetYearlyBalancesRow) ([]*record.CardYearlyBalance)
	ToMonthlyTopupAmount func(card *db.GetMonthlyTopupAmountRow) (*record.CardMonthAmount)
	ToMonthlyTopupAmounts func(cards []*db.GetMonthlyTopupAmountRow) ([]*record.CardMonthAmount)
	ToYearlyTopupAmount func(card *db.GetYearlyTopupAmountRow) (*record.CardYearAmount)
	ToYearlyTopupAmounts func(cards []*db.GetYearlyTopupAmountRow) ([]*record.CardYearAmount)
	ToMonthlyWithdrawAmount func(card *db.GetMonthlyWithdrawAmountRow) (*record.CardMonthAmount)
	ToMonthlyWithdrawAmounts func(cards []*db.GetMonthlyWithdrawAmountRow) ([]*record.CardMonthAmount)
	ToYearlyWithdrawAmount func(card *db.GetYearlyWithdrawAmountRow) (*record.CardYearAmount)
	ToYearlyWithdrawAmounts func(cards []*db.GetYearlyWithdrawAmountRow) ([]*record.CardYearAmount)
	ToMonthlyTransactionAmount func(card *db.GetMonthlyTransactionAmountRow) (*record.CardMonthAmount)
	ToMonthlyTransactionAmounts func(cards []*db.GetMonthlyTransactionAmountRow) ([]*record.CardMonthAmount)
	ToYearlyTransactionAmount func(card *db.GetYearlyTransactionAmountRow) (*record.CardYearAmount)
	ToYearlyTransactionAmounts func(cards []*db.GetYearlyTransactionAmountRow) ([]*record.CardYearAmount)
	ToMonthlyTransferSenderAmount func(card *db.GetMonthlyTransferAmountSenderRow) (*record.CardMonthAmount)
	ToMonthlyTransferSenderAmounts func(cards []*db.GetMonthlyTransferAmountSenderRow) ([]*record.CardMonthAmount)
	ToYearlyTransferSenderAmount func(card *db.GetYearlyTransferAmountSenderRow) (*record.CardYearAmount)
	ToYearlyTransferSenderAmounts func(cards []*db.GetYearlyTransferAmountSenderRow) ([]*record.CardYearAmount)
	ToMonthlyTransferReceiverAmount func(card *db.GetMonthlyTransferAmountReceiverRow) (*record.CardMonthAmount)
	ToMonthlyTransferReceiverAmounts func(cards []*db.GetMonthlyTransferAmountReceiverRow) ([]*record.CardMonthAmount)
	ToYearlyTransferReceiverAmount func(card *db.GetYearlyTransferAmountReceiverRow) (*record.CardYearAmount)
	ToYearlyTransferReceiverAmounts func(cards []*db.GetYearlyTransferAmountReceiverRow) ([]*record.CardYearAmount)
	ToMonthlyBalanceCardNumber func(card *db.GetMonthlyBalancesByCardNumberRow) (*record.CardMonthBalance)
	ToMonthlyBalancesCardNumber func(cards []*db.GetMonthlyBalancesByCardNumberRow) ([]*record.CardMonthBalance)
	ToYearlyBalanceCardNumber func(card *db.GetYearlyBalancesByCardNumberRow) (*record.CardYearlyBalance)
	ToYearlyBalancesCardNumber func(cards []*db.GetYearlyBalancesByCardNumberRow) ([]*record.CardYearlyBalance)
	ToMonthlyTopupAmountByCardNumber func(card *db.GetMonthlyTopupAmountByCardNumberRow) (*record.CardMonthAmount)
	ToMonthlyTopupAmountsByCardNumber func(cards []*db.GetMonthlyTopupAmountByCardNumberRow) ([]*record.CardMonthAmount)
	ToYearlyTopupAmountByCardNumber func(card *db.GetYearlyTopupAmountByCardNumberRow) (*record.CardYearAmount)
	ToYearlyTopupAmountsByCardNumber func(cards []*db.GetYearlyTopupAmountByCardNumberRow) ([]*record.CardYearAmount)
	ToMonthlyWithdrawAmountByCardNumber func(card *db.GetMonthlyWithdrawAmountByCardNumberRow) (*record.CardMonthAmount)
	ToMonthlyWithdrawAmountsByCardNumber func(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) ([]*record.CardMonthAmount)
	ToYearlyWithdrawAmountByCardNumber func(card *db.GetYearlyWithdrawAmountByCardNumberRow) (*record.CardYearAmount)
	ToYearlyWithdrawAmountsByCardNumber func(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) ([]*record.CardYearAmount)
	ToMonthlyTransactionAmountByCardNumber func(card *db.GetMonthlyTransactionAmountByCardNumberRow) (*record.CardMonthAmount)
	ToMonthlyTransactionAmountsByCardNumber func(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) ([]*record.CardMonthAmount)
	ToYearlyTransactionAmountByCardNumber func(card *db.GetYearlyTransactionAmountByCardNumberRow) (*record.CardYearAmount)
	ToYearlyTransactionAmountsByCardNumber func(cards []*db.GetYearlyTransactionAmountByCardNumberRow) ([]*record.CardYearAmount)
	ToMonthlyTransferSenderAmountByCardNumber func(card *db.GetMonthlyTransferAmountBySenderRow) (*record.CardMonthAmount)
	ToMonthlyTransferSenderAmountsByCardNumber func(cards []*db.GetMonthlyTransferAmountBySenderRow) ([]*record.CardMonthAmount)
	ToYearlyTransferSenderAmountByCardNumber func(card *db.GetYearlyTransferAmountBySenderRow) (*record.CardYearAmount)
	ToYearlyTransferSenderAmountsByCardNumber func(cards []*db.GetYearlyTransferAmountBySenderRow) ([]*record.CardYearAmount)
	ToMonthlyTransferReceiverAmountByCardNumber func(card *db.GetMonthlyTransferAmountByReceiverRow) (*record.CardMonthAmount)
	ToMonthlyTransferReceiverAmountsByCardNumber func(cards []*db.GetMonthlyTransferAmountByReceiverRow) ([]*record.CardMonthAmount)
	ToYearlyTransferReceiverAmountByCardNumber func(card *db.GetYearlyTransferAmountByReceiverRow) (*record.CardYearAmount)
	ToYearlyTransferReceiverAmountsByCardNumber func(cards []*db.GetYearlyTransferAmountByReceiverRow) ([]*record.CardYearAmount)
}
```

### `MerchantDocumentMapping`

```go
type MerchantDocumentMapping interface {
	ToGetMerchantDocument func(doc *db.MerchantDocument) (*record.MerchantDocumentRecord)
	ToMerchantDocumentRecord func(doc *db.GetMerchantDocumentsRow) (*record.MerchantDocumentRecord)
	ToMerchantDocumentsRecord func(docs []*db.GetMerchantDocumentsRow) ([]*record.MerchantDocumentRecord)
	ToMerchantDocumentActiveRecord func(doc *db.GetActiveMerchantDocumentsRow) (*record.MerchantDocumentRecord)
	ToMerchantDocumentsActiveRecord func(docs []*db.GetActiveMerchantDocumentsRow) ([]*record.MerchantDocumentRecord)
	ToMerchantDocumentTrashedRecord func(doc *db.GetTrashedMerchantDocumentsRow) (*record.MerchantDocumentRecord)
	ToMerchantDocumentsTrashedRecord func(docs []*db.GetTrashedMerchantDocumentsRow) ([]*record.MerchantDocumentRecord)
}
```

### `MerchantRecordMapping`

MerchantRecordMapping defines a set of methods for mapping merchant-related
database rows into internal application records, used for reporting, logging,
analytics, and API responses.

```go
type MerchantRecordMapping interface {
	ToMerchantRecord func(merchant *db.Merchant) (*record.MerchantRecord)
	ToMerchantsRecord func(merchants []*db.Merchant) ([]*record.MerchantRecord)
	ToMerchantMonthlyTotalAmount func(ms *db.GetMonthlyTotalAmountMerchantRow) (*record.MerchantMonthlyTotalAmount)
	ToMerchantMonthlyTotalAmounts func(ms []*db.GetMonthlyTotalAmountMerchantRow) ([]*record.MerchantMonthlyTotalAmount)
	ToMerchantYearlyTotalAmount func(ms *db.GetYearlyTotalAmountMerchantRow) (*record.MerchantYearlyTotalAmount)
	ToMerchantYearlyTotalAmounts func(ms []*db.GetYearlyTotalAmountMerchantRow) ([]*record.MerchantYearlyTotalAmount)
	ToMerchantTransactionRecord func(merchant *db.FindAllTransactionsRow) (*record.MerchantTransactionsRecord)
	ToMerchantsTransactionRecord func(merchants []*db.FindAllTransactionsRow) ([]*record.MerchantTransactionsRecord)
	ToMerchantGetAllRecord func(merchant *db.GetMerchantsRow) (*record.MerchantRecord)
	ToMerchantsGetAllRecord func(merchants []*db.GetMerchantsRow) ([]*record.MerchantRecord)
	ToMerchantMonthlyPaymentMethod func(ms *db.GetMonthlyPaymentMethodsMerchantRow) (*record.MerchantMonthlyPaymentMethod)
	ToMerchantMonthlyPaymentMethods func(ms []*db.GetMonthlyPaymentMethodsMerchantRow) ([]*record.MerchantMonthlyPaymentMethod)
	ToMerchantYearlyPaymentMethod func(ms *db.GetYearlyPaymentMethodMerchantRow) (*record.MerchantYearlyPaymentMethod)
	ToMerchantYearlyPaymentMethods func(ms []*db.GetYearlyPaymentMethodMerchantRow) ([]*record.MerchantYearlyPaymentMethod)
	ToMerchantMonthlyAmount func(ms *db.GetMonthlyAmountMerchantRow) (*record.MerchantMonthlyAmount)
	ToMerchantMonthlyAmounts func(ms []*db.GetMonthlyAmountMerchantRow) ([]*record.MerchantMonthlyAmount)
	ToMerchantYearlyAmount func(ms *db.GetYearlyAmountMerchantRow) (*record.MerchantYearlyAmount)
	ToMerchantYearlyAmounts func(ms []*db.GetYearlyAmountMerchantRow) ([]*record.MerchantYearlyAmount)
	ToMerchantTransactionByMerchantRecord func(merchant *db.FindAllTransactionsByMerchantRow) (*record.MerchantTransactionsRecord)
	ToMerchantsTransactionByMerchantRecord func(merchants []*db.FindAllTransactionsByMerchantRow) ([]*record.MerchantTransactionsRecord)
	ToMerchantMonthlyPaymentMethodByMerchant func(ms *db.GetMonthlyPaymentMethodByMerchantsRow) (*record.MerchantMonthlyPaymentMethod)
	ToMerchantMonthlyPaymentMethodsByMerchant func(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) ([]*record.MerchantMonthlyPaymentMethod)
	ToMerchantYearlyPaymentMethodByMerchant func(ms *db.GetYearlyPaymentMethodByMerchantsRow) (*record.MerchantYearlyPaymentMethod)
	ToMerchantYearlyPaymentMethodsByMerchant func(ms []*db.GetYearlyPaymentMethodByMerchantsRow) ([]*record.MerchantYearlyPaymentMethod)
	ToMerchantMonthlyAmountByMerchant func(ms *db.GetMonthlyAmountByMerchantsRow) (*record.MerchantMonthlyAmount)
	ToMerchantMonthlyAmountsByMerchant func(ms []*db.GetMonthlyAmountByMerchantsRow) ([]*record.MerchantMonthlyAmount)
	ToMerchantYearlyAmountByMerchant func(ms *db.GetYearlyAmountByMerchantsRow) (*record.MerchantYearlyAmount)
	ToMerchantYearlyAmountsByMerchant func(ms []*db.GetYearlyAmountByMerchantsRow) ([]*record.MerchantYearlyAmount)
	ToMerchantMonthlyTotalAmountByMerchant func(ms *db.GetMonthlyTotalAmountByMerchantRow) (*record.MerchantMonthlyTotalAmount)
	ToMerchantMonthlyTotalAmountsByMerchant func(ms []*db.GetMonthlyTotalAmountByMerchantRow) ([]*record.MerchantMonthlyTotalAmount)
	ToMerchantYearlyTotalAmountByMerchant func(ms *db.GetYearlyTotalAmountByMerchantRow) (*record.MerchantYearlyTotalAmount)
	ToMerchantYearlyTotalAmountsByMerchant func(ms []*db.GetYearlyTotalAmountByMerchantRow) ([]*record.MerchantYearlyTotalAmount)
	ToMerchantTransactionByApikeyRecord func(merchant *db.FindAllTransactionsByApikeyRow) (*record.MerchantTransactionsRecord)
	ToMerchantsTransactionByApikeyRecord func(merchants []*db.FindAllTransactionsByApikeyRow) ([]*record.MerchantTransactionsRecord)
	ToMerchantMonthlyPaymentMethodByApikey func(ms *db.GetMonthlyPaymentMethodByApikeyRow) (*record.MerchantMonthlyPaymentMethod)
	ToMerchantMonthlyPaymentMethodsByApikey func(ms []*db.GetMonthlyPaymentMethodByApikeyRow) ([]*record.MerchantMonthlyPaymentMethod)
	ToMerchantYearlyPaymentMethodByApikey func(ms *db.GetYearlyPaymentMethodByApikeyRow) (*record.MerchantYearlyPaymentMethod)
	ToMerchantYearlyPaymentMethodsByApikey func(ms []*db.GetYearlyPaymentMethodByApikeyRow) ([]*record.MerchantYearlyPaymentMethod)
	ToMerchantMonthlyAmountByApikey func(ms *db.GetMonthlyAmountByApikeyRow) (*record.MerchantMonthlyAmount)
	ToMerchantMonthlyAmountsByApikey func(ms []*db.GetMonthlyAmountByApikeyRow) ([]*record.MerchantMonthlyAmount)
	ToMerchantYearlyAmountByApikey func(ms *db.GetYearlyAmountByApikeyRow) (*record.MerchantYearlyAmount)
	ToMerchantYearlyAmountsByApikey func(ms []*db.GetYearlyAmountByApikeyRow) ([]*record.MerchantYearlyAmount)
	ToMerchantMonthlyTotalAmountByApikey func(ms *db.GetMonthlyTotalAmountByApikeyRow) (*record.MerchantMonthlyTotalAmount)
	ToMerchantMonthlyTotalAmountsByApikey func(ms []*db.GetMonthlyTotalAmountByApikeyRow) ([]*record.MerchantMonthlyTotalAmount)
	ToMerchantYearlyTotalAmountByApikey func(ms *db.GetYearlyTotalAmountByApikeyRow) (*record.MerchantYearlyTotalAmount)
	ToMerchantYearlyTotalAmountsByApikey func(ms []*db.GetYearlyTotalAmountByApikeyRow) ([]*record.MerchantYearlyTotalAmount)
	ToMerchantActiveRecord func(merchant *db.GetActiveMerchantsRow) (*record.MerchantRecord)
	ToMerchantsActiveRecord func(merchants []*db.GetActiveMerchantsRow) ([]*record.MerchantRecord)
	ToMerchantTrashedRecord func(merchant *db.GetTrashedMerchantsRow) (*record.MerchantRecord)
	ToMerchantsTrashedRecord func(merchants []*db.GetTrashedMerchantsRow) ([]*record.MerchantRecord)
}
```

### `RecordMapper`

RecordMapper provides a set of mapping functions between database rows and
domain models for records related to the user, role, user role, refresh
token, reset token, saldo, topup, transfer, withdraw, card, transaction,
merchant, and merchant document.

```go
type RecordMapper struct {
	UserRecordMapper UserRecordMapping
	RoleRecordMapper RoleRecordMapping
	UserRoleRecordMapper UserRoleRecordMapping
	RefreshTokenRecordMapper RefreshTokenRecordMapping
	ResetTokenRecordMapper ResetTokenRecordMapping
	SaldoRecordMapper SaldoRecordMapping
	TopupRecordMapper TopupRecordMapping
	TransferRecordMapper TransferRecordMapping
	WithdrawRecordMapper WithdrawRecordMapping
	CardRecordMapper CardRecordMapping
	TransactionRecordMapper TransactionRecordMapping
	MerchantRecordMapper MerchantRecordMapping
	MerchantDocumentRecordMapper MerchantDocumentMapping
}
```

### `RefreshTokenRecordMapping`

RefreshTokenRecordMapping defines a mapping function from a RefreshToken database row to a RefreshTokenRecord domain model.

```go
type RefreshTokenRecordMapping interface {
	ToRefreshTokenRecord func(refreshToken *db.RefreshToken) (*record.RefreshTokenRecord)
	ToRefreshTokensRecord func(refreshTokens []*db.RefreshToken) ([]*record.RefreshTokenRecord)
}
```

### `ResetTokenRecordMapping`

ResetTokenRecordMapping defines a mapping function from a ResetToken database row to a ResetTokenRecord domain model.

```go
type ResetTokenRecordMapping interface {
	ToResetTokenRecord func(resetToken *db.ResetToken) (*record.ResetTokenRecord)
}
```

### `RoleRecordMapping`

RoleRecordMapping defines a mapping function from a Role database row to a RoleRecord domain model.

```go
type RoleRecordMapping interface {
	ToRoleRecord func(role *db.Role) (*record.RoleRecord)
	ToRolesRecord func(roles []*db.Role) ([]*record.RoleRecord)
	ToRoleRecordAll func(role *db.GetRolesRow) (*record.RoleRecord)
	ToRolesRecordAll func(roles []*db.GetRolesRow) ([]*record.RoleRecord)
	ToRoleRecordActive func(role *db.GetActiveRolesRow) (*record.RoleRecord)
	ToRolesRecordActive func(roles []*db.GetActiveRolesRow) ([]*record.RoleRecord)
	ToRoleRecordTrashed func(role *db.GetTrashedRolesRow) (*record.RoleRecord)
	ToRolesRecordTrashed func(roles []*db.GetTrashedRolesRow) ([]*record.RoleRecord)
}
```

### `SaldoRecordMapping`

SaldoRecordMapping provides methods to map database rows to Saldo domain models.

```go
type SaldoRecordMapping interface {
	ToSaldoRecord func(saldo *db.Saldo) (*record.SaldoRecord)
	ToSaldosRecord func(saldos []*db.Saldo) ([]*record.SaldoRecord)
	ToSaldoMonthTotalBalance func(ss *db.GetMonthlyTotalSaldoBalanceRow) (*record.SaldoMonthTotalBalance)
	ToSaldoMonthTotalBalances func(ss []*db.GetMonthlyTotalSaldoBalanceRow) ([]*record.SaldoMonthTotalBalance)
	ToSaldoYearTotalBalance func(ss *db.GetYearlyTotalSaldoBalancesRow) (*record.SaldoYearTotalBalance)
	ToSaldoYearTotalBalances func(ss []*db.GetYearlyTotalSaldoBalancesRow) ([]*record.SaldoYearTotalBalance)
	ToSaldoMonthBalance func(ss *db.GetMonthlySaldoBalancesRow) (*record.SaldoMonthSaldoBalance)
	ToSaldoMonthBalances func(ss []*db.GetMonthlySaldoBalancesRow) ([]*record.SaldoMonthSaldoBalance)
	ToSaldoYearSaldoBalance func(ss *db.GetYearlySaldoBalancesRow) (*record.SaldoYearSaldoBalance)
	ToSaldoYearSaldoBalances func(ss []*db.GetYearlySaldoBalancesRow) ([]*record.SaldoYearSaldoBalance)
	ToSaldoRecordAll func(saldo *db.GetSaldosRow) (*record.SaldoRecord)
	ToSaldosRecordAll func(saldos []*db.GetSaldosRow) ([]*record.SaldoRecord)
	ToSaldoRecordActive func(saldo *db.GetActiveSaldosRow) (*record.SaldoRecord)
	ToSaldosRecordActive func(saldos []*db.GetActiveSaldosRow) ([]*record.SaldoRecord)
	ToSaldoRecordTrashed func(saldo *db.GetTrashedSaldosRow) (*record.SaldoRecord)
	ToSaldosRecordTrashed func(saldos []*db.GetTrashedSaldosRow) ([]*record.SaldoRecord)
}
```

### `TopupRecordMapping`

TopupRecordMapping provides methods to map database rows related to top-ups to domain models.

```go
type TopupRecordMapping interface {
	ToTopupRecord func(topup *db.Topup) (*record.TopupRecord)
	ToTopupRecords func(topups []*db.Topup) ([]*record.TopupRecord)
	ToTopupByCardNumberRecord func(topup *db.GetTopupsByCardNumberRow) (*record.TopupRecord)
	ToTopupByCardNumberRecords func(topups []*db.GetTopupsByCardNumberRow) ([]*record.TopupRecord)
	ToTopupRecordMonthStatusSuccess func(s *db.GetMonthTopupStatusSuccessRow) (*record.TopupRecordMonthStatusSuccess)
	ToTopupRecordsMonthStatusSuccess func(topups []*db.GetMonthTopupStatusSuccessRow) ([]*record.TopupRecordMonthStatusSuccess)
	ToTopupRecordYearStatusSuccess func(s *db.GetYearlyTopupStatusSuccessRow) (*record.TopupRecordYearStatusSuccess)
	ToTopupRecordsYearStatusSuccess func(topups []*db.GetYearlyTopupStatusSuccessRow) ([]*record.TopupRecordYearStatusSuccess)
	ToTopupRecordMonthStatusFailed func(s *db.GetMonthTopupStatusFailedRow) (*record.TopupRecordMonthStatusFailed)
	ToTopupRecordsMonthStatusFailed func(topups []*db.GetMonthTopupStatusFailedRow) ([]*record.TopupRecordMonthStatusFailed)
	ToTopupRecordYearStatusFailed func(s *db.GetYearlyTopupStatusFailedRow) (*record.TopupRecordYearStatusFailed)
	ToTopupRecordsYearStatusFailed func(topups []*db.GetYearlyTopupStatusFailedRow) ([]*record.TopupRecordYearStatusFailed)
	ToTopupRecordMonthStatusSuccessByCardNumber func(s *db.GetMonthTopupStatusSuccessCardNumberRow) (*record.TopupRecordMonthStatusSuccess)
	ToTopupRecordsMonthStatusSuccessByCardNumber func(topups []*db.GetMonthTopupStatusSuccessCardNumberRow) ([]*record.TopupRecordMonthStatusSuccess)
	ToTopupRecordYearStatusSuccessByCardNumber func(s *db.GetYearlyTopupStatusSuccessCardNumberRow) (*record.TopupRecordYearStatusSuccess)
	ToTopupRecordsYearStatusSuccessByCardNumber func(topups []*db.GetYearlyTopupStatusSuccessCardNumberRow) ([]*record.TopupRecordYearStatusSuccess)
	ToTopupRecordMonthStatusFailedByCardNumber func(s *db.GetMonthTopupStatusFailedCardNumberRow) (*record.TopupRecordMonthStatusFailed)
	ToTopupRecordsMonthStatusFailedByCardNumber func(topups []*db.GetMonthTopupStatusFailedCardNumberRow) ([]*record.TopupRecordMonthStatusFailed)
	ToTopupRecordYearStatusFailedByCardNumber func(s *db.GetYearlyTopupStatusFailedCardNumberRow) (*record.TopupRecordYearStatusFailed)
	ToTopupRecordsYearStatusFailedByCardNumber func(topups []*db.GetYearlyTopupStatusFailedCardNumberRow) ([]*record.TopupRecordYearStatusFailed)
	ToTopupMonthlyMethod func(s *db.GetMonthlyTopupMethodsRow) (*record.TopupMonthMethod)
	ToTopupMonthlyMethods func(s []*db.GetMonthlyTopupMethodsRow) ([]*record.TopupMonthMethod)
	ToTopupYearlyMethod func(s *db.GetYearlyTopupMethodsRow) (*record.TopupYearlyMethod)
	ToTopupYearlyMethods func(s []*db.GetYearlyTopupMethodsRow) ([]*record.TopupYearlyMethod)
	ToTopupMonthlyAmount func(s *db.GetMonthlyTopupAmountsRow) (*record.TopupMonthAmount)
	ToTopupMonthlyAmounts func(s []*db.GetMonthlyTopupAmountsRow) ([]*record.TopupMonthAmount)
	ToTopupYearlyAmount func(s *db.GetYearlyTopupAmountsRow) (*record.TopupYearlyAmount)
	ToTopupYearlyAmounts func(s []*db.GetYearlyTopupAmountsRow) ([]*record.TopupYearlyAmount)
	ToTopupMonthlyMethodByCardNumber func(s *db.GetMonthlyTopupMethodsByCardNumberRow) (*record.TopupMonthMethod)
	ToTopupMonthlyMethodsByCardNumber func(s []*db.GetMonthlyTopupMethodsByCardNumberRow) ([]*record.TopupMonthMethod)
	ToTopupYearlyMethodByCardNumber func(s *db.GetYearlyTopupMethodsByCardNumberRow) (*record.TopupYearlyMethod)
	ToTopupYearlyMethodsByCardNumber func(s []*db.GetYearlyTopupMethodsByCardNumberRow) ([]*record.TopupYearlyMethod)
	ToTopupMonthlyAmountByCardNumber func(s *db.GetMonthlyTopupAmountsByCardNumberRow) (*record.TopupMonthAmount)
	ToTopupMonthlyAmountsByCardNumber func(s []*db.GetMonthlyTopupAmountsByCardNumberRow) ([]*record.TopupMonthAmount)
	ToTopupYearlyAmountByCardNumber func(s *db.GetYearlyTopupAmountsByCardNumberRow) (*record.TopupYearlyAmount)
	ToTopupYearlyAmountsByCardNumber func(s []*db.GetYearlyTopupAmountsByCardNumberRow) ([]*record.TopupYearlyAmount)
	ToTopupRecordAll func(topup *db.GetTopupsRow) (*record.TopupRecord)
	ToTopupRecordsAll func(topups []*db.GetTopupsRow) ([]*record.TopupRecord)
	ToTopupRecordActive func(topup *db.GetActiveTopupsRow) (*record.TopupRecord)
	ToTopupRecordsActive func(topups []*db.GetActiveTopupsRow) ([]*record.TopupRecord)
	ToTopupRecordTrashed func(topup *db.GetTrashedTopupsRow) (*record.TopupRecord)
	ToTopupRecordsTrashed func(topups []*db.GetTrashedTopupsRow) ([]*record.TopupRecord)
}
```

### `TransactionRecordMapping`

TransactionRecordMapping defines a set of functions for converting transaction-related
database rows into application-level record structures used for processing,
logging, and API responses.

```go
type TransactionRecordMapping interface {
	ToTransactionRecord func(transaction *db.Transaction) (*record.TransactionRecord)
	ToTransactionsRecord func(transactions []*db.Transaction) ([]*record.TransactionRecord)
	ToTransactionByCardNumberRecord func(transaction *db.GetTransactionsByCardNumberRow) (*record.TransactionRecord)
	ToTransactionsByCardNumberRecord func(transactions []*db.GetTransactionsByCardNumberRow) ([]*record.TransactionRecord)
	ToTransactionRecordMonthStatusSuccess func(s *db.GetMonthTransactionStatusSuccessRow) (*record.TransactionRecordMonthStatusSuccess)
	ToTransactionRecordsMonthStatusSuccess func(transactions []*db.GetMonthTransactionStatusSuccessRow) ([]*record.TransactionRecordMonthStatusSuccess)
	ToTransactionRecordYearStatusSuccess func(s *db.GetYearlyTransactionStatusSuccessRow) (*record.TransactionRecordYearStatusSuccess)
	ToTransactionRecordsYearStatusSuccess func(transactions []*db.GetYearlyTransactionStatusSuccessRow) ([]*record.TransactionRecordYearStatusSuccess)
	ToTransactionRecordMonthStatusFailed func(s *db.GetMonthTransactionStatusFailedRow) (*record.TransactionRecordMonthStatusFailed)
	ToTransactionRecordsMonthStatusFailed func(transactions []*db.GetMonthTransactionStatusFailedRow) ([]*record.TransactionRecordMonthStatusFailed)
	ToTransactionRecordYearStatusFailed func(s *db.GetYearlyTransactionStatusFailedRow) (*record.TransactionRecordYearStatusFailed)
	ToTransactionRecordsYearStatusFailed func(transactions []*db.GetYearlyTransactionStatusFailedRow) ([]*record.TransactionRecordYearStatusFailed)
	ToTransactionRecordMonthStatusSuccessCardNumber func(s *db.GetMonthTransactionStatusSuccessCardNumberRow) (*record.TransactionRecordMonthStatusSuccess)
	ToTransactionRecordsMonthStatusSuccessCardNumber func(transactions []*db.GetMonthTransactionStatusSuccessCardNumberRow) ([]*record.TransactionRecordMonthStatusSuccess)
	ToTransactionRecordYearStatusSuccessCardNumber func(s *db.GetYearlyTransactionStatusSuccessCardNumberRow) (*record.TransactionRecordYearStatusSuccess)
	ToTransactionRecordsYearStatusSuccessCardNumber func(transactions []*db.GetYearlyTransactionStatusSuccessCardNumberRow) ([]*record.TransactionRecordYearStatusSuccess)
	ToTransactionRecordMonthStatusFailedCardNumber func(s *db.GetMonthTransactionStatusFailedCardNumberRow) (*record.TransactionRecordMonthStatusFailed)
	ToTransactionRecordsMonthStatusFailedCardNumber func(transactions []*db.GetMonthTransactionStatusFailedCardNumberRow) ([]*record.TransactionRecordMonthStatusFailed)
	ToTransactionRecordYearStatusFailedCardNumber func(s *db.GetYearlyTransactionStatusFailedCardNumberRow) (*record.TransactionRecordYearStatusFailed)
	ToTransactionRecordsYearStatusFailedCardNumber func(transactions []*db.GetYearlyTransactionStatusFailedCardNumberRow) ([]*record.TransactionRecordYearStatusFailed)
	ToTransactionMonthlyMethod func(ss *db.GetMonthlyPaymentMethodsRow) (*record.TransactionMonthMethod)
	ToTransactionMonthlyMethods func(ss []*db.GetMonthlyPaymentMethodsRow) ([]*record.TransactionMonthMethod)
	ToTransactionYearlyMethod func(ss *db.GetYearlyPaymentMethodsRow) (*record.TransactionYearMethod)
	ToTransactionYearlyMethods func(ss []*db.GetYearlyPaymentMethodsRow) ([]*record.TransactionYearMethod)
	ToTransactionMonthlyAmount func(ss *db.GetMonthlyAmountsRow) (*record.TransactionMonthAmount)
	ToTransactionMonthlyAmounts func(ss []*db.GetMonthlyAmountsRow) ([]*record.TransactionMonthAmount)
	ToTransactionYearlyAmount func(ss *db.GetYearlyAmountsRow) (*record.TransactionYearlyAmount)
	ToTransactionYearlyAmounts func(ss []*db.GetYearlyAmountsRow) ([]*record.TransactionYearlyAmount)
	ToTransactionMonthlyMethodByCardNumber func(ss *db.GetMonthlyPaymentMethodsByCardNumberRow) (*record.TransactionMonthMethod)
	ToTransactionMonthlyMethodsByCardNumber func(ss []*db.GetMonthlyPaymentMethodsByCardNumberRow) ([]*record.TransactionMonthMethod)
	ToTransactionYearlyMethodByCardNumber func(ss *db.GetYearlyPaymentMethodsByCardNumberRow) (*record.TransactionYearMethod)
	ToTransactionYearlyMethodsByCardNumber func(ss []*db.GetYearlyPaymentMethodsByCardNumberRow) ([]*record.TransactionYearMethod)
	ToTransactionMonthlyAmountByCardNumber func(ss *db.GetMonthlyAmountsByCardNumberRow) (*record.TransactionMonthAmount)
	ToTransactionMonthlyAmountsByCardNumber func(ss []*db.GetMonthlyAmountsByCardNumberRow) ([]*record.TransactionMonthAmount)
	ToTransactionYearlyAmountByCardNumber func(ss *db.GetYearlyAmountsByCardNumberRow) (*record.TransactionYearlyAmount)
	ToTransactionYearlyAmountsByCardNumber func(ss []*db.GetYearlyAmountsByCardNumberRow) ([]*record.TransactionYearlyAmount)
	ToTransactionRecordAll func(transaction *db.GetTransactionsRow) (*record.TransactionRecord)
	ToTransactionsRecordAll func(transactions []*db.GetTransactionsRow) ([]*record.TransactionRecord)
	ToTransactionRecordActive func(transaction *db.GetActiveTransactionsRow) (*record.TransactionRecord)
	ToTransactionsRecordActive func(transactions []*db.GetActiveTransactionsRow) ([]*record.TransactionRecord)
	ToTransactionRecordTrashed func(transaction *db.GetTrashedTransactionsRow) (*record.TransactionRecord)
	ToTransactionsRecordTrashed func(transactions []*db.GetTrashedTransactionsRow) ([]*record.TransactionRecord)
}
```

### `TransferRecordMapping`

```go
type TransferRecordMapping interface {
	ToTransferRecord func(transfer *db.Transfer) (*record.TransferRecord)
	ToTransfersRecord func(transfers []*db.Transfer) ([]*record.TransferRecord)
	ToTransferRecordMonthStatusSuccess func(s *db.GetMonthTransferStatusSuccessRow) (*record.TransferRecordMonthStatusSuccess)
	ToTransferRecordsMonthStatusSuccess func(Transfers []*db.GetMonthTransferStatusSuccessRow) ([]*record.TransferRecordMonthStatusSuccess)
	ToTransferRecordYearStatusSuccess func(s *db.GetYearlyTransferStatusSuccessRow) (*record.TransferRecordYearStatusSuccess)
	ToTransferRecordsYearStatusSuccess func(Transfers []*db.GetYearlyTransferStatusSuccessRow) ([]*record.TransferRecordYearStatusSuccess)
	ToTransferRecordMonthStatusFailed func(s *db.GetMonthTransferStatusFailedRow) (*record.TransferRecordMonthStatusFailed)
	ToTransferRecordsMonthStatusFailed func(Transfers []*db.GetMonthTransferStatusFailedRow) ([]*record.TransferRecordMonthStatusFailed)
	ToTransferRecordYearStatusFailed func(s *db.GetYearlyTransferStatusFailedRow) (*record.TransferRecordYearStatusFailed)
	ToTransferRecordsYearStatusFailed func(Transfers []*db.GetYearlyTransferStatusFailedRow) ([]*record.TransferRecordYearStatusFailed)
	ToTransferRecordMonthStatusSuccessCardNumber func(s *db.GetMonthTransferStatusSuccessCardNumberRow) (*record.TransferRecordMonthStatusSuccess)
	ToTransferRecordsMonthStatusSuccessCardNumber func(Transfers []*db.GetMonthTransferStatusSuccessCardNumberRow) ([]*record.TransferRecordMonthStatusSuccess)
	ToTransferRecordYearStatusSuccessCardNumber func(s *db.GetYearlyTransferStatusSuccessCardNumberRow) (*record.TransferRecordYearStatusSuccess)
	ToTransferRecordsYearStatusSuccessCardNumber func(Transfers []*db.GetYearlyTransferStatusSuccessCardNumberRow) ([]*record.TransferRecordYearStatusSuccess)
	ToTransferRecordMonthStatusFailedCardNumber func(s *db.GetMonthTransferStatusFailedCardNumberRow) (*record.TransferRecordMonthStatusFailed)
	ToTransferRecordsMonthStatusFailedCardNumber func(Transfers []*db.GetMonthTransferStatusFailedCardNumberRow) ([]*record.TransferRecordMonthStatusFailed)
	ToTransferRecordYearStatusFailedCardNumber func(s *db.GetYearlyTransferStatusFailedCardNumberRow) (*record.TransferRecordYearStatusFailed)
	ToTransferRecordsYearStatusFailedCardNumber func(Transfers []*db.GetYearlyTransferStatusFailedCardNumberRow) ([]*record.TransferRecordYearStatusFailed)
	ToTransferMonthAmount func(ss *db.GetMonthlyTransferAmountsRow) (*record.TransferMonthAmount)
	ToTransferMonthAmounts func(ss []*db.GetMonthlyTransferAmountsRow) ([]*record.TransferMonthAmount)
	ToTransferYearAmount func(ss *db.GetYearlyTransferAmountsRow) (*record.TransferYearAmount)
	ToTransferYearAmounts func(ss []*db.GetYearlyTransferAmountsRow) ([]*record.TransferYearAmount)
	ToTransferMonthAmountSender func(ss *db.GetMonthlyTransferAmountsBySenderCardNumberRow) (*record.TransferMonthAmount)
	ToTransferMonthAmountsSender func(ss []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) ([]*record.TransferMonthAmount)
	ToTransferYearAmountSender func(ss *db.GetYearlyTransferAmountsBySenderCardNumberRow) (*record.TransferYearAmount)
	ToTransferYearAmountsSender func(ss []*db.GetYearlyTransferAmountsBySenderCardNumberRow) ([]*record.TransferYearAmount)
	ToTransferMonthAmountReceiver func(ss *db.GetMonthlyTransferAmountsByReceiverCardNumberRow) (*record.TransferMonthAmount)
	ToTransferMonthAmountsReceiver func(ss []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) ([]*record.TransferMonthAmount)
	ToTransferYearAmountReceiver func(ss *db.GetYearlyTransferAmountsByReceiverCardNumberRow) (*record.TransferYearAmount)
	ToTransferYearAmountsReceiver func(ss []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) ([]*record.TransferYearAmount)
	ToTransferRecordAll func(transfer *db.GetTransfersRow) (*record.TransferRecord)
	ToTransfersRecordAll func(transfers []*db.GetTransfersRow) ([]*record.TransferRecord)
	ToTransferRecordActive func(transfer *db.GetActiveTransfersRow) (*record.TransferRecord)
	ToTransfersRecordActive func(transfers []*db.GetActiveTransfersRow) ([]*record.TransferRecord)
	ToTransferRecordTrashed func(transfer *db.GetTrashedTransfersRow) (*record.TransferRecord)
	ToTransfersRecordTrashed func(transfers []*db.GetTrashedTransfersRow) ([]*record.TransferRecord)
}
```

### `UserRecordMapping`

```go
type UserRecordMapping interface {
	ToUserRecord func(user *db.User) (*record.UserRecord)
	ToUserRecordPagination func(user *db.GetUsersWithPaginationRow) (*record.UserRecord)
	ToUsersRecordPagination func(users []*db.GetUsersWithPaginationRow) ([]*record.UserRecord)
	ToUserRecordActivePagination func(user *db.GetActiveUsersWithPaginationRow) (*record.UserRecord)
	ToUsersRecordActivePagination func(users []*db.GetActiveUsersWithPaginationRow) ([]*record.UserRecord)
	ToUserRecordTrashedPagination func(user *db.GetTrashedUsersWithPaginationRow) (*record.UserRecord)
	ToUsersRecordTrashedPagination func(users []*db.GetTrashedUsersWithPaginationRow) ([]*record.UserRecord)
}
```

### `UserRoleRecordMapping`

UserRoleRecordMapping defines a mapping function from a UserRole database row to a UserRoleRecord domain model.

```go
type UserRoleRecordMapping interface {
	ToUserRoleRecord func(userRole *db.UserRole) (*record.UserRoleRecord)
}
```

### `WithdrawRecordMapping`

WithdrawRecordMapping defines a set of functions for mapping database query results
into internal record representations used across different application layers
(e.g., response rendering, logging, business logic processing).

```go
type WithdrawRecordMapping interface {
	ToWithdrawRecord func(withdraw *db.Withdraw) (*record.WithdrawRecord)
	ToWithdrawsRecord func(withdraws []*db.Withdraw) ([]*record.WithdrawRecord)
	ToWithdrawByCardNumberRecord func(withdraw *db.GetWithdrawsByCardNumberRow) (*record.WithdrawRecord)
	ToWithdrawsByCardNumberRecord func(withdraws []*db.GetWithdrawsByCardNumberRow) ([]*record.WithdrawRecord)
	ToWithdrawRecordMonthStatusSuccess func(s *db.GetMonthWithdrawStatusSuccessRow) (*record.WithdrawRecordMonthStatusSuccess)
	ToWithdrawRecordsMonthStatusSuccess func(withdraws []*db.GetMonthWithdrawStatusSuccessRow) ([]*record.WithdrawRecordMonthStatusSuccess)
	ToWithdrawRecordYearStatusSuccess func(s *db.GetYearlyWithdrawStatusSuccessRow) (*record.WithdrawRecordYearStatusSuccess)
	ToWithdrawRecordsYearStatusSuccess func(withdraws []*db.GetYearlyWithdrawStatusSuccessRow) ([]*record.WithdrawRecordYearStatusSuccess)
	ToWithdrawRecordMonthStatusFailed func(s *db.GetMonthWithdrawStatusFailedRow) (*record.WithdrawRecordMonthStatusFailed)
	ToWithdrawRecordsMonthStatusFailed func(withdraws []*db.GetMonthWithdrawStatusFailedRow) ([]*record.WithdrawRecordMonthStatusFailed)
	ToWithdrawRecordYearStatusFailed func(s *db.GetYearlyWithdrawStatusFailedRow) (*record.WithdrawRecordYearStatusFailed)
	ToWithdrawRecordsYearStatusFailed func(withdraws []*db.GetYearlyWithdrawStatusFailedRow) ([]*record.WithdrawRecordYearStatusFailed)
	ToWithdrawRecordMonthStatusSuccessCardNumber func(s *db.GetMonthWithdrawStatusSuccessCardNumberRow) (*record.WithdrawRecordMonthStatusSuccess)
	ToWithdrawRecordsMonthStatusSuccessCardNumber func(withdraws []*db.GetMonthWithdrawStatusSuccessCardNumberRow) ([]*record.WithdrawRecordMonthStatusSuccess)
	ToWithdrawRecordYearStatusSuccessCardNumber func(s *db.GetYearlyWithdrawStatusSuccessCardNumberRow) (*record.WithdrawRecordYearStatusSuccess)
	ToWithdrawRecordsYearStatusSuccessCardNumber func(withdraws []*db.GetYearlyWithdrawStatusSuccessCardNumberRow) ([]*record.WithdrawRecordYearStatusSuccess)
	ToWithdrawRecordMonthStatusFailedCardNumber func(s *db.GetMonthWithdrawStatusFailedCardNumberRow) (*record.WithdrawRecordMonthStatusFailed)
	ToWithdrawRecordsMonthStatusFailedCardNumber func(withdraws []*db.GetMonthWithdrawStatusFailedCardNumberRow) ([]*record.WithdrawRecordMonthStatusFailed)
	ToWithdrawRecordYearStatusFailedCardNumber func(s *db.GetYearlyWithdrawStatusFailedCardNumberRow) (*record.WithdrawRecordYearStatusFailed)
	ToWithdrawRecordsYearStatusFailedCardNumber func(withdraws []*db.GetYearlyWithdrawStatusFailedCardNumberRow) ([]*record.WithdrawRecordYearStatusFailed)
	ToWithdrawAmountMonthly func(ss *db.GetMonthlyWithdrawsRow) (*record.WithdrawMonthlyAmount)
	ToWithdrawsAmountMonthly func(ss []*db.GetMonthlyWithdrawsRow) ([]*record.WithdrawMonthlyAmount)
	ToWithdrawAmountYearly func(ss *db.GetYearlyWithdrawsRow) (*record.WithdrawYearlyAmount)
	ToWithdrawsAmountYearly func(ss []*db.GetYearlyWithdrawsRow) ([]*record.WithdrawYearlyAmount)
	ToWithdrawAmountMonthlyByCardNumber func(ss *db.GetMonthlyWithdrawsByCardNumberRow) (*record.WithdrawMonthlyAmount)
	ToWithdrawsAmountMonthlyByCardNumber func(ss []*db.GetMonthlyWithdrawsByCardNumberRow) ([]*record.WithdrawMonthlyAmount)
	ToWithdrawAmountYearlyByCardNumber func(ss *db.GetYearlyWithdrawsByCardNumberRow) (*record.WithdrawYearlyAmount)
	ToWithdrawsAmountYearlyByCardNumber func(ss []*db.GetYearlyWithdrawsByCardNumberRow) ([]*record.WithdrawYearlyAmount)
	ToWithdrawRecordAll func(withdraw *db.GetWithdrawsRow) (*record.WithdrawRecord)
	ToWithdrawsRecordALl func(withdraws []*db.GetWithdrawsRow) ([]*record.WithdrawRecord)
	ToWithdrawRecordActive func(withdraw *db.GetActiveWithdrawsRow) (*record.WithdrawRecord)
	ToWithdrawsRecordActive func(withdraws []*db.GetActiveWithdrawsRow) ([]*record.WithdrawRecord)
	ToWithdrawRecordTrashed func(withdraw *db.GetTrashedWithdrawsRow) (*record.WithdrawRecord)
	ToWithdrawsRecordTrashed func(withdraws []*db.GetTrashedWithdrawsRow) ([]*record.WithdrawRecord)
}
```

### `cardRecordMapper`

cardRecordMapper provides methods to map Card database rows to CardRecord domain models.

```go
type cardRecordMapper struct {
}
```

#### Methods

##### `ToCardEmailRecord`

ToCardEmailRecord maps a GetUserEmailByCardNumberRow database row to a CardEmailRecord domain model.

Args:
  - card: A pointer to a GetUserEmailByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardEmailRecord containing the mapped data, including ID, UserID,
    Email, CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.

```go
func (s *cardRecordMapper) ToCardEmailRecord(card *db.GetUserEmailByCardNumberRow) *record.CardEmailRecord
```

##### `ToCardGetAll`

ToCardGetAll maps a GetCardsRow database row to a CardRecord domain model.

Args:
  - card: A pointer to a GetCardsRow representing the database row.

Returns:
  - A pointer to a CardRecord containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardRecordMapper) ToCardGetAll(card *db.GetCardsRow) *record.CardRecord
```

##### `ToCardRecord`

ToCardRecord maps a Card database row to a CardRecord domain model.

Args:
  - card: A pointer to a Card representing the database row.

Returns:
  - A pointer to a CardRecord containing the mapped data, including ID, UserID, CardNumber,
    CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardRecordMapper) ToCardRecord(card *db.Card) *record.CardRecord
```

##### `ToCardRecordActive`

ToCardRecordActive maps a GetActiveCardsWithCountRow database row to a CardRecord domain model.

Args:
  - card: A pointer to a GetActiveCardsWithCountRow representing the database row.

Returns:
  - A pointer to a CardRecord containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardRecordMapper) ToCardRecordActive(card *db.GetActiveCardsWithCountRow) *record.CardRecord
```

##### `ToCardRecordTrashed`

ToCardRecordTrashed maps a GetTrashedCardsWithCountRow database row to a CardRecord domain model.

Args:
  - card: A pointer to a GetTrashedCardsWithCountRow representing the database row.

Returns:
  - A pointer to a CardRecord containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardRecordMapper) ToCardRecordTrashed(card *db.GetTrashedCardsWithCountRow) *record.CardRecord
```

##### `ToCardRecordsActive`

ToCardRecordsActive maps a slice of GetActiveCardsWithCountRow database rows to a slice of CardRecord domain models.

Args:
  - cards: A slice of pointers to GetActiveCardsWithCountRow representing the database rows.

Returns:
  - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardRecordMapper) ToCardRecordsActive(cards []*db.GetActiveCardsWithCountRow) []*record.CardRecord
```

##### `ToCardRecordsTrashed`

ToCardRecordsTrashed maps a slice of GetTrashedCardsWithCountRow database rows to a slice of CardRecord domain models.

Args:
  - cards: A slice of pointers to GetTrashedCardsWithCountRow representing the database rows.

Returns:
  - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardRecordMapper) ToCardRecordsTrashed(cards []*db.GetTrashedCardsWithCountRow) []*record.CardRecord
```

##### `ToCardsRecord`

ToCardsRecord maps a slice of GetCardsRow database rows to a slice of CardRecord domain models.

Args:
  - cards: A slice of pointers to GetCardsRow representing the database rows.

Returns:
  - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
    CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *cardRecordMapper) ToCardsRecord(cards []*db.GetCardsRow) []*record.CardRecord
```

##### `ToMonthlyBalance`

ToMonthlyBalance maps a GetMonthlyBalancesRow database row to a CardMonthBalance domain model.

Args:
  - card: A pointer to a GetMonthlyBalancesRow representing the database row.

Returns:
  - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.

```go
func (s *cardRecordMapper) ToMonthlyBalance(card *db.GetMonthlyBalancesRow) *record.CardMonthBalance
```

##### `ToMonthlyBalanceCardNumber`

ToMonthlyBalanceCardNumber maps a GetMonthlyBalancesByCardNumberRow database row to a CardMonthBalance domain model.

Args:
  - card: A pointer to a GetMonthlyBalancesByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.

```go
func (s *cardRecordMapper) ToMonthlyBalanceCardNumber(card *db.GetMonthlyBalancesByCardNumberRow) *record.CardMonthBalance
```

##### `ToMonthlyBalances`

ToMonthlyBalances maps a slice of GetMonthlyBalancesRow database rows to a slice of CardMonthBalance domain models.

Args:
  - cards: A slice of pointers to GetMonthlyBalancesRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.

```go
func (s *cardRecordMapper) ToMonthlyBalances(cards []*db.GetMonthlyBalancesRow) []*record.CardMonthBalance
```

##### `ToMonthlyBalancesCardNumber`

ToMonthlyBalancesCardNumber maps a slice of GetMonthlyBalancesByCardNumberRow database rows to a slice of CardMonthBalance domain models.

Args:
  - cards: A slice of pointers to GetMonthlyBalancesByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.

```go
func (s *cardRecordMapper) ToMonthlyBalancesCardNumber(cards []*db.GetMonthlyBalancesByCardNumberRow) []*record.CardMonthBalance
```

##### `ToMonthlyTopupAmount`

ToMonthlyTopupAmount maps a GetMonthlyTopupAmountRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTopupAmountRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTopupAmount(card *db.GetMonthlyTopupAmountRow) *record.CardMonthAmount
```

##### `ToMonthlyTopupAmountByCardNumber`

ToMonthlyTopupAmountByCardNumber maps a GetMonthlyTopupAmountByCardNumberRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTopupAmountByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTopupAmountByCardNumber(card *db.GetMonthlyTopupAmountByCardNumberRow) *record.CardMonthAmount
```

##### `ToMonthlyTopupAmounts`

ToMonthlyTopupAmounts maps a slice of GetMonthlyTopupAmountRow database rows to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTopupAmountRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTopupAmounts(cards []*db.GetMonthlyTopupAmountRow) []*record.CardMonthAmount
```

##### `ToMonthlyTopupAmountsByCardNumber`

ToMonthlyTopupAmountsByCardNumber maps a slice of GetMonthlyTopupAmountByCardNumberRow database rows to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTopupAmountByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTopupAmountsByCardNumber(cards []*db.GetMonthlyTopupAmountByCardNumberRow) []*record.CardMonthAmount
```

##### `ToMonthlyTransactionAmount`

ToMonthlyTransactionAmount maps a GetMonthlyTransactionAmountRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTransactionAmountRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransactionAmount(card *db.GetMonthlyTransactionAmountRow) *record.CardMonthAmount
```

##### `ToMonthlyTransactionAmountByCardNumber`

ToMonthlyTransactionAmountByCardNumber maps a GetMonthlyTransactionAmountByCardNumberRow database row
to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTransactionAmountByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransactionAmountByCardNumber(card *db.GetMonthlyTransactionAmountByCardNumberRow) *record.CardMonthAmount
```

##### `ToMonthlyTransactionAmounts`

ToMonthlyTransactionAmounts maps a slice of GetMonthlyTransactionAmountRow database rows to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTransactionAmountRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransactionAmounts(cards []*db.GetMonthlyTransactionAmountRow) []*record.CardMonthAmount
```

##### `ToMonthlyTransactionAmountsByCardNumber`

ToMonthlyTransactionAmountsByCardNumber maps a slice of GetMonthlyTransactionAmountByCardNumberRow database rows
to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTransactionAmountByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransactionAmountsByCardNumber(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) []*record.CardMonthAmount
```

##### `ToMonthlyTransferReceiverAmount`

ToMonthlyTransferReceiverAmount maps a GetMonthlyTransferAmountReceiverRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTransferAmountReceiverRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmount(card *db.GetMonthlyTransferAmountReceiverRow) *record.CardMonthAmount
```

##### `ToMonthlyTransferReceiverAmountByCardNumber`

ToMonthlyTransferReceiverAmountByCardNumber maps a GetMonthlyTransferAmountByReceiverRow database row
to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTransferAmountByReceiverRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmountByCardNumber(card *db.GetMonthlyTransferAmountByReceiverRow) *record.CardMonthAmount
```

##### `ToMonthlyTransferReceiverAmounts`

ToMonthlyTransferReceiverAmounts maps a slice of GetMonthlyTransferAmountReceiverRow database rows to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTransferAmountReceiverRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmounts(cards []*db.GetMonthlyTransferAmountReceiverRow) []*record.CardMonthAmount
```

##### `ToMonthlyTransferReceiverAmountsByCardNumber`

ToMonthlyTransferReceiverAmountsByCardNumber maps a slice of GetMonthlyTransferAmountByReceiverRow database rows
to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTransferAmountByReceiverRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountByReceiverRow) []*record.CardMonthAmount
```

##### `ToMonthlyTransferSenderAmount`

ToMonthlyTransferSenderAmount maps a GetMonthlyTransferAmountSenderRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTransferAmountSenderRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferSenderAmount(card *db.GetMonthlyTransferAmountSenderRow) *record.CardMonthAmount
```

##### `ToMonthlyTransferSenderAmountByCardNumber`

ToMonthlyTransferSenderAmountByCardNumber maps a GetMonthlyTransferAmountBySenderRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyTransferAmountBySenderRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferSenderAmountByCardNumber(card *db.GetMonthlyTransferAmountBySenderRow) *record.CardMonthAmount
```

##### `ToMonthlyTransferSenderAmounts`

ToMonthlyTransferSenderAmounts maps a slice of GetMonthlyTransferAmountSenderRow database rows to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTransferAmountSenderRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferSenderAmounts(cards []*db.GetMonthlyTransferAmountSenderRow) []*record.CardMonthAmount
```

##### `ToMonthlyTransferSenderAmountsByCardNumber`

ToMonthlyTransferSenderAmountsByCardNumber maps a slice of GetMonthlyTransferAmountBySenderRow database rows
to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyTransferAmountBySenderRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyTransferSenderAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountBySenderRow) []*record.CardMonthAmount
```

##### `ToMonthlyWithdrawAmount`

ToMonthlyWithdrawAmount maps a GetMonthlyWithdrawAmountRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyWithdrawAmountRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyWithdrawAmount(card *db.GetMonthlyWithdrawAmountRow) *record.CardMonthAmount
```

##### `ToMonthlyWithdrawAmountByCardNumber`

ToMonthlyWithdrawAmountByCardNumber maps a GetMonthlyWithdrawAmountByCardNumberRow database row to a CardMonthAmount domain model.

Args:
  - card: A pointer to a GetMonthlyWithdrawAmountByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyWithdrawAmountByCardNumber(card *db.GetMonthlyWithdrawAmountByCardNumberRow) *record.CardMonthAmount
```

##### `ToMonthlyWithdrawAmounts`

ToMonthlyWithdrawAmounts maps a slice of GetMonthlyWithdrawAmountRow database rows to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyWithdrawAmountRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyWithdrawAmounts(cards []*db.GetMonthlyWithdrawAmountRow) []*record.CardMonthAmount
```

##### `ToMonthlyWithdrawAmountsByCardNumber`

ToMonthlyWithdrawAmountsByCardNumber maps a slice of GetMonthlyWithdrawAmountByCardNumberRow database rows
to a slice of CardMonthAmount domain models.

Args:
  - cards: A slice of pointers to GetMonthlyWithdrawAmountByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *cardRecordMapper) ToMonthlyWithdrawAmountsByCardNumber(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) []*record.CardMonthAmount
```

##### `ToYearlyBalance`

ToYearlyBalance maps a GetYearlyBalancesRow database row to a CardYearlyBalance domain model.

Args:
  - card: A pointer to a GetYearlyBalancesRow representing the database row.

Returns:
  - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.

```go
func (s *cardRecordMapper) ToYearlyBalance(card *db.GetYearlyBalancesRow) *record.CardYearlyBalance
```

##### `ToYearlyBalanceCardNumber`

ToYearlyBalanceCardNumber maps a GetYearlyBalancesByCardNumberRow database row to a CardYearlyBalance domain model.

Args:
  - card: A pointer to a GetYearlyBalancesByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.

```go
func (s *cardRecordMapper) ToYearlyBalanceCardNumber(card *db.GetYearlyBalancesByCardNumberRow) *record.CardYearlyBalance
```

##### `ToYearlyBalances`

ToYearlyBalances maps a slice of GetYearlyBalancesRow database rows to a slice of CardYearlyBalance domain models.

Args:
  - cards: A slice of pointers to GetYearlyBalancesRow representing the database rows.

Returns:
  - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.

```go
func (s *cardRecordMapper) ToYearlyBalances(cards []*db.GetYearlyBalancesRow) []*record.CardYearlyBalance
```

##### `ToYearlyBalancesCardNumber`

ToYearlyBalancesCardNumber maps a slice of GetYearlyBalancesByCardNumberRow database rows to a slice of CardYearlyBalance domain models.

Args:
  - cards: A slice of pointers to GetYearlyBalancesByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.

```go
func (s *cardRecordMapper) ToYearlyBalancesCardNumber(cards []*db.GetYearlyBalancesByCardNumberRow) []*record.CardYearlyBalance
```

##### `ToYearlyTopupAmount`

ToYearlyTopupAmount maps a GetYearlyTopupAmountRow database row to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTopupAmountRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTopupAmount(card *db.GetYearlyTopupAmountRow) *record.CardYearAmount
```

##### `ToYearlyTopupAmountByCardNumber`

ToYearlyTopupAmountByCardNumber maps a GetYearlyTopupAmountByCardNumberRow database row to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTopupAmountByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTopupAmountByCardNumber(card *db.GetYearlyTopupAmountByCardNumberRow) *record.CardYearAmount
```

##### `ToYearlyTopupAmounts`

ToYearlyTopupAmounts maps a slice of GetYearlyTopupAmountRow database rows to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTopupAmountRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTopupAmounts(cards []*db.GetYearlyTopupAmountRow) []*record.CardYearAmount
```

##### `ToYearlyTopupAmountsByCardNumber`

ToYearlyTopupAmountsByCardNumber maps a slice of GetYearlyTopupAmountByCardNumberRow database rows
to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTopupAmountByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTopupAmountsByCardNumber(cards []*db.GetYearlyTopupAmountByCardNumberRow) []*record.CardYearAmount
```

##### `ToYearlyTransactionAmount`

ToYearlyTransactionAmount maps a GetYearlyTransactionAmountRow database row to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTransactionAmountRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransactionAmount(card *db.GetYearlyTransactionAmountRow) *record.CardYearAmount
```

##### `ToYearlyTransactionAmountByCardNumber`

ToYearlyTransactionAmountByCardNumber maps a GetYearlyTransactionAmountByCardNumberRow database row
to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTransactionAmountByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransactionAmountByCardNumber(card *db.GetYearlyTransactionAmountByCardNumberRow) *record.CardYearAmount
```

##### `ToYearlyTransactionAmounts`

ToYearlyTransactionAmounts maps a slice of GetYearlyTransactionAmountRow database rows to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTransactionAmountRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransactionAmounts(cards []*db.GetYearlyTransactionAmountRow) []*record.CardYearAmount
```

##### `ToYearlyTransactionAmountsByCardNumber`

ToYearlyTransactionAmountsByCardNumber maps a slice of GetYearlyTransactionAmountByCardNumberRow database rows
to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTransactionAmountByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransactionAmountsByCardNumber(cards []*db.GetYearlyTransactionAmountByCardNumberRow) []*record.CardYearAmount
```

##### `ToYearlyTransferReceiverAmount`

ToYearlyTransferReceiverAmount maps a GetYearlyTransferAmountReceiverRow database row to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTransferAmountReceiverRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferReceiverAmount(card *db.GetYearlyTransferAmountReceiverRow) *record.CardYearAmount
```

##### `ToYearlyTransferReceiverAmountByCardNumber`

ToYearlyTransferReceiverAmountByCardNumber maps a GetYearlyTransferAmountByReceiverRow database row
to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTransferAmountByReceiverRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferReceiverAmountByCardNumber(card *db.GetYearlyTransferAmountByReceiverRow) *record.CardYearAmount
```

##### `ToYearlyTransferReceiverAmounts`

ToYearlyTransferReceiverAmounts maps a slice of GetYearlyTransferAmountReceiverRow database rows to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTransferAmountReceiverRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferReceiverAmounts(cards []*db.GetYearlyTransferAmountReceiverRow) []*record.CardYearAmount
```

##### `ToYearlyTransferReceiverAmountsByCardNumber`

ToYearlyTransferReceiverAmountsByCardNumber maps a slice of GetYearlyTransferAmountByReceiverRow database rows
to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTransferAmountByReceiverRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferReceiverAmountsByCardNumber(cards []*db.GetYearlyTransferAmountByReceiverRow) []*record.CardYearAmount
```

##### `ToYearlyTransferSenderAmount`

ToYearlyTransferSenderAmount maps a GetYearlyTransferAmountSenderRow database row to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTransferAmountSenderRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferSenderAmount(card *db.GetYearlyTransferAmountSenderRow) *record.CardYearAmount
```

##### `ToYearlyTransferSenderAmountByCardNumber`

ToYearlyTransferSenderAmountByCardNumber maps a GetYearlyTransferAmountBySenderRow database row
to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyTransferAmountBySenderRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferSenderAmountByCardNumber(card *db.GetYearlyTransferAmountBySenderRow) *record.CardYearAmount
```

##### `ToYearlyTransferSenderAmounts`

ToYearlyTransferSenderAmounts maps a slice of GetYearlyTransferAmountSenderRow database rows to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTransferAmountSenderRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferSenderAmounts(cards []*db.GetYearlyTransferAmountSenderRow) []*record.CardYearAmount
```

##### `ToYearlyTransferSenderAmountsByCardNumber`

ToYearlyTransferSenderAmountsByCardNumber maps a slice of GetYearlyTransferAmountBySenderRow database rows
to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyTransferAmountBySenderRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyTransferSenderAmountsByCardNumber(cards []*db.GetYearlyTransferAmountBySenderRow) []*record.CardYearAmount
```

##### `ToYearlyWithdrawAmount`

ToYearlyWithdrawAmount maps a GetYearlyWithdrawAmountRow database row to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyWithdrawAmountRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyWithdrawAmount(card *db.GetYearlyWithdrawAmountRow) *record.CardYearAmount
```

##### `ToYearlyWithdrawAmountByCardNumber`

ToYearlyWithdrawAmountByCardNumber maps a GetYearlyWithdrawAmountByCardNumberRow database row to a CardYearAmount domain model.

Args:
  - card: A pointer to a GetYearlyWithdrawAmountByCardNumberRow representing the database row.

Returns:
  - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyWithdrawAmountByCardNumber(card *db.GetYearlyWithdrawAmountByCardNumberRow) *record.CardYearAmount
```

##### `ToYearlyWithdrawAmounts`

ToYearlyWithdrawAmounts maps a slice of GetYearlyWithdrawAmountRow database rows to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyWithdrawAmountRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyWithdrawAmounts(cards []*db.GetYearlyWithdrawAmountRow) []*record.CardYearAmount
```

##### `ToYearlyWithdrawAmountsByCardNumber`

ToYearlyWithdrawAmountsByCardNumber maps a slice of GetYearlyWithdrawAmountByCardNumberRow database rows
to a slice of CardYearAmount domain models.

Args:
  - cards: A slice of pointers to GetYearlyWithdrawAmountByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *cardRecordMapper) ToYearlyWithdrawAmountsByCardNumber(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) []*record.CardYearAmount
```

### `merchantDocumentRecordMapper`

merchantDocumentRecordMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.

```go
type merchantDocumentRecordMapper struct {
}
```

#### Methods

##### `ToGetMerchantDocument`

MerchantDocumentRecord maps a MerchantDocument database row to a MerchantDocumentRecord domain model.

Args:
  - merchantdocument: A pointer to a MerchantDocument representing the database row.

Returns:
  - A pointer to a MerchantDocumentRecord containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantDocumentRecordMapper) ToGetMerchantDocument(doc *db.MerchantDocument) *record.MerchantDocumentRecord
```

##### `ToMerchantDocumentActiveRecord`

ToMerchantDocumentActiveRecord maps a GetActiveMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.

Args:
  - doc: A pointer to a GetActiveMerchantDocumentsRow representing the database row.

Returns:
  - A pointer to a MerchantDocumentRecord containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantDocumentRecordMapper) ToMerchantDocumentActiveRecord(doc *db.GetActiveMerchantDocumentsRow) *record.MerchantDocumentRecord
```

##### `ToMerchantDocumentRecord`

ToMerchantDocumentRecord maps a GetMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.

Args:
  - doc: A pointer to a GetMerchantDocumentsRow representing the database row.

Returns:
  - A pointer to a MerchantDocumentRecord containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantDocumentRecordMapper) ToMerchantDocumentRecord(doc *db.GetMerchantDocumentsRow) *record.MerchantDocumentRecord
```

##### `ToMerchantDocumentTrashedRecord`

ToMerchantDocumentTrashedRecord maps a GetTrashedMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.

Args:
  - doc: A pointer to a GetTrashedMerchantDocumentsRow representing the database row.

Returns:
  - A pointer to a MerchantDocumentRecord containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantDocumentRecordMapper) ToMerchantDocumentTrashedRecord(doc *db.GetTrashedMerchantDocumentsRow) *record.MerchantDocumentRecord
```

##### `ToMerchantDocumentsActiveRecord`

ToMerchantDocumentsActiveRecord maps a slice of GetActiveMerchantDocumentsRow database rows to a slice of MerchantDocumentRecord domain models.

Args:
  - docs: A slice of pointers to GetActiveMerchantDocumentsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantDocumentRecordMapper) ToMerchantDocumentsActiveRecord(docs []*db.GetActiveMerchantDocumentsRow) []*record.MerchantDocumentRecord
```

##### `ToMerchantDocumentsRecord`

ToMerchantDocumentsRecord maps a slice of GetMerchantDocumentsRow database rows to a slice of MerchantDocumentRecord domain models.

Args:
  - docs: A slice of pointers to GetMerchantDocumentsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantDocumentRecordMapper) ToMerchantDocumentsRecord(docs []*db.GetMerchantDocumentsRow) []*record.MerchantDocumentRecord
```

##### `ToMerchantDocumentsTrashedRecord`

ToMerchantDocumentsTrashedRecord maps a slice of GetTrashedMerchantDocumentsRow database rows to a slice of
MerchantDocumentRecord domain models.

Args:
  - docs: A slice of pointers to GetTrashedMerchantDocumentsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
    ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantDocumentRecordMapper) ToMerchantDocumentsTrashedRecord(docs []*db.GetTrashedMerchantDocumentsRow) []*record.MerchantDocumentRecord
```

### `merchantRecordMapper`

merchantRecordMapper provides methods to map Merchant database rows to MerchantRecord domain models.

```go
type merchantRecordMapper struct {
}
```

#### Methods

##### `ToMerchantActiveRecord`

ToMerchantActiveRecord maps a GetActiveMerchantsRow to a MerchantRecord domain model.

Args:
  - merchant: A pointer to a GetActiveMerchantsRow representing the database row.

Returns:
  - A pointer to a MerchantRecord containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord
```

##### `ToMerchantGetAllRecord`

ToMerchantGetAllRecord maps a GetMerchantsRow to a MerchantRecord domain model.

Args:
  - merchant: A pointer to a GetMerchantsRow representing the database row.

Returns:
  - A pointer to a MerchantRecord containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantGetAllRecord(merchant *db.GetMerchantsRow) *record.MerchantRecord
```

##### `ToMerchantMonthlyAmount`

ToMerchantMonthlyAmount maps a GetMonthlyAmountMerchantRow to a MerchantMonthlyAmount domain model.

Args:
  - ms: A pointer to a GetMonthlyAmountMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyAmount(ms *db.GetMonthlyAmountMerchantRow) *record.MerchantMonthlyAmount
```

##### `ToMerchantMonthlyAmountByApikey`

ToMerchantMonthlyAmountByApikey maps a GetMonthlyAmountByApikeyRow to a
MerchantMonthlyAmount domain model.

Args:
  - ms: A pointer to a GetMonthlyAmountByApikeyRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyAmount containing the mapped data, including
    Month, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyAmountByApikey(ms *db.GetMonthlyAmountByApikeyRow) *record.MerchantMonthlyAmount
```

##### `ToMerchantMonthlyAmountByMerchant`

ToMerchantMonthlyAmountByMerchant maps a GetMonthlyAmountByMerchantsRow to a
MerchantMonthlyAmount domain model.

Args:
  - ms: A pointer to a GetMonthlyAmountByMerchantsRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyAmountByMerchant(ms *db.GetMonthlyAmountByMerchantsRow) *record.MerchantMonthlyAmount
```

##### `ToMerchantMonthlyAmounts`

ToMerchantMonthlyAmounts maps a slice of GetMonthlyAmountMerchantRow to a slice of MerchantMonthlyAmount domain models.

Args:
  - ms: A slice of pointers to GetMonthlyAmountMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyAmounts(ms []*db.GetMonthlyAmountMerchantRow) []*record.MerchantMonthlyAmount
```

##### `ToMerchantMonthlyAmountsByApikey`

ToMerchantMonthlyAmountsByApikey maps a slice of GetMonthlyAmountByApikeyRow to a slice of
MerchantMonthlyAmount domain models.

Args:
  - ms: A slice of pointers to GetMonthlyAmountByApikeyRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyAmountsByApikey(ms []*db.GetMonthlyAmountByApikeyRow) []*record.MerchantMonthlyAmount
```

##### `ToMerchantMonthlyAmountsByMerchant`

ToMerchantMonthlyAmountsByMerchant maps a slice of GetMonthlyAmountByMerchantsRow
to a slice of MerchantMonthlyAmount domain models.

Args:
  - ms: A slice of pointers to GetMonthlyAmountByMerchantsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including
    Month and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyAmountsByMerchant(ms []*db.GetMonthlyAmountByMerchantsRow) []*record.MerchantMonthlyAmount
```

##### `ToMerchantMonthlyPaymentMethod`

ToMerchantMonthlyPaymentMethod maps a GetMonthlyPaymentMethodsMerchantRow to a
MerchantMonthlyPaymentMethod domain model.

Args:
  - ms: A pointer to a GetMonthlyPaymentMethodsMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
    Month, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethod(ms *db.GetMonthlyPaymentMethodsMerchantRow) *record.MerchantMonthlyPaymentMethod
```

##### `ToMerchantMonthlyPaymentMethodByApikey`

ToMerchantMonthlyPaymentMethodByApikey maps a GetMonthlyPaymentMethodByApikeyRow to a
MerchantMonthlyPaymentMethod domain model.

Args:
  - ms: A pointer to a GetMonthlyPaymentMethodByApikeyRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
    Month, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodByApikey(ms *db.GetMonthlyPaymentMethodByApikeyRow) *record.MerchantMonthlyPaymentMethod
```

##### `ToMerchantMonthlyPaymentMethodByMerchant`

ToMerchantMonthlyPaymentMethodByMerchant maps a GetMonthlyPaymentMethodByMerchantsRow to a
MerchantMonthlyPaymentMethod domain model.

Args:
  - ms: A pointer to a GetMonthlyPaymentMethodByMerchantsRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
    Month, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodByMerchant(ms *db.GetMonthlyPaymentMethodByMerchantsRow) *record.MerchantMonthlyPaymentMethod
```

##### `ToMerchantMonthlyPaymentMethods`

ToMerchantMonthlyPaymentMethods maps a slice of GetMonthlyPaymentMethodsMerchantRow
to a slice of MerchantMonthlyPaymentMethod domain models.

Args:
  - ms: A slice of pointers to GetMonthlyPaymentMethodsMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
    Month, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethods(ms []*db.GetMonthlyPaymentMethodsMerchantRow) []*record.MerchantMonthlyPaymentMethod
```

##### `ToMerchantMonthlyPaymentMethodsByApikey`

ToMerchantMonthlyPaymentMethodsByApikey maps a slice of GetMonthlyPaymentMethodByApikeyRow to a slice
of MerchantMonthlyPaymentMethod domain models.

Args:
  - ms: A slice of pointers to GetMonthlyPaymentMethodByApikeyRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
    Month, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodsByApikey(ms []*db.GetMonthlyPaymentMethodByApikeyRow) []*record.MerchantMonthlyPaymentMethod
```

##### `ToMerchantMonthlyPaymentMethodsByMerchant`

ToMerchantMonthlyPaymentMethodsByMerchant maps a slice of GetMonthlyPaymentMethodByMerchantsRow
to a slice of MerchantMonthlyPaymentMethod domain models.

Args:
  - ms: A slice of pointers to GetMonthlyPaymentMethodByMerchantsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
    Month, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodsByMerchant(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) []*record.MerchantMonthlyPaymentMethod
```

##### `ToMerchantMonthlyTotalAmount`

ToMerchantMonthlyTotalAmount maps a GetMonthlyTotalAmountMerchantRow to a
MerchantMonthlyTotalAmount domain model.

Args:
  - ms: A pointer to a GetMonthlyTotalAmountMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
    Month, Year, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmount(ms *db.GetMonthlyTotalAmountMerchantRow) *record.MerchantMonthlyTotalAmount
```

##### `ToMerchantMonthlyTotalAmountByApikey`

ToMerchantMonthlyTotalAmountByApikey maps a GetMonthlyTotalAmountByApikeyRow to a
MerchantMonthlyTotalAmount domain model.

Args:
  - ms: A pointer to a GetMonthlyTotalAmountByApikeyRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
    Month, Year, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountByApikey(ms *db.GetMonthlyTotalAmountByApikeyRow) *record.MerchantMonthlyTotalAmount
```

##### `ToMerchantMonthlyTotalAmountByMerchant`

ToMerchantMonthlyTotalAmountByMerchant maps a GetMonthlyTotalAmountByMerchantRow to a
MerchantMonthlyTotalAmount domain model.

Args:
  - ms: A pointer to a GetMonthlyTotalAmountByMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
    Month, Year, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountByMerchant(ms *db.GetMonthlyTotalAmountByMerchantRow) *record.MerchantMonthlyTotalAmount
```

##### `ToMerchantMonthlyTotalAmounts`

ToMerchantMonthlyTotalAmounts maps a slice of GetMonthlyTotalAmountMerchantRow to a slice
of MerchantMonthlyTotalAmount domain models.

Args:
  - ms: A slice of pointers to GetMonthlyTotalAmountMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
    Month, Year, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmounts(ms []*db.GetMonthlyTotalAmountMerchantRow) []*record.MerchantMonthlyTotalAmount
```

##### `ToMerchantMonthlyTotalAmountsByApikey`

ToMerchantMonthlyTotalAmountsByApikey maps a slice of GetMonthlyTotalAmountByApikeyRow
to a slice of MerchantMonthlyTotalAmount domain models.

Args:
  - ms: A slice of pointers to GetMonthlyTotalAmountByApikeyRow representing the
    database rows.

Returns:
  - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data,
    including Month, Year, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountsByApikey(ms []*db.GetMonthlyTotalAmountByApikeyRow) []*record.MerchantMonthlyTotalAmount
```

##### `ToMerchantMonthlyTotalAmountsByMerchant`

ToMerchantMonthlyTotalAmountsByMerchant maps a slice of GetMonthlyTotalAmountByMerchantRow to a
slice of MerchantMonthlyTotalAmount domain models.

Args:
  - ms: A slice of pointers to GetMonthlyTotalAmountByMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
    Month, Year, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountsByMerchant(ms []*db.GetMonthlyTotalAmountByMerchantRow) []*record.MerchantMonthlyTotalAmount
```

##### `ToMerchantRecord`

ToMerchantRecord maps a Merchant to a MerchantRecord domain model.

Args:
  - merchant: A pointer to a Merchant representing the database row.

Returns:
  - A pointer to a MerchantRecord containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord
```

##### `ToMerchantTransactionByApikeyRecord`

ToMerchantTransactionByApikeyRecord maps a FindAllTransactionsByApikeyRow to a MerchantTransactionsRecord domain model.

Args:
  - merchant: A pointer to a FindAllTransactionsByApikeyRow representing the database row.

Returns:
  - A pointer to a MerchantTransactionsRecord containing the mapped data, including
    TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
    TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantTransactionByApikeyRecord(merchant *db.FindAllTransactionsByApikeyRow) *record.MerchantTransactionsRecord
```

##### `ToMerchantTransactionByMerchantRecord`

ToMerchantTransactionByMerchantRecord maps a FindAllTransactionsByMerchantRow to a MerchantTransactionsRecord domain model.

Args:
  - merchant: A pointer to a FindAllTransactionsByMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantTransactionsRecord containing the mapped data, including
    TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
    TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantTransactionByMerchantRecord(merchant *db.FindAllTransactionsByMerchantRow) *record.MerchantTransactionsRecord
```

##### `ToMerchantTransactionRecord`

ToMerchantTransactionRecord maps a FindAllTransactionsRow to a MerchantTransactionsRecord domain model.

Args:
  - merchant: A pointer to a FindAllTransactionsRow representing the database row.

Returns:
  - A pointer to a MerchantTransactionsRecord containing the mapped data, including
    TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
    TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantTransactionRecord(merchant *db.FindAllTransactionsRow) *record.MerchantTransactionsRecord
```

##### `ToMerchantTrashedRecord`

ToMerchantTrashedRecord maps a GetTrashedMerchantsRow to a MerchantRecord domain model.

Args:
  - merchant: A pointer to a GetTrashedMerchantsRow representing the database row.

Returns:
  - A pointer to a MerchantRecord containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord
```

##### `ToMerchantYearlyAmount`

ToMerchantYearlyAmount maps a GetYearlyAmountMerchantRow to a MerchantYearlyAmount domain model.

Args:
  - ms: A pointer to a GetYearlyAmountMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyAmount containing the mapped data, including Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyAmount(ms *db.GetYearlyAmountMerchantRow) *record.MerchantYearlyAmount
```

##### `ToMerchantYearlyAmountByApikey`

ToMerchantYearlyAmountByApikey maps a GetYearlyAmountByApikeyRow to a
MerchantYearlyAmount domain model.

Args:
  - ms: A pointer to a GetYearlyAmountByApikeyRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyAmountByApikey(ms *db.GetYearlyAmountByApikeyRow) *record.MerchantYearlyAmount
```

##### `ToMerchantYearlyAmountByMerchant`

ToMerchantYearlyAmountByMerchant maps a GetYearlyAmountByMerchantsRow to a
MerchantYearlyAmount domain model.

Args:
  - ms: A pointer to a GetYearlyAmountByMerchantsRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyAmountByMerchant(ms *db.GetYearlyAmountByMerchantsRow) *record.MerchantYearlyAmount
```

##### `ToMerchantYearlyAmounts`

ToMerchantYearlyAmounts maps a slice of GetYearlyAmountMerchantRow to a slice
of MerchantYearlyAmount domain models.

Args:
  - ms: A slice of pointers to GetYearlyAmountMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyAmounts(ms []*db.GetYearlyAmountMerchantRow) []*record.MerchantYearlyAmount
```

##### `ToMerchantYearlyAmountsByApikey`

ToMerchantYearlyAmountsByApikey maps a slice of GetYearlyAmountByApikeyRow to a slice
of MerchantYearlyAmount domain models.

Args:
  - ms: A slice of pointers to GetYearlyAmountByApikeyRow representing the database rows.

Returns:
  - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyAmountsByApikey(ms []*db.GetYearlyAmountByApikeyRow) []*record.MerchantYearlyAmount
```

##### `ToMerchantYearlyAmountsByMerchant`

ToMerchantYearlyAmountsByMerchant maps a slice of GetYearlyAmountByMerchantsRow
to a slice of MerchantYearlyAmount domain models.

Args:
  - ms: A slice of pointers to GetYearlyAmountByMerchantsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyAmountsByMerchant(ms []*db.GetYearlyAmountByMerchantsRow) []*record.MerchantYearlyAmount
```

##### `ToMerchantYearlyPaymentMethod`

ToMerchantYearlyPaymentMethod maps a GetYearlyPaymentMethodMerchantRow to a
MerchantYearlyPaymentMethod domain model.

Args:
  - ms: A pointer to a GetYearlyPaymentMethodMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
    Year, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethod(ms *db.GetYearlyPaymentMethodMerchantRow) *record.MerchantYearlyPaymentMethod
```

##### `ToMerchantYearlyPaymentMethodByApikey`

ToMerchantYearlyPaymentMethodByApikey maps a GetYearlyPaymentMethodByApikeyRow to a
MerchantYearlyPaymentMethod domain model.

Args:
  - ms: A pointer to a GetYearlyPaymentMethodByApikeyRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
    Year, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodByApikey(ms *db.GetYearlyPaymentMethodByApikeyRow) *record.MerchantYearlyPaymentMethod
```

##### `ToMerchantYearlyPaymentMethodByMerchant`

ToMerchantYearlyPaymentMethodByMerchant maps a GetYearlyPaymentMethodByMerchantsRow to a
MerchantYearlyPaymentMethod domain model.

Args:
  - ms: A pointer to a GetYearlyPaymentMethodByMerchantsRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
    Year, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodByMerchant(ms *db.GetYearlyPaymentMethodByMerchantsRow) *record.MerchantYearlyPaymentMethod
```

##### `ToMerchantYearlyPaymentMethods`

ToMerchantYearlyPaymentMethods maps a slice of GetYearlyPaymentMethodMerchantRow to a slice
of MerchantYearlyPaymentMethod domain models.

Args:
  - ms: A slice of pointers to GetYearlyPaymentMethodMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
    Year, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethods(ms []*db.GetYearlyPaymentMethodMerchantRow) []*record.MerchantYearlyPaymentMethod
```

##### `ToMerchantYearlyPaymentMethodsByApikey`

ToMerchantYearlyPaymentMethodsByApikey maps a slice of GetYearlyPaymentMethodByApikeyRow
to a slice of MerchantYearlyPaymentMethod domain models.

Args:
  - ms: A slice of pointers to GetYearlyPaymentMethodByApikeyRow representing the database rows.

Returns:
  - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
    Year, PaymentMethod, and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodsByApikey(ms []*db.GetYearlyPaymentMethodByApikeyRow) []*record.MerchantYearlyPaymentMethod
```

##### `ToMerchantYearlyPaymentMethodsByMerchant`

```go
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodsByMerchant(ms []*db.GetYearlyPaymentMethodByMerchantsRow) []*record.MerchantYearlyPaymentMethod
```

##### `ToMerchantYearlyTotalAmount`

ToMerchantYearlyTotalAmount maps a GetYearlyTotalAmountMerchantRow to a
MerchantYearlyTotalAmount domain model.

Args:
  - ms: A pointer to a GetYearlyTotalAmountMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmount(ms *db.GetYearlyTotalAmountMerchantRow) *record.MerchantYearlyTotalAmount
```

##### `ToMerchantYearlyTotalAmountByApikey`

ToMerchantYearlyTotalAmountByApikey maps a GetYearlyTotalAmountByApikeyRow to a
MerchantYearlyTotalAmount domain model.

Args:
  - ms: A pointer to a GetYearlyTotalAmountByApikeyRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountByApikey(ms *db.GetYearlyTotalAmountByApikeyRow) *record.MerchantYearlyTotalAmount
```

##### `ToMerchantYearlyTotalAmountByMerchant`

ToMerchantYearlyTotalAmountByMerchant maps a GetYearlyTotalAmountByMerchantRow to a
MerchantYearlyTotalAmount domain model.

Args:
  - ms: A pointer to a GetYearlyTotalAmountByMerchantRow representing the database row.

Returns:
  - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountByMerchant(ms *db.GetYearlyTotalAmountByMerchantRow) *record.MerchantYearlyTotalAmount
```

##### `ToMerchantYearlyTotalAmounts`

ToMerchantYearlyTotalAmounts maps a slice of GetYearlyTotalAmountMerchantRow to a slice
of MerchantYearlyTotalAmount domain models.

Args:
  - ms: A slice of pointers to GetYearlyTotalAmountMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmounts(ms []*db.GetYearlyTotalAmountMerchantRow) []*record.MerchantYearlyTotalAmount
```

##### `ToMerchantYearlyTotalAmountsByApikey`

ToMerchantYearlyTotalAmountsByApikey maps a slice of GetYearlyTotalAmountByApikeyRow
to a slice of MerchantYearlyTotalAmount domain models.

Args:
  - ms: A slice of pointers to GetYearlyTotalAmountByApikeyRow representing the
    database rows.

Returns:
  - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data,
    including Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountsByApikey(ms []*db.GetYearlyTotalAmountByApikeyRow) []*record.MerchantYearlyTotalAmount
```

##### `ToMerchantYearlyTotalAmountsByMerchant`

ToMerchantYearlyTotalAmountsByMerchant maps a slice of GetYearlyTotalAmountByMerchantRow to a
slice of MerchantYearlyTotalAmount domain models.

Args:
  - ms: A slice of pointers to GetYearlyTotalAmountByMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
    Year and TotalAmount.

```go
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountsByMerchant(ms []*db.GetYearlyTotalAmountByMerchantRow) []*record.MerchantYearlyTotalAmount
```

##### `ToMerchantsActiveRecord`

ToMerchantsActiveRecord maps a slice of GetActiveMerchantsRow to a slice of
MerchantRecord domain models.

Args:
  - merchants: A slice of pointers to GetActiveMerchantsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantRecord containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord
```

##### `ToMerchantsGetAllRecord`

ToMerchantsGetAllRecord maps a slice of GetMerchantsRow to a slice of
MerchantRecord domain models.

Args:
  - merchants: A slice of pointers to GetMerchantsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantRecord containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantsGetAllRecord(merchants []*db.GetMerchantsRow) []*record.MerchantRecord
```

##### `ToMerchantsRecord`

ToMerchantsRecord maps a slice of Merchant database rows to a slice of MerchantRecord domain models.

Args:
  - merchants: A slice of Merchant pointers representing the database rows.

Returns:
  - A slice of MerchantRecord pointers containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord
```

##### `ToMerchantsTransactionByApikeyRecord`

ToMerchantsTransactionByApikeyRecord maps a slice of FindAllTransactionsByApikeyRow to a slice of
MerchantTransactionsRecord domain models.

Args:
  - merchants: A slice of pointers to FindAllTransactionsByApikeyRow representing the database rows.

Returns:
  - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
    TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
    TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantsTransactionByApikeyRecord(merchants []*db.FindAllTransactionsByApikeyRow) []*record.MerchantTransactionsRecord
```

##### `ToMerchantsTransactionByMerchantRecord`

ToMerchantsTransactionByMerchantRecord maps a slice of FindAllTransactionsByMerchantRow to a slice
of MerchantTransactionsRecord domain models.

Args:
  - merchants: A slice of pointers to FindAllTransactionsByMerchantRow representing the database rows.

Returns:
  - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
    TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
    TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantsTransactionByMerchantRecord(merchants []*db.FindAllTransactionsByMerchantRow) []*record.MerchantTransactionsRecord
```

##### `ToMerchantsTransactionRecord`

ToMerchantsTransactionRecord maps multiple FindAllTransactionsRow to multiple
MerchantTransactionsRecord domain models.

Args:
  - merchants: A slice of pointers to FindAllTransactionsRow representing the database rows.

Returns:
  - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
    TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
    TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantsTransactionRecord(merchants []*db.FindAllTransactionsRow) []*record.MerchantTransactionsRecord
```

##### `ToMerchantsTrashedRecord`

ToMerchantsTrashedRecord maps a slice of GetTrashedMerchantsRow to a slice of
MerchantRecord domain models.

Args:
  - merchants: A slice of pointers to GetTrashedMerchantsRow representing the trashed database rows.

Returns:
  - A slice of pointers to MerchantRecord containing the mapped data, including
    ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (m *merchantRecordMapper) ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord
```

### `refreshTokenRecordMapper`

refreshTokenRecordMapper provides methods to map RefreshToken database rows to RefreshTokenRecord domain models.

```go
type refreshTokenRecordMapper struct {
}
```

#### Methods

##### `ToRefreshTokenRecord`

ToRefreshTokenRecord maps a RefreshToken database row to a RefreshTokenRecord domain model.
Args:
  - refreshToken: A pointer to a RefreshToken representing the database row.

Returns:
  - A pointer to a RefreshTokenRecord containing the mapped data, including
    ID, UserID, Token, and ExpiredAt, CreatedAt, and UpdatedAt.

```go
func (r *refreshTokenRecordMapper) ToRefreshTokenRecord(refreshToken *db.RefreshToken) *record.RefreshTokenRecord
```

##### `ToRefreshTokensRecord`

ToRefreshTokensRecord maps a slice of RefreshToken database rows to a slice of RefreshTokenRecord
domain models.

Args:
  - refreshTokens: A slice of pointers to RefreshToken representing the database rows.

Returns:
  - A slice of pointers to RefreshTokenRecord containing the mapped data, including
    ID, UserID, Token, and ExpiredAt, CreatedAt, and UpdatedAt.

```go
func (r *refreshTokenRecordMapper) ToRefreshTokensRecord(refreshTokens []*db.RefreshToken) []*record.RefreshTokenRecord
```

### `resetTokenRecordMapper`

resetTokenRecordMapper provides methods to map ResetToken database rows to ResetTokenRecord domain models.

```go
type resetTokenRecordMapper struct {
}
```

#### Methods

##### `ToResetTokenRecord`

ToResetTokenRecord maps a ResetToken database row to a ResetTokenRecord domain model.

Args:
  - resetToken: A pointer to a ResetToken representing the database row.

Returns:
  - A pointer to a ResetTokenRecord containing the mapped data, including ID, UserID, Token, and ExpiredAt.

```go
func (r *resetTokenRecordMapper) ToResetTokenRecord(resetToken *db.ResetToken) *record.ResetTokenRecord
```

### `roleRecordMapper`

roleRecordMapper provides methods to map Role database rows to RoleRecord domain models.

```go
type roleRecordMapper struct {
}
```

#### Methods

##### `ToRoleRecord`

ToRoleRecord maps a Role database row to a RoleRecord domain model.

Args:
  - role: A pointer to a Role representing the database row.

Returns:
  - A pointer to a RoleRecord containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRoleRecord(role *db.Role) *record.RoleRecord
```

##### `ToRoleRecordActive`

ToRoleRecordActive maps a GetActiveRolesRow to a RoleRecord domain model.

Args:
  - role: A pointer to a GetActiveRolesRow representing the database row.

Returns:
  - A pointer to a RoleRecord containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRoleRecordActive(role *db.GetActiveRolesRow) *record.RoleRecord
```

##### `ToRoleRecordAll`

ToRoleRecordAll maps a GetRolesRow to a RoleRecord domain model.

Args:
  - role: A pointer to a GetRolesRow representing the database row.

Returns:
  - A pointer to a RoleRecord containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRoleRecordAll(role *db.GetRolesRow) *record.RoleRecord
```

##### `ToRoleRecordTrashed`

ToRoleRecordTrashed maps a GetTrashedRolesRow to a RoleRecord domain model.

Args:
  - role: A pointer to a GetTrashedRolesRow representing the database row.

Returns:
  - A pointer to a RoleRecord containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRoleRecordTrashed(role *db.GetTrashedRolesRow) *record.RoleRecord
```

##### `ToRolesRecord`

ToRolesRecord maps a slice of Role database rows to a slice of RoleRecord domain models.

Args:
  - roles: A slice of pointers to Role structs representing the database rows.

Returns:
  - A slice of pointers to RoleRecord structs containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRolesRecord(roles []*db.Role) []*record.RoleRecord
```

##### `ToRolesRecordActive`

ToRolesRecordActive maps a slice of GetActiveRolesRow to a slice of RoleRecord domain models.

Args:
  - roles: A slice of pointers to GetActiveRolesRow structs representing the database rows.

Returns:
  - A slice of pointers to RoleRecord structs containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRolesRecordActive(roles []*db.GetActiveRolesRow) []*record.RoleRecord
```

##### `ToRolesRecordAll`

ToRolesRecordAll maps a slice of GetRolesRow to a slice of RoleRecord domain models.

Args:
  - roles: A slice of pointers to GetRolesRow structs representing the database rows.

Returns:
  - A slice of pointers to RoleRecord structs containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRolesRecordAll(roles []*db.GetRolesRow) []*record.RoleRecord
```

##### `ToRolesRecordTrashed`

ToRolesRecordTrashed maps a slice of GetTrashedRolesRow to a slice of RoleRecord domain models.

Args:
  - roles: A slice of pointers to GetTrashedRolesRow structs representing the database rows.

Returns:
  - A slice of pointers to RoleRecord structs containing the mapped data, including
    ID, Name, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *roleRecordMapper) ToRolesRecordTrashed(roles []*db.GetTrashedRolesRow) []*record.RoleRecord
```

### `saldoRecordMapper`

saldoRecordMapper provides methods to map Saldo database rows to SaldoRecord domain models.

```go
type saldoRecordMapper struct {
}
```

#### Methods

##### `ToSaldoMonthBalance`

ToSaldoMonthBalance maps a MonthlySaldoBalancesRow database row to a SaldoMonthSaldoBalance domain model.

Args:
  - ss: A pointer to a MonthlySaldoBalancesRow representing the database row.

Returns:
  - A pointer to a SaldoMonthSaldoBalance containing the mapped data, including Month, and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoMonthBalance(ss *db.GetMonthlySaldoBalancesRow) *record.SaldoMonthSaldoBalance
```

##### `ToSaldoMonthBalances`

ToSaldoMonthBalances maps a slice of MonthlySaldoBalancesRow database rows to a slice of SaldoMonthSaldoBalance domain models.

Args:
  - ss: A slice of pointers to MonthlySaldoBalancesRow representing the database rows.

Returns:
  - A slice of pointers to SaldoMonthSaldoBalance containing the mapped data, including Month, and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoMonthBalances(ss []*db.GetMonthlySaldoBalancesRow) []*record.SaldoMonthSaldoBalance
```

##### `ToSaldoMonthTotalBalance`

ToSaldoMonthTotalBalance maps a GetMonthlyTotalSaldoBalanceRow database row to a SaldoMonthTotalBalance domain model.

Args:
  - ss: A pointer to a GetMonthlyTotalSaldoBalanceRow representing the database row.

Returns:
  - A pointer to a SaldoMonthTotalBalance containing the mapped data, including Month, Year, and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoMonthTotalBalance(ss *db.GetMonthlyTotalSaldoBalanceRow) *record.SaldoMonthTotalBalance
```

##### `ToSaldoMonthTotalBalances`

ToSaldoMonthTotalBalances maps a slice of GetMonthlyTotalSaldoBalanceRow database rows to a slice of SaldoMonthTotalBalance domain models.

Args:
  - ss: A slice of pointers to GetMonthlyTotalSaldoBalanceRow representing the database rows.

Returns:
  - A slice of pointers to SaldoMonthTotalBalance containing the mapped data, including Month, Year, and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoMonthTotalBalances(ss []*db.GetMonthlyTotalSaldoBalanceRow) []*record.SaldoMonthTotalBalance
```

##### `ToSaldoRecord`

ToSaldoRecord maps a Saldo database row to a SaldoRecord domain model.

Args:
  - saldo: A pointer to a Saldo representing the database row.

Returns:
  - A pointer to a SaldoRecord containing the mapped data, including
    ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord
```

##### `ToSaldoRecordActive`

ToSaldoRecordActive maps a GetActiveSaldosRow database row to a SaldoRecord domain model.

Args:
  - saldo: A pointer to a GetActiveSaldosRow representing the database row.

Returns:
  - A pointer to a SaldoRecord containing the mapped data, including
    ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldoRecordActive(saldo *db.GetActiveSaldosRow) *record.SaldoRecord
```

##### `ToSaldoRecordAll`

ToSaldoRecordAll maps a GetSaldosRow database row to a SaldoRecord domain model.

Args:
  - saldo: A pointer to a GetSaldosRow representing the database row.

Returns:
  - A pointer to a SaldoRecord containing the mapped data, including
    ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldoRecordAll(saldo *db.GetSaldosRow) *record.SaldoRecord
```

##### `ToSaldoRecordTrashed`

ToSaldoRecordTrashed maps a GetTrashedSaldosRow database row to a SaldoRecord domain model.

Args:
  - saldo: A pointer to a GetTrashedSaldosRow representing the trashed saldo database row.

Returns:
  - A pointer to a SaldoRecord containing the mapped data, including ID, CardNumber,
    TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldoRecordTrashed(saldo *db.GetTrashedSaldosRow) *record.SaldoRecord
```

##### `ToSaldoYearSaldoBalance`

ToSaldoYearSaldoBalance maps a GetYearlySaldoBalancesRow database row to a SaldoYearSaldoBalance domain model.

Args:
  - ss: A pointer to a GetYearlySaldoBalancesRow representing the database row.

Returns:
  - A pointer to a SaldoYearSaldoBalance containing the mapped data, including Year and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoYearSaldoBalance(ss *db.GetYearlySaldoBalancesRow) *record.SaldoYearSaldoBalance
```

##### `ToSaldoYearSaldoBalances`

ToSaldoYearSaldoBalances maps a slice of GetYearlySaldoBalancesRow database rows to a slice of SaldoYearSaldoBalance domain models.

Args:
  - ss: A slice of pointers to GetYearlySaldoBalancesRow representing the database rows.

Returns:
  - A slice of pointers to SaldoYearSaldoBalance containing the mapped data, including Year and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoYearSaldoBalances(ss []*db.GetYearlySaldoBalancesRow) []*record.SaldoYearSaldoBalance
```

##### `ToSaldoYearTotalBalance`

ToSaldoYearTotalBalance maps a GetYearlyTotalSaldoBalancesRow database row to a SaldoYearTotalBalance domain model.

Args:
  - ss: A pointer to a GetYearlyTotalSaldoBalancesRow representing the database row.

Returns:
  - A pointer to a SaldoYearTotalBalance containing the mapped data, including Year, and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoYearTotalBalance(ss *db.GetYearlyTotalSaldoBalancesRow) *record.SaldoYearTotalBalance
```

##### `ToSaldoYearTotalBalances`

ToSaldoYearTotalBalances maps a slice of GetYearlyTotalSaldoBalancesRow database rows to a slice of SaldoYearTotalBalance domain models.

Args:
  - ss: A slice of pointers to GetYearlyTotalSaldoBalancesRow representing the database rows.

Returns:
  - A slice of pointers to SaldoYearTotalBalance containing the mapped data, including Year, and TotalBalance.

```go
func (s *saldoRecordMapper) ToSaldoYearTotalBalances(ss []*db.GetYearlyTotalSaldoBalancesRow) []*record.SaldoYearTotalBalance
```

##### `ToSaldosRecord`

ToSaldosRecord maps a slice of Saldo database rows to a slice of SaldoRecord domain models.

Args:
  - saldos: A slice of pointers to Saldo representing the database rows.

Returns:
  - A slice of pointers to SaldoRecord containing the mapped data, including
    ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldosRecord(saldos []*db.Saldo) []*record.SaldoRecord
```

##### `ToSaldosRecordActive`

ToSaldosRecordActive maps a slice of GetActiveSaldosRow database rows to a slice of SaldoRecord domain models.

Args:
  - saldos: A slice of pointers to GetActiveSaldosRow representing the database rows.

Returns:
  - A slice of pointers to SaldoRecord containing the mapped data, including
    ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldosRecordActive(saldos []*db.GetActiveSaldosRow) []*record.SaldoRecord
```

##### `ToSaldosRecordAll`

ToSaldosRecordAll maps a slice of GetSaldosRow database rows to a slice of SaldoRecord domain models.

Args:
  - saldos: A slice of pointers to GetSaldosRow representing the database rows.

Returns:
  - A slice of pointers to SaldoRecord containing the mapped data, including
    ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldosRecordAll(saldos []*db.GetSaldosRow) []*record.SaldoRecord
```

##### `ToSaldosRecordTrashed`

ToSaldosRecordTrashed maps a slice of GetTrashedSaldosRow database rows to a slice of SaldoRecord domain models.

Args:
  - saldos: A slice of pointers to GetTrashedSaldosRow representing the trashed saldo database rows.

Returns:
  - A slice of pointers to SaldoRecord containing the mapped data, including ID, CardNumber,
    TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *saldoRecordMapper) ToSaldosRecordTrashed(saldos []*db.GetTrashedSaldosRow) []*record.SaldoRecord
```

### `topupRecordMapper`

```go
type topupRecordMapper struct {
}
```

#### Methods

##### `ToTopupByCardNumberRecord`

ToTopupByCardNumberRecord converts a db.GetTopupsByCardNumberRow to a record.TopupRecord.
It takes a pointer to a db.GetTopupsByCardNumberRow as a parameter and returns a pointer to a record.TopupRecord.
The function maps the fields of the db.GetTopupsByCardNumberRow to the corresponding fields of the record.TopupRecord.
If the DeletedAt field of the db.GetTopupsByCardNumberRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupByCardNumberRecord(topup *db.GetTopupsByCardNumberRow) *record.TopupRecord
```

##### `ToTopupByCardNumberRecords`

ToTopupByCardNumberRecords converts a slice of db.GetTopupsByCardNumberRow to a slice of record.TopupRecord.
It takes a slice of pointers to db.GetTopupsByCardNumberRow as a parameter and returns a slice of pointers to record.TopupRecord.
The function iterates over the provided slice of db.GetTopupsByCardNumberRow, converting each element
using the ToTopupByCardNumberRecord method and appending the result to a new slice.
The function returns a slice of pointers to record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupByCardNumberRecords(topups []*db.GetTopupsByCardNumberRow) []*record.TopupRecord
```

##### `ToTopupMonthlyAmount`

ToTopupMonthlyAmount maps a GetMonthlyTopupAmountsRow database row to a TopupMonthAmount domain model.
It takes a pointer to a GetMonthlyTopupAmountsRow as a parameter and returns a pointer to a TopupMonthAmount.
The function maps the fields of GetMonthlyTopupAmountsRow to the corresponding fields of TopupMonthAmount.

```go
func (t *topupRecordMapper) ToTopupMonthlyAmount(s *db.GetMonthlyTopupAmountsRow) *record.TopupMonthAmount
```

##### `ToTopupMonthlyAmountByCardNumber`

ToTopupMonthlyAmountByCardNumber maps a GetMonthlyTopupAmountsByCardNumberRow to a TopupMonthAmount domain model.

Args:
  - s: A pointer to a GetMonthlyTopupAmountsByCardNumberRow representing the database row.

Returns:
  - A pointer to a TopupMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (t *topupRecordMapper) ToTopupMonthlyAmountByCardNumber(s *db.GetMonthlyTopupAmountsByCardNumberRow) *record.TopupMonthAmount
```

##### `ToTopupMonthlyAmounts`

ToTopupMonthlyAmounts maps a slice of GetMonthlyTopupAmountsRow database rows to a slice of TopupMonthAmount domain models.
It iterates over the provided slice, converting each element using the ToTopupMonthlyAmount method,
and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.

```go
func (t *topupRecordMapper) ToTopupMonthlyAmounts(s []*db.GetMonthlyTopupAmountsRow) []*record.TopupMonthAmount
```

##### `ToTopupMonthlyAmountsByCardNumber`

ToTopupMonthlyAmountsByCardNumber maps a slice of GetMonthlyTopupAmountsByCardNumberRow database rows to a slice of TopupMonthAmount domain models.
It iterates over the provided slice, converting each element using the ToTopupMonthlyAmountByCardNumber method,
and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.

```go
func (t *topupRecordMapper) ToTopupMonthlyAmountsByCardNumber(s []*db.GetMonthlyTopupAmountsByCardNumberRow) []*record.TopupMonthAmount
```

##### `ToTopupMonthlyMethod`

ToTopupMonthlyMethod maps a GetMonthlyTopupMethodsRow database row to a TopupMonthMethod domain model.
It takes a pointer to a GetMonthlyTopupMethodsRow as a parameter and returns a pointer to a TopupMonthMethod.
The function maps the fields of GetMonthlyTopupMethodsRow to the corresponding fields of TopupMonthMethod.

```go
func (t *topupRecordMapper) ToTopupMonthlyMethod(s *db.GetMonthlyTopupMethodsRow) *record.TopupMonthMethod
```

##### `ToTopupMonthlyMethodByCardNumber`

ToTopupMonthlyMethodByCardNumber maps a GetMonthlyTopupMethodsByCardNumberRow database row
to a TopupMonthMethod domain model. It takes a pointer to a
GetMonthlyTopupMethodsByCardNumberRow as a parameter and returns a pointer to a
TopupMonthMethod. The function maps the fields of GetMonthlyTopupMethodsByCardNumberRow
to the corresponding fields of TopupMonthMethod.

```go
func (t *topupRecordMapper) ToTopupMonthlyMethodByCardNumber(s *db.GetMonthlyTopupMethodsByCardNumberRow) *record.TopupMonthMethod
```

##### `ToTopupMonthlyMethods`

ToTopupMonthlyMethods maps a slice of GetMonthlyTopupMethodsRow database rows to a slice of TopupMonthMethod domain models.
It iterates over the provided slice, converting each element using the ToTopupMonthlyMethod method,
and appends the result to a new slice. The function returns a slice of pointers to TopupMonthMethod.

```go
func (t *topupRecordMapper) ToTopupMonthlyMethods(s []*db.GetMonthlyTopupMethodsRow) []*record.TopupMonthMethod
```

##### `ToTopupMonthlyMethodsByCardNumber`

ToTopupMonthlyMethodsByCardNumber maps a slice of GetMonthlyTopupMethodsByCardNumberRow database rows
to a slice of TopupMonthMethod domain models. It iterates over the provided slice, converting
each element using the ToTopupMonthlyMethodByCardNumber method, and appends the result to a new
slice. The function returns a slice of pointers to TopupMonthMethod.

```go
func (t *topupRecordMapper) ToTopupMonthlyMethodsByCardNumber(s []*db.GetMonthlyTopupMethodsByCardNumberRow) []*record.TopupMonthMethod
```

##### `ToTopupRecord`

ToTopupRecord converts a db.Topup to a record.TopupRecord.
It takes a pointer to a db.Topup as a parameter and returns a pointer to a record.TopupRecord.
The function maps the fields of the db.Topup to the corresponding fields of the record.TopupRecord.
If the DeletedAt field of the db.Topup is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecord(topup *db.Topup) *record.TopupRecord
```

##### `ToTopupRecordActive`

ToTopupRecordActive converts a db.GetActiveTopupsRow to a record.TopupRecord.
It takes a pointer to a db.GetActiveTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
The function maps the fields of the db.GetActiveTopupsRow to the corresponding fields of the record.TopupRecord.
If the DeletedAt field of the db.GetActiveTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecordActive(topup *db.GetActiveTopupsRow) *record.TopupRecord
```

##### `ToTopupRecordAll`

ToTopupRecordAll converts a db.GetTopupsRow to a record.TopupRecord.
It takes a pointer to a db.GetTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
The function maps the fields of the db.GetTopupsRow to the corresponding fields of the record.TopupRecord.
If the DeletedAt field of the db.GetTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecordAll(topup *db.GetTopupsRow) *record.TopupRecord
```

##### `ToTopupRecordMonthStatusFailed`

ToTopupRecordMonthStatusFailed maps a GetMonthTopupStatusFailedRow database row
to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
GetMonthTopupStatusFailedRow as a parameter and returns a pointer to a
TopupRecordMonthStatusFailed. The function maps the fields of the
GetMonthTopupStatusFailedRow to the corresponding fields of the
TopupRecordMonthStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordMonthStatusFailed(s *db.GetMonthTopupStatusFailedRow) *record.TopupRecordMonthStatusFailed
```

##### `ToTopupRecordMonthStatusFailedByCardNumber`

ToTopupRecordMonthStatusFailedByCardNumber maps a GetMonthTopupStatusFailedCardNumberRow database row
to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
GetMonthTopupStatusFailedCardNumberRow as a parameter and returns a pointer to a
TopupRecordMonthStatusFailed. The function maps the fields of the
GetMonthTopupStatusFailedCardNumberRow to the corresponding fields of the
TopupRecordMonthStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordMonthStatusFailedByCardNumber(s *db.GetMonthTopupStatusFailedCardNumberRow) *record.TopupRecordMonthStatusFailed
```

##### `ToTopupRecordMonthStatusSuccess`

ToTopupRecordMonthStatusSuccess maps a GetMonthTopupStatusSuccessRow database row to a TopupRecordMonthStatusSuccess domain model.
It takes a pointer to a GetMonthTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
The function maps the fields of the GetMonthTopupStatusSuccessRow to the corresponding fields of the TopupRecordMonthStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordMonthStatusSuccess(s *db.GetMonthTopupStatusSuccessRow) *record.TopupRecordMonthStatusSuccess
```

##### `ToTopupRecordMonthStatusSuccessByCardNumber`

ToTopupRecordMonthStatusSuccessByCardNumber maps a GetMonthTopupStatusSuccessCardNumberRow database row to a TopupRecordMonthStatusSuccess domain model.
It takes a pointer to a GetMonthTopupStatusSuccessCardNumberRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
The function maps the fields of the GetMonthTopupStatusSuccessCardNumberRow to the corresponding fields of the TopupRecordMonthStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordMonthStatusSuccessByCardNumber(s *db.GetMonthTopupStatusSuccessCardNumberRow) *record.TopupRecordMonthStatusSuccess
```

##### `ToTopupRecordTrashed`

ToTopupRecordTrashed converts a db.GetTrashedTopupsRow to a record.TopupRecord.
It takes a pointer to a db.GetTrashedTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
The function maps the fields of the db.GetTrashedTopupsRow to the corresponding fields of the record.TopupRecord.
If the DeletedAt field of the db.GetTrashedTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecordTrashed(topup *db.GetTrashedTopupsRow) *record.TopupRecord
```

##### `ToTopupRecordYearStatusFailed`

ToTopupRecordYearStatusFailed maps a GetYearlyTopupStatusFailedRow database row to a TopupRecordYearStatusFailed domain model.
It takes a pointer to a GetYearlyTopupStatusFailedRow as a parameter and returns a pointer to a TopupRecordYearStatusFailed.
The function maps the fields of the GetYearlyTopupStatusFailedRow to the corresponding fields of the TopupRecordYearStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordYearStatusFailed(s *db.GetYearlyTopupStatusFailedRow) *record.TopupRecordYearStatusFailed
```

##### `ToTopupRecordYearStatusFailedByCardNumber`

ToTopupRecordYearStatusFailedByCardNumber maps a GetYearlyTopupStatusFailedCardNumberRow database row
to a TopupRecordYearStatusFailed domain model. It takes a pointer to a
GetYearlyTopupStatusFailedCardNumberRow as a parameter and returns a pointer
to a TopupRecordYearStatusFailed. The function maps the fields of
GetYearlyTopupStatusFailedCardNumberRow to the corresponding fields of
TopupRecordYearStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordYearStatusFailedByCardNumber(s *db.GetYearlyTopupStatusFailedCardNumberRow) *record.TopupRecordYearStatusFailed
```

##### `ToTopupRecordYearStatusSuccess`

ToTopupRecordYearStatusSuccess maps a GetYearlyTopupStatusSuccessRow database row to a TopupRecordYearStatusSuccess domain model.
It takes a pointer to a GetYearlyTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordYearStatusSuccess.
The function maps the fields of the GetYearlyTopupStatusSuccessRow to the corresponding fields of the TopupRecordYearStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordYearStatusSuccess(s *db.GetYearlyTopupStatusSuccessRow) *record.TopupRecordYearStatusSuccess
```

##### `ToTopupRecordYearStatusSuccessByCardNumber`

ToTopupRecordYearStatusSuccessByCardNumber maps a GetYearlyTopupStatusSuccessCardNumberRow database row
to a TopupRecordYearStatusSuccess domain model. It takes a pointer to a
GetYearlyTopupStatusSuccessCardNumberRow as a parameter and returns a pointer
to a TopupRecordYearStatusSuccess. The function maps the fields of
GetYearlyTopupStatusSuccessCardNumberRow to the corresponding fields of
TopupRecordYearStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordYearStatusSuccessByCardNumber(s *db.GetYearlyTopupStatusSuccessCardNumberRow) *record.TopupRecordYearStatusSuccess
```

##### `ToTopupRecords`

ToTopupRecords converts a slice of db.Topup to a slice of record.TopupRecord.
It iterates over the provided slice of db.Topup, converting each element
using the ToTopupRecord method and appending the result to a new slice.
The function returns a slice of pointers to record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecords(topups []*db.Topup) []*record.TopupRecord
```

##### `ToTopupRecordsActive`

ToTopupRecordsActive maps a slice of GetActiveTopupsRow database rows to a slice of TopupRecord domain models.
It iterates over the provided slice of GetActiveTopupsRow, converting each element
using the ToTopupRecordActive method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecordsActive(topups []*db.GetActiveTopupsRow) []*record.TopupRecord
```

##### `ToTopupRecordsAll`

ToTopupRecordsAll converts a slice of db.GetTopupsRow to a slice of record.TopupRecord.
It iterates over the provided slice of db.GetTopupsRow, converting each element
using the ToTopupRecordAll method and appending the result to a new slice.
The function returns a slice of pointers to record.TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecordsAll(topups []*db.GetTopupsRow) []*record.TopupRecord
```

##### `ToTopupRecordsMonthStatusFailed`

ToTopupRecordsMonthStatusFailed maps a slice of GetMonthTopupStatusFailedRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
It iterates over the provided slice of GetMonthTopupStatusFailedRow, converting each element
using the ToTopupRecordMonthStatusFailed method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecordMonthStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordsMonthStatusFailed(topups []*db.GetMonthTopupStatusFailedRow) []*record.TopupRecordMonthStatusFailed
```

##### `ToTopupRecordsMonthStatusFailedByCardNumber`

ToTopupRecordsMonthStatusFailedByCardNumber maps a slice of GetMonthTopupStatusFailedCardNumberRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
It iterates over the provided slice of GetMonthTopupStatusFailedCardNumberRow, converting each element
using the ToTopupRecordMonthStatusFailedByCardNumber method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecordMonthStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordsMonthStatusFailedByCardNumber(topups []*db.GetMonthTopupStatusFailedCardNumberRow) []*record.TopupRecordMonthStatusFailed
```

##### `ToTopupRecordsMonthStatusSuccess`

ToTopupRecordsMonthStatusSuccess maps a slice of GetMonthTopupStatusSuccessRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
It iterates over the provided slice of GetMonthTopupStatusSuccessRow, converting each element
using the ToTopupRecordMonthStatusSuccess method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecordMonthStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordsMonthStatusSuccess(topups []*db.GetMonthTopupStatusSuccessRow) []*record.TopupRecordMonthStatusSuccess
```

##### `ToTopupRecordsMonthStatusSuccessByCardNumber`

ToTopupRecordsMonthStatusSuccessByCardNumber maps a slice of GetMonthTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
It iterates over the provided slice of GetMonthTopupStatusSuccessCardNumberRow, converting each element
using the ToTopupRecordMonthStatusSuccessByCardNumber method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecordMonthStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordsMonthStatusSuccessByCardNumber(topups []*db.GetMonthTopupStatusSuccessCardNumberRow) []*record.TopupRecordMonthStatusSuccess
```

##### `ToTopupRecordsTrashed`

ToTopupRecordsTrashed maps a slice of GetTrashedTopupsRow database rows to a slice of TopupRecord domain models.
It iterates over the provided slice of GetTrashedTopupsRow, converting each element
using the ToTopupRecordTrashed method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecord.

```go
func (t *topupRecordMapper) ToTopupRecordsTrashed(topups []*db.GetTrashedTopupsRow) []*record.TopupRecord
```

##### `ToTopupRecordsYearStatusFailed`

ToTopupRecordsYearStatusFailed maps a slice of GetYearlyTopupStatusFailedRow database rows to a slice of TopupRecordYearStatusFailed domain models.
It iterates over the provided slice of GetYearlyTopupStatusFailedRow, converting each element
using the ToTopupRecordYearStatusFailed method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecordYearStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordsYearStatusFailed(topups []*db.GetYearlyTopupStatusFailedRow) []*record.TopupRecordYearStatusFailed
```

##### `ToTopupRecordsYearStatusFailedByCardNumber`

ToTopupRecordsYearStatusFailedByCardNumber maps a slice of GetYearlyTopupStatusFailedCardNumberRow database rows
to a slice of TopupRecordYearStatusFailed domain models. It iterates over the provided slice, converting each element
using the ToTopupRecordYearStatusFailedByCardNumber method, and appends the result to a new slice.
The function returns a slice of pointers to TopupRecordYearStatusFailed.

```go
func (t *topupRecordMapper) ToTopupRecordsYearStatusFailedByCardNumber(topups []*db.GetYearlyTopupStatusFailedCardNumberRow) []*record.TopupRecordYearStatusFailed
```

##### `ToTopupRecordsYearStatusSuccess`

ToTopupRecordsYearStatusSuccess maps a slice of GetYearlyTopupStatusSuccessRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
It iterates over the provided slice of GetYearlyTopupStatusSuccessRow, converting each element
using the ToTopupRecordYearStatusSuccess method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecordYearStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordsYearStatusSuccess(topups []*db.GetYearlyTopupStatusSuccessRow) []*record.TopupRecordYearStatusSuccess
```

##### `ToTopupRecordsYearStatusSuccessByCardNumber`

ToTopupRecordsYearStatusSuccessByCardNumber maps a slice of GetYearlyTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
It iterates over the provided slice of GetYearlyTopupStatusSuccessCardNumberRow, converting each element
using the ToTopupRecordYearStatusSuccessByCardNumber method and appending the result to a new slice.
The function returns a slice of pointers to TopupRecordYearStatusSuccess.

```go
func (t *topupRecordMapper) ToTopupRecordsYearStatusSuccessByCardNumber(topups []*db.GetYearlyTopupStatusSuccessCardNumberRow) []*record.TopupRecordYearStatusSuccess
```

##### `ToTopupYearlyAmount`

ToTopupYearlyAmount maps a GetYearlyTopupAmountsRow database row to a TopupYearlyAmount domain model.
It takes a pointer to a GetYearlyTopupAmountsRow as a parameter and returns a pointer to a TopupYearlyAmount.
The function maps the fields of GetYearlyTopupAmountsRow to the corresponding fields of TopupYearlyAmount.

```go
func (t *topupRecordMapper) ToTopupYearlyAmount(s *db.GetYearlyTopupAmountsRow) *record.TopupYearlyAmount
```

##### `ToTopupYearlyAmountByCardNumber`

ToTopupYearlyAmountByCardNumber maps a GetYearlyTopupAmountsByCardNumberRow to a TopupYearlyAmount domain model.

Args:
  - s: A pointer to a GetYearlyTopupAmountsByCardNumberRow representing the database row.

Returns:
  - A pointer to a TopupYearlyAmount containing the mapped data, including Year and TotalAmount.

```go
func (t *topupRecordMapper) ToTopupYearlyAmountByCardNumber(s *db.GetYearlyTopupAmountsByCardNumberRow) *record.TopupYearlyAmount
```

##### `ToTopupYearlyAmounts`

ToTopupYearlyAmounts maps a slice of GetYearlyTopupAmountsRow database rows to a slice of TopupYearlyAmount domain models.
It iterates over the provided slice, converting each element using the ToTopupYearlyAmount method,
and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyAmount.

```go
func (t *topupRecordMapper) ToTopupYearlyAmounts(s []*db.GetYearlyTopupAmountsRow) []*record.TopupYearlyAmount
```

##### `ToTopupYearlyAmountsByCardNumber`

ToTopupYearlyAmountsByCardNumber maps a slice of GetYearlyTopupAmountsByCardNumberRow to a
slice of TopupYearlyAmount domain models.

Args:
  - s: A slice of pointers to GetYearlyTopupAmountsByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TopupYearlyAmount containing the mapped data, including Year and TotalAmount.

```go
func (t *topupRecordMapper) ToTopupYearlyAmountsByCardNumber(s []*db.GetYearlyTopupAmountsByCardNumberRow) []*record.TopupYearlyAmount
```

##### `ToTopupYearlyMethod`

ToTopupYearlyMethod maps a GetYearlyTopupMethodsRow database row to a TopupYearlyMethod domain model.
It takes a pointer to a GetYearlyTopupMethodsRow as a parameter and returns a pointer to a TopupYearlyMethod.
The function maps the fields of GetYearlyTopupMethodsRow to the corresponding fields of TopupYearlyMethod.

```go
func (t *topupRecordMapper) ToTopupYearlyMethod(s *db.GetYearlyTopupMethodsRow) *record.TopupYearlyMethod
```

##### `ToTopupYearlyMethodByCardNumber`

ToTopupYearlyMethodByCardNumber maps a GetYearlyTopupMethodsByCardNumberRow database row to a TopupYearlyMethod domain model.
It takes a pointer to a GetYearlyTopupMethodsByCardNumberRow as a parameter and returns a pointer to a TopupYearlyMethod.
The function maps the fields of GetYearlyTopupMethodsByCardNumberRow to the corresponding fields of TopupYearlyMethod.

```go
func (t *topupRecordMapper) ToTopupYearlyMethodByCardNumber(s *db.GetYearlyTopupMethodsByCardNumberRow) *record.TopupYearlyMethod
```

##### `ToTopupYearlyMethods`

ToTopupYearlyMethods maps a slice of GetYearlyTopupMethodsRow database rows to a slice of TopupYearlyMethod domain models.
It iterates over the provided slice, converting each element using the ToTopupYearlyMethod method,
and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyMethod.

```go
func (t *topupRecordMapper) ToTopupYearlyMethods(s []*db.GetYearlyTopupMethodsRow) []*record.TopupYearlyMethod
```

##### `ToTopupYearlyMethodsByCardNumber`

ToTopupYearlyMethodsByCardNumber maps a slice of GetYearlyTopupMethodsByCardNumberRow database rows
to a slice of TopupYearlyMethod domain models. It iterates over the provided slice, converting
each element using the ToTopupYearlyMethodByCardNumber method, and appends the result to a new
slice. The function returns a slice of pointers to TopupYearlyMethod.

```go
func (t *topupRecordMapper) ToTopupYearlyMethodsByCardNumber(s []*db.GetYearlyTopupMethodsByCardNumberRow) []*record.TopupYearlyMethod
```

### `transactionRecordMapper`

transactionRecordMapper is a struct that provides methods to map
database rows to transaction domain model records.

```go
type transactionRecordMapper struct {
}
```

#### Methods

##### `ToTransactionByCardNumberRecord`

ToTransactionByCardNumberRecord maps a GetTransactionsByCardNumberRow database row to a TransactionRecord domain model.

Args:
  - transaction: A pointer to a GetTransactionsByCardNumberRow representing the database row.

Returns:
  - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionByCardNumberRecord(transaction *db.GetTransactionsByCardNumberRow) *record.TransactionRecord
```

##### `ToTransactionMonthlyAmount`

ToTransactionMonthlyAmount maps a GetMonthlyAmountsRow database row to a TransactionMonthAmount
domain model.

Args:
  - ss: A pointer to GetMonthlyAmountsRow representing the database row.

Returns:
  - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyAmount(ss *db.GetMonthlyAmountsRow) *record.TransactionMonthAmount
```

##### `ToTransactionMonthlyAmountByCardNumber`

ToTransactionMonthlyAmountByCardNumber maps a GetMonthlyAmountsByCardNumberRow database row
to a TransactionMonthAmount domain model.

Args:
  - ss: A pointer to GetMonthlyAmountsByCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyAmountByCardNumber(ss *db.GetMonthlyAmountsByCardNumberRow) *record.TransactionMonthAmount
```

##### `ToTransactionMonthlyAmounts`

ToTransactionMonthlyAmounts maps a slice of GetMonthlyAmountsRow database rows
to a slice of TransactionMonthAmount domain models.

Args:
  - ss: A slice of pointers to GetMonthlyAmountsRow representing the database rows.

Returns:
  - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyAmounts(ss []*db.GetMonthlyAmountsRow) []*record.TransactionMonthAmount
```

##### `ToTransactionMonthlyAmountsByCardNumber`

ToTransactionMonthlyAmountsByCardNumber maps a slice of GetMonthlyAmountsByCardNumberRow database rows
to a slice of TransactionMonthAmount domain models.

Args:
  - ss: A slice of pointers to GetMonthlyAmountsByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyAmountsByCardNumber(ss []*db.GetMonthlyAmountsByCardNumberRow) []*record.TransactionMonthAmount
```

##### `ToTransactionMonthlyMethod`

ToTransactionMonthlyMethod maps a GetMonthlyPaymentMethodsRow database row to a TransactionMonthMethod
domain model.

Args:
  - s: A pointer to GetMonthlyPaymentMethodsRow representing the database row.

Returns:
  - A pointer to TransactionMonthMethod containing the mapped data, including Month,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyMethod(ss *db.GetMonthlyPaymentMethodsRow) *record.TransactionMonthMethod
```

##### `ToTransactionMonthlyMethodByCardNumber`

ToTransactionMonthlyMethodByCardNumber maps a GetMonthlyPaymentMethodsByCardNumberRow database row to a TransactionMonthMethod
domain model.

Args:
  - ss: A pointer to GetMonthlyPaymentMethodsByCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionMonthMethod containing the mapped data, including Month,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyMethodByCardNumber(ss *db.GetMonthlyPaymentMethodsByCardNumberRow) *record.TransactionMonthMethod
```

##### `ToTransactionMonthlyMethods`

ToTransactionMonthlyMethods maps a slice of GetMonthlyPaymentMethodsRow database rows
to a slice of TransactionMonthMethod domain models.

Args:
  - ss: A slice of pointers to GetMonthlyPaymentMethodsRow representing the database rows.

Returns:
  - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyMethods(ss []*db.GetMonthlyPaymentMethodsRow) []*record.TransactionMonthMethod
```

##### `ToTransactionMonthlyMethodsByCardNumber`

ToTransactionMonthlyMethodsByCardNumber maps a slice of GetMonthlyPaymentMethodsByCardNumberRow database rows
to a slice of TransactionMonthMethod domain models.

Args:
  - ss: A slice of pointers to GetMonthlyPaymentMethodsByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionMonthlyMethodsByCardNumber(ss []*db.GetMonthlyPaymentMethodsByCardNumberRow) []*record.TransactionMonthMethod
```

##### `ToTransactionRecord`

ToTransactionRecord maps a Transaction database row to a TransactionRecord domain model.

Args:
  - transaction: A pointer to a Transaction representing the database row.

Returns:
  - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord
```

##### `ToTransactionRecordActive`

ToTransactionRecordActive maps a GetActiveTransactionsRow database row to a TransactionRecord domain model.
It is intended for use with database rows that contain active transaction records.
It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionRecordActive(transaction *db.GetActiveTransactionsRow) *record.TransactionRecord
```

##### `ToTransactionRecordAll`

ToTransactionRecordAll maps a GetTransactionsRow database row to a TransactionRecord domain model.

Args:
  - transaction: A pointer to a GetTransactionsRow representing the database row.

Returns:
  - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord
```

##### `ToTransactionRecordMonthStatusFailed`

ToTransactionRecordMonthStatusFailed maps a GetMonthTransactionStatusFailedRow database row
to a TransactionRecordMonthStatusFailed domain model.

Args:
  - s: A pointer to GetMonthTransactionStatusFailedRow representing the database row.

Returns:
  - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusFailed(s *db.GetMonthTransactionStatusFailedRow) *record.TransactionRecordMonthStatusFailed
```

##### `ToTransactionRecordMonthStatusFailedCardNumber`

ToTransactionRecordMonthStatusFailedCardNumber maps a GetMonthTransactionStatusFailedCardNumberRow
database row to a TransactionRecordMonthStatusFailed domain model.

Args:
  - s: A pointer to GetMonthTransactionStatusFailedCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusFailedCardNumber(s *db.GetMonthTransactionStatusFailedCardNumberRow) *record.TransactionRecordMonthStatusFailed
```

##### `ToTransactionRecordMonthStatusSuccess`

ToTransactionRecordMonthStatusSuccess maps a single GetMonthTransactionStatusSuccessRow database row
to a TransactionRecordMonthStatusSuccess domain model.

Args:
  - s: A pointer to GetMonthTransactionStatusSuccessRow representing the database row.

Returns:
  - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusSuccess(s *db.GetMonthTransactionStatusSuccessRow) *record.TransactionRecordMonthStatusSuccess
```

##### `ToTransactionRecordMonthStatusSuccessCardNumber`

ToTransactionRecordMonthStatusSuccessCardNumber maps a GetMonthTransactionStatusSuccessCardNumberRow database row
to a TransactionRecordMonthStatusSuccess domain model.

Args:
  - s: A pointer to GetMonthTransactionStatusSuccessCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransactionStatusSuccessCardNumberRow) *record.TransactionRecordMonthStatusSuccess
```

##### `ToTransactionRecordTrashed`

ToTransactionRecordTrashed maps a GetTrashedTransactionsRow database row to a TransactionRecord domain model.
It is intended for use with database rows that contain trashed transaction records.
It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionRecordTrashed(transaction *db.GetTrashedTransactionsRow) *record.TransactionRecord
```

##### `ToTransactionRecordYearStatusFailed`

ToTransactionRecordYearStatusFailed maps a GetYearlyTransactionStatusFailedRow database row
to a TransactionRecordYearStatusFailed domain model.

Args:
  - s: A pointer to GetYearlyTransactionStatusFailedRow representing the database row.

Returns:
  - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordYearStatusFailed(s *db.GetYearlyTransactionStatusFailedRow) *record.TransactionRecordYearStatusFailed
```

##### `ToTransactionRecordYearStatusFailedCardNumber`

ToTransactionRecordYearStatusFailedCardNumber maps a GetYearlyTransactionStatusFailedCardNumberRow
database row to a TransactionRecordYearStatusFailed domain model.

Args:
  - s: A pointer to GetYearlyTransactionStatusFailedCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordYearStatusFailedCardNumber(s *db.GetYearlyTransactionStatusFailedCardNumberRow) *record.TransactionRecordYearStatusFailed
```

##### `ToTransactionRecordYearStatusSuccess`

ToTransactionRecordYearStatusSuccess maps a GetYearlyTransactionStatusSuccessRow database row
to a TransactionRecordYearStatusSuccess domain model.

Args:
  - s: A pointer to GetYearlyTransactionStatusSuccessRow representing the database row.

Returns:
  - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordYearStatusSuccess(s *db.GetYearlyTransactionStatusSuccessRow) *record.TransactionRecordYearStatusSuccess
```

##### `ToTransactionRecordYearStatusSuccessCardNumber`

ToTransactionRecordYearStatusSuccessCardNumber maps a GetYearlyTransactionStatusSuccessCardNumberRow database row
to a TransactionRecordYearStatusSuccess domain model.

Args:
  - s: A pointer to GetYearlyTransactionStatusSuccessCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransactionStatusSuccessCardNumberRow) *record.TransactionRecordYearStatusSuccess
```

##### `ToTransactionRecordsMonthStatusFailed`

ToTransactionRecordsMonthStatusFailed maps a slice of GetMonthTransactionStatusFailedRow database rows
to a slice of TransactionRecordMonthStatusFailed domain models.

Args:
  - Transactions: A slice of pointers to GetMonthTransactionStatusFailedRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusFailed(Transactions []*db.GetMonthTransactionStatusFailedRow) []*record.TransactionRecordMonthStatusFailed
```

##### `ToTransactionRecordsMonthStatusFailedCardNumber`

ToTransactionRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransactionStatusFailedCardNumberRow
database rows to a slice of TransactionRecordMonthStatusFailed domain models.

Args:
  - Transactions: A slice of pointers to GetMonthTransactionStatusFailedCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusFailedCardNumber(Transactions []*db.GetMonthTransactionStatusFailedCardNumberRow) []*record.TransactionRecordMonthStatusFailed
```

##### `ToTransactionRecordsMonthStatusSuccess`

ToTransactionRecordsMonthStatusSuccess maps a slice of GetMonthTransactionStatusSuccessRow database rows to a slice of
TransactionRecordMonthStatusSuccess domain models.

Args:
  - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusSuccess(Transactions []*db.GetMonthTransactionStatusSuccessRow) []*record.TransactionRecordMonthStatusSuccess
```

##### `ToTransactionRecordsMonthStatusSuccessCardNumber`

ToTransactionRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransactionStatusSuccessCardNumberRow database rows
to a slice of TransactionRecordMonthStatusSuccess domain models.

Args:
  - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusSuccessCardNumber(Transactions []*db.GetMonthTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordMonthStatusSuccess
```

##### `ToTransactionRecordsYearStatusFailed`

ToTransactionRecordsYearStatusFailed maps a slice of GetYearlyTransactionStatusFailedRow database rows
to a slice of TransactionRecordYearStatusFailed domain models.

Args:
  - Transactions: A slice of pointers to GetYearlyTransactionStatusFailedRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusFailed(Transactions []*db.GetYearlyTransactionStatusFailedRow) []*record.TransactionRecordYearStatusFailed
```

##### `ToTransactionRecordsYearStatusFailedCardNumber`

ToTransactionRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransactionStatusFailedCardNumberRow
database rows to a slice of TransactionRecordYearStatusFailed domain models.

Args:
  - Transactions: A slice of pointers to GetYearlyTransactionStatusFailedCardNumberRow representing the
    database rows.

Returns:
  - A slice of pointers to TransactionRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusFailedCardNumber(Transactions []*db.GetYearlyTransactionStatusFailedCardNumberRow) []*record.TransactionRecordYearStatusFailed
```

##### `ToTransactionRecordsYearStatusSuccess`

ToTransactionRecordsYearStatusSuccess maps a slice of GetYearlyTransactionStatusSuccessRow database rows
to a slice of TransactionRecordYearStatusSuccess domain models.

Args:
  - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusSuccess(Transactions []*db.GetYearlyTransactionStatusSuccessRow) []*record.TransactionRecordYearStatusSuccess
```

##### `ToTransactionRecordsYearStatusSuccessCardNumber`

ToTransactionRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransactionStatusSuccessCardNumberRow database rows
to a slice of TransactionRecordYearStatusSuccess domain models.

Args:
  - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusSuccessCardNumber(Transactions []*db.GetYearlyTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordYearStatusSuccess
```

##### `ToTransactionYearlyAmount`

ToTransactionYearlyAmount maps a GetYearlyAmountsRow database row to a TransactionYearlyAmount
domain model.

Args:
  - ss: A pointer to GetYearlyAmountsRow representing the database row.

Returns:
  - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyAmount(ss *db.GetYearlyAmountsRow) *record.TransactionYearlyAmount
```

##### `ToTransactionYearlyAmountByCardNumber`

ToTransactionYearlyAmountByCardNumber maps a GetYearlyAmountsByCardNumberRow database row
to a TransactionYearlyAmount domain model.

Args:
  - ss: A pointer to GetYearlyAmountsByCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyAmountByCardNumber(ss *db.GetYearlyAmountsByCardNumberRow) *record.TransactionYearlyAmount
```

##### `ToTransactionYearlyAmounts`

ToTransactionYearlyAmounts maps a slice of GetYearlyAmountsRow database rows
to a slice of TransactionYearlyAmount domain models.

Args:
  - ss: A slice of pointers to GetYearlyAmountsRow representing the database rows.

Returns:
  - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyAmounts(ss []*db.GetYearlyAmountsRow) []*record.TransactionYearlyAmount
```

##### `ToTransactionYearlyAmountsByCardNumber`

ToTransactionYearlyAmountsByCardNumber maps a slice of GetYearlyAmountsByCardNumberRow database rows
to a slice of TransactionYearlyAmount domain models.

Args:
  - ss: A slice of pointers to GetYearlyAmountsByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyAmountsByCardNumber(ss []*db.GetYearlyAmountsByCardNumberRow) []*record.TransactionYearlyAmount
```

##### `ToTransactionYearlyMethod`

ToTransactionYearlyMethod maps a GetYearlyPaymentMethodsRow database row to a TransactionYearMethod
domain model.

Args:
  - s: A pointer to GetYearlyPaymentMethodsRow representing the database row.

Returns:
  - A pointer to TransactionYearMethod containing the mapped data, including Year,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyMethod(ss *db.GetYearlyPaymentMethodsRow) *record.TransactionYearMethod
```

##### `ToTransactionYearlyMethodByCardNumber`

ToTransactionYearlyMethodByCardNumber maps a GetYearlyPaymentMethodsByCardNumberRow database row to a TransactionYearMethod
domain model.

Args:
  - ss: A pointer to GetYearlyPaymentMethodsByCardNumberRow representing the database row.

Returns:
  - A pointer to TransactionYearMethod containing the mapped data, including Year,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyMethodByCardNumber(ss *db.GetYearlyPaymentMethodsByCardNumberRow) *record.TransactionYearMethod
```

##### `ToTransactionYearlyMethods`

ToTransactionYearlyMethods maps a slice of GetYearlyPaymentMethodsRow database rows
to a slice of TransactionYearMethod domain models.

Args:
  - ss: A slice of pointers to GetYearlyPaymentMethodsRow representing the database rows.

Returns:
  - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyMethods(ss []*db.GetYearlyPaymentMethodsRow) []*record.TransactionYearMethod
```

##### `ToTransactionYearlyMethodsByCardNumber`

ToTransactionYearlyMethodsByCardNumber maps a slice of GetYearlyPaymentMethodsByCardNumberRow database rows
to a slice of TransactionYearMethod domain models.

Args:
  - ss: A slice of pointers to GetYearlyPaymentMethodsByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
    PaymentMethod, TotalTransactions, and TotalAmount.

```go
func (s *transactionRecordMapper) ToTransactionYearlyMethodsByCardNumber(ss []*db.GetYearlyPaymentMethodsByCardNumberRow) []*record.TransactionYearMethod
```

##### `ToTransactionsByCardNumberRecord`

ToTransactionsByCardNumberRecord maps a slice of GetTransactionsByCardNumberRow database rows
to a slice of TransactionRecord domain models.

Args:
  - transactions: A slice of pointers to GetTransactionsByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionsByCardNumberRecord(transactions []*db.GetTransactionsByCardNumberRow) []*record.TransactionRecord
```

##### `ToTransactionsRecord`

ToTransactionsRecord maps a slice of Transaction database rows to a slice of TransactionRecord domain models.

Args:
  - transactions: A slice of pointers to Transaction representing the database rows.

Returns:
  - A slice of TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord
```

##### `ToTransactionsRecordActive`

ToTransactionsRecordActive maps a slice of GetActiveTransactionsRow database rows
to a slice of TransactionRecord domain models.

Args:
  - transactions: A slice of pointers to GetActiveTransactionsRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionsRecordActive(transactions []*db.GetActiveTransactionsRow) []*record.TransactionRecord
```

##### `ToTransactionsRecordAll`

ToTransactionsRecordAll maps a slice of GetTransactionsRow database rows to a slice of TransactionRecord domain models.

Args:
  - transactions: A slice of pointers to GetTransactionsRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionsRecordAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord
```

##### `ToTransactionsRecordTrashed`

ToTransactionsRecordTrashed maps a slice of GetTrashedTransactionsRow database rows to a slice of
TransactionRecord domain models.

Args:
  - transactions: A slice of pointers to GetTrashedTransactionsRow representing the database rows.

Returns:
  - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
    CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *transactionRecordMapper) ToTransactionsRecordTrashed(transactions []*db.GetTrashedTransactionsRow) []*record.TransactionRecord
```

### `transferRecordMapper`

transferRecordMapper maps a Transfer entity to a TransferRecord struct.

```go
type transferRecordMapper struct {
}
```

#### Methods

##### `ToTransferMonthAmount`

ToTransferMonthAmount maps a GetMonthlyTransferAmountsRow database row
to a TransferMonthAmount domain model.

Args:
  - ss: A pointer to a GetMonthlyTransferAmountsRow representing the database row.

Returns:
  - A pointer to a TransferMonthAmount containing the mapped data, including Month
    and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferMonthAmount(ss *db.GetMonthlyTransferAmountsRow) *record.TransferMonthAmount
```

##### `ToTransferMonthAmountReceiver`

ToTransferMonthAmountReceiver maps a GetMonthlyTransferAmountsByReceiverCardNumberRow database row
to a TransferMonthAmount domain model.

Args:
  - ss: A pointer to a GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database row.

Returns:
  - A pointer to a TransferMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferMonthAmountReceiver(ss *db.GetMonthlyTransferAmountsByReceiverCardNumberRow) *record.TransferMonthAmount
```

##### `ToTransferMonthAmountSender`

ToTransferMonthAmountSender maps a GetMonthlyTransferAmountsBySenderCardNumberRow database row
to a TransferMonthAmount domain model.

Args:
  - ss: A pointer to a GetMonthlyTransferAmountsBySenderCardNumberRow representing the database row.

Returns:
  - A pointer to a TransferMonthAmount containing the mapped data, including Month
    and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferMonthAmountSender(ss *db.GetMonthlyTransferAmountsBySenderCardNumberRow) *record.TransferMonthAmount
```

##### `ToTransferMonthAmounts`

ToTransferMonthAmounts maps a slice of GetMonthlyTransferAmountsRow database rows
to a slice of TransferMonthAmount domain models.

Args:
  - ss: A slice of pointers to GetMonthlyTransferAmountsRow representing the database rows.

Returns:
  - A slice of pointers to TransferMonthAmount containing the mapped data, including Month
    and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferMonthAmounts(ss []*db.GetMonthlyTransferAmountsRow) []*record.TransferMonthAmount
```

##### `ToTransferMonthAmountsReceiver`

ToTransferMonthAmountsReceiver maps a slice of GetMonthlyTransferAmountsByReceiverCardNumberRow database rows
to a slice of TransferMonthAmount domain models.

Args:
  - ss: A slice of pointers to GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferMonthAmountsReceiver(ss []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) []*record.TransferMonthAmount
```

##### `ToTransferMonthAmountsSender`

ToTransferMonthAmountsSender maps a slice of GetMonthlyTransferAmountsBySenderCardNumberRow
database rows to a slice of TransferMonthAmount domain models.

Args:
  - ss: A slice of pointers to GetMonthlyTransferAmountsBySenderCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferMonthAmountsSender(ss []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) []*record.TransferMonthAmount
```

##### `ToTransferRecord`

ToTransferRecord maps a Transfer database row to a TransferRecord domain model.

Args:
  - transfer: A pointer to a Transfer representing the database row.

Returns:
  - A pointer to a TransferRecord containing the mapped data, including
    ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
    CreatedAt, UpdatedAt, and DeletedAt.

```go
func (t *transferRecordMapper) ToTransferRecord(transfer *db.Transfer) *record.TransferRecord
```

##### `ToTransferRecordActive`

ToTransferRecordActive maps a GetActiveTransfersRow database row to a TransferRecord domain model.

Args:
  - transfer: A pointer to a GetActiveTransfersRow representing the active transfer database row.

Returns:
  - A pointer to a TransferRecord containing the mapped data, including
    ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
    CreatedAt, UpdatedAt, and DeletedAt.

```go
func (t *transferRecordMapper) ToTransferRecordActive(transfer *db.GetActiveTransfersRow) *record.TransferRecord
```

##### `ToTransferRecordAll`

ToTransferRecordAll maps a GetTransfersRow database row to a TransferRecord domain model.

Args:
  - transfer: A pointer to a GetTransfersRow representing the database row.

Returns:
  - A pointer to a TransferRecord containing the mapped data, including
    ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
    CreatedAt, UpdatedAt, and DeletedAt.

```go
func (t *transferRecordMapper) ToTransferRecordAll(transfer *db.GetTransfersRow) *record.TransferRecord
```

##### `ToTransferRecordMonthStatusFailed`

ToTransferRecordMonthStatusFailed maps a GetMonthTransferStatusFailedRow database row
to a TransferRecordMonthStatusFailed domain model.

Args:
  - s: A pointer to a GetMonthTransferStatusFailedRow representing the database row.

Returns:
  - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordMonthStatusFailed(s *db.GetMonthTransferStatusFailedRow) *record.TransferRecordMonthStatusFailed
```

##### `ToTransferRecordMonthStatusFailedCardNumber`

ToTransferRecordMonthStatusFailedCardNumber maps a GetMonthTransferStatusFailedCardNumberRow database row
to a TransferRecordMonthStatusFailed domain model.

Args:
  - s: A pointer to a GetMonthTransferStatusFailedCardNumberRow representing the database row.

Returns:
  - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordMonthStatusFailedCardNumber(s *db.GetMonthTransferStatusFailedCardNumberRow) *record.TransferRecordMonthStatusFailed
```

##### `ToTransferRecordMonthStatusSuccess`

ToTransferRecordMonthStatusSuccess maps a GetMonthTransferStatusSuccessRow database row
to a TransferRecordMonthStatusSuccess domain model.

Args:
  - s: A pointer to a GetMonthTransferStatusSuccessRow representing the database row.

Returns:
  - A pointer to a TransferRecordMonthStatusSuccess containing the mapped data, including
    Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordMonthStatusSuccess(s *db.GetMonthTransferStatusSuccessRow) *record.TransferRecordMonthStatusSuccess
```

##### `ToTransferRecordMonthStatusSuccessCardNumber`

ToTransferRecordMonthStatusSuccessCardNumber maps a GetMonthTransferStatusSuccessCardNumberRow database row
to a TransferRecordMonthStatusSuccess domain model.

Args:
  - s: A pointer to GetMonthTransferStatusSuccessCardNumberRow representing the database row.

Returns:
  - A pointer to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransferStatusSuccessCardNumberRow) *record.TransferRecordMonthStatusSuccess
```

##### `ToTransferRecordTrashed`

ToTransferRecordTrashed maps a GetTrashedTransfersRow database row to a TransferRecord domain model.

Args:
  - transfer: A pointer to a GetTrashedTransfersRow representing the active transfer database row.

Returns:
  - A pointer to a TransferRecord containing the mapped data, including
    ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
    CreatedAt, UpdatedAt, and DeletedAt.

```go
func (t *transferRecordMapper) ToTransferRecordTrashed(transfer *db.GetTrashedTransfersRow) *record.TransferRecord
```

##### `ToTransferRecordYearStatusFailed`

ToTransferRecordYearStatusFailed maps a GetYearlyTransferStatusFailedRow database row
to a TransferRecordYearStatusFailed domain model.

Args:
  - s: A pointer to a GetYearlyTransferStatusFailedRow representing the database row.

Returns:
  - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordYearStatusFailed(s *db.GetYearlyTransferStatusFailedRow) *record.TransferRecordYearStatusFailed
```

##### `ToTransferRecordYearStatusFailedCardNumber`

ToTransferRecordYearStatusFailedCardNumber maps a GetYearlyTransferStatusFailedCardNumberRow database row
to a TransferRecordYearStatusFailed domain model.

Args:
  - s: A pointer to a GetYearlyTransferStatusFailedCardNumberRow representing the database row.

Returns:
  - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordYearStatusFailedCardNumber(s *db.GetYearlyTransferStatusFailedCardNumberRow) *record.TransferRecordYearStatusFailed
```

##### `ToTransferRecordYearStatusSuccess`

ToTransferRecordYearStatusSuccess maps a GetYearlyTransferStatusSuccessRow database row
to a TransferRecordYearStatusSuccess domain model.

Args:
  - s: A pointer to a GetYearlyTransferStatusSuccessRow representing the database row.

Returns:
  - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordYearStatusSuccess(s *db.GetYearlyTransferStatusSuccessRow) *record.TransferRecordYearStatusSuccess
```

##### `ToTransferRecordYearStatusSuccessCardNumber`

ToTransferRecordYearStatusSuccessCardNumber maps a GetYearlyTransferStatusSuccessCardNumberRow database row
to a TransferRecordYearStatusSuccess domain model.

Args:
  - s: A pointer to a GetYearlyTransferStatusSuccessCardNumberRow representing the database row.

Returns:
  - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransferStatusSuccessCardNumberRow) *record.TransferRecordYearStatusSuccess
```

##### `ToTransferRecordsMonthStatusFailed`

ToTransferRecordsMonthStatusFailed maps a slice of GetMonthTransferStatusFailedRow database rows
to a slice of TransferRecordMonthStatusFailed domain models.

Args:
  - Transfers: A slice of pointers to GetMonthTransferStatusFailedRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsMonthStatusFailed(Transfers []*db.GetMonthTransferStatusFailedRow) []*record.TransferRecordMonthStatusFailed
```

##### `ToTransferRecordsMonthStatusFailedCardNumber`

ToTransferRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransferStatusFailedCardNumberRow database rows
to a slice of TransferRecordMonthStatusFailed domain models.

Args:
  - Transfers: A slice of pointers to GetMonthTransferStatusFailedCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
    Month, TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsMonthStatusFailedCardNumber(Transfers []*db.GetMonthTransferStatusFailedCardNumberRow) []*record.TransferRecordMonthStatusFailed
```

##### `ToTransferRecordsMonthStatusSuccess`

ToTransferRecordsMonthStatusSuccess maps a slice of GetMonthTransferStatusSuccessRow database rows
to a slice of TransferRecordMonthStatusSuccess domain models.

Args:
  - Transfers: A slice of pointers to GetMonthTransferStatusSuccessRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsMonthStatusSuccess(Transfers []*db.GetMonthTransferStatusSuccessRow) []*record.TransferRecordMonthStatusSuccess
```

##### `ToTransferRecordsMonthStatusSuccessCardNumber`

ToTransferRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransferStatusSuccessCardNumberRow database rows
to a slice of TransferRecordMonthStatusSuccess domain models.

Args:
  - Transfers: A slice of pointers to GetMonthTransferStatusSuccessCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
    Month, TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsMonthStatusSuccessCardNumber(Transfers []*db.GetMonthTransferStatusSuccessCardNumberRow) []*record.TransferRecordMonthStatusSuccess
```

##### `ToTransferRecordsYearStatusFailed`

ToTransferRecordsYearStatusFailed maps a slice of GetYearlyTransferStatusFailedRow database rows
to a slice of TransferRecordYearStatusFailed domain models.

Args:
  - Transfers: A slice of pointers to GetYearlyTransferStatusFailedRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsYearStatusFailed(Transfers []*db.GetYearlyTransferStatusFailedRow) []*record.TransferRecordYearStatusFailed
```

##### `ToTransferRecordsYearStatusFailedCardNumber`

ToTransferRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransferStatusFailedCardNumberRow database rows
to a slice of TransferRecordYearStatusFailed domain models.

Args:
  - Transfers: A slice of pointers to GetYearlyTransferStatusFailedCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
    TotalFailed, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsYearStatusFailedCardNumber(Transfers []*db.GetYearlyTransferStatusFailedCardNumberRow) []*record.TransferRecordYearStatusFailed
```

##### `ToTransferRecordsYearStatusSuccess`

ToTransferRecordsYearStatusSuccess maps a slice of GetYearlyTransferStatusSuccessRow database rows
to a slice of TransferRecordYearStatusSuccess domain models.

Args:
  - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsYearStatusSuccess(Transfers []*db.GetYearlyTransferStatusSuccessRow) []*record.TransferRecordYearStatusSuccess
```

##### `ToTransferRecordsYearStatusSuccessCardNumber`

ToTransferRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransferStatusSuccessCardNumberRow database rows
to a slice of TransferRecordYearStatusSuccess domain models.

Args:
  - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
    TotalSuccess, and TotalAmount.

```go
func (t *transferRecordMapper) ToTransferRecordsYearStatusSuccessCardNumber(Transfers []*db.GetYearlyTransferStatusSuccessCardNumberRow) []*record.TransferRecordYearStatusSuccess
```

##### `ToTransferYearAmount`

ToTransferYearAmount maps a GetYearlyTransferAmountsRow database row to a TransferYearAmount domain model.

Args:
  - ss: A pointer to a GetYearlyTransferAmountsRow representing the database row.

Returns:
  - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferYearAmount(ss *db.GetYearlyTransferAmountsRow) *record.TransferYearAmount
```

##### `ToTransferYearAmountReceiver`

ToTransferYearAmountReceiver maps a GetYearlyTransferAmountsByReceiverCardNumberRow database row
to a TransferYearAmount domain model.

Args:
  - ss: A pointer to a GetYearlyTransferAmountsByReceiverCardNumberRow representing the database row.

Returns:
  - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferYearAmountReceiver(ss *db.GetYearlyTransferAmountsByReceiverCardNumberRow) *record.TransferYearAmount
```

##### `ToTransferYearAmountSender`

ToTransferYearAmountSender maps a GetYearlyTransferAmountsBySenderCardNumberRow database row
to a TransferYearAmount domain model.

Args:
  - ss: A pointer to a GetYearlyTransferAmountsBySenderCardNumberRow representing the database row.

Returns:
  - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferYearAmountSender(ss *db.GetYearlyTransferAmountsBySenderCardNumberRow) *record.TransferYearAmount
```

##### `ToTransferYearAmounts`

ToTransferYearAmounts maps a slice of GetYearlyTransferAmountsRow database rows
to a slice of TransferYearAmount domain models.

Args:
  - ss: A slice of pointers to GetYearlyTransferAmountsRow representing the database rows.

Returns:
  - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferYearAmounts(ss []*db.GetYearlyTransferAmountsRow) []*record.TransferYearAmount
```

##### `ToTransferYearAmountsReceiver`

ToTransferYearAmountsReceiver maps a slice of GetYearlyTransferAmountsByReceiverCardNumberRow database rows
to a slice of TransferYearAmount domain models.

Args:
  - ss: A slice of pointers to GetYearlyTransferAmountsByReceiverCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferYearAmountsReceiver(ss []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) []*record.TransferYearAmount
```

##### `ToTransferYearAmountsSender`

ToTransferYearAmountsSender maps a slice of GetYearlyTransferAmountsBySenderCardNumberRow database rows
to a slice of TransferYearAmount domain models.

Args:
  - ss: A slice of pointers to GetYearlyTransferAmountsBySenderCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.

```go
func (s *transferRecordMapper) ToTransferYearAmountsSender(ss []*db.GetYearlyTransferAmountsBySenderCardNumberRow) []*record.TransferYearAmount
```

##### `ToTransfersRecord`

ToTransfersRecord maps a slice of Transfer database rows to a slice of TransferRecord domain models.

Args:
  - transfers: A slice of pointers to Transfer structs representing the database rows.

Returns:
  - A slice of pointers to TransferRecord structs containing the mapped data.

```go
func (s *transferRecordMapper) ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord
```

##### `ToTransfersRecordActive`

ToTransfersRecordActive maps a slice of GetActiveTransfersRow database rows to a slice of TransferRecord domain models.

Args:
  - transfers: A slice of pointers to GetActiveTransfersRow structs representing the database rows.

Returns:
  - A slice of pointers to TransferRecord structs containing the mapped data.

```go
func (s *transferRecordMapper) ToTransfersRecordActive(transfers []*db.GetActiveTransfersRow) []*record.TransferRecord
```

##### `ToTransfersRecordAll`

```go
func (s *transferRecordMapper) ToTransfersRecordAll(transfers []*db.GetTransfersRow) []*record.TransferRecord
```

##### `ToTransfersRecordTrashed`

ToTransfersRecordTrashed maps a slice of GetTrashedTransfersRow database rows to a slice of TransferRecord domain models.

Args:
  - transfers: A slice of pointers to GetTrashedTransfersRow structs representing the database rows.

Returns:
  - A slice of pointers to TransferRecord structs containing the mapped data.

```go
func (s *transferRecordMapper) ToTransfersRecordTrashed(transfers []*db.GetTrashedTransfersRow) []*record.TransferRecord
```

### `userRecordMapper`

userRecordMapper provides methods to map User database rows to UserRecord domain models.

```go
type userRecordMapper struct {
}
```

#### Methods

##### `ToUserRecord`

ToUserRecord maps a User database row to a UserRecord domain model.

Args:
  - user: A pointer to a User representing the database row.

Returns:
  - A pointer to a UserRecord containing the mapped data, including
    ID, FirstName, LastName, VerifiedCode, IsVerified, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userRecordMapper) ToUserRecord(user *db.User) *record.UserRecord
```

##### `ToUserRecordActivePagination`

ToUserRecordActivePagination maps a GetActiveUsersWithPaginationRow database row to a UserRecord domain model.

Args:
  - user: A pointer to a GetActiveUsersWithPaginationRow representing the database row.

Returns:
  - A pointer to a UserRecord containing the mapped data, including
    ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userRecordMapper) ToUserRecordActivePagination(user *db.GetActiveUsersWithPaginationRow) *record.UserRecord
```

##### `ToUserRecordPagination`

ToUserRecordPagination maps a GetUsersWithPaginationRow database row to a UserRecord domain model.

Args:
  - user: A pointer to a GetUsersWithPaginationRow representing the database row.

Returns:
  - A pointer to a UserRecord containing the mapped data, including
    ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userRecordMapper) ToUserRecordPagination(user *db.GetUsersWithPaginationRow) *record.UserRecord
```

##### `ToUserRecordTrashedPagination`

ToUserRecordTrashedPagination maps a GetTrashedUsersWithPaginationRow database row to a UserRecord domain model.

Args:
  - user: A pointer to a GetTrashedUsersWithPaginationRow representing the database row.

Returns:
  - A pointer to a UserRecord containing the mapped data, including
    ID, FirstName, LastName, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userRecordMapper) ToUserRecordTrashedPagination(user *db.GetTrashedUsersWithPaginationRow) *record.UserRecord
```

##### `ToUsersRecordActivePagination`

ToUsersRecordActivePagination maps a slice of GetActiveUsersWithPaginationRow database rows to a slice of UserRecord domain models.

Args:
  - users: A slice of pointers to GetActiveUsersWithPaginationRow representing the database rows.

Returns:
  - A slice of pointers to UserRecord containing the mapped data for each user, including
    ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userRecordMapper) ToUsersRecordActivePagination(users []*db.GetActiveUsersWithPaginationRow) []*record.UserRecord
```

##### `ToUsersRecordPagination`

ToUsersRecordPagination maps a slice of GetUsersWithPaginationRow database rows to a slice of UserRecord domain models.

Args:
  - users: A slice of pointers to GetUsersWithPaginationRow representing the database rows.

Returns:
  - A slice of pointers to UserRecord containing the mapped data for each user, including
    ID, FirstName, LastName, Email, Password, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userRecordMapper) ToUsersRecordPagination(users []*db.GetUsersWithPaginationRow) []*record.UserRecord
```

##### `ToUsersRecordTrashedPagination`

ToUsersRecordTrashedPagination maps a slice of GetTrashedUsersWithPaginationRow database rows
to a slice of UserRecord domain models.

Args:
  - users: A slice of pointers to GetTrashedUsersWithPaginationRow representing the database rows.

Returns:
  - A slice of pointers to UserRecord containing the mapped data for each user, including
    ID, FirstName, LastName, Email, Password, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *userRecordMapper) ToUsersRecordTrashedPagination(users []*db.GetTrashedUsersWithPaginationRow) []*record.UserRecord
```

### `userRoleRecordMapper`

userRoleRecordMapper provides methods to map UserRole database rows to UserRoleRecord domain models.

```go
type userRoleRecordMapper struct {
}
```

#### Methods

##### `ToUserRoleRecord`

ToUserRoleRecord maps a UserRole database row to a UserRoleRecord domain model.

Args:
  - user: A pointer to a UserRole representing the database row.

Returns:
  - A pointer to a UserRoleRecord containing the mapped data, including
    UserRoleID, UserID, RoleID, CreatedAt, and UpdatedAt.

```go
func (t *userRoleRecordMapper) ToUserRoleRecord(userRole *db.UserRole) *record.UserRoleRecord
```

### `withdrawRecordMapper`

withdrawRecordMapper provides methods to map Withdraw database rows to WithdrawRecord domain models.

```go
type withdrawRecordMapper struct {
}
```

#### Methods

##### `ToWithdrawAmountMonthly`

ToWithdrawAmountMonthly maps a single database row to a WithdrawMonthlyAmount
domain model. It is intended for use with database rows that contain monthly
withdraw statistics.

Args:
  - ss: A pointer to a GetMonthlyWithdrawsRow representing the database row
    with monthly withdraw statistics.

Returns:
  - A pointer to a WithdrawMonthlyAmount containing the mapped data,
    including Month and TotalAmount.

```go
func (r *withdrawRecordMapper) ToWithdrawAmountMonthly(ss *db.GetMonthlyWithdrawsRow) *record.WithdrawMonthlyAmount
```

##### `ToWithdrawAmountMonthlyByCardNumber`

ToWithdrawAmountMonthlyByCardNumber maps a database row representing monthly
withdraw statistics, filtered by card number, to a WithdrawMonthlyAmount
domain model.

Args:
  - ss: A pointer to a GetMonthlyWithdrawsByCardNumberRow representing the
    database row with monthly withdraw statistics filtered by card number.

Returns:
  - A pointer to a WithdrawMonthlyAmount containing the mapped data,
    including Month and TotalAmount.

```go
func (r *withdrawRecordMapper) ToWithdrawAmountMonthlyByCardNumber(ss *db.GetMonthlyWithdrawsByCardNumberRow) *record.WithdrawMonthlyAmount
```

##### `ToWithdrawAmountYearly`

ToWithdrawAmountYearly maps a single database row to a WithdrawYearlyAmount
domain model. It is intended for use with database rows that contain yearly
withdraw statistics.

Args:
  - ss: A pointer to a GetYearlyWithdrawsRow representing the database row
    with yearly withdraw statistics.

Returns:
  - A pointer to a WithdrawYearlyAmount containing the mapped data,
    including Year and TotalAmount.

```go
func (r *withdrawRecordMapper) ToWithdrawAmountYearly(ss *db.GetYearlyWithdrawsRow) *record.WithdrawYearlyAmount
```

##### `ToWithdrawAmountYearlyByCardNumber`

ToWithdrawAmountYearlyByCardNumber maps a database row representing yearly
withdraw statistics, filtered by card number, to a WithdrawYearlyAmount
domain model.

Args:
  - ss: A pointer to a GetYearlyWithdrawsByCardNumberRow representing the
    database row with yearly withdraw statistics filtered by card number.

Returns:
  - A pointer to a WithdrawYearlyAmount containing the mapped data,
    including Year and TotalAmount.

```go
func (r *withdrawRecordMapper) ToWithdrawAmountYearlyByCardNumber(ss *db.GetYearlyWithdrawsByCardNumberRow) *record.WithdrawYearlyAmount
```

##### `ToWithdrawByCardNumberRecord`

ToWithdrawByCardNumberRecord maps a database row representing a withdraw record
associated with a given card number to a WithdrawRecord domain model.

Args:
  - withdraw: A pointer to a GetWithdrawsByCardNumberRow representing the database row.

Returns:
  - A pointer to a WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *withdrawRecordMapper) ToWithdrawByCardNumberRecord(withdraw *db.GetWithdrawsByCardNumberRow) *record.WithdrawRecord
```

##### `ToWithdrawRecord`

ToWithdrawRecord maps a Withdraw database row to a WithdrawRecord domain model.

Args:
  - withdraw: A pointer to a Withdraw representing the database row.

Returns:
  - A pointer to a WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *withdrawRecordMapper) ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord
```

##### `ToWithdrawRecordActive`

ToWithdrawRecordActive maps a single GetActiveWithdrawsRow database row to a WithdrawRecord
domain model.

Args:
  - withdraw: A pointer to a GetActiveWithdrawsRow representing the database row.

Returns:
  - A pointer to a WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and
    optionally DeletedAt if it is valid.

```go
func (s *withdrawRecordMapper) ToWithdrawRecordActive(withdraw *db.GetActiveWithdrawsRow) *record.WithdrawRecord
```

##### `ToWithdrawRecordAll`

ToWithdrawRecordAll maps a GetWithdrawsRow database row to a WithdrawRecord domain model.

Args:
  - withdraw: A pointer to a GetWithdrawsRow representing the database row.

Returns:
  - A pointer to a WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and optionally DeletedAt if it is valid.

```go
func (s *withdrawRecordMapper) ToWithdrawRecordAll(withdraw *db.GetWithdrawsRow) *record.WithdrawRecord
```

##### `ToWithdrawRecordMonthStatusFailed`

ToWithdrawRecordMonthStatusFailed maps a database row representing monthly
failed withdraw statistics to a WithdrawRecordMonthStatusFailed domain model.

Args:
  - s: A pointer to a GetMonthWithdrawStatusFailedRow representing the database row.

Returns:
  - A pointer to a WithdrawRecordMonthStatusFailed containing the mapped data,
    including Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusFailed(s *db.GetMonthWithdrawStatusFailedRow) *record.WithdrawRecordMonthStatusFailed
```

##### `ToWithdrawRecordMonthStatusFailedCardNumber`

ToWithdrawRecordMonthStatusFailedCardNumber maps a GetMonthWithdrawStatusFailedCardNumberRow database row
to a WithdrawRecordMonthStatusFailed domain model, filtered by card number.

Args:
  - s: A pointer to a GetMonthWithdrawStatusFailedCardNumberRow representing the
    database row.

Returns:
  - A pointer to a WithdrawRecordMonthStatusFailed containing the mapped data,
    including Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusFailedCardNumber(s *db.GetMonthWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordMonthStatusFailed
```

##### `ToWithdrawRecordMonthStatusSuccess`

ToWithdrawRecordMonthStatusSuccess maps a database row representing monthly
successful withdraw statistics to a WithdrawRecordMonthStatusSuccess domain model.

Args:
  - s: A pointer to a GetMonthWithdrawStatusSuccessRow representing the database row.

Returns:
  - A pointer to a WithdrawRecordMonthStatusSuccess containing the mapped data,
    including Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusSuccess(s *db.GetMonthWithdrawStatusSuccessRow) *record.WithdrawRecordMonthStatusSuccess
```

##### `ToWithdrawRecordMonthStatusSuccessCardNumber`

ToWithdrawRecordMonthStatusSuccessCardNumber maps a database row representing
monthly successful withdraw statistics, filtered by card number, to a
WithdrawRecordMonthStatusSuccess domain model.

Args:
  - s: A pointer to a GetMonthWithdrawStatusSuccessCardNumberRow representing
    the database row.

Returns:
  - A pointer to a WithdrawRecordMonthStatusSuccess containing the mapped data,
    including Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusSuccessCardNumber(s *db.GetMonthWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordMonthStatusSuccess
```

##### `ToWithdrawRecordTrashed`

ToWithdrawRecordTrashed maps a GetTrashedWithdrawsRow database row to a WithdrawRecord domain model.
It is intended for use with database rows that contain trashed withdraw records.
It returns a pointer to a WithdrawRecord containing the mapped data, including ID, WithdrawNo,
CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *withdrawRecordMapper) ToWithdrawRecordTrashed(withdraw *db.GetTrashedWithdrawsRow) *record.WithdrawRecord
```

##### `ToWithdrawRecordYearStatusFailed`

ToWithdrawRecordYearStatusFailed maps a database row representing yearly
failed withdraw statistics to a WithdrawRecordYearStatusFailed domain model.

Args:
  - s: A pointer to a GetYearlyWithdrawStatusFailedRow representing the database row.

Returns:
  - A pointer to a WithdrawRecordYearStatusFailed containing the mapped data,
    including Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusFailed(s *db.GetYearlyWithdrawStatusFailedRow) *record.WithdrawRecordYearStatusFailed
```

##### `ToWithdrawRecordYearStatusFailedCardNumber`

ToWithdrawRecordYearStatusFailedCardNumber maps a single
GetYearlyWithdrawStatusFailedCardNumberRow database row to a
WithdrawRecordYearStatusFailed domain model, filtered by card number.

Args:
  - s: A pointer to a GetYearlyWithdrawStatusFailedCardNumberRow
    representing the database row with yearly failed withdraw statistics
    filtered by card number.

Returns:
  - A pointer to a WithdrawRecordYearStatusFailed containing the mapped
    data, including Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusFailedCardNumber(s *db.GetYearlyWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordYearStatusFailed
```

##### `ToWithdrawRecordYearStatusSuccess`

ToWithdrawRecordYearStatusSuccess maps a database row representing yearly
successful withdraw statistics to a WithdrawRecordYearStatusSuccess domain model.

Args:
  - s: A pointer to a GetYearlyWithdrawStatusSuccessRow representing the database row.

Returns:
  - A pointer to a WithdrawRecordYearStatusSuccess containing the mapped data,
    including Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusSuccess(s *db.GetYearlyWithdrawStatusSuccessRow) *record.WithdrawRecordYearStatusSuccess
```

##### `ToWithdrawRecordYearStatusSuccessCardNumber`

ToWithdrawRecordYearStatusSuccessCardNumber maps a database row representing yearly
successful withdraw statistics, filtered by card number, to a
WithdrawRecordYearStatusSuccess domain model.

Args:
  - s: A pointer to a GetYearlyWithdrawStatusSuccessCardNumberRow representing
    the database row.

Returns:
  - A pointer to a WithdrawRecordYearStatusSuccess containing the mapped data,
    including Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusSuccessCardNumber(s *db.GetYearlyWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordYearStatusSuccess
```

##### `ToWithdrawRecordsMonthStatusFailed`

ToWithdrawRecordsMonthStatusFailed maps a slice of GetMonthWithdrawStatusFailedRow
database rows to a slice of WithdrawRecordMonthStatusFailed domain models.

Args:
  - Withdraws: A slice of pointers to GetMonthWithdrawStatusFailedRow representing
    the database rows containing monthly failed withdraw statistics.

Returns:
  - A slice of pointers to WithdrawRecordMonthStatusFailed, each containing
    the mapped data including Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusFailed(Withdraws []*db.GetMonthWithdrawStatusFailedRow) []*record.WithdrawRecordMonthStatusFailed
```

##### `ToWithdrawRecordsMonthStatusFailedCardNumber`

ToWithdrawRecordsMonthStatusFailedCardNumber maps a slice of
GetMonthWithdrawStatusFailedCardNumberRow database rows to a slice of
WithdrawRecordMonthStatusFailed domain models, filtered by card number.

Args:
  - Withdraws: A slice of pointers to GetMonthWithdrawStatusFailedCardNumberRow
    representing the database rows with monthly failed withdraw statistics
    filtered by card number.

Returns:
  - A slice of pointers to WithdrawRecordMonthStatusFailed, each containing
    the mapped data including Year, Month, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusFailedCardNumber(Withdraws []*db.GetMonthWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordMonthStatusFailed
```

##### `ToWithdrawRecordsMonthStatusSuccess`

ToWithdrawRecordsMonthStatusSuccess converts a slice of GetMonthWithdrawStatusSuccessRow
database rows to a slice of WithdrawRecordMonthStatusSuccess domain models.

Args:
  - Withdraws: A slice of pointers to GetMonthWithdrawStatusSuccessRow representing
    the database rows containing monthly successful withdraw statistics.

Returns:
  - A slice of pointers to WithdrawRecordMonthStatusSuccess, each containing
    the mapped data including Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusSuccess(Withdraws []*db.GetMonthWithdrawStatusSuccessRow) []*record.WithdrawRecordMonthStatusSuccess
```

##### `ToWithdrawRecordsMonthStatusSuccessCardNumber`

ToWithdrawRecordsMonthStatusSuccessCardNumber maps a slice of
GetMonthWithdrawStatusSuccessCardNumberRow database rows to a slice of
WithdrawRecordMonthStatusSuccess domain models, filtered by card number.

Args:
  - Withdraws: A slice of pointers to GetMonthWithdrawStatusSuccessCardNumberRow
    representing the database rows with monthly successful withdraw statistics
    filtered by card number.

Returns:
  - A slice of pointers to WithdrawRecordMonthStatusSuccess, each containing
    the mapped data including Year, Month, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusSuccessCardNumber(Withdraws []*db.GetMonthWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordMonthStatusSuccess
```

##### `ToWithdrawRecordsYearStatusFailed`

ToWithdrawRecordsYearStatusFailed maps a slice of GetYearlyWithdrawStatusFailedRow
database rows to a slice of WithdrawRecordYearStatusFailed domain models.

Args:
  - Withdraws: A slice of pointers to GetYearlyWithdrawStatusFailedRow representing
    the database rows containing yearly failed withdraw statistics.

Returns:
  - A slice of pointers to WithdrawRecordYearStatusFailed, each containing
    the mapped data including Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusFailed(Withdraws []*db.GetYearlyWithdrawStatusFailedRow) []*record.WithdrawRecordYearStatusFailed
```

##### `ToWithdrawRecordsYearStatusFailedCardNumber`

ToWithdrawRecordsYearStatusFailedCardNumber maps a slice of
GetYearlyWithdrawStatusFailedCardNumberRow database rows to a slice of
WithdrawRecordYearStatusFailed domain models, filtered by card number.

Args:
  - Withdraws: A slice of pointers to GetYearlyWithdrawStatusFailedCardNumberRow
    representing the database rows with yearly failed withdraw statistics
    filtered by card number.

Returns:
  - A slice of pointers to WithdrawRecordYearStatusFailed, each containing
    the mapped data including Year, TotalFailed, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusFailedCardNumber(Withdraws []*db.GetYearlyWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordYearStatusFailed
```

##### `ToWithdrawRecordsYearStatusSuccess`

ToWithdrawRecordsYearStatusSuccess maps a slice of GetYearlyWithdrawStatusSuccessRow
database rows to a slice of WithdrawRecordYearStatusSuccess domain models.

Args:
  - Withdraws: A slice of pointers to GetYearlyWithdrawStatusSuccessRow representing
    the database rows containing yearly successful withdraw statistics.

Returns:
  - A slice of pointers to WithdrawRecordYearStatusSuccess, each containing
    the mapped data including Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusSuccess(Withdraws []*db.GetYearlyWithdrawStatusSuccessRow) []*record.WithdrawRecordYearStatusSuccess
```

##### `ToWithdrawRecordsYearStatusSuccessCardNumber`

ToWithdrawRecordsYearStatusSuccessCardNumber maps a slice of
GetYearlyWithdrawStatusSuccessCardNumberRow database rows to a slice of
WithdrawRecordYearStatusSuccess domain models, filtered by card number.

Args:
  - Withdraws: A slice of pointers to GetYearlyWithdrawStatusSuccessCardNumberRow
    representing the database rows with yearly successful withdraw statistics
    filtered by card number.

Returns:
  - A slice of pointers to WithdrawRecordYearStatusSuccess, each containing
    the mapped data including Year, TotalSuccess, and TotalAmount.

```go
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusSuccessCardNumber(Withdraws []*db.GetYearlyWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordYearStatusSuccess
```

##### `ToWithdrawsAmountMonthly`

ToWithdrawsAmountMonthly maps a slice of GetMonthlyWithdrawsRow database rows
to a slice of WithdrawMonthlyAmount domain models.

Args:
  - ss: A slice of pointers to GetMonthlyWithdrawsRow representing the database
    rows with monthly withdraw statistics.

Returns:
  - A slice of pointers to WithdrawMonthlyAmount, each containing the mapped
    data including Month and TotalAmount.

```go
func (s *withdrawRecordMapper) ToWithdrawsAmountMonthly(ss []*db.GetMonthlyWithdrawsRow) []*record.WithdrawMonthlyAmount
```

##### `ToWithdrawsAmountMonthlyByCardNumber`

ToWithdrawsAmountMonthlyByCardNumber maps a slice of
GetMonthlyWithdrawsByCardNumberRow database rows to a slice of
WithdrawMonthlyAmount domain models, filtered by card number.

Args:
  - ss: A slice of pointers to GetMonthlyWithdrawsByCardNumberRow
    representing the database rows with monthly withdraw statistics
    filtered by card number.

Returns:
  - A slice of pointers to WithdrawMonthlyAmount, each containing the
    mapped data including Month and TotalAmount.

```go
func (s *withdrawRecordMapper) ToWithdrawsAmountMonthlyByCardNumber(ss []*db.GetMonthlyWithdrawsByCardNumberRow) []*record.WithdrawMonthlyAmount
```

##### `ToWithdrawsAmountYearly`

ToWithdrawsAmountYearly maps a slice of GetYearlyWithdrawsRow database rows to a slice
of WithdrawYearlyAmount domain models.

Args:
  - ss: A slice of pointers to GetYearlyWithdrawsRow representing the database rows
    with yearly withdraw statistics.

Returns:
  - A slice of pointers to WithdrawYearlyAmount, each containing the mapped
    data including Year and TotalAmount.

```go
func (s *withdrawRecordMapper) ToWithdrawsAmountYearly(ss []*db.GetYearlyWithdrawsRow) []*record.WithdrawYearlyAmount
```

##### `ToWithdrawsAmountYearlyByCardNumber`

ToWithdrawsAmountYearlyByCardNumber maps a slice of
GetYearlyWithdrawsByCardNumberRow database rows to a slice of
WithdrawYearlyAmount domain models, filtered by card number.

Args:
  - ss: A slice of pointers to GetYearlyWithdrawsByCardNumberRow
    representing the database rows with yearly withdraw statistics
    filtered by card number.

Returns:
  - A slice of pointers to WithdrawYearlyAmount, each containing the
    mapped data including Year and TotalAmount.

```go
func (s *withdrawRecordMapper) ToWithdrawsAmountYearlyByCardNumber(ss []*db.GetYearlyWithdrawsByCardNumberRow) []*record.WithdrawYearlyAmount
```

##### `ToWithdrawsByCardNumberRecord`

ToWithdrawsByCardNumberRecord maps a slice of GetWithdrawsByCardNumberRow database rows to a slice of WithdrawRecord domain models.

Args:
  - withdraws: A slice of pointers to GetWithdrawsByCardNumberRow representing the database rows.

Returns:
  - A slice of pointers to WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *withdrawRecordMapper) ToWithdrawsByCardNumberRecord(withdraws []*db.GetWithdrawsByCardNumberRow) []*record.WithdrawRecord
```

##### `ToWithdrawsRecord`

ToWithdrawsRecord maps a slice of Withdraw database rows to a slice of WithdrawRecord
domain models.

Args:
  - withdraws: A slice of pointers to Withdraw representing the database rows.

Returns:
  - A slice of pointers to WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.

```go
func (s *withdrawRecordMapper) ToWithdrawsRecord(withdraws []*db.Withdraw) []*record.WithdrawRecord
```

##### `ToWithdrawsRecordALl`

ToWithdrawsRecordALl maps a slice of GetWithdrawsRow database rows to a slice of WithdrawRecord domain models.

Args:
  - withdraws: A slice of pointers to GetWithdrawsRow representing the database rows.

Returns:
  - A slice of pointers to WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and optionally DeletedAt if it is valid.

```go
func (s *withdrawRecordMapper) ToWithdrawsRecordALl(withdraws []*db.GetWithdrawsRow) []*record.WithdrawRecord
```

##### `ToWithdrawsRecordActive`

ToWithdrawsRecordActive maps a slice of GetActiveWithdrawsRow database rows to a slice of WithdrawRecord domain models.

Args:
  - withdraws: A slice of pointers to GetActiveWithdrawsRow representing the database rows.

Returns:
  - A slice of pointers to WithdrawRecord containing the mapped data, including
    ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and
    optionally DeletedAt if it is valid.

```go
func (s *withdrawRecordMapper) ToWithdrawsRecordActive(withdraws []*db.GetActiveWithdrawsRow) []*record.WithdrawRecord
```

##### `ToWithdrawsRecordTrashed`

ToWithdrawsRecordTrashed maps a slice of GetTrashedWithdrawsRow database rows to a slice of WithdrawRecord
domain models. It is intended for use with database rows that contain trashed withdraw records.

Args:
  - withdraws: A slice of pointers to GetTrashedWithdrawsRow representing the database rows.

Returns:
  - A slice of pointers to WithdrawRecord containing the mapped data, including ID, WithdrawNo,
    CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
```go
func (s *withdrawRecordMapper) ToWithdrawsRecordTrashed(withdraws []*db.GetTrashedWithdrawsRow) []*record.WithdrawRecord
```