package withdrawrecordmapper

import (
	withdrawstatsrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/withdraw/stats"
	withdrawstatsbycardrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/withdraw/statsbycard"
)

// WithdrawRecordMapper provides access to different mapping layers for withdraw data.
type WithdrawRecordMapper interface {
	QueryMapper() WithdrawQueryRecordMapping
	CommandMapper() WithdrawCommandRecordMapping
	StatsMapper() withdrawstatsrecordmapper.WithdrawStatisticRecordMapper
	StatsByCardMapper() withdrawstatsbycardrecordmapper.WithdrawStatisticByCardRecordMapper
}

// withdrawRecordMapper is the concrete implementation of WithdrawRecordMapper.
type withdrawRecordMapper struct {
	queryMapper       WithdrawQueryRecordMapping
	commandMapper     WithdrawCommandRecordMapping
	statsMapper       withdrawstatsrecordmapper.WithdrawStatisticRecordMapper
	statsByCardMapper withdrawstatsbycardrecordmapper.WithdrawStatisticByCardRecordMapper
}

// NewWithdrawRecordMapper initializes and returns a new WithdrawRecordMapper with
// concrete mapper implementations for query, command, and statistics mapping.
func NewWithdrawRecordMapper() WithdrawRecordMapper {
	return &withdrawRecordMapper{
		queryMapper:       NewWithdrawQueryRecordMapper(),
		commandMapper:     NewWithdrawCommandRecordMapper(),
		statsMapper:       withdrawstatsrecordmapper.NewWithdrawStatisticRecordMapper(),
		statsByCardMapper: withdrawstatsbycardrecordmapper.NewWithdrawStatisticByCardRecordMapper(),
	}
}

// QueryMapper returns the query mapper implementation.
func (w *withdrawRecordMapper) QueryMapper() WithdrawQueryRecordMapping {
	return w.queryMapper
}

// CommandMapper returns the command mapper implementation.
func (w *withdrawRecordMapper) CommandMapper() WithdrawCommandRecordMapping {
	return w.commandMapper
}

// StatsMapper returns the statistics mapper implementation.
func (w *withdrawRecordMapper) StatsMapper() withdrawstatsrecordmapper.WithdrawStatisticRecordMapper {
	return w.statsMapper
}

// StatsByCardMapper returns the card-based statistics mapper implementation.
func (w *withdrawRecordMapper) StatsByCardMapper() withdrawstatsbycardrecordmapper.WithdrawStatisticByCardRecordMapper {
	return w.statsByCardMapper
}
