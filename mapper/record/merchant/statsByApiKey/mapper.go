package merchantstatbyapikeyrecordmapper

type MerchantStatisticByApiKeyMapper interface {
	MerchantStatisticAmountByApiKeyMapper
	MerchantStatisticMethodByApiKeyMapper
	MerchantStatisticTotalAmountByApiKeyMapper
}

// MerchantStatisticByApiKeyMapper is an interface that provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticByApiKeyMapper struct {
	MerchantStatisticAmountByApiKeyMapper
	MerchantStatisticMethodByApiKeyMapper
	MerchantStatisticTotalAmountByApiKeyMapper
}

// NewMerchantStatisticByApiKeyMapper creates a new instance of
// MerchantStatisticByApiKeyMapper, providing methods to map database rows
// to MerchantMonthlyPaymentMethod, MerchantMonthlyAmount, and
// MerchantMonthlyTotalAmount domain models filtered by API key.
func NewMerchantStatisticByApiKeyMapper() MerchantStatisticByApiKeyMapper {
	return &merchantStatisticByApiKeyMapper{
		MerchantStatisticMethodByApiKeyMapper:      NewMerchantStatisticMethodByApiKeyMapper(),
		MerchantStatisticAmountByApiKeyMapper:      NewMerchantStatisticAmountByApiKeyMapper(),
		MerchantStatisticTotalAmountByApiKeyMapper: NewMerchantStatisticTotalAmountByApiKeyMapper(),
	}
}
