package resettokenrecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// resetTokenRecordMapper provides methods to map ResetToken database rows to ResetTokenRecord domain models.
type resetTokenRecordMapper struct {
}

// NewResetTokenRecordMapper returns a new instance of resetTokenRecordMapper which provides methods to map ResetToken database rows to ResetTokenRecord domain models.
func NewResetTokenRecordMapper() ResetTokenRecordMapping {
	return &resetTokenRecordMapper{}
}

// ToResetTokenRecord maps a ResetToken database row to a ResetTokenRecord domain model.
//
// Parameters:
//   - resetToken: A pointer to a ResetToken representing the database row.
//
// Returns:
//   - A pointer to a ResetTokenRecord containing the mapped data, including ID, UserID, Token, and ExpiredAt.
func (r *resetTokenRecordMapper) ToResetTokenRecord(resetToken *db.ResetToken) *record.ResetTokenRecord {
	return &record.ResetTokenRecord{
		ID:        int64(resetToken.ID),
		UserID:    resetToken.UserID,
		Token:     resetToken.Token,
		ExpiredAt: resetToken.ExpiryDate.String(),
	}
}
