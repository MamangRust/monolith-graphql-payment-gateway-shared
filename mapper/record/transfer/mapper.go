package transferrecordmapper

import (
	transferstatsrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/transfer/stats"
	transferstatsbycardrecordmapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/record/transfer/statsbycard"
)

// TransferRecordMapper provides access to all mappers related to Transfer.
type TransferRecordMapper interface {
	QueryMapper() TransferQueryRecordMapper
	CommandMapper() TransferCommandRecordMapper
	StatsMapper() transferstatsrecordmapper.TransferStatisticRecordMapper
	StatsByCardMapper() transferstatsbycardrecordmapper.TransferStatisticByCardRecordMapper
}

// transferRecordMapper is the concrete implementation of TransferRecordMapper.
type transferRecordMapper struct {
	queryMapper       TransferQueryRecordMapper
	commandMapper     TransferCommandRecordMapper
	statsMapper       transferstatsrecordmapper.TransferStatisticRecordMapper
	statsByCardMapper transferstatsbycardrecordmapper.TransferStatisticByCardRecordMapper
}

// NewTransferRecordMapper initializes and returns a new TransferRecordMapper.
func NewTransferRecordMapper() TransferRecordMapper {
	return &transferRecordMapper{
		queryMapper:       NewTransferQueryRecordMapper(),
		commandMapper:     NewTransferCommandRecordMapper(),
		statsMapper:       transferstatsrecordmapper.NewTransferStatisticRecordMapper(),
		statsByCardMapper: transferstatsbycardrecordmapper.NewTransferStatisticByCardRecordMapper(),
	}
}

// QueryMapper returns the query mapper for transfers.
func (t *transferRecordMapper) QueryMapper() TransferQueryRecordMapper {
	return t.queryMapper
}

// CommandMapper returns the command mapper for transfers.
func (t *transferRecordMapper) CommandMapper() TransferCommandRecordMapper {
	return t.commandMapper
}

// StatsMapper returns the general statistics mapper for transfers.
func (t *transferRecordMapper) StatsMapper() transferstatsrecordmapper.TransferStatisticRecordMapper {
	return t.statsMapper
}

// StatsByCardMapper returns the card-specific statistics mapper for transfers.
func (t *transferRecordMapper) StatsByCardMapper() transferstatsbycardrecordmapper.TransferStatisticByCardRecordMapper {
	return t.statsByCardMapper
}
