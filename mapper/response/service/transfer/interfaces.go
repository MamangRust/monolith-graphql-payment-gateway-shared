package transferresponsemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// TransferBaseResponseMapper provides methods to map TransferRecord domain models to TransferResponse API-compatible response types.
type TransferBaseResponseMapper interface {
	// Converts a single transfer record into a TransferResponse.
	ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse
}

// TransferQueryResponseMapper provides methods to map TransferRecord domain models to TransferResponse API-compatible response types.
type TransferQueryResponseMapper interface {
	TransferBaseResponseMapper

	// Converts multiple transfer records into a slice of TransferResponse.
	ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse

	// Converts multiple soft-deleted transfer records into a slice of TransferResponseDeleteAt.
	ToTransfersResponseDeleteAt(transfers []*record.TransferRecord) []*response.TransferResponseDeleteAt
}

// TransferCommandResponseMapper provides methods to map TransferRecord domain models to TransferResponse API-compatible response types.
type TransferCommandResponseMapper interface {
	TransferBaseResponseMapper

	// Converts a soft-deleted transfer record into a TransferResponseDeleteAt.
	ToTransferResponseDeleteAt(transfer *record.TransferRecord) *response.TransferResponseDeleteAt
}

type TransferStatsStatusResponseMapper interface {
	// Converts a monthly success status transfer record into a TransferResponseMonthStatusSuccess.
	ToTransferResponseMonthStatusSuccess(s *record.TransferRecordMonthStatusSuccess) *response.TransferResponseMonthStatusSuccess

	// Converts multiple monthly success status transfer records into a slice of TransferResponseMonthStatusSuccess.
	ToTransferResponsesMonthStatusSuccess(Transfers []*record.TransferRecordMonthStatusSuccess) []*response.TransferResponseMonthStatusSuccess

	// Converts a yearly success status transfer record into a TransferResponseYearStatusSuccess.
	ToTransferResponseYearStatusSuccess(s *record.TransferRecordYearStatusSuccess) *response.TransferResponseYearStatusSuccess

	// Converts multiple yearly success status transfer records into a slice of TransferResponseYearStatusSuccess.
	ToTransferResponsesYearStatusSuccess(Transfers []*record.TransferRecordYearStatusSuccess) []*response.TransferResponseYearStatusSuccess

	// Converts a monthly failed status transfer record into a TransferResponseMonthStatusFailed.
	ToTransferResponseMonthStatusFailed(s *record.TransferRecordMonthStatusFailed) *response.TransferResponseMonthStatusFailed

	// Converts multiple monthly failed status transfer records into a slice of TransferResponseMonthStatusFailed.
	ToTransferResponsesMonthStatusFailed(Transfers []*record.TransferRecordMonthStatusFailed) []*response.TransferResponseMonthStatusFailed

	// Converts a yearly failed status transfer record into a TransferResponseYearStatusFailed.
	ToTransferResponseYearStatusFailed(s *record.TransferRecordYearStatusFailed) *response.TransferResponseYearStatusFailed

	// Converts multiple yearly failed status transfer records into a slice of TransferResponseYearStatusFailed.
	ToTransferResponsesYearStatusFailed(Transfers []*record.TransferRecordYearStatusFailed) []*response.TransferResponseYearStatusFailed
}

// TransferAmountResponseMapper provides methods to map TransferAmountRecord domain models to TransferAmountResponse API-compatible response types.
type TransferAmountResponseMapper interface {
	// Converts a monthly transfer amount record into a TransferMonthAmountResponse.
	ToTransferResponseMonthAmount(s *record.TransferMonthAmount) *response.TransferMonthAmountResponse

	// Converts multiple monthly transfer amount records into a slice of TransferMonthAmountResponse.
	ToTransferResponsesMonthAmount(s []*record.TransferMonthAmount) []*response.TransferMonthAmountResponse

	// Converts a yearly transfer amount record into a TransferYearAmountResponse.
	ToTransferResponseYearAmount(s *record.TransferYearAmount) *response.TransferYearAmountResponse

	// Converts multiple yearly transfer amount records into a slice of TransferYearAmountResponse.
	ToTransferResponsesYearAmount(s []*record.TransferYearAmount) []*response.TransferYearAmountResponse
}
