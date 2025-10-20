package topupservicemapper

type TopupResponseMapper interface {
	QueryMapper() TopupQueryResponseMapper
	CommandMapper() TopupCommandResponseMapper
	StatusStatsMapper() TopupStatsStatusResponseMapper
	AmountStatsMapper() TopupStatsAmountResponseMapper
	MethodStatsMapper() TopupStatsMethodResponseMapper
}

type topupResponseMapper struct {
	queryMapper       TopupQueryResponseMapper
	commandMapper     TopupCommandResponseMapper
	statusStatsMapper TopupStatsStatusResponseMapper
	amountStatsMapper TopupStatsAmountResponseMapper
	methodStatsMapper TopupStatsMethodResponseMapper
}

func NewTopupResponseMapper() TopupResponseMapper {
	return &topupResponseMapper{
		queryMapper:       NewTopupQueryResponseMapper(),
		commandMapper:     NewTopupCommandResponseMapper(),
		statusStatsMapper: NewTopupStatsStatusResponseMapper(),
		amountStatsMapper: NewTopupStatsAmountResponseMapper(),
		methodStatsMapper: NewTopupMethodResponseMapper(),
	}
}

func (t *topupResponseMapper) QueryMapper() TopupQueryResponseMapper {
	return t.queryMapper
}

func (t *topupResponseMapper) CommandMapper() TopupCommandResponseMapper {
	return t.commandMapper
}

func (t *topupResponseMapper) StatusStatsMapper() TopupStatsStatusResponseMapper {
	return t.statusStatsMapper
}

func (t *topupResponseMapper) AmountStatsMapper() TopupStatsAmountResponseMapper {
	return t.amountStatsMapper
}

func (t *topupResponseMapper) MethodStatsMapper() TopupStatsMethodResponseMapper {
	return t.methodStatsMapper
}
