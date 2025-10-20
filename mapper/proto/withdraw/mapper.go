package withdrawprotomapper

type withdrawProtoMapper struct {
	WithdrawQueryProtoMapper        WithdrawQueryProtoMapper
	WithdrawCommandProtoMapper      WithdrawCommandProtoMapper
	WithdrawStatsStatusProtoMapper  WithdrawStatsStatusProtoMapper
	WithdrawaStatsAmountProtoMapper WithdrawaStatsAmountProtoMapper
}

func NewWithdrawProtoMapper() *withdrawProtoMapper {
	return &withdrawProtoMapper{
		WithdrawQueryProtoMapper:        NewWithdrawQueryProtoMapper(),
		WithdrawCommandProtoMapper:      NewWithdrawCommandProtoMapper(),
		WithdrawStatsStatusProtoMapper:  NewWithdrawStatsStatusProtoMapper(),
		WithdrawaStatsAmountProtoMapper: NewWithdrawStatsAmountMapper(),
	}
}
