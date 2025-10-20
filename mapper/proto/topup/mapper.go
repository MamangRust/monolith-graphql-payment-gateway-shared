package topupprotomapper

type topupProtoMapper struct {
	TopupQueryProtoMapper       TopupQueryProtoMapper
	TopupCommandProtoMapper     TopupCommandProtoMapper
	TopupStatsAmountProtoMapper TopupStatsAmountProtoMapper
	TopupStatsMethodProtoMapper TopupStatsMethodProtoMapper
	TopupStatsStatusProtoMapper TopupStatsStatusProtoMapper
}

func NewTopupProtoMapper() *topupProtoMapper {
	return &topupProtoMapper{
		TopupQueryProtoMapper:       NewTopupQueryProtoMapper(),
		TopupCommandProtoMapper:     NewTopupCommandProtoMapper(),
		TopupStatsAmountProtoMapper: NewTopupStatsAmountProtoMapper(),
		TopupStatsMethodProtoMapper: NewTopupStatsMethodProtoMapper(),
		TopupStatsStatusProtoMapper: NewTopupStatsStatusProtoMapper(),
	}
}
