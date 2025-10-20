package transferrepositoryerrors

import "errors"

// ErrGetMonthlyTransferAmountsFailed indicates a failure when retrieving the total amount of monthly transfers.
var ErrGetMonthlyTransferAmountsFailed = errors.New("failed to get monthly transfer amounts")

// ErrGetYearlyTransferAmountsFailed indicates a failure when retrieving the total amount of yearly transfers.
var ErrGetYearlyTransferAmountsFailed = errors.New("failed to get yearly transfer amounts")

// ErrGetMonthlyTransferAmountsBySenderCardFailed indicates a failure when retrieving monthly transfer amounts filtered by sender card number.
var ErrGetMonthlyTransferAmountsBySenderCardFailed = errors.New("failed to get monthly transfer amounts by sender card number")

// ErrGetYearlyTransferAmountsBySenderCardFailed indicates a failure when retrieving yearly transfer amounts filtered by sender card number.
var ErrGetYearlyTransferAmountsBySenderCardFailed = errors.New("failed to get yearly transfer amounts by sender card number")

// ErrGetMonthlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving monthly transfer amounts filtered by receiver card number.
var ErrGetMonthlyTransferAmountsByReceiverCardFailed = errors.New("failed to get monthly transfer amounts by receiver card number")

// ErrGetYearlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving yearly transfer amounts filtered by receiver card number.
var ErrGetYearlyTransferAmountsByReceiverCardFailed = errors.New("failed to get yearly transfer amounts by receiver card number")
