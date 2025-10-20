package merchantservicemapper

type MerchantResponseMapper interface {
	QueryMapper() MerchantQueryResponseMapper
	CommandMapper() MerchantCommandResponseMapper
	TotalAmountMapper() MerchantTotalAmountResponseMapper
	AmountMapper() MerchantAmountResponseMapper
	MethodMapper() MerchantPaymentMethodResponseMapper
	TransactionMapper() MerchantTransactionResponseMapper
}

type merchantResponseMapper struct {
	queryMapper       MerchantQueryResponseMapper
	commandMapper     MerchantCommandResponseMapper
	totalAmountMapper MerchantTotalAmountResponseMapper
	amountMapper      MerchantAmountResponseMapper
	methodMapper      MerchantPaymentMethodResponseMapper
	transactionMapper MerchantTransactionResponseMapper
}

func NewMerchantResponseMapper() MerchantResponseMapper {
	return &merchantResponseMapper{
		queryMapper:       NewMerchantQueryResponseMapper(),
		commandMapper:     NewMerchantCommandResponseMapper(),
		totalAmountMapper: NewMerchantTotalAmountResponseMapper(),
		amountMapper:      NewMerchantAmountResponseMapper(),
		methodMapper:      NewMerchantPaymentMethodResponseMapper(),
		transactionMapper: NewMerchantTransactionResponseMapper(),
	}
}

func (m *merchantResponseMapper) QueryMapper() MerchantQueryResponseMapper {
	return m.queryMapper
}

func (m *merchantResponseMapper) CommandMapper() MerchantCommandResponseMapper {
	return m.commandMapper
}

func (m *merchantResponseMapper) TotalAmountMapper() MerchantTotalAmountResponseMapper {
	return m.totalAmountMapper
}

func (m *merchantResponseMapper) AmountMapper() MerchantAmountResponseMapper {
	return m.amountMapper
}

func (m *merchantResponseMapper) MethodMapper() MerchantPaymentMethodResponseMapper {
	return m.methodMapper
}

func (m *merchantResponseMapper) TransactionMapper() MerchantTransactionResponseMapper {
	return m.transactionMapper
}
