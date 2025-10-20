package merchantdocumentrepositoryerrors

import "errors"

// ErrCreateMerchantDocumentFailed is returned when failing to create a new merchant document.
var ErrCreateMerchantDocumentFailed = errors.New("failed to create merchant document")

// ErrUpdateMerchantDocumentFailed is returned when failing to update an existing merchant document.
var ErrUpdateMerchantDocumentFailed = errors.New("failed to update merchant document")

// ErrUpdateMerchantDocumentStatusFailed is returned when failing to update the status of a merchant document.
var ErrUpdateMerchantDocumentStatusFailed = errors.New("failed to update merchant document status")

// ErrTrashedMerchantDocumentFailed is returned when failing to move a merchant document to trash.
var ErrTrashedMerchantDocumentFailed = errors.New("failed to soft-delete (trash) merchant document")

// ErrRestoreMerchantDocumentFailed is returned when failing to restore a trashed merchant document.
var ErrRestoreMerchantDocumentFailed = errors.New("failed to restore merchant document")

// ErrDeleteMerchantDocumentPermanentFailed is returned when failing to permanently delete a merchant document.
var ErrDeleteMerchantDocumentPermanentFailed = errors.New("failed to permanently delete merchant document")

// ErrRestoreAllMerchantDocumentsFailed is returned when failing to restore all trashed merchant documents.
var ErrRestoreAllMerchantDocumentsFailed = errors.New("failed to restore all merchant documents")

// ErrDeleteAllMerchantDocumentsPermanentFailed is returned when failing to permanently delete all merchant documents.
var ErrDeleteAllMerchantDocumentsPermanentFailed = errors.New("failed to permanently delete all merchant documents")
