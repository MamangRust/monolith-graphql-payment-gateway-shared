package topuprecordmapper

import (
	topupstatsrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/topup/stats"
	topupstatsbycardrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/topup/statsbycard"
)

type TopupRecordMapper interface {
	QueryMapper() TopupQueryRecordMapping
	CommandMapper() TopupCommandRecordMapping
	StatsMapper() topupstatsrecordmapper.TopupStatisticRecordMapper
	StatsByCardMapper() topupstatsbycardrecordmapper.TopupStatisticByCardRecordMapper
}

type topupRecordMapper struct {
	queryMapper       TopupQueryRecordMapping
	commandMapper     TopupCommandRecordMapping
	statsMapper       topupstatsrecordmapper.TopupStatisticRecordMapper
	statsByCardMapper topupstatsbycardrecordmapper.TopupStatisticByCardRecordMapper
}

func NewTopupRecordMapper() TopupRecordMapper {
	return &topupRecordMapper{
		queryMapper:       NewTopupQueryRecordMapper(),
		commandMapper:     NewTopupCommandRecordMapper(),
		statsMapper:       topupstatsrecordmapper.NewTopupStatisticRecordMapper(),
		statsByCardMapper: topupstatsbycardrecordmapper.NewTopupStatisticByCardRecordMapper(),
	}
}

func (t *topupRecordMapper) QueryMapper() TopupQueryRecordMapping {
	return t.queryMapper
}

func (t *topupRecordMapper) CommandMapper() TopupCommandRecordMapping {
	return t.commandMapper
}

func (t *topupRecordMapper) StatsMapper() topupstatsrecordmapper.TopupStatisticRecordMapper {
	return t.statsMapper
}

func (t *topupRecordMapper) StatsByCardMapper() topupstatsbycardrecordmapper.TopupStatisticByCardRecordMapper {
	return t.statsByCardMapper
}
