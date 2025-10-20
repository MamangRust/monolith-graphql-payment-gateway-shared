package merchantstatbyidrecordmapper

type MerchantStatisticByMerchantMapper interface {
	MerchantStatisticMethodByMerchantMapper
	MerchantStatisticAmountByMerchantMapper
	MerchantStatisticTotalAmountByMerchantMapper
}

// merchantStatisticByMerchantRecordMapper provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticByMerchantRecordMapper struct {
	MerchantStatisticMethodByMerchantMapper
	MerchantStatisticAmountByMerchantMapper
	MerchantStatisticTotalAmountByMerchantMapper
}

// NewMerchantStatisticByMerchantRecordMapper returns a new instance of
// MerchantStatisticByMerchantMapper, which provides methods to map Merchant
// database rows to MerchantRecord domain models for statistic operations.
func NewMerchantStatisticByMerchantRecordMapper() MerchantStatisticByMerchantMapper {
	return &merchantStatisticByMerchantRecordMapper{
		MerchantStatisticMethodByMerchantMapper:      NewMerchantStatisticMethodByMerchantRecordMapper(),
		MerchantStatisticAmountByMerchantMapper:      NewMerchantStatisticAmountByMerchantRecordMapper(),
		MerchantStatisticTotalAmountByMerchantMapper: NewMerchantStatisticTotalAmountByMerchantRecordMapper(),
	}
}
