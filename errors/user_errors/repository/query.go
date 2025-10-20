package userrepositoryerrors

import "errors"

// ErrUserNotFound is returned when a user is not found.
var ErrUserNotFound = errors.New("user not found")

// ErrFindAllUsers is returned when fetching all users fails.
var ErrFindAllUsers = errors.New("failed to find all users")

// ErrFindActiveUsers is returned when fetching active users fails.
var ErrFindActiveUsers = errors.New("failed to find active users")

// ErrFindTrashedUsers is returned when fetching trashed users fails.
var ErrFindTrashedUsers = errors.New("failed to find trashed users")
