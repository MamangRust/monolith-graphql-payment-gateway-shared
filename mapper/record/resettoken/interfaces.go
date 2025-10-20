package resettokenrecord

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// ResetTokenRecordMapping defines a mapping function from a ResetToken database row to a ResetTokenRecord domain model.
type ResetTokenRecordMapping interface {
	// ToResetTokenRecord maps a ResetToken database row to a ResetTokenRecord domain model.
	//
	// Parameters:
	//   - resetToken: A pointer to a ResetToken representing the database row.
	//
	// Returns:
	//   - A pointer to a ResetTokenRecord containing the mapped data, including ID, UserID, Token, and ExpiredAt.
	ToResetTokenRecord(resetToken *db.ResetToken) *record.ResetTokenRecord
}
