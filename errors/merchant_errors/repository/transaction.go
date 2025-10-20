package merchantrepositoryerrors

import "errors"

// ErrFindAllTransactionsFailed indicates failure fetching all transactions.
var ErrFindAllTransactionsFailed = errors.New("failed to find all merchant transactions")

// ErrFindAllTransactionsByMerchantFailed indicates failure fetching transactions by merchant ID.
var ErrFindAllTransactionsByMerchantFailed = errors.New("failed to find merchant transactions by merchant ID")

// ErrFindAllTransactionsByApiKeyFailed indicates failure fetching transactions by API key.
var ErrFindAllTransactionsByApiKeyFailed = errors.New("failed to find merchant transactions by API key")
