package merchantdocumentprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchantdocument"
)

type merchantDocumentCommandProtoMapper struct {
}

func NewMerchantDocumentCommandProtoMapper() MerchantDocumentCommandProtoMapper {
	return &merchantDocumentCommandProtoMapper{}

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
func (m *merchantDocumentCommandProtoMapper) ToProtoResponseMerchantDocument(status string, message string, doc *response.MerchantDocumentResponse) *pb.ApiResponseMerchantDocument {
	return &pb.ApiResponseMerchantDocument{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantDocument(doc),
	}
}

func (m *merchantDocumentCommandProtoMapper) ToProtoResponseMerchantDocumentDeleteAt(status string, message string, doc *response.MerchantDocumentResponseDeleteAt) *pb.ApiResponseMerchantDocumentDeleteAt {
	return &pb.ApiResponseMerchantDocumentDeleteAt{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantDocumentDeleteAt(doc),
	}
}

// ToProtoResponseMerchantDocumentDelete converts a response message to its protobuf representation
// specifically for merchant document deletion operations.
// It includes the status and message of the API response.
// Parameters:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseMerchantDocumentDelete containing the status and message.
func (m *merchantDocumentCommandProtoMapper) ToProtoResponseMerchantDocumentDelete(status string, message string) *pb.ApiResponseMerchantDocumentDelete {
	return &pb.ApiResponseMerchantDocumentDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseMerchantDocumentAll converts a status and message to its protobuf representation
// specifically for bulk merchant document operations.
// It includes the status and message of the API response.
// Parameters:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseMerchantDocumentAll containing the status and message.
func (m *merchantDocumentCommandProtoMapper) ToProtoResponseMerchantDocumentAll(status string, message string) *pb.ApiResponseMerchantDocumentAll {
	return &pb.ApiResponseMerchantDocumentAll{
		Status:  status,
		Message: message,
	}
}

// mapMerchantDocument maps a MerchantDocumentResponse to its protobuf representation.
// Parameters:
//   - doc: The MerchantDocumentResponse that needs to be converted.
//
// Returns:
//
//	A pointer to MerchantDocument containing the document data.
func (m *merchantDocumentCommandProtoMapper) mapMerchantDocument(doc *response.MerchantDocumentResponse) *pb.MerchantDocument {
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

func (m *merchantDocumentCommandProtoMapper) mapMerchantDocumentDeleteAt(doc *response.MerchantDocumentResponseDeleteAt) *pb.MerchantDocumentDeleteAt {
	res := &pb.MerchantDocumentDeleteAt{
		DocumentId:   int32(doc.ID),
		MerchantId:   int32(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentURL,
		Status:       doc.Status,
		Note:         doc.Note,
		UploadedAt:   doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
	}

	if doc.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(doc.DeletedAt)
	}

	return res
}
