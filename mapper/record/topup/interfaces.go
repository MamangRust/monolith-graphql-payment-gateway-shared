package topuprecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TopupBaseRecordMapper provides methods to map Topup database rows to TopupRecord domain models.
type TopupBaseRecordMapper interface {
	// ToTopupRecord maps a Topup database row to a TopupRecord domain model.
	// It takes a pointer to a db.Topup as a parameter and returns a pointer to a record.TopupRecord.
	// The function maps the fields of the db.Topup to the corresponding fields of the record.TopupRecord.
	// If the DeletedAt field of the db.Topup is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
	ToTopupRecord(topup *db.Topup) *record.TopupRecord
}

// TopupCommandRecordMapping maps raw topup rows for command operations.
type TopupCommandRecordMapping interface {
	TopupBaseRecordMapper
}

// TopupQueryRecordMapping maps query result rows (non-statistical) to Topup domain models.
type TopupQueryRecordMapping interface {
	TopupBaseRecordMapper

	// ToTopupByCardNumberRecord converts a db.GetTopupsByCardNumberRow to a record.TopupRecord.
	// It takes a pointer to a db.GetTopupsByCardNumberRow as a parameter and returns a pointer to a record.TopupRecord.
	// The function maps the fields of the db.GetTopupsByCardNumberRow to the corresponding fields of the record.TopupRecord.
	// If the DeletedAt field of the db.GetTopupsByCardNumberRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
	ToTopupByCardNumberRecord(topup *db.GetTopupsByCardNumberRow) *record.TopupRecord
	// ToTopupByCardNumberRecords converts a slice of db.GetTopupsByCardNumberRow to a slice of record.TopupRecord.
	// It takes a slice of pointers to db.GetTopupsByCardNumberRow as a parameter and returns a slice of pointers to record.TopupRecord.
	// The function iterates over the provided slice of db.GetTopupsByCardNumberRow, converting each element
	// using the ToTopupByCardNumberRecord method and appending the result to a new slice.
	// The function returns a slice of pointers to record.TopupRecord.
	ToTopupByCardNumberRecords(topups []*db.GetTopupsByCardNumberRow) []*record.TopupRecord

	// ToTopupRecordAll converts a db.GetTopupsRow to a record.TopupRecord.
	// It takes a pointer to a db.GetTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
	// The function maps the fields of the db.GetTopupsRow to the corresponding fields of the record.TopupRecord.
	// If the DeletedAt field of the db.GetTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
	ToTopupRecordAll(topup *db.GetTopupsRow) *record.TopupRecord
	// ToTopupRecordsAll converts a slice of db.GetTopupsRow to a slice of record.TopupRecord.
	// It iterates over the provided slice of db.GetTopupsRow, converting each element
	// using the ToTopupRecordAll method and appending the result to a new slice.
	// The function returns a slice of pointers to record.TopupRecord.
	ToTopupRecordsAll(topups []*db.GetTopupsRow) []*record.TopupRecord

	// ToTopupRecordActive converts a db.GetActiveTopupsRow to a record.TopupRecord.
	// It takes a pointer to a db.GetActiveTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
	// The function maps the fields of the db.GetActiveTopupsRow to the corresponding fields of the record.TopupRecord.
	// If the DeletedAt field of the db.GetActiveTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
	ToTopupRecordActive(topup *db.GetActiveTopupsRow) *record.TopupRecord
	// ToTopupRecordsActive maps a slice of GetActiveTopupsRow database rows to a slice of TopupRecord domain models.
	// It iterates over the provided slice of GetActiveTopupsRow, converting each element
	// using the ToTopupRecordActive method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecord.
	ToTopupRecordsActive(topups []*db.GetActiveTopupsRow) []*record.TopupRecord

	// ToTopupRecordTrashed converts a db.GetTrashedTopupsRow to a record.TopupRecord.
	// It takes a pointer to a db.GetTrashedTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
	// The function maps the fields of the db.GetTrashedTopupsRow to the corresponding fields of the record.TopupRecord.
	// If the DeletedAt field of the db.GetTrashedTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
	ToTopupRecordTrashed(topup *db.GetTrashedTopupsRow) *record.TopupRecord
	// ToTopupRecordsTrashed maps a slice of GetTrashedTopupsRow database rows to a slice of TopupRecord domain models.
	// It iterates over the provided slice of GetTrashedTopupsRow, converting each element
	// using the ToTopupRecordTrashed method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecord.
	ToTopupRecordsTrashed(topups []*db.GetTrashedTopupsRow) []*record.TopupRecord
}
