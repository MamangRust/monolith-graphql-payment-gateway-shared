package merchantdocumentprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type MerchantDocumentBaseProtoMapper interface {
	// ToProtoResponseMerchantDocument maps a single merchant document to a protobuf response.
	ToProtoResponseMerchantDocument(status string, message string, doc *response.MerchantDocumentResponse) *pb.ApiResponseMerchantDocument
}

type MerchantDocumentQueryProtoMapper interface {
	MerchantDocumentBaseProtoMapper

	// ToProtoResponsePaginationMerchantDocument maps paginated merchant documents to a protobuf response.
	ToProtoResponsePaginationMerchantDocument(pagination *pbhelpers.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponse) *pb.ApiResponsePaginationMerchantDocument

	// ToProtoResponsePaginationMerchantDocumentDeleteAt maps paginated soft-deleted merchant documents to a protobuf response.
	ToProtoResponsePaginationMerchantDocumentDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponseDeleteAt) *pb.ApiResponsePaginationMerchantDocumentAt
}

type MerchantDocumentCommandProtoMapper interface {
	MerchantDocumentBaseProtoMapper

	ToProtoResponseMerchantDocumentDeleteAt(status string, message string, doc *response.MerchantDocumentResponseDeleteAt) *pb.ApiResponseMerchantDocumentDeleteAt

	// ToProtoResponseMerchantDocumentDelete returns a response indicating a merchant document has been deleted.
	ToProtoResponseMerchantDocumentDelete(status string, message string) *pb.ApiResponseMerchantDocumentDelete

	// ToProtoResponseMerchantDocumentAll maps all merchant documents to a protobuf response.
	ToProtoResponseMerchantDocumentAll(status string, message string) *pb.ApiResponseMerchantDocumentAll
}
