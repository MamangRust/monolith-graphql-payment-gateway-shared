package cardprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
)

type cardStatsAmountProtoMapper struct {
}

func NewCardStatsAmountProtoMapper() CardStatsAmountProtoMapper {
	return &cardStatsAmountProtoMapper{}
}

// ToProtoResponseMonthlyAmounts maps a status, message and a list of *response.CardResponseMonthAmount
// to a *pb.ApiResponseMonthlyAmount proto response.
//
// It is used to generate the response for the CardService.GetMonthlyAmount rpc method.
func (s *cardStatsAmountProtoMapper) ToProtoResponseMonthlyAmounts(status string, message string, cards []*response.CardResponseMonthAmount) *pb.ApiResponseMonthlyAmount {

	return &pb.ApiResponseMonthlyAmount{
		Status:  status,
		Message: message,
		Data:    s.mapMonthlyAmounts(cards),
	}
}

// ToProtoResponseYearlyAmounts maps a status, message and a list of *response.CardResponseYearAmount
// to a *pb.ApiResponseYearlyAmount proto response.
//
// It is used to generate the response for the CardService.GetYearlyAmount rpc method.
func (s *cardStatsAmountProtoMapper) ToProtoResponseYearlyAmounts(status string, message string, cards []*response.CardResponseYearAmount) *pb.ApiResponseYearlyAmount {
	return &pb.ApiResponseYearlyAmount{
		Status:  status,
		Message: message,
		Data:    s.mapYearlyAmounts(cards),
	}
}

// mapMonthlyAmount maps a *response.CardResponseMonthAmount to a *pb.CardResponseMonthlyAmount proto response.
//
// It is used to generate the response for the CardService.GetMonthlyAmount rpc method.
func (s *cardStatsAmountProtoMapper) mapMonthlyAmount(card *response.CardResponseMonthAmount) *pb.CardResponseMonthlyAmount {
	return &pb.CardResponseMonthlyAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

// mapMonthlyAmounts maps a list of *response.CardResponseMonthAmount to a list of
// *pb.CardResponseMonthlyAmount proto responses.
//
// It iterates over each CardResponseMonthAmount in the input slice, converting
// them to their protobuf equivalent using the mapMonthlyAmount function. This
// function is used to generate the response data for monthly amount RPC methods.
func (s *cardStatsAmountProtoMapper) mapMonthlyAmounts(roles []*response.CardResponseMonthAmount) []*pb.CardResponseMonthlyAmount {
	var responseRoles []*pb.CardResponseMonthlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMonthlyAmount(role))
	}

	return responseRoles
}

// mapYearlyAmount maps a *response.CardResponseYearAmount to a *pb.CardResponseYearlyAmount proto response.
//
// This function takes a CardResponseYearAmount from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total amount fields are accurately mapped.
func (s *cardStatsAmountProtoMapper) mapYearlyAmount(card *response.CardResponseYearAmount) *pb.CardResponseYearlyAmount {
	return &pb.CardResponseYearlyAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

// mapYearlyAmounts maps a list of *response.CardResponseYearAmount to a list of
// *pb.CardResponseYearlyAmount proto responses.
//
// It iterates over each CardResponseYearAmount in the input slice, converting
// them to their protobuf equivalent using the mapYearlyAmount function. This
// function is used to generate the response data for yearly amount RPC methods.
func (s *cardStatsAmountProtoMapper) mapYearlyAmounts(roles []*response.CardResponseYearAmount) []*pb.CardResponseYearlyAmount {
	var responseRoles []*pb.CardResponseYearlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapYearlyAmount(role))
	}

	return responseRoles
}
