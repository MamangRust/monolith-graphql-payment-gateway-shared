package cardservicemapper

type CardResponseMapper interface {
	QueryMapper() CardQueryResponseMapper
	CommandMapper() CardCommandResponseMapper
	AmountStatsMapper() CardStatisticAmountResponseMapper
	BalanceStatsMapper() CardStatisticBalanceResponseMapper
}

type cardResponseMapper struct {
	queryMapper        CardQueryResponseMapper
	commandMapper      CardCommandResponseMapper
	amountStatsMapper  CardStatisticAmountResponseMapper
	balanceStatsMapper CardStatisticBalanceResponseMapper
}

func NewCardResponseMapper() CardResponseMapper {
	return &cardResponseMapper{
		queryMapper:        NewCardQueryResponseMapper(),
		commandMapper:      NewCardCommandResponseMapper(),
		amountStatsMapper:  NewCardStatsAmountResponseMapper(),
		balanceStatsMapper: NewCardStatsBalanceResponseMapper(),
	}
}

func (c *cardResponseMapper) QueryMapper() CardQueryResponseMapper {
	return c.queryMapper
}

func (c *cardResponseMapper) CommandMapper() CardCommandResponseMapper {
	return c.commandMapper
}

func (c *cardResponseMapper) AmountStatsMapper() CardStatisticAmountResponseMapper {
	return c.amountStatsMapper
}

func (c *cardResponseMapper) BalanceStatsMapper() CardStatisticBalanceResponseMapper {
	return c.balanceStatsMapper
}
