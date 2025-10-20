package merchantrepositoryerrors

import "errors"

// ErrGetMonthlyPaymentMethodsMerchantFailed indicates failure fetching monthly payment methods for a merchant.
var ErrGetMonthlyPaymentMethodsMerchantFailed = errors.New("failed to get monthly payment methods of merchant")

// ErrGetYearlyPaymentMethodMerchantFailed indicates failure fetching yearly payment methods for a merchant.
var ErrGetYearlyPaymentMethodMerchantFailed = errors.New("failed to get yearly payment method of merchant")

// ErrGetMonthlyPaymentMethodByMerchantsFailed indicates failure fetching monthly payment methods for all merchants.
var ErrGetMonthlyPaymentMethodByMerchantsFailed = errors.New("failed to get monthly payment method by merchants")

// ErrGetYearlyPaymentMethodByMerchantsFailed indicates failure fetching yearly payment methods for all merchants.
var ErrGetYearlyPaymentMethodByMerchantsFailed = errors.New("failed to get yearly payment method by merchants")

// ErrGetMonthlyPaymentMethodByApikeyFailed indicates failure fetching monthly payment methods by API key.
var ErrGetMonthlyPaymentMethodByApikeyFailed = errors.New("failed to get monthly payment method by API key")

// ErrGetYearlyPaymentMethodByApikeyFailed indicates failure fetching yearly payment methods by API key.
var ErrGetYearlyPaymentMethodByApikeyFailed = errors.New("failed to get yearly payment method by API key")
