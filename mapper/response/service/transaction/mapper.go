package transactionservicemapper

type TransactionResponseMapper interface {
	QueryMapper() TransactionQueryResponseMapper
	CommandMapper() TransactionCommandResponseMapper
	AmountStatsMapper() TransactionStatsAmountResponseMapper
	MethodStatsMapper() TransactionStatsMethodResponseMapper
	StatusStatsMapper() TransactionStatsStatusResponseMapper
}

type transactionResponseMapper struct {
	queryMapper       TransactionQueryResponseMapper
	commandMapper     TransactionCommandResponseMapper
	amountStatsMapper TransactionStatsAmountResponseMapper
	methodStatsMapper TransactionStatsMethodResponseMapper
	statusStatsMapper TransactionStatsStatusResponseMapper
}

func NewTransactionResponseMapper() TransactionResponseMapper {
	return &transactionResponseMapper{
		queryMapper:       NewTransactionQueryResponseMapper(),
		commandMapper:     NewTransactionCommandResponseMapper(),
		amountStatsMapper: NewTransactionStatsAmountResponseMapper(),
		methodStatsMapper: NewTransactionStatsMethodResponseMapper(),
		statusStatsMapper: NewTransactionStatsStatusResponseMapper(),
	}
}

func (m *transactionResponseMapper) QueryMapper() TransactionQueryResponseMapper {
	return m.queryMapper
}

func (m *transactionResponseMapper) CommandMapper() TransactionCommandResponseMapper {
	return m.commandMapper
}

func (m *transactionResponseMapper) AmountStatsMapper() TransactionStatsAmountResponseMapper {
	return m.amountStatsMapper
}

func (m *transactionResponseMapper) MethodStatsMapper() TransactionStatsMethodResponseMapper {
	return m.methodStatsMapper
}

func (m *transactionResponseMapper) StatusStatsMapper() TransactionStatsStatusResponseMapper {
	return m.statusStatsMapper
}
