package transferresponsemapper

type TransferResponseMapper interface {
	QueryMapper() TransferQueryResponseMapper
	CommandMapper() TransferCommandResponseMapper
	AmountStatsMapper() TransferAmountResponseMapper
	StatusStatsMapper() TransferStatsStatusResponseMapper
}

type transferResponseMapper struct {
	queryMapper       TransferQueryResponseMapper
	commandMapper     TransferCommandResponseMapper
	amountStatsMapper TransferAmountResponseMapper
	statusStatsMapper TransferStatsStatusResponseMapper
}

func NewTransferResponseMapper() TransferResponseMapper {
	return &transferResponseMapper{
		queryMapper:       NewTransferQueryResponseMapper(),
		commandMapper:     NewTransferCommandResponseMapper(),
		amountStatsMapper: NewTransferStatsAmountResponseMapper(),
		statusStatsMapper: NewTransferStatsStatusResponseMapper(),
	}
}

func (m *transferResponseMapper) QueryMapper() TransferQueryResponseMapper {
	return m.queryMapper
}

func (m *transferResponseMapper) CommandMapper() TransferCommandResponseMapper {
	return m.commandMapper
}

func (m *transferResponseMapper) AmountStatsMapper() TransferAmountResponseMapper {
	return m.amountStatsMapper
}

func (m *transferResponseMapper) StatusStatsMapper() TransferStatsStatusResponseMapper {
	return m.statusStatsMapper
}
