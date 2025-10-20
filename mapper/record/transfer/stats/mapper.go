package transferstatsrecordmapper

type TransferStatisticRecordMapper interface {
	TransferStatisticAmountRecordMapper
	TransferStatisticStatusRecordMapper
}

type transferStatisticRecordMapper struct {
	TransferStatisticAmountRecordMapper
	TransferStatisticStatusRecordMapper
}

// NewTransferStatisticRecordMapper returns a new instance of transferStatisticRecordMapper
// which provides methods to map database rows to TransferStatisticAmount and
// TransferStatisticStatus domain models.
func NewTransferStatisticRecordMapper() TransferStatisticRecordMapper {
	return &transferStatisticRecordMapper{
		TransferStatisticAmountRecordMapper: NewTransferStatisticAmountRecordMapper(),
		TransferStatisticStatusRecordMapper: NewTransferStatisticStatusRecordMapper(),
	}
}
