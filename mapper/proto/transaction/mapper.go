package transactionprotomapper

type transactonProtoMapper struct {
	TransactionQueryProtoMapper       TransactionQueryProtoMapper
	TransactionCommandProtoMapper     TransactionCommandProtoMapper
	TransactionStatsAmountProtoMapper TransactionStatsAmountProtoMapper
	TransactionStatsMethodProtoMapper TransactionStatsMethodProtoMapper
	TransactionStatsStatusProtoMapper TransactionStatsStatusProtoMapper
}

func NewTransactionProtoMapper() *transactonProtoMapper {
	return &transactonProtoMapper{
		TransactionQueryProtoMapper:       NewTransactonQueryProtoMapper(),
		TransactionCommandProtoMapper:     NewTransactionCommandProtoMapper(),
		TransactionStatsAmountProtoMapper: NewTransactionStatsAmountProtoMapper(),
		TransactionStatsMethodProtoMapper: NewTransactionStatsMethodProtoMapper(),
		TransactionStatsStatusProtoMapper: NewTransactionStatsStatusProtoMapper(),
	}
}
