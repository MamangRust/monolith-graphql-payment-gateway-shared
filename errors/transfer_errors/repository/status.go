package transferrepositoryerrors

import "errors"

// ErrGetMonthTransferStatusSuccessFailed indicates a failure when retrieving monthly successful transfer statistics.
var ErrGetMonthTransferStatusSuccessFailed = errors.New("failed to get monthly transfer status success")

// ErrGetYearlyTransferStatusSuccessFailed indicates a failure when retrieving yearly successful transfer statistics.
var ErrGetYearlyTransferStatusSuccessFailed = errors.New("failed to get yearly transfer status success")

// ErrGetMonthTransferStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transfers filtered by card number.
var ErrGetMonthTransferStatusSuccessByCardFailed = errors.New("failed to get monthly transfer status success by card number")

// ErrGetYearlyTransferStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transfers filtered by card number.
var ErrGetYearlyTransferStatusSuccessByCardFailed = errors.New("failed to get yearly transfer status success by card number")

// ErrGetMonthTransferStatusFailedFailed indicates a failure when retrieving monthly failed transfer statistics.
var ErrGetMonthTransferStatusFailedFailed = errors.New("failed to get monthly transfer status failed")

// ErrGetYearlyTransferStatusFailedFailed indicates a failure when retrieving yearly failed transfer statistics.
var ErrGetYearlyTransferStatusFailedFailed = errors.New("failed to get yearly transfer status failed")

// ErrGetMonthTransferStatusFailedByCardFailed indicates a failure when retrieving monthly failed transfers filtered by card number.
var ErrGetMonthTransferStatusFailedByCardFailed = errors.New("failed to get monthly transfer status failed by card number")

// ErrGetYearlyTransferStatusFailedByCardFailed indicates a failure when retrieving yearly failed transfers filtered by card number.
var ErrGetYearlyTransferStatusFailedByCardFailed = errors.New("failed to get yearly transfer status failed by card number")
