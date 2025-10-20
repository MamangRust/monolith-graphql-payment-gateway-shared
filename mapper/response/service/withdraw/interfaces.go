package withdrawresponseservice

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type WithdrawBaseResponseMapper interface {
	// Converts a single WithdrawRecord into a WithdrawResponse.
	ToWithdrawResponse(withdraw *record.WithdrawRecord) *response.WithdrawResponse
}

type WithdrawQueryResponseMapper interface {
	WithdrawBaseResponseMapper
	// Converts multiple WithdrawRecords into a slice of WithdrawResponse.
	ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse

	// Converts multiple soft-deleted WithdrawRecords into a slice of WithdrawResponseDeleteAt.
	ToWithdrawsResponseDeleteAt(withdraws []*record.WithdrawRecord) []*response.WithdrawResponseDeleteAt
}

type WithdrawCommandResponseMapper interface {
	WithdrawBaseResponseMapper

	// Converts a soft-deleted WithdrawRecord into a WithdrawResponseDeleteAt.
	ToWithdrawResponseDeleteAt(withdraw *record.WithdrawRecord) *response.WithdrawResponseDeleteAt
}

type WithdrawStatsStatusResponseMapper interface {
	// Converts a monthly success-status WithdrawRecord into a WithdrawResponseMonthStatusSuccess.
	ToWithdrawResponseMonthStatusSuccess(s *record.WithdrawRecordMonthStatusSuccess) *response.WithdrawResponseMonthStatusSuccess

	// Converts multiple monthly success-status WithdrawRecords into a slice of WithdrawResponseMonthStatusSuccess.
	ToWithdrawResponsesMonthStatusSuccess(Withdraws []*record.WithdrawRecordMonthStatusSuccess) []*response.WithdrawResponseMonthStatusSuccess

	// Converts a yearly success-status WithdrawRecord into a WithdrawResponseYearStatusSuccess.
	ToWithdrawResponseYearStatusSuccess(s *record.WithdrawRecordYearStatusSuccess) *response.WithdrawResponseYearStatusSuccess

	// Converts multiple yearly success-status WithdrawRecords into a slice of WithdrawResponseYearStatusSuccess.
	ToWithdrawResponsesYearStatusSuccess(Withdraws []*record.WithdrawRecordYearStatusSuccess) []*response.WithdrawResponseYearStatusSuccess

	// Converts a monthly failed-status WithdrawRecord into a WithdrawResponseMonthStatusFailed.
	ToWithdrawResponseMonthStatusFailed(s *record.WithdrawRecordMonthStatusFailed) *response.WithdrawResponseMonthStatusFailed

	// Converts multiple monthly failed-status WithdrawRecords into a slice of WithdrawResponseMonthStatusFailed.
	ToWithdrawResponsesMonthStatusFailed(Withdraws []*record.WithdrawRecordMonthStatusFailed) []*response.WithdrawResponseMonthStatusFailed

	// Converts a yearly failed-status WithdrawRecord into a WithdrawResponseYearStatusFailed.
	ToWithdrawResponseYearStatusFailed(s *record.WithdrawRecordYearStatusFailed) *response.WithdrawResponseYearStatusFailed

	// Converts multiple yearly failed-status WithdrawRecords into a slice of WithdrawResponseYearStatusFailed.
	ToWithdrawResponsesYearStatusFailed(Withdraws []*record.WithdrawRecordYearStatusFailed) []*response.WithdrawResponseYearStatusFailed
}

type WithdrawStatsAmountResponseMapper interface {
	// Converts a monthly amount summary into a WithdrawMonthlyAmountResponse.
	ToWithdrawAmountMonthlyResponse(s *record.WithdrawMonthlyAmount) *response.WithdrawMonthlyAmountResponse

	// Converts multiple monthly amount summaries into a slice of WithdrawMonthlyAmountResponse.
	ToWithdrawsAmountMonthlyResponses(s []*record.WithdrawMonthlyAmount) []*response.WithdrawMonthlyAmountResponse

	// Converts a yearly amount summary into a WithdrawYearlyAmountResponse.
	ToWithdrawAmountYearlyResponse(s *record.WithdrawYearlyAmount) *response.WithdrawYearlyAmountResponse

	// Converts multiple yearly amount summaries into a slice of WithdrawYearlyAmountResponse.
	ToWithdrawsAmountYearlyResponses(s []*record.WithdrawYearlyAmount) []*response.WithdrawYearlyAmountResponse
}
