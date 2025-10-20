package merchantrepositoryerrors

import "errors"

// ErrGetMonthlyAmountMerchantFailed indicates failure fetching monthly amount for a merchant.
var ErrGetMonthlyAmountMerchantFailed = errors.New("failed to get monthly amount of merchant")

// ErrGetYearlyAmountMerchantFailed indicates failure fetching yearly amount for a merchant.
var ErrGetYearlyAmountMerchantFailed = errors.New("failed to get yearly amount of merchant")

// ErrGetMonthlyAmountByMerchantsFailed indicates failure fetching monthly amount for all merchants.
var ErrGetMonthlyAmountByMerchantsFailed = errors.New("failed to get monthly amount by merchants")

// ErrGetYearlyAmountByMerchantsFailed indicates failure fetching yearly amount for all merchants.
var ErrGetYearlyAmountByMerchantsFailed = errors.New("failed to get yearly amount by merchants")

// ErrGetMonthlyAmountByApikeyFailed indicates failure fetching monthly amount by API key.
var ErrGetMonthlyAmountByApikeyFailed = errors.New("failed to get monthly amount by API key")

// ErrGetYearlyAmountByApikeyFailed indicates failure fetching yearly amount by API key.TotalAmount
var ErrGetYearlyAmountByApikeyFailed = errors.New("failed to get yearly amount by API key")
