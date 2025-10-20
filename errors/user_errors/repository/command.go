package userrepositoryerrors

import "errors"

// ErrUserConflict is returned when a user with the same email address already exists.
var ErrUserConflict = errors.New("failed user already exists")

// ErrCreateUser is returned when creating a user fails.
var ErrCreateUser = errors.New("failed to create user")

// ErrUpdateUser is returned when updating a user fails.
var ErrUpdateUser = errors.New("failed to update user")

// ErrUpdateUserVerificationCode is returned when updating a user verification code fails.
var ErrUpdateUserVerificationCode = errors.New("failed to update user verification code")

// ErrUpdateUserPassword is returned when updating a user password fails.
var ErrUpdateUserPassword = errors.New("failed to update user password")

// ErrTrashedUser is returned when moving a user to trash fails.
var ErrTrashedUser = errors.New("failed to move user to trash")

// ErrRestoreUser is returned when restoring a user from trash fails.
var ErrRestoreUser = errors.New("failed to restore user from trash")

// ErrDeleteUserPermanent is returned when permanently deleting a user fails.
var ErrDeleteUserPermanent = errors.New("failed to permanently delete user")

// ErrRestoreAllUsers is returned when restoring all users fails.
var ErrRestoreAllUsers = errors.New("failed to restore all users")

// ErrDeleteAllUsers is returned when permanently deleting all users fails.
var ErrDeleteAllUsers = errors.New("failed to permanently delete all users")
