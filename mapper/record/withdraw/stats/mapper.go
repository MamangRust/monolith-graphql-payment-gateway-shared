package withdrawstatsrecordmapper

type WithdrawStatisticRecordMapper interface {
	WithdrawStatisticAmountRecordMapper
	WithdrawStatisticStatusRecordMapper
}

// withdrawStatisticRecordMapper is an interface that provides methods to map
// database rows related to withdraw statistics (amount & status) into domain models.
type withdrawStatisticRecordMapper struct {
	WithdrawStatisticAmountRecordMapper
	WithdrawStatisticStatusRecordMapper
}

// NewWithdrawStatisticRecordMapper creates and returns a new instance of
// withdrawStatisticRecordMapper which provides methods to map database rows
// related to withdraw statistics (amount & status) into domain models.
func NewWithdrawStatisticRecordMapper() WithdrawStatisticRecordMapper {
	return &withdrawStatisticRecordMapper{
		WithdrawStatisticAmountRecordMapper: NewWithdrawStatisticAmountRecordMapper(),
		WithdrawStatisticStatusRecordMapper: NewWithdrawStatisticStatusRecordMapper(),
	}
}
