package withdrawresponseservice

type WithdrawResponseMapper interface {
	QueryMapper() WithdrawQueryResponseMapper
	CommandMapper() WithdrawCommandResponseMapper
	StatusStatsMapper() WithdrawStatsStatusResponseMapper
	AmountStatsMapper() WithdrawStatsAmountResponseMapper
}

type withdrawResponseMapper struct {
	queryMapper       WithdrawQueryResponseMapper
	commandMapper     WithdrawCommandResponseMapper
	statusStatsMapper WithdrawStatsStatusResponseMapper
	amountStatsMapper WithdrawStatsAmountResponseMapper
}

func NewWithdrawResponseMapper() WithdrawResponseMapper {
	return &withdrawResponseMapper{
		queryMapper:       NewWithdrawQueryResponseMapper(),
		commandMapper:     NewWithdrawCommandResponseMapper(),
		statusStatsMapper: NewWithdrawStatsStatusResponseMapper(),
		amountStatsMapper: NewWithdrawStatsAmountResponseMapper(),
	}
}

func (m *withdrawResponseMapper) QueryMapper() WithdrawQueryResponseMapper {
	return m.queryMapper
}

func (m *withdrawResponseMapper) CommandMapper() WithdrawCommandResponseMapper {
	return m.commandMapper
}

func (m *withdrawResponseMapper) StatusStatsMapper() WithdrawStatsStatusResponseMapper {
	return m.statusStatsMapper
}

func (m *withdrawResponseMapper) AmountStatsMapper() WithdrawStatsAmountResponseMapper {
	return m.amountStatsMapper
}
