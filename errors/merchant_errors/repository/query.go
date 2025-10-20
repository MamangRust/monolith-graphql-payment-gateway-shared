package merchantrepositoryerrors

import "errors"

// ErrFindAllMerchantsFailed is returned when fetching all merchants fails.
var ErrFindAllMerchantsFailed = errors.New("failed to find all merchants")

// ErrFindActiveMerchantsFailed is returned when fetching active merchants fails.
var ErrFindActiveMerchantsFailed = errors.New("failed to find active merchants")

// ErrFindTrashedMerchantsFailed is returned when fetching trashed merchants fails.
var ErrFindTrashedMerchantsFailed = errors.New("failed to find trashed merchants")

// ErrFindMerchantByIdFailed is returned when a merchant cannot be found by ID.
var ErrFindMerchantByIdFailed = errors.New("failed to find merchant by ID")

// ErrFindMerchantByApiKeyFailed is returned when a merchant cannot be found by API key.
var ErrFindMerchantByApiKeyFailed = errors.New("failed to find merchant by API key")

// ErrFindMerchantByNameFailed is returned when a merchant cannot be found by name.
var ErrFindMerchantByNameFailed = errors.New("failed to find merchant by name")

// ErrFindMerchantByUserIdFailed is returned when a merchant cannot be found by user ID.
var ErrFindMerchantByUserIdFailed = errors.New("failed to find merchant by user ID")
