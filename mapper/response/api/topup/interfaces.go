package topupapimapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
)

type TopupBaseResponseMapper interface {
	// Converts a single top-up response into an API format.
	ToApiResponseTopup(s *pb.ApiResponseTopup) *response.ApiResponseTopup
}

type TopupQueryResponseMapper interface {
	TopupBaseResponseMapper

	// Converts a paginated list of top-ups into an API response.
	ToApiResponsePaginationTopup(s *pb.ApiResponsePaginationTopup) *response.ApiResponsePaginationTopup

	// Converts a paginated list of soft-deleted top-ups into an API response.
	ToApiResponsePaginationTopupDeleteAt(s *pb.ApiResponsePaginationTopupDeleteAt) *response.ApiResponsePaginationTopupDeleteAt
}

type TopupCommandResponseMapper interface {
	TopupBaseResponseMapper

	// Converts a soft-deleted top-up response into an API format.
	ToApiResponseTopupDeleteAt(s *pb.ApiResponseTopupDeleteAt) *response.ApiResponseTopupDeleteAt

	// Converts all top-up records into a general API response.
	ToApiResponseTopupAll(s *pb.ApiResponseTopupAll) *response.ApiResponseTopupAll

	// Converts a permanently deleted top-up response into an API format.
	ToApiResponseTopupDelete(s *pb.ApiResponseTopupDelete) *response.ApiResponseTopupDelete
}

type TopupStatsStatusResponseMapper interface {
	// Converts monthly successful top-up stats into an API response.
	ToApiResponseTopupMonthStatusSuccess(s *pb.ApiResponseTopupMonthStatusSuccess) *response.ApiResponseTopupMonthStatusSuccess

	// Converts yearly successful top-up stats into an API response.
	ToApiResponseTopupYearStatusSuccess(s *pb.ApiResponseTopupYearStatusSuccess) *response.ApiResponseTopupYearStatusSuccess

	// Converts monthly failed top-up stats into an API response.
	ToApiResponseTopupMonthStatusFailed(s *pb.ApiResponseTopupMonthStatusFailed) *response.ApiResponseTopupMonthStatusFailed

	// Converts yearly failed top-up stats into an API response.
	ToApiResponseTopupYearStatusFailed(s *pb.ApiResponseTopupYearStatusFailed) *response.ApiResponseTopupYearStatusFailed
}

type TopupStatsMethodResponseMapper interface {
	// Converts monthly top-up statistics by payment method into an API response.
	ToApiResponseTopupMonthMethod(s *pb.ApiResponseTopupMonthMethod) *response.ApiResponseTopupMonthMethod

	// Converts yearly top-up statistics by payment method into an API response.
	ToApiResponseTopupYearMethod(s *pb.ApiResponseTopupYearMethod) *response.ApiResponseTopupYearMethod
}

type TopupStatsAmountResponseMapper interface {
	// Converts monthly top-up amount statistics into an API response.
	ToApiResponseTopupMonthAmount(s *pb.ApiResponseTopupMonthAmount) *response.ApiResponseTopupMonthAmount

	// Converts yearly top-up amount statistics into an API response.
	ToApiResponseTopupYearAmount(s *pb.ApiResponseTopupYearAmount) *response.ApiResponseTopupYearAmount
}
