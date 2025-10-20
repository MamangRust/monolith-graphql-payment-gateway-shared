package transactionrepositoryerrors

import "errors"

// ErrGetMonthTransactionStatusSuccessFailed indicates a failure when retrieving monthly successful transaction status.
var ErrGetMonthTransactionStatusSuccessFailed = errors.New("failed to get monthly transaction status success")

// ErrGetYearlyTransactionStatusSuccessFailed indicates a failure when retrieving yearly successful transaction status.
var ErrGetYearlyTransactionStatusSuccessFailed = errors.New("failed to get yearly transaction status success")

// ErrGetMonthTransactionStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transactions by card number.
var ErrGetMonthTransactionStatusSuccessByCardFailed = errors.New("failed to get monthly transaction status success by card number")

// ErrGetYearlyTransactionStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transactions by card number.
var ErrGetYearlyTransactionStatusSuccessByCardFailed = errors.New("failed to get yearly transaction status success by card number")

// ErrGetMonthTransactionStatusFailedFailed indicates a failure when retrieving monthly failed transaction status.
var ErrGetMonthTransactionStatusFailedFailed = errors.New("failed to get monthly transaction status failed")

// ErrGetYearlyTransactionStatusFailedFailed indicates a failure when retrieving yearly failed transaction status.
var ErrGetYearlyTransactionStatusFailedFailed = errors.New("failed to get yearly transaction status failed")

// ErrGetMonthTransactionStatusFailedByCardFailed indicates a failure when retrieving monthly failed transactions by card number.
var ErrGetMonthTransactionStatusFailedByCardFailed = errors.New("failed to get monthly transaction status failed by card number")

// ErrGetYearlyTransactionStatusFailedByCardFailed indicates a failure when retrieving yearly failed transactions by card number.
var ErrGetYearlyTransactionStatusFailedByCardFailed = errors.New("failed to get yearly transaction status failed by card number")
