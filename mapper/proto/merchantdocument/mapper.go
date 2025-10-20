package merchantdocumentprotomapper

type merchantDocumentProtoMapper struct {
	MerchantDocumentQueryProtoMapper   MerchantDocumentQueryProtoMapper
	MerchantDocumentCommandProtoMapper MerchantDocumentCommandProtoMapper
}

func NewMerchantDocumentProtoMapper() *merchantDocumentProtoMapper {
	return &merchantDocumentProtoMapper{
		MerchantDocumentQueryProtoMapper:   NewMerchantDocumentQueryProtoMapper(),
		MerchantDocumentCommandProtoMapper: NewMerchantDocumentCommandProtoMapper(),
	}
}
