package rolerepositoryerrors

import "errors"

// ErrRoleNotFound indicates that the requested Role was not found in the database.
var ErrRoleNotFound = errors.New("role not found")

// ErrFindAllRoles is returned when retrieving all Roles from the database fails.
var ErrFindAllRoles = errors.New("failed to find all Roles")

// ErrFindActiveRoles is returned when retrieving all active Roles fails.
var ErrFindActiveRoles = errors.New("failed to find active Roles")

// ErrFindTrashedRoles is returned when retrieving trashed (soft-deleted) Roles fails.
var ErrFindTrashedRoles = errors.New("failed to find trashed Roles")

// ErrRoleConflict indicates a conflict where a Role already exists.
var ErrRoleConflict = errors.New("failed Role already exists")
