package topupservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type TopupBaseResponseMapper interface {
	// Converts a single TopupRecord into a TopupResponse.
	ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse
}

type TopupQueryResponseMapper interface {
	TopupBaseResponseMapper

	// Converts a slice of TopupRecord into a slice of TopupResponse.
	ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse

	// Converts multiple soft-deleted TopupRecord into a slice of TopupResponseDeleteAt.
	ToTopupResponsesDeleteAt(topups []*record.TopupRecord) []*response.TopupResponseDeleteAt
}

type TopupCommandResponseMapper interface {
	TopupBaseResponseMapper

	// Converts a soft-deleted TopupRecord into a TopupResponseDeleteAt.
	ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt
}

type TopupStatsStatusResponseMapper interface {
	// Converts a monthly top-up success record into a TopupResponseMonthStatusSuccess.
	ToTopupResponseMonthStatusSuccess(s *record.TopupRecordMonthStatusSuccess) *response.TopupResponseMonthStatusSuccess

	// Converts multiple monthly top-up success records into a slice of TopupResponseMonthStatusSuccess.
	ToTopupResponsesMonthStatusSuccess(topups []*record.TopupRecordMonthStatusSuccess) []*response.TopupResponseMonthStatusSuccess

	// Converts a yearly top-up success record into a TopupResponseYearStatusSuccess.
	ToTopupResponseYearStatusSuccess(s *record.TopupRecordYearStatusSuccess) *response.TopupResponseYearStatusSuccess

	// Converts multiple yearly top-up success records into a slice of TopupResponseYearStatusSuccess.
	ToTopupResponsesYearStatusSuccess(topups []*record.TopupRecordYearStatusSuccess) []*response.TopupResponseYearStatusSuccess

	// Converts a monthly top-up failed record into a TopupResponseMonthStatusFailed.
	ToTopupResponseMonthStatusFailed(s *record.TopupRecordMonthStatusFailed) *response.TopupResponseMonthStatusFailed

	// Converts multiple monthly top-up failed records into a slice of TopupResponseMonthStatusFailed.
	ToTopupResponsesMonthStatusFailed(topups []*record.TopupRecordMonthStatusFailed) []*response.TopupResponseMonthStatusFailed

	// Converts a yearly top-up failed record into a TopupResponseYearStatusFailed.
	ToTopupResponseYearStatusFailed(s *record.TopupRecordYearStatusFailed) *response.TopupResponseYearStatusFailed

	// Converts multiple yearly top-up failed records into a slice of TopupResponseYearStatusFailed.
	ToTopupResponsesYearStatusFailed(topups []*record.TopupRecordYearStatusFailed) []*response.TopupResponseYearStatusFailed
}

type TopupStatsAmountResponseMapper interface {
	// Converts a monthly amount record into a TopupMonthAmountResponse.
	ToTopupMonthlyAmountResponse(s *record.TopupMonthAmount) *response.TopupMonthAmountResponse

	// Converts multiple monthly amount records into a slice of TopupMonthAmountResponse.
	ToTopupMonthlyAmountResponses(s []*record.TopupMonthAmount) []*response.TopupMonthAmountResponse

	// Converts a yearly amount record into a TopupYearlyAmountResponse.
	ToTopupYearlyAmountResponse(s *record.TopupYearlyAmount) *response.TopupYearlyAmountResponse

	// Converts multiple yearly amount records into a slice of TopupYearlyAmountResponse.
	ToTopupYearlyAmountResponses(s []*record.TopupYearlyAmount) []*response.TopupYearlyAmountResponse
}

type TopupStatsMethodResponseMapper interface {
	// Converts a monthly method record into a TopupMonthMethodResponse.
	ToTopupMonthlyMethodResponse(s *record.TopupMonthMethod) *response.TopupMonthMethodResponse

	// Converts multiple monthly method records into a slice of TopupMonthMethodResponse.
	ToTopupMonthlyMethodResponses(s []*record.TopupMonthMethod) []*response.TopupMonthMethodResponse

	// Converts a yearly method record into a TopupYearlyMethodResponse.
	ToTopupYearlyMethodResponse(s *record.TopupYearlyMethod) *response.TopupYearlyMethodResponse

	// Converts multiple yearly method records into a slice of TopupYearlyMethodResponse.
	ToTopupYearlyMethodResponses(s []*record.TopupYearlyMethod) []*response.TopupYearlyMethodResponse
}
