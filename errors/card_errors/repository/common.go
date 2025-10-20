package cardrepositoryerrors

import "errors"

// ErrInvalidCardRequest is returned when the card request data is invalid.
var ErrInvalidCardRequest = errors.New("invalid card request data")

// ErrInvalidCardId is returned when the card ID is invalid.
var ErrInvalidCardId = errors.New("invalid card ID")

// ErrInvalidUserId is returned when the user ID is invalid.
var ErrInvalidUserId = errors.New("invalid user ID")

// ErrInvalidCardNumber is returned when the card number is invalid.
var ErrInvalidCardNumber = errors.New("invalid card number")
