package cardprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type cardDashboardProtoMapper struct {
}

func NewCardDashboardProtoMapper() *cardDashboardProtoMapper {
	return &cardDashboardProtoMapper{}
}

// ToProtoResponseDashboardCard maps a status, message and a *response.DashboardCard to a *pb.ApiResponseDashboardCard proto response.
//
// It is used to generate the response for the CardService.GetDashboardCard rpc method.
func (s *cardDashboardProtoMapper) ToProtoResponseDashboardCard(status string, message string, dash *response.DashboardCard) *pb.ApiResponseDashboardCard {
	return &pb.ApiResponseDashboardCard{
		Status:  status,
		Message: message,
		Data:    s.mapDashboardCard(dash),
	}
}

// ToProtoResponseDashboardCardCardNumber maps a status, message, and a *response.DashboardCardCardNumber
// to a *pb.ApiResponseDashboardCardNumber proto response.
//
// It is used to generate the response for the CardService.GetDashboardCardCardNumber rpc method.
func (s *cardDashboardProtoMapper) ToProtoResponseDashboardCardCardNumber(status string, message string, dash *response.DashboardCardCardNumber) *pb.ApiResponseDashboardCardNumber {
	return &pb.ApiResponseDashboardCardNumber{
		Status:  status,
		Message: message,
		Data:    s.mapDashboardCardCardNumber(dash),
	}
}

// mapDashboardCard maps a *response.DashboardCard to a *pb.CardResponseDashboard proto response.
//
// This function converts a DashboardCard from the response domain model to its protobuf
// representation, handling the mapping of all financial metrics fields.
func (s *cardDashboardProtoMapper) mapDashboardCard(dash *response.DashboardCard) *pb.CardResponseDashboard {
	return &pb.CardResponseDashboard{
		TotalBalance:     *dash.TotalBalance,
		TotalWithdraw:    *dash.TotalWithdraw,
		TotalTopup:       *dash.TotalTopup,
		TotalTransfer:    *dash.TotalTransfer,
		TotalTransaction: *dash.TotalTransaction,
	}
}

// mapDashboardCardCardNumber maps a *response.DashboardCardCardNumber to a *pb.CardResponseDashboardCardNumber proto response.
//
// This function converts a DashboardCardCardNumber from the response domain model to its protobuf
// representation, handling the mapping of all financial metrics fields for a specific card.
func (s *cardDashboardProtoMapper) mapDashboardCardCardNumber(dash *response.DashboardCardCardNumber) *pb.CardResponseDashboardCardNumber {
	return &pb.CardResponseDashboardCardNumber{
		TotalBalance:          *dash.TotalBalance,
		TotalWithdraw:         *dash.TotalWithdraw,
		TotalTopup:            *dash.TotalTopup,
		TotalTransferSend:     *dash.TotalTransferSend,
		TotalTransferReceiver: *dash.TotalTransferReceiver,
		TotalTransaction:      *dash.TotalTransaction,
	}
}
