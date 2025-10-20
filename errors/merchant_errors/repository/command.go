package merchantrepositoryerrors

import "errors"

// ErrCreateMerchantFailed indicates failure when creating a merchant.
var ErrCreateMerchantFailed = errors.New("failed to create merchant")

// ErrUpdateMerchantFailed indicates failure when updating a merchant.
var ErrUpdateMerchantFailed = errors.New("failed to update merchant")

// ErrUpdateMerchantStatusFailed indicates failure when updating merchant status.
var ErrUpdateMerchantStatusFailed = errors.New("failed to update merchant status")

// ErrTrashedMerchantFailed indicates failure when soft-deleting a merchant.
var ErrTrashedMerchantFailed = errors.New("failed to soft-delete (trash) merchant")

// ErrRestoreMerchantFailed indicates failure when restoring a trashed merchant.
var ErrRestoreMerchantFailed = errors.New("failed to restore merchant")

// ErrDeleteMerchantPermanentFailed indicates failure when permanently deleting a merchant.
var ErrDeleteMerchantPermanentFailed = errors.New("failed to permanently delete merchant")

// ErrRestoreAllMerchantFailed indicates failure when restoring all trashed merchants.
var ErrRestoreAllMerchantFailed = errors.New("failed to restore all merchants")

// ErrDeleteAllMerchantPermanentFailed indicates failure when permanently deleting all merchants.
var ErrDeleteAllMerchantPermanentFailed = errors.New("failed to permanently delete all merchants")
