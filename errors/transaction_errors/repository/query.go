package transactionrepositoryerrors

import "errors"

// ErrFindAllTransactionsFailed indicates a failure when attempting to retrieve all transactions.
var ErrFindAllTransactionsFailed = errors.New("failed to find all transactions")

// ErrFindActiveTransactionsFailed indicates a failure when retrieving active (non-deleted) transactions.
var ErrFindActiveTransactionsFailed = errors.New("failed to find active transactions")

// ErrFindTrashedTransactionsFailed indicates a failure when retrieving soft-deleted (trashed) transactions.
var ErrFindTrashedTransactionsFailed = errors.New("failed to find trashed transactions")

// ErrFindTransactionsByCardNumberFailed indicates a failure when retrieving transactions by card number.
var ErrFindTransactionsByCardNumberFailed = errors.New("failed to find transactions by card number")

// ErrFindTransactionByIdFailed indicates a failure when retrieving a transaction by its ID.
var ErrFindTransactionByIdFailed = errors.New("failed to find transaction by ID")

// ErrFindTransactionByMerchantIdFailed indicates a failure when retrieving transactions by merchant ID.
var ErrFindTransactionByMerchantIdFailed = errors.New("failed to find transaction by merchant ID")
