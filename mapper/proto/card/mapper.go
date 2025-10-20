package cardprotomapper

type cardProtoMapper struct {
	CardQueryProtoMapper        CardQueryProtoMapper
	CardCommandProtoMapper      CardCommandProtoMapper
	CardDashboardProtoMapper    CardDashboardProtoMapper
	CardStatsAmountProtoMapper  CardStatsAmountProtoMapper
	CardStatsBalanceProtoMapper CardStatsBalanceProtoMapper
}

func NewCardProtoMapper() *cardProtoMapper {
	return &cardProtoMapper{
		CardQueryProtoMapper:        NewCardQueryProtoMapper(),
		CardCommandProtoMapper:      NewCardCommandProtoMapper(),
		CardDashboardProtoMapper:    NewCardDashboardProtoMapper(),
		CardStatsAmountProtoMapper:  NewCardStatsAmountProtoMapper(),
		CardStatsBalanceProtoMapper: NewCardStatsBalanceProtoMapper(),
	}
}
