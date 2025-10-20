package cardstatsrecordmapper

type CardStatsRecordMapper interface {
	CardStatisticBalanceRecordMapper
	CardStatisticTopupRecordMapper
	CardStatisticTransactionRecordMapper
	CardStatisticTransferRecordMapper
	CardStatisticWithdrawRecordMapper
}

// cardStatisticRecordMapper provides methods to map database rows to CardStatisticRecord domain models.
type cardStatisticRecordMapper struct {
	CardStatisticBalanceRecordMapper
	CardStatisticTopupRecordMapper
	CardStatisticTransactionRecordMapper
	CardStatisticTransferRecordMapper
	CardStatisticWithdrawRecordMapper
}

// NewCardStatisticRecordMapper returns a new instance of cardStatisticRecordMapper, which provides methods to map database rows to CardStatisticRecord domain models.
func NewCardStatisticRecordMapper() CardStatsRecordMapper {
	return &cardStatisticRecordMapper{
		CardStatisticBalanceRecordMapper:     NewCardStatisticBalanceRecordMapper(),
		CardStatisticTopupRecordMapper:       NewCardStatisticTopupRecordMapper(),
		CardStatisticTransactionRecordMapper: NewCardStatisticTransactionRecordMapper(),
		CardStatisticTransferRecordMapper:    NewCardStatisticTransferRecordMapper(),
		CardStatisticWithdrawRecordMapper:    NewCardStatisticWithdrawRecordMapper(),
	}
}
