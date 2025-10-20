package refreshtokenrecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// refreshTokenRecordMapper provides methods to map RefreshToken database rows to RefreshTokenRecord domain models.
type refreshTokenRecordMapper struct {
}

// NewRefreshTokenRecordMapper returns a new instance of refreshTokenRecordMapper which provides methods to map RefreshToken database rows to RefreshTokenRecord domain models.
func NewRefreshTokenRecordMapper() RefreshTokenRecordMapping {
	return &refreshTokenRecordMapper{}
}

// ToRefreshTokenRecord maps a RefreshToken database row to a RefreshTokenRecord domain model.
// Parameters:
//   - refreshToken: A pointer to a RefreshToken representing the database row.
//
// Returns:
//   - A pointer to a RefreshTokenRecord containing the mapped data, including
//     ID, UserID, Token, and ExpiredAt, CreatedAt, and UpdatedAt.
func (r *refreshTokenRecordMapper) ToRefreshTokenRecord(refreshToken *db.RefreshToken) *record.RefreshTokenRecord {
	return &record.RefreshTokenRecord{
		ID:        int(refreshToken.RefreshTokenID),
		UserID:    int(refreshToken.UserID),
		Token:     refreshToken.Token,
		ExpiredAt: refreshToken.Expiration.String(),
		CreatedAt: refreshToken.CreatedAt.Time.String(),
		UpdatedAt: refreshToken.UpdatedAt.Time.String(),
	}
}

// ToRefreshTokensRecord maps a slice of RefreshToken database rows to a slice of RefreshTokenRecord
// domain models.
//
// Parameters:
//   - refreshTokens: A slice of pointers to RefreshToken representing the database rows.
//
// Returns:
//   - A slice of pointers to RefreshTokenRecord containing the mapped data, including
//     ID, UserID, Token, and ExpiredAt, CreatedAt, and UpdatedAt.
func (r *refreshTokenRecordMapper) ToRefreshTokensRecord(refreshTokens []*db.RefreshToken) []*record.RefreshTokenRecord {
	var refreshTokenRecords []*record.RefreshTokenRecord
	for _, refreshToken := range refreshTokens {
		refreshTokenRecords = append(refreshTokenRecords, r.ToRefreshTokenRecord(refreshToken))
	}
	return refreshTokenRecords
}
