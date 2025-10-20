package cardrepositoryerrors

import "errors"

// ErrCreateCardFailed is returned when creating a new card fails.
var ErrCreateCardFailed = errors.New("failed to create card")

// ErrUpdateCardFailed is returned when updating a card fails.
var ErrUpdateCardFailed = errors.New("failed to update card")

// ErrTrashCardFailed is returned when trashing a card fails.
var ErrTrashCardFailed = errors.New("failed to trash card")

// ErrRestoreCardFailed is returned when restoring a trashed card fails.
var ErrRestoreCardFailed = errors.New("failed to restore card")

// ErrDeleteCardPermanentFailed is returned when permanently deleting a card fails.
var ErrDeleteCardPermanentFailed = errors.New("failed to delete card permanently")

// ErrRestoreAllCardsFailed is returned when restoring all trashed cards fails.
var ErrRestoreAllCardsFailed = errors.New("failed to restore all cards")

// ErrDeleteAllCardsPermanentFailed is returned when permanently deleting all cards fails.
var ErrDeleteAllCardsPermanentFailed = errors.New("failed to delete all cards permanently")
