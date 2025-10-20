package withdrawrepositoryerrors

import "errors"

// ErrGetMonthlyWithdrawsFailed is used when the system fails to get monthly withdraw amounts
var ErrGetMonthlyWithdrawsFailed = errors.New("failed to get monthly withdraw amounts")

// ErrGetYearlyWithdrawsFailed is used when the system fails to get yearly withdraw amounts
var ErrGetYearlyWithdrawsFailed = errors.New("failed to get yearly withdraw amounts")

// ErrGetMonthlyWithdrawsByCardFailed indicates a failure when retrieving monthly withdraw amounts by card number.
var ErrGetMonthlyWithdrawsByCardFailed = errors.New("failed to get monthly withdraw amounts by card number")

// ErrGetYearlyWithdrawsByCardFailed indicates a failure when retrieving yearly withdraw amounts by card number.
var ErrGetYearlyWithdrawsByCardFailed = errors.New("failed to get yearly withdraw amounts by card number")
