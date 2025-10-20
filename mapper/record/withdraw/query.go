package withdrawrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// withdrawQueryRecordMapper provides methods to map Withdraw database rows to WithdrawRecord domain models.
type withdrawQueryRecordMapper struct {
}

// NewWithdrawQueryRecordMapper creates a new instance of withdrawQueryRecordMapper.
// It returns a pointer to a withdrawQueryRecordMapper, which provides methods
// for mapping Withdraw database rows to WithdrawRecord domain models.
func NewWithdrawQueryRecordMapper() WithdrawQueryRecordMapping {
	return &withdrawQueryRecordMapper{}
}

// ToWithdrawRecord maps a Withdraw database row to a WithdrawRecord domain model.
//
// Args:
//   - withdraw: A pointer to a Withdraw representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawQueryRecordMapper) ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord {
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

// ToWithdrawByCardNumberRecord maps a database row representing a withdraw record
// associated with a given card number to a WithdrawRecord domain model.
//
// Args:
//   - withdraw: A pointer to a GetWithdrawsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawQueryRecordMapper) ToWithdrawByCardNumberRecord(withdraw *db.GetWithdrawsByCardNumberRow) *record.WithdrawRecord {
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

// ToWithdrawsByCardNumberRecord maps a slice of GetWithdrawsByCardNumberRow database rows to a slice of WithdrawRecord domain models.
//
// Args:
//   - withdraws: A slice of pointers to GetWithdrawsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawQueryRecordMapper) ToWithdrawsByCardNumberRecord(withdraws []*db.GetWithdrawsByCardNumberRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawByCardNumberRecord(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawRecordAll maps a GetWithdrawsRow database row to a WithdrawRecord domain model.
//
// Args:
//   - withdraw: A pointer to a GetWithdrawsRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and optionally DeletedAt if it is valid.
func (s *withdrawQueryRecordMapper) ToWithdrawRecordAll(withdraw *db.GetWithdrawsRow) *record.WithdrawRecord {
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

// ToWithdrawsRecordALl maps a slice of GetWithdrawsRow database rows to a slice of WithdrawRecord domain models.
//
// Args:
//   - withdraws: A slice of pointers to GetWithdrawsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and optionally DeletedAt if it is valid.
func (s *withdrawQueryRecordMapper) ToWithdrawsRecordAll(withdraws []*db.GetWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordAll(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawRecordActive maps a single GetActiveWithdrawsRow database row to a WithdrawRecord
// domain model.
//
// Args:
//   - withdraw: A pointer to a GetActiveWithdrawsRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and
//     optionally DeletedAt if it is valid.
func (s *withdrawQueryRecordMapper) ToWithdrawRecordActive(withdraw *db.GetActiveWithdrawsRow) *record.WithdrawRecord {
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

// ToWithdrawsRecordActive maps a slice of GetActiveWithdrawsRow database rows to a slice of WithdrawRecord domain models.
//
// Args:
//   - withdraws: A slice of pointers to GetActiveWithdrawsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and
//     optionally DeletedAt if it is valid.
func (s *withdrawQueryRecordMapper) ToWithdrawsRecordActive(withdraws []*db.GetActiveWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordActive(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawRecordTrashed maps a GetTrashedWithdrawsRow database row to a WithdrawRecord domain model.
// It is intended for use with database rows that contain trashed withdraw records.
// It returns a pointer to a WithdrawRecord containing the mapped data, including ID, WithdrawNo,
// CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawQueryRecordMapper) ToWithdrawRecordTrashed(withdraw *db.GetTrashedWithdrawsRow) *record.WithdrawRecord {
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

// ToWithdrawsRecordTrashed maps a slice of GetTrashedWithdrawsRow database rows to a slice of WithdrawRecord
// domain models. It is intended for use with database rows that contain trashed withdraw records.
//
// Args:
//   - withdraws: A slice of pointers to GetTrashedWithdrawsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawQueryRecordMapper) ToWithdrawsRecordTrashed(withdraws []*db.GetTrashedWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordTrashed(withdraw))
	}

	return withdrawRecords
}
