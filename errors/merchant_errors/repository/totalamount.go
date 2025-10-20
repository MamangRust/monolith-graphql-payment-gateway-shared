package merchantrepositoryerrors

import "errors"

// ErrGetMonthlyTotalAmountMerchantFailed indicates failure fetching monthly total amount for a merchant.
var ErrGetMonthlyTotalAmountMerchantFailed = errors.New("failed to get monthly total amount of merchant")

// ErrGetYearlyTotalAmountMerchantFailed indicates failure fetching yearly total amount for a merchant.
var ErrGetYearlyTotalAmountMerchantFailed = errors.New("failed to get yearly total amount of merchant")

// ErrGetMonthlyTotalAmountByMerchantsFailed indicates failure fetching monthly total amount for all merchants.
var ErrGetMonthlyTotalAmountByMerchantsFailed = errors.New("failed to get monthly total amount by merchants")

// ErrGetYearlyTotalAmountByMerchantsFailed indicates failure fetching yearly total amount for all merchants.
var ErrGetYearlyTotalAmountByMerchantsFailed = errors.New("failed to get yearly total amount by merchants")

// ErrGetMonthlyTotalAmountByApikeyFailed indicates failure fetching monthly total amount by API key.
var ErrGetMonthlyTotalAmountByApikeyFailed = errors.New("failed to get monthly total amount by API key")

// ErrGetYearlyTotalAmountByApikeyFailed indicates failure fetching yearly total amount by API key.
var ErrGetYearlyTotalAmountByApikeyFailed = errors.New("failed to get yearly total amount by API key")
