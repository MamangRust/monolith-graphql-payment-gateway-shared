package withdrawrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// WithdrawBaseRecordMapping maps raw withdraw entities (insert/update/delete) to domain records.
type WithdrawBaseRecordMapping interface {
	ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord
}

// WithdrawCommandRecordMapping maps raw withdraw entities (insert/update/delete) to domain records.
type WithdrawCommandRecordMapping interface {
	WithdrawBaseRecordMapping
}

// WithdrawQueryRecordMapping maps query results for listing, active, trashed, and by-card data.
type WithdrawQueryRecordMapping interface {
	WithdrawBaseRecordMapping

	ToWithdrawByCardNumberRecord(withdraw *db.GetWithdrawsByCardNumberRow) *record.WithdrawRecord
	ToWithdrawsByCardNumberRecord(withdraws []*db.GetWithdrawsByCardNumberRow) []*record.WithdrawRecord

	ToWithdrawRecordAll(withdraw *db.GetWithdrawsRow) *record.WithdrawRecord
	ToWithdrawsRecordAll(withdraws []*db.GetWithdrawsRow) []*record.WithdrawRecord

	ToWithdrawRecordActive(withdraw *db.GetActiveWithdrawsRow) *record.WithdrawRecord
	ToWithdrawsRecordActive(withdraws []*db.GetActiveWithdrawsRow) []*record.WithdrawRecord

	ToWithdrawRecordTrashed(withdraw *db.GetTrashedWithdrawsRow) *record.WithdrawRecord
	ToWithdrawsRecordTrashed(withdraws []*db.GetTrashedWithdrawsRow) []*record.WithdrawRecord
}
