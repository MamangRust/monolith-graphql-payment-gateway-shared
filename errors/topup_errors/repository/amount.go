package topuprepositoryerrors

import "errors"

// ErrGetMonthlyTopupAmountsFailed indicates failure in retrieving the monthly top-up amounts.
var ErrGetMonthlyTopupAmountsFailed = errors.New("failed to get monthly topup amounts")

// ErrGetYearlyTopupAmountsFailed indicates failure in retrieving the yearly top-up amounts.
var ErrGetYearlyTopupAmountsFailed = errors.New("failed to get yearly topup amounts")

// ErrGetMonthlyTopupAmountsByCardFailed indicates failure in retrieving monthly top-up amount stats by card number.
var ErrGetMonthlyTopupAmountsByCardFailed = errors.New("failed to get monthly topup amounts by card number")

// ErrGetYearlyTopupAmountsByCardFailed indicates failure in retrieving yearly top-up amount stats by card number.
var ErrGetYearlyTopupAmountsByCardFailed = errors.New("failed to get yearly topup amounts by card number")
