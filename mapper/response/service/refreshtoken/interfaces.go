package refreshtokenservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type RefreshTokenResponseMapper interface {
	// Converts a single RefreshTokenRecord into RefreshTokenResponse.
	ToRefreshTokenResponse(refresh *record.RefreshTokenRecord) *response.RefreshTokenResponse

	// Converts multiple RefreshTokenRecord into a slice of RefreshTokenResponse.
	ToRefreshTokenResponses(refreshs []*record.RefreshTokenRecord) []*response.RefreshTokenResponse
}
