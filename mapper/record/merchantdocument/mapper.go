package merchantdocumentrecordmapper

type merchantDocumentRecordMapper struct {
	queryMapper   MerchantDocumentQueryMapper
	commandMapper MerchantDocumentCommandMapper
}

type MerchantDocumentMapper interface {
	QueryMapper() MerchantDocumentQueryMapper
	CommandMapper() MerchantDocumentCommandMapper
}


func NewMerchantDocumentRecordMapper() MerchantDocumentMapper {
	return &merchantDocumentRecordMapper{
		queryMapper:   NewMerchantDocumentQueryMapper(),
		commandMapper: NewMerchantDocumentCommandMapper(),
	}
}


func (m *merchantDocumentRecordMapper) QueryMapper() MerchantDocumentQueryMapper {
	return m.queryMapper
}

func (m *merchantDocumentRecordMapper) CommandMapper() MerchantDocumentCommandMapper {
	return m.commandMapper
}
