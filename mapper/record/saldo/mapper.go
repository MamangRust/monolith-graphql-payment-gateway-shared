package saldorecordmapper

type saldoRecordMapper struct {
	queryMapper     SaldoQueryRecordMapping
	commandMapper   SaldoCommandRecordMapping
	statisticMapper SaldoStatisticRecordMapping
}

type SaldoRecordMapper interface {
	QueryMapper() SaldoQueryRecordMapping
	CommandMapper() SaldoCommandRecordMapping
	StatisticMapper() SaldoStatisticRecordMapping
}

func NewSaldoRecordMapper() SaldoRecordMapper {
	return &saldoRecordMapper{
		queryMapper:     NewSaldoQueryRecordMapper(),
		commandMapper:   NewSaldoCommandRecordMapper(),
		statisticMapper: NewSaldoStatisticRecordMapper(),
	}
}

func (s *saldoRecordMapper) QueryMapper() SaldoQueryRecordMapping {
	return s.queryMapper
}

func (s *saldoRecordMapper) CommandMapper() SaldoCommandRecordMapping {
	return s.commandMapper
}

func (s *saldoRecordMapper) StatisticMapper() SaldoStatisticRecordMapping {
	return s.statisticMapper
}
