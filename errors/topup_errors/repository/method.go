package topuprepositoryerrors

import "errors"

// ErrGetMonthlyTopupMethodsFailed indicates failure in retrieving monthly top-up payment methods statistics.
var ErrGetMonthlyTopupMethodsFailed = errors.New("failed to get monthly topup methods")

// ErrGetYearlyTopupMethodsFailed indicates failure in retrieving yearly top-up payment methods statistics.
var ErrGetYearlyTopupMethodsFailed = errors.New("failed to get yearly topup methods")

// ErrGetMonthlyTopupMethodsByCardFailed indicates failure in retrieving monthly payment method stats by card number.
var ErrGetMonthlyTopupMethodsByCardFailed = errors.New("failed to get monthly topup methods by card number")

// ErrGetYearlyTopupMethodsByCardFailed indicates failure in retrieving yearly payment method stats by card number.
var ErrGetYearlyTopupMethodsByCardFailed = errors.New("failed to get yearly topup methods by card number")
