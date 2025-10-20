package topupprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type TopupBaseProtoMapper interface {
	// ToProtoResponseTopup maps a single top-up record to a protobuf response.
	ToProtoResponseTopup(status string, message string, s *response.TopupResponse) *pb.ApiResponseTopup
}

type TopupQueryProtoMapper interface {
	TopupBaseProtoMapper

	// ToProtoResponsePaginationTopup maps paginated top-up records to a protobuf response.
	ToProtoResponsePaginationTopup(pagination *pbhelpers.PaginationMeta, status string, message string, s []*response.TopupResponse) *pb.ApiResponsePaginationTopup

	// ToProtoResponsePaginationTopupDeleteAt maps paginated soft-deleted top-up records to a protobuf response.
	ToProtoResponsePaginationTopupDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, s []*response.TopupResponseDeleteAt) *pb.ApiResponsePaginationTopupDeleteAt
}

type TopupCommandProtoMapper interface {
	TopupBaseProtoMapper

	// ToProtoResponseTopupDeletAt maps a soft-deleted top-up record to a protobuf response.
	ToProtoResponseTopupDeleteAt(status string, message string, s *response.TopupResponseDeleteAt) *pb.ApiResponseTopupDeleteAt

	// ToProtoResponseTopupAll returns all top-up records in a protobuf response.
	ToProtoResponseTopupAll(status string, message string) *pb.ApiResponseTopupAll

	// ToProtoResponseTopupDelete returns a response indicating a top-up record has been deleted.
	ToProtoResponseTopupDelete(status string, message string) *pb.ApiResponseTopupDelete
}

type TopupStatsStatusProtoMapper interface {
	// ToProtoResponseTopupMonthStatusSuccess maps monthly successful top-ups to a protobuf response.
	ToProtoResponseTopupMonthStatusSuccess(status string, message string, s []*response.TopupResponseMonthStatusSuccess) *pb.ApiResponseTopupMonthStatusSuccess

	// ToProtoResponseTopupYearStatusSuccess maps yearly successful top-ups to a protobuf response.
	ToProtoResponseTopupYearStatusSuccess(status string, message string, s []*response.TopupResponseYearStatusSuccess) *pb.ApiResponseTopupYearStatusSuccess

	// ToProtoResponseTopupMonthStatusFailed maps monthly failed top-ups to a protobuf response.
	ToProtoResponseTopupMonthStatusFailed(status string, message string, s []*response.TopupResponseMonthStatusFailed) *pb.ApiResponseTopupMonthStatusFailed

	// ToProtoResponseTopupYearStatusFailed maps yearly failed top-ups to a protobuf response.
	ToProtoResponseTopupYearStatusFailed(status string, message string, s []*response.TopupResponseYearStatusFailed) *pb.ApiResponseTopupYearStatusFailed
}

type TopupStatsMethodProtoMapper interface {
	// ToProtoResponseTopupMonthMethod maps monthly top-up methods to a protobuf response.
	ToProtoResponseTopupMonthMethod(status string, message string, s []*response.TopupMonthMethodResponse) *pb.ApiResponseTopupMonthMethod

	// ToProtoResponseTopupYearMethod maps yearly top-up methods to a protobuf response.
	ToProtoResponseTopupYearMethod(status string, message string, s []*response.TopupYearlyMethodResponse) *pb.ApiResponseTopupYearMethod
}

// TopupStatsAmountProtoMapper maps top-up amounts to a protobuf response.
type TopupStatsAmountProtoMapper interface {
	// ToProtoResponseTopupMonthAmount maps monthly top-up amounts to a protobuf response.
	ToProtoResponseTopupMonthAmount(status string, message string, s []*response.TopupMonthAmountResponse) *pb.ApiResponseTopupMonthAmount

	// ToProtoResponseTopupYearAmount maps yearly top-up amounts to a protobuf response.
	ToProtoResponseTopupYearAmount(status string, message string, s []*response.TopupYearlyAmountResponse) *pb.ApiResponseTopupYearAmount
}
