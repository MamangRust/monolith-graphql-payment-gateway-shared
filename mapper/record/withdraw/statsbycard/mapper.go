package withdrawstatsbycardrecordmapper

type WithdrawStatisticByCardRecordMapper interface {
	WithdrawStatisticAmountByCardRecordMapper
	WithdrawStatisticStatusByCardRecordMapper
}

// withdrawStatisticRecordMapper is an interface that provides methods to map
// database rows related to withdraw statistics (amount & status) into domain models.
type withdrawStatisticByCardRecordMapper struct {
	WithdrawStatisticAmountByCardRecordMapper
	WithdrawStatisticStatusByCardRecordMapper
}

// NewWithdrawStatisticRecordMapper creates and returns a new instance of
// withdrawStatisticRecordMapper which provides methods to map database rows
// related to withdraw statistics (amount & status) filtered by card number
// into domain models.
func NewWithdrawStatisticByCardRecordMapper() WithdrawStatisticByCardRecordMapper {
	return &withdrawStatisticByCardRecordMapper{
		WithdrawStatisticAmountByCardRecordMapper: NewWithdrawStatisticAmountByCardRecordMapper(),
		WithdrawStatisticStatusByCardRecordMapper: NewWithdrawStatisticStatusByCardRecordMapper(),
	}
}
