package transferprotomapper

type transfrProtoMapper struct {
	TransferQueryProtoMapper       TransferQueryProtoMapper
	TransferCommandProtoMapper     TransferCommandProtoMapper
	TransferStatsStatusProtoMapper TransferStatsStatusProtoMapper
	TransferStatsAmountProtoMapper TransferStatsAmountProtoMapper
}

func NewTransferProtoMapper() *transfrProtoMapper {
	return &transfrProtoMapper{
		TransferQueryProtoMapper:       NewTransferQueryProtoMapper(),
		TransferCommandProtoMapper:     NewTransferCommandProtoMapper(),
		TransferStatsStatusProtoMapper: NewTransferStatsStatusProtoMapper(),
		TransferStatsAmountProtoMapper: NewTransferStatsAmountProtoMapper(),
	}
}
