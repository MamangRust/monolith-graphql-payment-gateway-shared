package topuprepositoryerrors

import "errors"

// ErrGetMonthTopupStatusSuccessFailed indicates failure in getting the monthly count of successful top-ups.
var ErrGetMonthTopupStatusSuccessFailed = errors.New("failed to get monthly topup status success")

// ErrGetYearlyTopupStatusSuccessFailed indicates failure in getting the yearly count of successful top-ups.
var ErrGetYearlyTopupStatusSuccessFailed = errors.New("failed to get yearly topup status success")

// ErrGetMonthTopupStatusSuccessByCardFailed indicates failure in getting monthly successful top-up status by card number.
var ErrGetMonthTopupStatusSuccessByCardFailed = errors.New("failed to get monthly topup status success by card number")

// ErrGetYearlyTopupStatusSuccessByCardFailed indicates failure in getting yearly successful top-up status by card number.
var ErrGetYearlyTopupStatusSuccessByCardFailed = errors.New("failed to get yearly topup status success by card number")

// ErrGetMonthTopupStatusFailedFailed indicates failure in getting the monthly count of failed top-ups.
var ErrGetMonthTopupStatusFailedFailed = errors.New("failed to get monthly topup status failed")

// ErrGetYearlyTopupStatusFailedFailed indicates failure in getting the yearly count of failed top-ups.
var ErrGetYearlyTopupStatusFailedFailed = errors.New("failed to get yearly topup status failed")

// ErrGetMonthTopupStatusFailedByCardFailed indicates failure in getting monthly failed top-up status by card number.
var ErrGetMonthTopupStatusFailedByCardFailed = errors.New("failed to get monthly topup status failed by card number")

// ErrGetYearlyTopupStatusFailedByCardFailed indicates failure in getting yearly failed top-up status by card number.
var ErrGetYearlyTopupStatusFailedByCardFailed = errors.New("failed to get yearly topup status failed by card number")
