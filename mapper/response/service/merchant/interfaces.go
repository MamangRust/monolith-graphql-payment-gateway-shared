package merchantservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type MerchantBaseResponseMapper interface {
	// Converts a single MerchantRecord to a MerchantResponse.
	ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse
}

type MerchantQueryResponseMapper interface {
	MerchantBaseResponseMapper

	// Converts multiple MerchantRecords to a slice of MerchantResponse.
	ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse

	// Converts multiple soft-deleted MerchantRecords into a slice of MerchantResponseDeleteAt.
	ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt
}

type MerchantCommandResponseMapper interface {
	MerchantBaseResponseMapper

	// Converts a soft-deleted MerchantRecord into a MerchantResponseDeleteAt.
	ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt
}

type MerchantTransactionResponseMapper interface {
	// Converts a merchant transaction record into a MerchantTransactionResponse.
	ToMerchantTransactionResponse(merchant *record.MerchantTransactionsRecord) *response.MerchantTransactionResponse

	// Converts multiple merchant transaction records into a slice of MerchantTransactionResponse.
	ToMerchantsTransactionResponse(merchants []*record.MerchantTransactionsRecord) []*response.MerchantTransactionResponse
}

type MerchantTotalAmountResponseMapper interface {
	// Converts a monthly total amount record into a MerchantResponseMonthlyTotalAmount.
	ToMerchantMonthlyTotalAmount(ms *record.MerchantMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount

	// Converts multiple monthly total amount records into a slice of MerchantResponseMonthlyTotalAmount.
	ToMerchantMonthlyTotalAmounts(ms []*record.MerchantMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount

	// Converts a yearly total amount record into a MerchantResponseYearlyTotalAmount.
	ToMerchantYearlyTotalAmount(ms *record.MerchantYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount

	// Converts multiple yearly total amount records into a slice of MerchantResponseYearlyTotalAmount.
	ToMerchantYearlyTotalAmounts(ms []*record.MerchantYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount
}

type MerchantPaymentMethodResponseMapper interface {
	// Converts a monthly payment method record into a MerchantResponseMonthlyPaymentMethod.
	ToMerchantMonthlyPaymentMethod(ms *record.MerchantMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod

	// Converts multiple monthly payment method records into a slice of MerchantResponseMonthlyPaymentMethod.
	ToMerchantMonthlyPaymentMethods(ms []*record.MerchantMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod

	// Converts a yearly payment method record into a MerchantResponseYearlyPaymentMethod.
	ToMerchantYearlyPaymentMethod(ms *record.MerchantYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod

	// Converts multiple yearly payment method records into a slice of MerchantResponseYearlyPaymentMethod.
	ToMerchantYearlyPaymentMethods(ms []*record.MerchantYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod
}

type MerchantAmountResponseMapper interface {
	// Converts a monthly amount record into a MerchantResponseMonthlyAmount.
	ToMerchantMonthlyAmount(ms *record.MerchantMonthlyAmount) *response.MerchantResponseMonthlyAmount

	// Converts multiple monthly amount records into a slice of MerchantResponseMonthlyAmount.
	ToMerchantMonthlyAmounts(ms []*record.MerchantMonthlyAmount) []*response.MerchantResponseMonthlyAmount

	// Converts a yearly amount record into a MerchantResponseYearlyAmount.
	ToMerchantYearlyAmount(ms *record.MerchantYearlyAmount) *response.MerchantResponseYearlyAmount

	// Converts multiple yearly amount records into a slice of MerchantResponseYearlyAmount.
	ToMerchantYearlyAmounts(ms []*record.MerchantYearlyAmount) []*response.MerchantResponseYearlyAmount
}
