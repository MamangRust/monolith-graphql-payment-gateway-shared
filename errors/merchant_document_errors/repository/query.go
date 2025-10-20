package merchantdocumentrepositoryerrors

import "errors"

// ErrFindAllMerchantDocumentsFailed is returned when failing to retrieve all merchant documents.
var ErrFindAllMerchantDocumentsFailed = errors.New("failed to find all merchant documents")

// ErrFindActiveMerchantDocumentsFailed is returned when failing to retrieve active merchant documents.
var ErrFindActiveMerchantDocumentsFailed = errors.New("failed to find active merchant documents")

// ErrFindTrashedMerchantDocumentsFailed is returned when failing to retrieve trashed merchant documents.
var ErrFindTrashedMerchantDocumentsFailed = errors.New("failed to find trashed merchant documents")

// ErrFindMerchantDocumentByIdFailed is returned when failing to retrieve a merchant document by ID.
var ErrFindMerchantDocumentByIdFailed = errors.New("failed to find merchant document by ID")
