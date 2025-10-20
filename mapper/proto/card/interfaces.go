package cardprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
)

type CardBaseProtoMapper interface {
	// ToProtoResponseCard maps a single card response to a protobuf message.
	ToProtoResponseCard(status string, message string, card *response.CardResponse) *pb.ApiResponseCard
}

type CardQueryProtoMapper interface {
	CardBaseProtoMapper

	// ToProtoResponsePaginationCard maps a paginated list of cards to a protobuf message.
	ToProtoResponsePaginationCard(pagination *pbhelpers.PaginationMeta, status string, message string, cards []*response.CardResponse) *pb.ApiResponsePaginationCard

	// ToProtoResponsePaginationCardDeletedAt maps paginated soft-deleted cards to a protobuf message.
	ToProtoResponsePaginationCardDeletedAt(pagination *pbhelpers.PaginationMeta, status string, message string, cards []*response.CardResponseDeleteAt) *pb.ApiResponsePaginationCardDeleteAt
}

type CardCommandProtoMapper interface {
	CardBaseProtoMapper

	ToProtoResponseCardDeleteAt(status string, message string, card *response.CardResponseDeleteAt) *pb.ApiResponseCardDeleteAt

	ToProtoResponseCardDelete(status string, message string) *pb.ApiResponseCardDelete

	// ToProtoResponseCardAll returns all cards without pagination in a protobuf message.
	ToProtoResponseCardAll(status string, message string) *pb.ApiResponseCardAll
}

type CardDashboardProtoMapper interface {
	// ToProtoResponseDashboardCard maps dashboard statistics for cards to a protobuf message.
	ToProtoResponseDashboardCard(status string, message string, dash *response.DashboardCard) *pb.ApiResponseDashboardCard

	// ToProtoResponseDashboardCardCardNumber maps card number statistics to a protobuf message.
	ToProtoResponseDashboardCardCardNumber(status string, message string, dash *response.DashboardCardCardNumber) *pb.ApiResponseDashboardCardNumber
}

type CardStatsBalanceProtoMapper interface {
	// ToProtoResponseMonthlyBalances maps monthly card balances to a protobuf message.
	ToProtoResponseMonthlyBalances(status string, message string, cards []*response.CardResponseMonthBalance) *pb.ApiResponseMonthlyBalance

	// ToProtoResponseYearlyBalances maps yearly card balances to a protobuf message.
	ToProtoResponseYearlyBalances(status string, message string, cards []*response.CardResponseYearlyBalance) *pb.ApiResponseYearlyBalance
}

type CardStatsAmountProtoMapper interface {
	// ToProtoResponseMonthlyAmounts maps monthly card amounts to a protobuf message.
	ToProtoResponseMonthlyAmounts(status string, message string, cards []*response.CardResponseMonthAmount) *pb.ApiResponseMonthlyAmount

	// ToProtoResponseYearlyAmounts maps yearly card amounts to a protobuf message.
	ToProtoResponseYearlyAmounts(status string, message string, cards []*response.CardResponseYearAmount) *pb.ApiResponseYearlyAmount
}
