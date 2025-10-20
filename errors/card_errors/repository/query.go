package cardrepositoryerrors

import "errors"

// ErrFindAllCardsFailed is returned when fetching all card records fails.
var ErrFindAllCardsFailed = errors.New("failed to find all cards")

// ErrFindActiveCardsFailed is returned when fetching active card records fails.
var ErrFindActiveCardsFailed = errors.New("failed to find active cards")

// ErrFindTrashedCardsFailed is returned when fetching trashed card records fails.
var ErrFindTrashedCardsFailed = errors.New("failed to find trashed cards")

// ErrFindCardByIdFailed is returned when fetching a card by its ID fails.
var ErrFindCardByIdFailed = errors.New("failed to find card by ID")

// ErrFindCardByUserIdFailed is returned when fetching a card by user ID fails.
var ErrFindCardByUserIdFailed = errors.New("failed to find card by user ID")

// ErrFindCardByCardNumberFailed is returned when fetching a card by card number fails.
var ErrFindCardByCardNumberFailed = errors.New("failed to find card by card number")

// ErrCardAlreadyExists is returned when a card already exists.
var ErrCardAlreadyExists = errors.New("card already exists")

// ErrCardNotFound is returned when a card is not found.
var ErrCardNotFound = errors.New("card not found")
