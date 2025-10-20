package withdrawapimapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
)

type WithdrawBaseResponseMapper interface {
	// Converts a single withdraw response into an API response.
	ToApiResponseWithdraw(pbResponse *pb.ApiResponseWithdraw) *response.ApiResponseWithdraw
}

type WithdrawQueryResponseMapper interface {
	WithdrawBaseResponseMapper

	// Converts a list of withdraw responses into a grouped API response.
	ToApiResponsesWithdraw(pbResponse *pb.ApiResponsesWithdraw) *response.ApiResponsesWithdraw

	// Converts paginated withdraw records into an API response.
	ToApiResponsePaginationWithdraw(pbResponse *pb.ApiResponsePaginationWithdraw) *response.ApiResponsePaginationWithdraw

	// Converts paginated soft-deleted withdraw records into an API response.
	ToApiResponsePaginationWithdrawDeleteAt(pbResponse *pb.ApiResponsePaginationWithdrawDeleteAt) *response.ApiResponsePaginationWithdrawDeleteAt
}

type WithdrawCommandResponseMapper interface {
	WithdrawBaseResponseMapper

	ToApiResponseWithdrawDeleteAt(pbResponse *pb.ApiResponseWithdrawDeleteAt) *response.ApiResponseWithdrawDeleteAt

	// Converts a permanently deleted withdraw response into an API response.
	ToApiResponseWithdrawDelete(pbResponse *pb.ApiResponseWithdrawDelete) *response.ApiResponseWithdrawDelete

	// Converts all withdraw records into an API response.
	ToApiResponseWithdrawAll(pbResponse *pb.ApiResponseWithdrawAll) *response.ApiResponseWithdrawAll
}

type WithdrawStatsStatusResponseMapper interface {
	// Converts monthly successful withdraw statistics into an API response.
	ToApiResponseWithdrawMonthStatusSuccess(pbResponse *pb.ApiResponseWithdrawMonthStatusSuccess) *response.ApiResponseWithdrawMonthStatusSuccess

	// Converts yearly successful withdraw statistics into an API response.
	ToApiResponseWithdrawYearStatusSuccess(pbResponse *pb.ApiResponseWithdrawYearStatusSuccess) *response.ApiResponseWithdrawYearStatusSuccess

	// Converts monthly failed withdraw statistics into an API response.
	ToApiResponseWithdrawMonthStatusFailed(pbResponse *pb.ApiResponseWithdrawMonthStatusFailed) *response.ApiResponseWithdrawMonthStatusFailed

	// Converts yearly failed withdraw statistics into an API response.
	ToApiResponseWithdrawYearStatusFailed(pbResponse *pb.ApiResponseWithdrawYearStatusFailed) *response.ApiResponseWithdrawYearStatusFailed
}

type WithdrawStatsAmountResponseMapper interface {
	// Converts monthly total withdraw amount statistics into an API response.
	ToApiResponseWithdrawMonthAmount(pbResponse *pb.ApiResponseWithdrawMonthAmount) *response.ApiResponseWithdrawMonthAmount

	// Converts yearly total withdraw amount statistics into an API response.
	ToApiResponseWithdrawYearAmount(pbResponse *pb.ApiResponseWithdrawYearAmount) *response.ApiResponseWithdrawYearAmount
}
