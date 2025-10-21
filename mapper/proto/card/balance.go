package cardprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/card/stats"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type cardStatsBalanceProtoMapper struct {
}

func NewCardStatsBalanceProtoMapper() CardStatsBalanceProtoMapper {
	return &cardStatsBalanceProtoMapper{}
}

// ToProtoResponseMonthlyBalances maps a status, message and a list of *response.CardResponseMonthBalance
// to a *pb.ApiResponseMonthlyBalance proto response.
//
// It is used to generate the response for the CardService.GetMonthlyBalance rpc method.
func (s *cardStatsBalanceProtoMapper) ToProtoResponseMonthlyBalances(status string, message string, cards []*response.CardResponseMonthBalance) *pb.ApiResponseMonthlyBalance {
	return &pb.ApiResponseMonthlyBalance{
		Status:  status,
		Message: message,
		Data:    s.mapMonthlyBalances(cards),
	}
}

// ToProtoResponseYearlyBalances maps a status, message and a list of *response.CardResponseYearlyBalance
// to a *pb.ApiResponseYearlyBalance proto response.
//
// It is used to generate the response for the CardService.GetYearlyBalance rpc method.
func (s *cardStatsBalanceProtoMapper) ToProtoResponseYearlyBalances(status string, message string, cards []*response.CardResponseYearlyBalance) *pb.ApiResponseYearlyBalance {

	return &pb.ApiResponseYearlyBalance{
		Status:  status,
		Message: message,
		Data:    s.mapYearlyBalances(cards),
	}
}

// mapMonthlyBalance maps a *response.CardResponseMonthBalance to a *pb.CardResponseMonthlyBalance
// proto response.
//
// It is used to generate the response for the CardService.GetMonthlyBalance rpc method.
func (s *cardStatsBalanceProtoMapper) mapMonthlyBalance(card *response.CardResponseMonthBalance) *pb.CardResponseMonthlyBalance {
	return &pb.CardResponseMonthlyBalance{
		Month:        card.Month,
		TotalBalance: card.TotalBalance,
	}
}

// mapMonthlyBalances maps a list of *response.CardResponseMonthBalance to a list of
// *pb.CardResponseMonthlyBalance proto responses.
//
// It iterates over each CardResponseMonthBalance in the input slice, converting
// them to their protobuf equivalent using the mapMonthlyBalance function. This
// function is used to generate the response data for monthly balance RPC methods.
func (s *cardStatsBalanceProtoMapper) mapMonthlyBalances(roles []*response.CardResponseMonthBalance) []*pb.CardResponseMonthlyBalance {
	var responseRoles []*pb.CardResponseMonthlyBalance

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMonthlyBalance(role))
	}

	return responseRoles
}

// mapYearlyBalance maps a *response.CardResponseYearlyBalance to a *pb.CardResponseYearlyBalance proto response.
//
// This function takes a CardResponseYearlyBalance from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.
func (s *cardStatsBalanceProtoMapper) mapYearlyBalance(card *response.CardResponseYearlyBalance) *pb.CardResponseYearlyBalance {
	return &pb.CardResponseYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

// mapYearlyBalances maps a list of *response.CardResponseYearlyBalance to a list of
// *pb.CardResponseYearlyBalance proto responses.
//
// It iterates over each CardResponseYearlyBalance in the input slice, converting
// them to their protobuf equivalent using the mapYearlyBalance function. This
// function is used to generate the response data for yearly balance RPC methods.
func (s *cardStatsBalanceProtoMapper) mapYearlyBalances(roles []*response.CardResponseYearlyBalance) []*pb.CardResponseYearlyBalance {
	var responseRoles []*pb.CardResponseYearlyBalance

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapYearlyBalance(role))
	}

	return responseRoles
}
