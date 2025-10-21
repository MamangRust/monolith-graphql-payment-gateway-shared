package withdrawprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/withdraw"
	pbStats "github.com/MamangRust/monolith-graphql-payment-gateway-pb/withdraw/stats"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type WithdrawBaseProtoMapper interface {
	// ToProtoResponseWithdraw maps a single withdraw to a protobuf response.
	ToProtoResponseWithdraw(status string, message string, withdraw *response.WithdrawResponse) *pb.ApiResponseWithdraw
}

type WithdrawQueryProtoMapper interface {
	WithdrawBaseProtoMapper

	// ToProtoResponsePaginationWithdraw maps paginated withdraw records to a protobuf response.
	ToProtoResponsePaginationWithdraw(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsePaginationWithdraw

	// ToProtoResponsePaginationWithdrawDeleteAt maps paginated soft-deleted withdraw records to a protobuf response.
	ToProtoResponsePaginationWithdrawDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponseDeleteAt) *pb.ApiResponsePaginationWithdrawDeleteAt
}

type WithdrawCommandProtoMapper interface {
	WithdrawBaseProtoMapper

	ToProtoResponseWithdrawDeleteAt(status string, message string, withdraw *response.WithdrawResponseDeleteAt) *pb.ApIResponseWithdrawDeleteAt

	// ToProtoResponseWithdrawDelete returns a response indicating a withdraw has been deleted.
	ToProtoResponseWithdrawDelete(status string, message string) *pb.ApiResponseWithdrawDelete

	// ToProtoResponseWithdrawAll returns all withdraws in a protobuf response.
	ToProtoResponseWithdrawAll(status string, message string) *pb.ApiResponseWithdrawAll
}

type WithdrawStatsStatusProtoMapper interface {
	// ToProtoResponseWithdrawMonthStatusSuccess maps monthly successful withdraws to a protobuf response.
	ToProtoResponseWithdrawMonthStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusSuccess) *pbStats.ApiResponseWithdrawMonthStatusSuccess

	// ToProtoResponseWithdrawYearStatusSuccess maps yearly successful withdraws to a protobuf response.
	ToProtoResponseWithdrawYearStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseYearStatusSuccess) *pbStats.ApiResponseWithdrawYearStatusSuccess

	// ToProtoResponseWithdrawMonthStatusFailed maps monthly failed withdraws to a protobuf response.
	ToProtoResponseWithdrawMonthStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusFailed) *pbStats.ApiResponseWithdrawMonthStatusFailed

	// ToProtoResponseWithdrawYearStatusFailed maps yearly failed withdraws to a protobuf response.
	ToProtoResponseWithdrawYearStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseYearStatusFailed) *pbStats.ApiResponseWithdrawYearStatusFailed
}

type WithdrawaStatsAmountProtoMapper interface {
	// ToProtoResponseWithdrawMonthAmount maps monthly withdraw amounts to a protobuf response.
	ToProtoResponseWithdrawMonthAmount(status string, message string, pbResponse []*response.WithdrawMonthlyAmountResponse) *pbStats.ApiResponseWithdrawMonthAmount

	// ToProtoResponseWithdrawYearAmount maps yearly withdraw amounts to a protobuf response.
	ToProtoResponseWithdrawYearAmount(status string, message string, pbResponse []*response.WithdrawYearlyAmountResponse) *pbStats.ApiResponseWithdrawYearAmount
}
