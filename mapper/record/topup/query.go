package topuprecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// topupQueryRecordMapper provides methods to map Topup database rows to TopupRecord domain models for query operations.
type topupQueryRecordMapper struct {
}

func NewTopupQueryRecordMapper() TopupQueryRecordMapping {
	return &topupQueryRecordMapper{}
}

// ToTopupRecord maps a Topup database row to a TopupRecord domain model.
// It takes a pointer to a db.Topup as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.Topup to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.Topup is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupQueryRecordMapper) ToTopupRecord(topup *db.Topup) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupByCardNumberRecord converts a db.GetTopupsByCardNumberRow to a record.TopupRecord.
// It takes a pointer to a db.GetTopupsByCardNumberRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetTopupsByCardNumberRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetTopupsByCardNumberRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupQueryRecordMapper) ToTopupByCardNumberRecord(topup *db.GetTopupsByCardNumberRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupByCardNumberRecords converts a slice of db.GetTopupsByCardNumberRow to a slice of record.TopupRecord.
// It takes a slice of pointers to db.GetTopupsByCardNumberRow as a parameter and returns a slice of pointers to record.TopupRecord.
// The function iterates over the provided slice of db.GetTopupsByCardNumberRow, converting each element
// using the ToTopupByCardNumberRecord method and appending the result to a new slice.
// The function returns a slice of pointers to record.TopupRecord.
func (t *topupQueryRecordMapper) ToTopupByCardNumberRecords(topups []*db.GetTopupsByCardNumberRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupByCardNumberRecord(topup))
	}

	return topupRecords
}

// ToTopupRecordAll converts a db.GetTopupsRow to a record.TopupRecord.
// It takes a pointer to a db.GetTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetTopupsRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupQueryRecordMapper) ToTopupRecordAll(topup *db.GetTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupRecordsAll converts a slice of db.GetTopupsRow to a slice of record.TopupRecord.
// It iterates over the provided slice of db.GetTopupsRow, converting each element
// using the ToTopupRecordAll method and appending the result to a new slice.
// The function returns a slice of pointers to record.TopupRecord.
func (t *topupQueryRecordMapper) ToTopupRecordsAll(topups []*db.GetTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordAll(topup))
	}

	return topupRecords
}

// ToTopupRecordActive converts a db.GetActiveTopupsRow to a record.TopupRecord.
// It takes a pointer to a db.GetActiveTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetActiveTopupsRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetActiveTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupQueryRecordMapper) ToTopupRecordActive(topup *db.GetActiveTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupRecordsActive maps a slice of GetActiveTopupsRow database rows to a slice of TopupRecord domain models.
// It iterates over the provided slice of GetActiveTopupsRow, converting each element
// using the ToTopupRecordActive method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecord.
func (t *topupQueryRecordMapper) ToTopupRecordsActive(topups []*db.GetActiveTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordActive(topup))
	}

	return topupRecords
}

// ToTopupRecordTrashed converts a db.GetTrashedTopupsRow to a record.TopupRecord.
// It takes a pointer to a db.GetTrashedTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetTrashedTopupsRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetTrashedTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupQueryRecordMapper) ToTopupRecordTrashed(topup *db.GetTrashedTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupRecordsTrashed maps a slice of GetTrashedTopupsRow database rows to a slice of TopupRecord domain models.
// It iterates over the provided slice of GetTrashedTopupsRow, converting each element
// using the ToTopupRecordTrashed method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecord.
func (t *topupQueryRecordMapper) ToTopupRecordsTrashed(topups []*db.GetTrashedTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordTrashed(topup))
	}

	return topupRecords
}
