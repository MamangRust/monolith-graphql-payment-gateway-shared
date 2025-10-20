package saldoprotomapper

type saldoProtoMapper struct {
	SaldoQueryProtoMapper           SaldoQueryProtoMapper
	SaldoCommandProtoMapper         SaldoCommandProtoMapper
	SaldoStatsBalanceProtoMapper    SaldoStatsBalanceProtoMapper
	SaldoStatsTotalSaldoProtoMapper SaldoStatsTotalSaldoProtoMapper
}

func NewSaldoProtoMapper() *saldoProtoMapper {
	return &saldoProtoMapper{
		SaldoQueryProtoMapper:           NewSaldoQueryProtoMapper(),
		SaldoCommandProtoMapper:         NewSaldoCommandProtoMapper(),
		SaldoStatsBalanceProtoMapper:    NewSaldoStatsBalanceProtoMapper(),
		SaldoStatsTotalSaldoProtoMapper: NewSaldoStatsTotalProtoMapper(),
	}
}
