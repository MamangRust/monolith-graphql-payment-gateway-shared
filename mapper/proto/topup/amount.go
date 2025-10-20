package topupprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type topupStatsAmountProtoMapper struct {
}

func NewTopupStatsAmountProtoMapper() TopupStatsAmountProtoMapper {
	return &topupStatsAmountProtoMapper{}
}

// ToProtoResponseTopupMonthAmount maps monthly top-up amounts to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to ApiResponseTopupMonthAmount containing the status, message, and mapped top-up amount data.
func (t *topupStatsAmountProtoMapper) ToProtoResponseTopupMonthAmount(status string, message string, s []*response.TopupMonthAmountResponse) *pb.ApiResponseTopupMonthAmount {
	return &pb.ApiResponseTopupMonthAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupMonthlyAmounts(s),
	}
}

// ToProtoResponseTopupYearAmount maps yearly top-up amounts to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to ApiResponseTopupYearAmount containing the status, message, and mapped top-up amount data.
func (t *topupStatsAmountProtoMapper) ToProtoResponseTopupYearAmount(status string, message string, s []*response.TopupYearlyAmountResponse) *pb.ApiResponseTopupYearAmount {
	return &pb.ApiResponseTopupYearAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupYearlyAmounts(s),
	}
}

// mapResponseTopupMonthlyAmount maps a TopupMonthAmountResponse to a TopupMonthAmountResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthAmountResponse containing the mapped monthly top-up amount data.
func (t *topupStatsAmountProtoMapper) mapResponseTopupMonthlyAmount(s *response.TopupMonthAmountResponse) *pb.TopupMonthAmountResponse {
	return &pb.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyAmounts maps a slice of TopupMonthAmountResponse to a slice of
// TopupMonthAmountResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupMonthAmountResponse containing the mapped monthly top-up amount data.
func (t *topupStatsAmountProtoMapper) mapResponseTopupMonthlyAmounts(s []*response.TopupMonthAmountResponse) []*pb.TopupMonthAmountResponse {
	var topupProtos []*pb.TopupMonthAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyAmount(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyAmount maps a TopupYearlyAmountResponse to a TopupYearlyAmountResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearlyAmountResponse containing the mapped yearly top-up amount data.
func (t *topupStatsAmountProtoMapper) mapResponseTopupYearlyAmount(s *response.TopupYearlyAmountResponse) *pb.TopupYearlyAmountResponse {
	return &pb.TopupYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupYearlyAmounts maps a slice of TopupYearlyAmountResponse to a slice of
// TopupYearlyAmountResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupYearlyAmountResponse containing the mapped yearly top-up amount data.
func (t *topupStatsAmountProtoMapper) mapResponseTopupYearlyAmounts(s []*response.TopupYearlyAmountResponse) []*pb.TopupYearlyAmountResponse {
	var topupProtos []*pb.TopupYearlyAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyAmount(topup))
	}
	return topupProtos
}
