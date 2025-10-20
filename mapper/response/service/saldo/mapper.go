package saldoservicemapper

type SaldoResponseMapper interface {
	QueryMapper() SaldoQueryResponseMapper
	CommandMapper() SaldoCommandResponseMapper
	StatisticBalanceMapper() SaldoStatisticBalanceResponseMapper
	TotalBalanceMapper() SaldoStatisticTotalBalanceResponseMapper
}

type saldoResponseMapper struct {
	queryMapper           SaldoQueryResponseMapper
	commandMapper         SaldoCommandResponseMapper
	statisticBalance      SaldoStatisticBalanceResponseMapper
	statisticTotalBalance SaldoStatisticTotalBalanceResponseMapper
}

func NewSaldoResponseMapper() SaldoResponseMapper {
	return &saldoResponseMapper{
		queryMapper:           NewSaldoQueryResponseMapper(),
		commandMapper:         NewSaldoCommandResponseMapper(),
		statisticBalance:      NewSaldoStatsBalanceResponseMapper(),
		statisticTotalBalance: NewSaldoTotalBalanceResponseMapper(),
	}
}

func (s *saldoResponseMapper) QueryMapper() SaldoQueryResponseMapper {
	return s.queryMapper
}

func (s *saldoResponseMapper) CommandMapper() SaldoCommandResponseMapper {
	return s.commandMapper
}

func (s *saldoResponseMapper) StatisticBalanceMapper() SaldoStatisticBalanceResponseMapper {
	return s.statisticBalance
}

func (s *saldoResponseMapper) TotalBalanceMapper() SaldoStatisticTotalBalanceResponseMapper {
	return s.statisticTotalBalance
}
