package transferrepositoryerrors

import "errors"

// ErrFindAllTransfersFailed indicates a failure when retrieving all transfer records.
var ErrFindAllTransfersFailed = errors.New("failed to find all transfers")

// ErrFindActiveTransfersFailed indicates a failure when retrieving active (non-trashed) transfer records.
var ErrFindActiveTransfersFailed = errors.New("failed to find active transfers")

// ErrFindTrashedTransfersFailed indicates a failure when retrieving trashed (soft-deleted) transfer records.
var ErrFindTrashedTransfersFailed = errors.New("failed to find trashed transfers")

// ErrFindTransferByIdFailed indicates a failure when retrieving a transfer record by its ID.
var ErrFindTransferByIdFailed = errors.New("failed to find transfer by ID")

// ErrFindTransferByTransferFromFailed indicates a failure when retrieving transfers by the sender (transfer from).
var ErrFindTransferByTransferFromFailed = errors.New("failed to find transfer by transfer from")

// ErrFindTransferByTransferToFailed indicates a failure when retrieving transfers by the receiver (transfer to).
var ErrFindTransferByTransferToFailed = errors.New("failed to find transfer by transfer to")
