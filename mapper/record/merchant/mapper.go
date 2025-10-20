package merchantrecordmapper

import (
	merchantstatsrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/merchant/stats"
	merchantstatbyapikeyrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/merchant/statsByApiKey"
	merchantstatbyidrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/merchant/statsByMerchant"
)

type MerchantRecordMapper interface {
	QueryMapper() MerchantQueryRecordMapper
	CommandMapper() MerchantCommandRecordMapper
	TransactionMapper() MerchantTransactionRecordMapper
	StatisticMapper() merchantstatsrecordmapper.MerchantStatisticRecordMapper
	ByMerchantMapper() merchantstatbyidrecordmapper.MerchantStatisticByMerchantMapper
	ByApiKeyMapper() merchantstatbyapikeyrecordmapper.MerchantStatisticByApiKeyMapper
}

type merchantRecordMapper struct {
	queryMapper       MerchantQueryRecordMapper
	commandMapper     MerchantCommandRecordMapper
	transactionMapper MerchantTransactionRecordMapper
	statsMapper       merchantstatsrecordmapper.MerchantStatisticRecordMapper
	byIdMapper        merchantstatbyidrecordmapper.MerchantStatisticByMerchantMapper
	byApiKeyMapper    merchantstatbyapikeyrecordmapper.MerchantStatisticByApiKeyMapper
}

func NewMerchantRecordMapper() MerchantRecordMapper {
	return &merchantRecordMapper{
		queryMapper:       NewMerchantQueryRecordMapper(),
		commandMapper:     NewMerchantCommandMapper(),
		transactionMapper: NewMerchantTransactionRecordMapper(),
		statsMapper:       merchantstatsrecordmapper.NewMerchantStatisticRecordMapper(),
		byIdMapper:        merchantstatbyidrecordmapper.NewMerchantStatisticByMerchantRecordMapper(),
		byApiKeyMapper:    merchantstatbyapikeyrecordmapper.NewMerchantStatisticByApiKeyMapper(),
	}
}

func (m *merchantRecordMapper) QueryMapper() MerchantQueryRecordMapper {
	return m.queryMapper
}

func (m *merchantRecordMapper) CommandMapper() MerchantCommandRecordMapper {
	return m.commandMapper
}

func (m *merchantRecordMapper) TransactionMapper() MerchantTransactionRecordMapper {
	return m.transactionMapper
}

func (m *merchantRecordMapper) StatisticMapper() merchantstatsrecordmapper.MerchantStatisticRecordMapper {
	return m.statsMapper
}

func (m *merchantRecordMapper) ByMerchantMapper() merchantstatbyidrecordmapper.MerchantStatisticByMerchantMapper {
	return m.byIdMapper
}

func (m *merchantRecordMapper) ByApiKeyMapper() merchantstatbyapikeyrecordmapper.MerchantStatisticByApiKeyMapper {
	return m.byApiKeyMapper
}
