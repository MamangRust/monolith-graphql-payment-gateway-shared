package transactionrepositoryerrors

import "errors"

// ErrGetMonthlyPaymentMethodsFailed indicates a failure when retrieving monthly payment method statistics.
var ErrGetMonthlyPaymentMethodsFailed = errors.New("failed to get monthly payment methods")

// ErrGetYearlyPaymentMethodsFailed indicates a failure when retrieving yearly payment method statistics.
var ErrGetYearlyPaymentMethodsFailed = errors.New("failed to get yearly payment methods")

// ErrGetMonthlyPaymentMethodsByCardFailed indicates a failure when retrieving monthly payment methods by card number.
var ErrGetMonthlyPaymentMethodsByCardFailed = errors.New("failed to get monthly payment methods by card number")

// ErrGetYearlyPaymentMethodsByCardFailed indicates a failure when retrieving yearly payment methods by card number.
var ErrGetYearlyPaymentMethodsByCardFailed = errors.New("failed to get yearly payment methods by card number")
