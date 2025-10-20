package transactionrecordmapper

import (
	transactionstatsrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/transaction/stats"
	transactionstatbycardrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/transaction/statsbycard"
)

// TransactionRecordMapper provides methods to access different mapping strategies
// for query, command, and statistics mapping of Transaction records.
type TransactionRecordMapper interface {
	QueryMapper() TransactionQueryRecordMapper
	CommandMapper() TransactionCommandRecordMapper
	StatsMapper() transactionstatsrecordmapper.TransactonStatisticsRecordMapper
	StatsByCardMapper() transactionstatbycardrecordmapper.TransactonStatisticByCardMapper
}

// transactionRecordMapper is the internal struct that implements TransactionRecordMapper.
// It holds concrete implementations of various record mappers for transaction data.
type transactionRecordMapper struct {
	queryMapper       TransactionQueryRecordMapper
	commandMapper     TransactionCommandRecordMapper
	statsMapper       transactionstatsrecordmapper.TransactonStatisticsRecordMapper
	statsByCardMapper transactionstatbycardrecordmapper.TransactonStatisticByCardMapper
}

// NewTransactionRecordMapper initializes and returns a new TransactionRecordMapper,
// wiring up the necessary mappers for different transaction-related record transformations.
func NewTransactionRecordMapper() TransactionRecordMapper {
	return &transactionRecordMapper{
		queryMapper:       NewTransactionQueryRecordMapper(),
		commandMapper:     NewTransactionCommandRecordMapper(),
		statsMapper:       transactionstatsrecordmapper.NewTransactionStatisticsRecordMapper(),
		statsByCardMapper: transactionstatbycardrecordmapper.NewTransactionStatisticsByCardRecordMapper(),
	}
}

// QueryMapper returns the mapper responsible for transforming transaction records in query operations.
func (t *transactionRecordMapper) QueryMapper() TransactionQueryRecordMapper {
	return t.queryMapper
}

// CommandMapper returns the mapper responsible for transforming transaction records in command operations.
func (t *transactionRecordMapper) CommandMapper() TransactionCommandRecordMapper {
	return t.commandMapper
}

// StatsMapper returns the mapper responsible for transforming transaction statistics records.
func (t *transactionRecordMapper) StatsMapper() transactionstatsrecordmapper.TransactonStatisticsRecordMapper {
	return t.statsMapper
}

// StatsByCardMapper returns the mapper responsible for transforming card-specific transaction statistics records.
func (t *transactionRecordMapper) StatsByCardMapper() transactionstatbycardrecordmapper.TransactonStatisticByCardMapper {
	return t.statsByCardMapper
}
