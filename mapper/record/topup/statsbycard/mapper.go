package topupstatsbycardrecordmapper

type TopupStatisticByCardRecordMapper interface {
	TopupStatisticStatusByCardNumberMapper
	TopupStatisticMethodByCardNumberMapper
	TopupStatisticAmountByCardNumberMapper
}

type topStatisticByCardRecordMapper struct {
	TopupStatisticStatusByCardNumberMapper
	TopupStatisticMethodByCardNumberMapper
	TopupStatisticAmountByCardNumberMapper
}

// NewTopupStatisticByCardRecordMapper returns a new instance of topStatisticByCardRecordMapper which provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
func NewTopupStatisticByCardRecordMapper() TopupStatisticByCardRecordMapper {
	return &topStatisticByCardRecordMapper{
		TopupStatisticStatusByCardNumberMapper: NewTopupStatisticStatusByCardNumberMapper(),
		TopupStatisticMethodByCardNumberMapper: NewTopupStatisticMethodByCardNumberMapper(),
		TopupStatisticAmountByCardNumberMapper: NewTopupStatisticAmountByCardNumberMapper(),
	}
}
