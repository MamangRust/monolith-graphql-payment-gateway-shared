package withdrawrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// WithdrawCommandRecordMapping is an interface that provides methods to map Withdraw database rows to WithdrawRecord domain models for command operations.
type withdrawCommandRecordMapper struct {
}

// NewWithdrawCommandRecordMapper returns a new instance of withdrawCommandRecordMapper,
// initializing it with methods to map Withdraw database rows to WithdrawRecord domain models
// for command operations.
func NewWithdrawCommandRecordMapper() WithdrawCommandRecordMapping {
	return &withdrawCommandRecordMapper{}
}

// ToWithdrawRecord maps a Withdraw database row to a WithdrawRecord domain model.
//
// Args:
//   - withdraw: A pointer to a Withdraw representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawCommandRecordMapper) ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord {
	var deletedAt *string

	if withdraw.DeletedAt.Valid {
		formatedDeletedAt := withdraw.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.WithdrawRecord{
		ID:             int(withdraw.WithdrawID),
		WithdrawNo:     withdraw.WithdrawNo.String(),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime.String(),
		CreatedAt:      withdraw.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      withdraw.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}
