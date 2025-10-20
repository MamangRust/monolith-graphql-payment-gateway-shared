package merchantdocumentprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchantdocument"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type merchantDocumentQueryProtoMapper struct {
}

// NewMerchantDocumentQueryProtoMapper returns a new instance of
// merchantDocumentQueryProtoMapper which provides methods to map
// MerchantDocumentResponse API-compatible responses to their protobuf representation.
func NewMerchantDocumentQueryProtoMapper() MerchantDocumentQueryProtoMapper {
	return &merchantDocumentQueryProtoMapper{}
}

// ToProtoResponseMerchantDocument converts a MerchantDocumentResponse to its protobuf representation.
// It includes the status, message, and mapped document data.
// Parameters:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - doc: The MerchantDocumentResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseMerchantDocument containing the status, message, and document data.
func (m *merchantDocumentQueryProtoMapper) ToProtoResponseMerchantDocument(status string, message string, doc *response.MerchantDocumentResponse) *pb.ApiResponseMerchantDocument {
	return &pb.ApiResponseMerchantDocument{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantDocument(doc),
	}
}

// ToProtoResponsePaginationMerchantDocument converts a list of MerchantDocumentResponse to its protobuf representation
// with pagination data.
// It includes the status, message, mapped document data, and pagination data.
// Parameters:
//   - pagination: The pagination data for the paginated response.
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - docs: The list of MerchantDocumentResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationMerchantDocument containing the status, message, document data, and pagination data.
func (m *merchantDocumentQueryProtoMapper) ToProtoResponsePaginationMerchantDocument(pagination *pbhelpers.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponse) *pb.ApiResponsePaginationMerchantDocument {
	return &pb.ApiResponsePaginationMerchantDocument{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantDocuments(docs),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationMerchantDocumentDeleteAt converts a list of MerchantDocumentResponseDeleteAt to its protobuf representation
// with pagination data.
// It includes the status, message, mapped document data, and pagination data.
// Parameters:
//   - pagination: The pagination data for the paginated response.
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - docs: The list of MerchantDocumentResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationMerchantDocumentAt containing the status, message, document data, and pagination data.
func (m *merchantDocumentQueryProtoMapper) ToProtoResponsePaginationMerchantDocumentDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponseDeleteAt) *pb.ApiResponsePaginationMerchantDocumentAt {
	return &pb.ApiResponsePaginationMerchantDocumentAt{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantDocumentsDeleteAt(docs),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// mapMerchantDocument maps a MerchantDocumentResponse to its protobuf representation.
// Parameters:
//   - doc: The MerchantDocumentResponse that needs to be converted.
//
// Returns:
//
//	A pointer to MerchantDocument containing the document data.
func (m *merchantDocumentQueryProtoMapper) mapMerchantDocument(doc *response.MerchantDocumentResponse) *pb.MerchantDocument {
	return &pb.MerchantDocument{
		DocumentId:   int32(doc.ID),
		MerchantId:   int32(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentURL,
		Status:       doc.Status,
		Note:         doc.Note,
		UploadedAt:   doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
	}
}

// mapMerchantDocuments maps a list of MerchantDocumentResponse to its protobuf representation.
// Parameters:
//   - docs: The list of MerchantDocumentResponse that needs to be converted.
//
// Returns:
//
//	A slice of pointers to MerchantDocument containing the document data.
func (m *merchantDocumentQueryProtoMapper) mapMerchantDocuments(docs []*response.MerchantDocumentResponse) []*pb.MerchantDocument {
	var res []*pb.MerchantDocument
	for _, doc := range docs {
		res = append(res, m.mapMerchantDocument(doc))
	}
	return res
}

// mapMerchantDocumentDeleteAt maps a MerchantDocumentResponseDeleteAt to its protobuf representation.
// It handles the conversion of document metadata including deletion information.
// Parameters:
//   - doc: The MerchantDocumentResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to MerchantDocumentDeleteAt containing the converted document data.
func (m *merchantDocumentQueryProtoMapper) mapMerchantDocumentDeleteAt(doc *response.MerchantDocumentResponseDeleteAt) *pb.MerchantDocumentDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if doc.DeletedAt != nil {
		deletedAt = wrapperspb.String(*doc.DeletedAt)
	}

	return &pb.MerchantDocumentDeleteAt{
		DocumentId:   int32(doc.ID),
		MerchantId:   int32(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentURL,
		Status:       doc.Status,
		Note:         doc.Note,
		UploadedAt:   doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

// mapMerchantDocumentsDeleteAt maps a list of MerchantDocumentResponseDeleteAt to their protobuf representation.
// It processes each document in the input slice and converts it to a MerchantDocumentDeleteAt,
// including all necessary metadata and deletion information.
// Parameters:
//   - docs: A slice of MerchantDocumentResponseDeleteAt to be converted.
//
// Returns:
//
//	A slice of pointers to MerchantDocumentDeleteAt containing the converted document data.
func (m *merchantDocumentQueryProtoMapper) mapMerchantDocumentsDeleteAt(docs []*response.MerchantDocumentResponseDeleteAt) []*pb.MerchantDocumentDeleteAt {
	var res []*pb.MerchantDocumentDeleteAt
	for _, doc := range docs {
		res = append(res, m.mapMerchantDocumentDeleteAt(doc))
	}
	return res
}
