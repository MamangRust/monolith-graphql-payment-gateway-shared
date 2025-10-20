package refreshtokenrecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// RefreshTokenRecordMapping defines a mapping function from a RefreshToken database row to a RefreshTokenRecord domain model.
type RefreshTokenRecordMapping interface {
	// ToRefreshTokenRecord maps a RefreshToken database row to a RefreshTokenRecord domain model.
	// Parameters:
	//   - refreshToken: A pointer to a RefreshToken representing the database row.
	//
	// Returns:
	//   - A pointer to a RefreshTokenRecord containing the mapped data, including
	//     ID, UserID, Token, and ExpiredAt, CreatedAt, and UpdatedAt.
	ToRefreshTokenRecord(refreshToken *db.RefreshToken) *record.RefreshTokenRecord

	// ToRefreshTokensRecord maps a slice of RefreshToken database rows to a slice of RefreshTokenRecord
	// domain models.
	//
	// Parameters:
	//   - refreshTokens: A slice of pointers to RefreshToken representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to RefreshTokenRecord containing the mapped data, including
	//     ID, UserID, Token, and ExpiredAt, CreatedAt, and UpdatedAt.
	ToRefreshTokensRecord(refreshTokens []*db.RefreshToken) []*record.RefreshTokenRecord
}
