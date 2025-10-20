package rolerepositoryerrors

import "errors"

// ErrCreateRole is returned when creating a new Role fails.
var ErrCreateRole = errors.New("failed to create Role")

// ErrUpdateRole is returned when updating an existing Role fails.
var ErrUpdateRole = errors.New("failed to update Role")

// ErrTrashedRole is returned when moving a Role to trash (soft-delete) fails.
var ErrTrashedRole = errors.New("failed to move Role to trash")

// ErrRestoreRole is returned when restoring a trashed Role fails.
var ErrRestoreRole = errors.New("failed to restore Role from trash")

// ErrDeleteRolePermanent is returned when permanently deleting a Role fails.
var ErrDeleteRolePermanent = errors.New("failed to permanently delete Role")

// ErrRestoreAllRoles is returned when restoring all trashed Roles fails.
var ErrRestoreAllRoles = errors.New("failed to restore all Roles")

// ErrDeleteAllRoles is returned when permanently deleting all Roles fails.
var ErrDeleteAllRoles = errors.New("failed to permanently delete all Roles")
