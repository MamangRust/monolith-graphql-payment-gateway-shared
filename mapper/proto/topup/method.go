package topupprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type topupStatsMethodProtoMapper struct{}

func NewTopupStatsMethodProtoMapper() TopupStatsMethodProtoMapper {
	return &topupStatsMethodProtoMapper{}
}

// ToProtoResponseTopupMonthMethod maps monthly top-up methods to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupMonthMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupMonthMethod containing the status, message, and top-up data.
func (t *topupStatsMethodProtoMapper) ToProtoResponseTopupMonthMethod(status string, message string, s []*response.TopupMonthMethodResponse) *pb.ApiResponseTopupMonthMethod {
	return &pb.ApiResponseTopupMonthMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupMonthlyMethods(s),
	}
}

// ToProtoResponseTopupYearMethod maps yearly top-up methods to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupYearlyMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupYearMethod containing the status, message, and top-up data.
func (t *topupStatsMethodProtoMapper) ToProtoResponseTopupYearMethod(status string, message string, s []*response.TopupYearlyMethodResponse) *pb.ApiResponseTopupYearMethod {
	return &pb.ApiResponseTopupYearMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupYearlyMethods(s),
	}
}

// mapResponseTopupMonthlyMethod maps a TopupMonthMethodResponse to a TopupMonthMethodResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupMonthMethodResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthMethodResponse containing the mapped monthly top-up method data.
func (t *topupStatsMethodProtoMapper) mapResponseTopupMonthlyMethod(s *response.TopupMonthMethodResponse) *pb.TopupMonthMethodResponse {
	return &pb.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyMethods maps a slice of TopupMonthMethodResponse to a slice of
// TopupMonthMethodResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupMonthMethodResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupMonthMethodResponse containing the mapped monthly top-up method data.
func (t *topupStatsMethodProtoMapper) mapResponseTopupMonthlyMethods(s []*response.TopupMonthMethodResponse) []*pb.TopupMonthMethodResponse {
	var topupProtos []*pb.TopupMonthMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyMethod(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyMethod maps a TopupYearlyMethodResponse to a TopupYearlyMethodResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupYearlyMethodResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearlyMethodResponse containing the mapped yearly top-up method data.
func (t *topupStatsMethodProtoMapper) mapResponseTopupYearlyMethod(s *response.TopupYearlyMethodResponse) *pb.TopupYearlyMethodResponse {
	return &pb.TopupYearlyMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupYearlyMethods maps a slice of TopupYearlyMethodResponse to a slice of
// TopupYearlyMethodResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupYearlyMethodResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupYearlyMethodResponse containing the mapped yearly top-up method data.
func (t *topupStatsMethodProtoMapper) mapResponseTopupYearlyMethods(s []*response.TopupYearlyMethodResponse) []*pb.TopupYearlyMethodResponse {
	var topupProtos []*pb.TopupYearlyMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyMethod(topup))
	}
	return topupProtos
}
