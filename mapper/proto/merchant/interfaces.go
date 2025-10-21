package merchantprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant"
	pbMerchantStats "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type MerchantBaseProtoMapper interface {
	// ToProtoResponseMerchant maps a single merchant response to a protobuf message.
	ToProtoResponseMerchant(status string, message string, res *response.MerchantResponse) *pb.ApiResponseMerchant
}

type MerchantQueryProtoMapper interface {
	MerchantBaseProtoMapper

	// ToProtoResponseMerchants maps a list of merchant responses to a protobuf message.
	ToProtoResponseMerchants(status string, message string, res []*response.MerchantResponse) *pb.ApiResponsesMerchant

	// ToProtoResponsePaginationMerchant maps paginated merchant responses to a protobuf message.
	ToProtoResponsePaginationMerchant(pagination *pbhelpers.PaginationMeta, status string, message string, merchants []*response.MerchantResponse) *pb.ApiResponsePaginationMerchant

	// ToProtoResponsePaginationMerchantDeleteAt maps paginated soft-deleted merchants to a protobuf message.
	ToProtoResponsePaginationMerchantDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, merchants []*response.MerchantResponseDeleteAt) *pb.ApiResponsePaginationMerchantDeleteAt
}

type MerchantCommandProtoMapper interface {
	MerchantBaseProtoMapper

	ToProtoResponseMerchantDeleteAt(status string, message string, res *response.MerchantResponseDeleteAt) *pb.ApiResponseMerchantDeleteAt

	// ToProtoResponseMerchantAll maps all merchant responses (non-paginated) to a protobuf message.
	ToProtoResponseMerchantAll(status string, message string) *pb.ApiResponseMerchantAll

	// ToProtoResponseMerchantDelete returns a protobuf message indicating a merchant has been deleted.
	ToProtoResponseMerchantDelete(status string, message string) *pb.ApiResponseMerchantDelete
}

type MerchantTransactionProtoMapper interface {
	// ToProtoResponsePaginationMerchantTransaction maps paginated merchant transaction summaries to a protobuf message.
	ToProtoResponsePaginationMerchantTransaction(pagination *pbhelpers.PaginationMeta, status string, message string, merchants []*response.MerchantTransactionResponse) *pb.ApiResponsePaginationMerchantTransaction
}

type MerchantStatsMethodProtoMapper interface {
	// ToProtoResponseMonthlyPaymentMethods maps monthly payment method statistics to a protobuf message.
	ToProtoResponseMonthlyPaymentMethods(status string, message string, ms []*response.MerchantResponseMonthlyPaymentMethod) *pbMerchantStats.ApiResponseMerchantMonthlyPaymentMethod

	// ToProtoResponseYearlyPaymentMethods maps yearly payment method statistics to a protobuf message.
	ToProtoResponseYearlyPaymentMethods(status string, message string, ms []*response.MerchantResponseYearlyPaymentMethod) *pbMerchantStats.ApiResponseMerchantYearlyPaymentMethod
}

type MerchantStatsAmountProtoMapper interface {
	// ToProtoResponseMonthlyAmounts maps monthly merchant transaction amounts to a protobuf message.
	ToProtoResponseMonthlyAmounts(status string, message string, ms []*response.MerchantResponseMonthlyAmount) *pbMerchantStats.ApiResponseMerchantMonthlyAmount

	// ToProtoResponseYearlyAmounts maps yearly merchant transaction amounts to a protobuf message.
	ToProtoResponseYearlyAmounts(status string, message string, ms []*response.MerchantResponseYearlyAmount) *pbMerchantStats.ApiResponseMerchantYearlyAmount
}

// MerchantStatsTotalAmountProtoMapper maps total transaction amounts across all merchants to a protobuf message.
type MerchantStatsTotalAmountProtoMapper interface {
	// ToProtoResponseMonthlyTotalAmounts maps monthly total transaction amounts across all merchants to a protobuf message.
	ToProtoResponseMonthlyTotalAmounts(status string, message string, ms []*response.MerchantResponseMonthlyTotalAmount) *pbMerchantStats.ApiResponseMerchantMonthlyTotalAmount

	// ToProtoResponseYearlyTotalAmounts maps yearly total transaction amounts across all merchants to a protobuf message.
	ToProtoResponseYearlyTotalAmounts(status string, message string, ms []*response.MerchantResponseYearlyTotalAmount) *pbMerchantStats.ApiResponseMerchantYearlyTotalAmount
}
