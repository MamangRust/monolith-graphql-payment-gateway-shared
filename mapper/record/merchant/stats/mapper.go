package merchantstatsrecordmapper

type MerchantStatisticRecordMapper interface {
	MerchantStatisticAmountRecordMapper
	MerchantStatisticMethodRecordMapper
	MerchantStatisticTotalAmountRecordMapper
}

// merchantStatisticRecordMapper maps database rows to domain models for statistic operations.
type merchantStatisticRecordMapper struct {
	MerchantStatisticTotalAmountRecordMapper
	MerchantStatisticMethodRecordMapper
	MerchantStatisticAmountRecordMapper
}

// NewMerchantStatisticRecordMapper creates a new instance of
// MerchantStatisticRecordMapper, providing methods to map database rows
// to MerchantMonthlyPaymentMethod, MerchantMonthlyAmount, and
// MerchantMonthlyTotalAmount domain models.
func NewMerchantStatisticRecordMapper() MerchantStatisticRecordMapper {
	return &merchantStatisticRecordMapper{
		MerchantStatisticTotalAmountRecordMapper: NewMerchantStatisticTotalAmountRecordMapper(),
		MerchantStatisticMethodRecordMapper:      NewMerchantStatisticMethodRecordMapper(),
		MerchantStatisticAmountRecordMapper:      NewMerchantStatisticAmountRecordMapper(),
	}
}
