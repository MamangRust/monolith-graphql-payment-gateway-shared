package transactionstatbycardrecordmapper

type TransactonStatisticByCardMapper interface {
	TransactionStatisticByCardStatusMapper
	TransactionStatisticByCardAmountMapper
	TransactionStatisticByCardMethodMapper
}

// transactionStatisticByCardRecordMapper is a struct that implements the TransactionStatisticByCardRecordMapper interface.
type transactionStatisticByCardRecordMapper struct {
	TransactionStatisticByCardStatusMapper
	TransactionStatisticByCardAmountMapper
	TransactionStatisticByCardMethodMapper
}

// NewTransactionStatisticsByCardRecordMapper returns a new instance of
// transactionStatisticByCardRecordMapper, which provides methods to map
// database rows related to transaction statistics by card number to domain
// models.
func NewTransactionStatisticsByCardRecordMapper() TransactonStatisticByCardMapper {
	return &transactionStatisticByCardRecordMapper{
		TransactionStatisticByCardStatusMapper: NewTransactionStatisticStatusByCardMapper(),
		TransactionStatisticByCardAmountMapper: NewTransactionStatisticAmountByCardMapper(),
		TransactionStatisticByCardMethodMapper: NewTransactionStatisticMethodByCardMapper(),
	}
}
