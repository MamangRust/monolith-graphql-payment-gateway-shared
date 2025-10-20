package transactionrepositoryerrors

import "errors"

// ErrGetMonthlyAmountsFailed indicates a failure when retrieving total monthly transaction amounts.
var ErrGetMonthlyAmountsFailed = errors.New("failed to get monthly amounts")

// ErrGetYearlyAmountsFailed indicates a failure when retrieving total yearly transaction amounts.
var ErrGetYearlyAmountsFailed = errors.New("failed to get yearly amounts")

// ErrGetMonthlyAmountsByCardFailed indicates a failure when retrieving monthly amounts by card number.
var ErrGetMonthlyAmountsByCardFailed = errors.New("failed to get monthly amounts by card number")

// ErrGetYearlyAmountsByCardFailed indicates a failure when retrieving yearly amounts by card number.
var ErrGetYearlyAmountsByCardFailed = errors.New("failed to get yearly amounts by card number")
