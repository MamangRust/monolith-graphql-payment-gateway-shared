package userroleserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrFailedAssignRoleToUser is an error that is returned from the service layer
// when an error occurs while trying to assign a role to the user.
var ErrFailedAssignRoleToUser = response.NewErrorResponse(
	"Failed to assign role to user",
	http.StatusInternalServerError,
)

// ErrFailedRemoveRole is an error that is returned from the service layer
// when an error occurs while trying to remove a role from the user.
var ErrFailedRemoveRole = response.NewErrorResponse(
	"Failed to remove role from user",
	http.StatusInternalServerError,
)
