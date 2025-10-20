package topupprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type topupStatsStatusProtoMapper struct {
}

func NewTopupStatsStatusProtoMapper() TopupStatsStatusProtoMapper {
	return &topupStatsStatusProtoMapper{}
}

// ToProtoResponseTopupMonthStatusSuccess maps monthly successful top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupMonthStatusSuccess containing the status, message, and top-up data.
func (t *topupStatsStatusProtoMapper) ToProtoResponseTopupMonthStatusSuccess(status string, message string, s []*response.TopupResponseMonthStatusSuccess) *pb.ApiResponseTopupMonthStatusSuccess {
	return &pb.ApiResponseTopupMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthStatusSuccess(s),
	}
}

// ToProtoResponseTopupYearStatusSuccess maps yearly successful top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupYearStatusSuccess containing the status, message, and top-up data.
func (t *topupStatsStatusProtoMapper) ToProtoResponseTopupYearStatusSuccess(status string, message string, s []*response.TopupResponseYearStatusSuccess) *pb.ApiResponseTopupYearStatusSuccess {
	return &pb.ApiResponseTopupYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapTopupResponsesYearStatusSuccess(s),
	}
}

// ToProtoResponseTopupMonthStatusFailed maps monthly failed top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupMonthStatusFailed containing the status, message, and top-up data.
func (t *topupStatsStatusProtoMapper) ToProtoResponseTopupMonthStatusFailed(status string, message string, s []*response.TopupResponseMonthStatusFailed) *pb.ApiResponseTopupMonthStatusFailed {
	return &pb.ApiResponseTopupMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthStatusFailed(s),
	}
}

// ToProtoResponseTopupYearStatusFailed maps yearly failed top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseYearStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupYearStatusFailed containing the status, message, and top-up data.
func (t *topupStatsStatusProtoMapper) ToProtoResponseTopupYearStatusFailed(status string, message string, s []*response.TopupResponseYearStatusFailed) *pb.ApiResponseTopupYearStatusFailed {
	return &pb.ApiResponseTopupYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapTopupResponsesYearStatusFailed(s),
	}
}

// mapResponseTopupMonthStatusSuccess maps a TopupResponseMonthStatusSuccess to a TopupMonthStatusSuccessResponse.
//
// Args:
//   - s: A pointer to TopupResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthStatusSuccessResponse containing the mapped top-up data.
func (t *topupStatsStatusProtoMapper) mapResponseTopupMonthStatusSuccess(s *response.TopupResponseMonthStatusSuccess) *pb.TopupMonthStatusSuccessResponse {
	return &pb.TopupMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusSuccess maps a slice of TopupResponseMonthStatusSuccess to a slice of
// TopupMonthStatusSuccessResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A slice of TopupMonthStatusSuccessResponse containing the mapped monthly top-up success data.
func (t *topupStatsStatusProtoMapper) mapResponsesTopupMonthStatusSuccess(topups []*response.TopupResponseMonthStatusSuccess) []*pb.TopupMonthStatusSuccessResponse {
	var topupRecords []*pb.TopupMonthStatusSuccessResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusSuccess(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusSuccess maps a TopupResponseYearStatusSuccess to a TopupYearStatusSuccessResponse.
//
// Args:
//   - s: A pointer to TopupResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearStatusSuccessResponse containing the mapped yearly top-up success data.
func (t *topupStatsStatusProtoMapper) mapTopupResponseYearStatusSuccess(s *response.TopupResponseYearStatusSuccess) *pb.TopupYearStatusSuccessResponse {
	return &pb.TopupYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusSuccess maps a slice of TopupResponseYearStatusSuccess to a slice of
// TopupYearStatusSuccessResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A slice of TopupYearStatusSuccessResponse containing the mapped yearly top-up success data.
func (t *topupStatsStatusProtoMapper) mapTopupResponsesYearStatusSuccess(topups []*response.TopupResponseYearStatusSuccess) []*pb.TopupYearStatusSuccessResponse {
	var topupRecords []*pb.TopupYearStatusSuccessResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusSuccess(topup))
	}

	return topupRecords
}

// mapResponseTopupMonthStatusFailed maps a TopupResponseMonthStatusFailed to a TopupMonthStatusFailedResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthStatusFailedResponse containing the mapped monthly top-up failure data.
func (t *topupStatsStatusProtoMapper) mapResponseTopupMonthStatusFailed(s *response.TopupResponseMonthStatusFailed) *pb.TopupMonthStatusFailedResponse {
	return &pb.TopupMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusFailed converts a slice of TopupResponseMonthStatusFailed
// into a slice of TopupMonthStatusFailedResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseMonthStatusFailed to be converted.
//
// Returns:
//   - A slice of TopupMonthStatusFailedResponse containing the mapped data
//     for each monthly top-up failure entry.
func (t *topupStatsStatusProtoMapper) mapResponsesTopupMonthStatusFailed(topups []*response.TopupResponseMonthStatusFailed) []*pb.TopupMonthStatusFailedResponse {
	var topupRecords []*pb.TopupMonthStatusFailedResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusFailed(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusFailed maps a TopupResponseYearStatusFailed to a TopupYearStatusFailedResponse.
//
// Args:
//   - s: A pointer to TopupResponseYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearStatusFailedResponse containing the mapped yearly top-up failure data.
func (t *topupStatsStatusProtoMapper) mapTopupResponseYearStatusFailed(s *response.TopupResponseYearStatusFailed) *pb.TopupYearStatusFailedResponse {
	return &pb.TopupYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusFailed maps a slice of TopupResponseYearStatusFailed
// into a slice of TopupYearStatusFailedResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseYearStatusFailed to be converted.
//
// Returns:
//   - A slice of TopupYearStatusFailedResponse containing the mapped yearly top-up failure data
//     for each yearly top-up failure entry.
func (t *topupStatsStatusProtoMapper) mapTopupResponsesYearStatusFailed(topups []*response.TopupResponseYearStatusFailed) []*pb.TopupYearStatusFailedResponse {
	var topupRecords []*pb.TopupYearStatusFailedResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusFailed(topup))
	}

	return topupRecords
}
