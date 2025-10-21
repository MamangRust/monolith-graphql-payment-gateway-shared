package transferprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transfer"
	pbStats "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transfer/stats"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type TransferBaseProtoMapper interface {
	// ToProtoResponseTransfer maps a single transfer to a protobuf response.
	ToProtoResponseTransfer(status string, message string, pbResponse *response.TransferResponse) *pb.ApiResponseTransfer
}

type TransferQueryProtoMapper interface {
	TransferBaseProtoMapper

	// ToProtoResponseTransfers maps a list of transfers to a protobuf response.
	ToProtoResponseTransfers(status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponseTransfers

	// ToProtoResponsePaginationTransfer maps paginated transfer records to a protobuf response.
	ToProtoResponsePaginationTransfer(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponsePaginationTransfer
	// ToProtoResponsePaginationTransferDeleteAt maps paginated soft-deleted transfers to a protobuf response.
	ToProtoResponsePaginationTransferDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransferResponseDeleteAt) *pb.ApiResponsePaginationTransferDeleteAt
}

type TransferCommandProtoMapper interface {
	TransferBaseProtoMapper

	ToProtoResponseTransferDeleteAt(status string, message string, pbResponse *response.TransferResponseDeleteAt) *pb.ApIResponseTransferDeleteAt

	// ToProtoResponseTransferDelete returns a response indicating a transfer has been deleted.
	ToProtoResponseTransferDelete(status string, message string) *pb.ApiResponseTransferDelete
	// ToProtoResponseTransferAll returns all transfers in a protobuf response.
	ToProtoResponseTransferAll(status string, message string) *pb.ApiResponseTransferAll
}

type TransferStatsStatusProtoMapper interface {
	// ToProtoResponseTransferMonthStatusSuccess maps monthly successful transfers to a protobuf response.
	ToProtoResponseTransferMonthStatusSuccess(status string, message string, pbResponse []*response.TransferResponseMonthStatusSuccess) *pbStats.ApiResponseTransferMonthStatusSuccess
	// ToProtoResponseTransferYearStatusSuccess maps yearly successful transfers to a protobuf response.
	ToProtoResponseTransferYearStatusSuccess(status string, message string, pbResponse []*response.TransferResponseYearStatusSuccess) *pbStats.ApiResponseTransferYearStatusSuccess
	// ToProtoResponseTransferMonthStatusFailed maps monthly failed transfers to a protobuf response.
	ToProtoResponseTransferMonthStatusFailed(status string, message string, pbResponse []*response.TransferResponseMonthStatusFailed) *pbStats.ApiResponseTransferMonthStatusFailed
	// ToProtoResponseTransferYearStatusFailed maps yearly failed transfers to a protobuf response.
	ToProtoResponseTransferYearStatusFailed(status string, message string, pbResponse []*response.TransferResponseYearStatusFailed) *pbStats.ApiResponseTransferYearStatusFailed
}

type TransferStatsAmountProtoMapper interface {
	// ToProtoResponseTransferMonthAmount maps monthly transfer amounts to a protobuf response.
	ToProtoResponseTransferMonthAmount(status string, message string, pbResponse []*response.TransferMonthAmountResponse) *pbStats.ApiResponseTransferMonthAmount

	// ToProtoResponseTransferYearAmount maps yearly transfer amounts to a protobuf response.
	ToProtoResponseTransferYearAmount(status string, message string, pbResponse []*response.TransferYearAmountResponse) *pbStats.ApiResponseTransferYearAmount
}
