package cardrecordmapper

import (
	cardstatsrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/card/stats"
	cardstatsbycardrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/card/statsbycard"
)

type cardRecordMapper struct {
	queryMapper   CardQueryRecordMapper
	commandMapper CardCommandRecordMapper

	statsMapper  cardstatsrecordmapper.CardStatsRecordMapper
	byCardMapper cardstatsbycardrecordmapper.CardStatsByCardRecordMapper
}

type CardRecordMapper interface {
	QueryMapper() CardQueryRecordMapper
	CommandMapper() CardCommandRecordMapper
	StatsMapper() cardstatsrecordmapper.CardStatsRecordMapper
	StatsByCardMapper() cardstatsbycardrecordmapper.CardStatsByCardRecordMapper
}

func NewCardRecordMapper() CardRecordMapper {
	return &cardRecordMapper{
		queryMapper:   NewCardQueryRecordMapper(),
		commandMapper: NewCardCommandRecordMapper(),
		statsMapper:   cardstatsrecordmapper.NewCardStatisticRecordMapper(),
		byCardMapper:  cardstatsbycardrecordmapper.NewCardStatisticByCardRecordMapper(),
	}
}

func (c *cardRecordMapper) QueryMapper() CardQueryRecordMapper {
	return c.queryMapper
}

func (c *cardRecordMapper) CommandMapper() CardCommandRecordMapper {
	return c.commandMapper
}

func (c *cardRecordMapper) StatsMapper() cardstatsrecordmapper.CardStatsRecordMapper {
	return c.statsMapper
}

func (c *cardRecordMapper) StatsByCardMapper() cardstatsbycardrecordmapper.CardStatsByCardRecordMapper {
	return c.byCardMapper
}
