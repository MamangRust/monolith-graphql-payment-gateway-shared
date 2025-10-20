package userrolerepositoryerrors

import "errors"

// ErrAssignRoleToUser is an error that is returned from the repository layer
// when an error occurs while trying to assign a role to the user.
var ErrAssignRoleToUser = errors.New("failed to assign role to user")

// ErrRemoveRole is an error that is returned from the repository layer
// when an error occurs while trying to remove a role from the user.
var ErrRemoveRole = errors.New("failed to remove role from user")
