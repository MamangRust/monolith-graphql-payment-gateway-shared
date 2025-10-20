package refreshtokenservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// refreshTokenResponseMapper provides methods to map RefreshTokenRecord domain models to RefreshTokenResponse API-compatible response types.
type refreshTokenResponseMapper struct {
}

// NewRefreshTokenResponseMapper creates a new instance of refreshTokenResponseMapper,
// which provides methods to map RefreshTokenRecord domain models to RefreshTokenResponse
// API-compatible response types.
func NewRefreshTokenResponseMapper() RefreshTokenResponseMapper {
	return &refreshTokenResponseMapper{}
}

// ToRefreshTokenResponse maps a RefreshTokenRecord domain model to a RefreshTokenResponse API-compatible response type.
// Args:
//   - refresh: A pointer to a RefreshTokenRecord representing the domain model.
//
// Returns:
//   - A pointer to a RefreshTokenResponse representing the API-compatible response type.
func (r *refreshTokenResponseMapper) ToRefreshTokenResponse(refresh *record.RefreshTokenRecord) *response.RefreshTokenResponse {
	return &response.RefreshTokenResponse{
		UserID:    refresh.UserID,
		Token:     refresh.Token,
		ExpiredAt: refresh.ExpiredAt,
		CreatedAt: refresh.CreatedAt,
		UpdatedAt: refresh.UpdatedAt,
	}
}

// ToRefreshTokenResponses maps a slice of RefreshTokenRecord domain models to a slice of RefreshTokenResponse API-compatible response types.
// Args:
//   - refreshs: A slice of pointers to RefreshTokenRecord representing the domain models.
//
// Returns:
//   - A slice of pointers to RefreshTokenResponse representing the API-compatible response types.
func (r *refreshTokenResponseMapper) ToRefreshTokenResponses(refreshs []*record.RefreshTokenRecord) []*response.RefreshTokenResponse {
	var responses []*response.RefreshTokenResponse

	for _, response := range refreshs {
		responses = append(responses, r.ToRefreshTokenResponse(response))
	}

	return responses
}
