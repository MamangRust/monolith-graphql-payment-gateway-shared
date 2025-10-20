package topupstatsrecordmapper

type TopupStatisticRecordMapper interface {
	TopupStatisticStatusMapper
	TopupStatisticMethodMapper
	TopupStatisticAmountMapper
}

// topupStatisticRecordMapper provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
type topupStatisticRecordMapper struct {
	TopupStatisticStatusMapper
	TopupStatisticMethodMapper
	TopupStatisticAmountMapper
}

// NewTopupStatisticRecordMapper returns a new instance of topupStatisticRecordMapper which provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
func NewTopupStatisticRecordMapper() TopupStatisticRecordMapper {
	return &topupStatisticRecordMapper{
		TopupStatisticStatusMapper: NewTopupStatisticStatusRecordMapper(),
		TopupStatisticMethodMapper: NewTopupStatisticMethodRecordMapper(),
		TopupStatisticAmountMapper: NewTopupStatisticAmountRecordMapper(),
	}
}
