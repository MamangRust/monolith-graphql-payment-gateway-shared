package merchantprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type merchantQueryProtoMapper struct {
}

func NewMerchantQueryProtoMapper() MerchantQueryProtoMapper {
	return &merchantQueryProtoMapper{}
}

// ToProtoResponsePaginationMerchant converts a list of MerchantResponse and pagination metadata
// into a protobuf ApiResponsePaginationMerchant. It includes the status and message for the API response.
// Parameters:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - merchants: A list of MerchantResponse objects to be included in the response.
//
// Returns:
//
//	A pointer to ApiResponsePaginationMerchant containing the status, message, merchant data, and pagination data.
func (m *merchantQueryProtoMapper) ToProtoResponsePaginationMerchant(pagination *pbhelpers.PaginationMeta, status string, message string, merchants []*response.MerchantResponse) *pb.ApiResponsePaginationMerchant {
	return &pb.ApiResponsePaginationMerchant{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantResponses(merchants),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationMerchantDeleteAt maps a pagination meta, status, message and a list of *response.MerchantResponseDeleteAt
// to a *pb.ApiResponsePaginationMerchantDeleteAt proto response.
//
// It is used to generate the response for the MerchantService.ListDeletedMerchant rpc method.
func (m *merchantQueryProtoMapper) ToProtoResponsePaginationMerchantDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, merchants []*response.MerchantResponseDeleteAt) *pb.ApiResponsePaginationMerchantDeleteAt {
	return &pb.ApiResponsePaginationMerchantDeleteAt{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantResponsesDeleteAt(merchants),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponseMerchant maps a status, message and a *response.MerchantResponse
// to a *pb.ApiResponseMerchant proto response.
//
// It is used to generate the response for the MerchantService.GetMerchant rpc method.
func (m *merchantQueryProtoMapper) ToProtoResponseMerchant(status string, message string, res *response.MerchantResponse) *pb.ApiResponseMerchant {
	return &pb.ApiResponseMerchant{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponse(res),
	}

}

// ToProtoResponseMerchants maps a status, message and a list of *response.MerchantResponse
// to a *pb.ApiResponsesMerchant proto response.
//
// It is used to generate the response for the MerchantService.ListMerchants rpc method.
func (m *merchantQueryProtoMapper) ToProtoResponseMerchants(status string, message string, res []*response.MerchantResponse) *pb.ApiResponsesMerchant {
	return &pb.ApiResponsesMerchant{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponses(res),
	}

}

// mapMerchantResponse maps a *response.MerchantResponse to a *pb.MerchantResponse proto message.
//
// It is used to generate the response for the MerchantService.GetMerchant rpc method.
func (m *merchantQueryProtoMapper) mapMerchantResponse(merchant *response.MerchantResponse) *pb.MerchantResponse {
	return &pb.MerchantResponse{
		Id:        int32(merchant.ID),
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

// mapMerchantResponses maps a list of *response.MerchantResponse to a list of
// *pb.MerchantResponse proto responses.
//
// It iterates over each MerchantResponse in the input slice, converting
// them to their protobuf equivalent using the mapMerchantResponse function.
// This function is used to generate the response data for merchant-related RPC methods.
func (s *merchantQueryProtoMapper) mapMerchantResponses(roles []*response.MerchantResponse) []*pb.MerchantResponse {
	var responseRoles []*pb.MerchantResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMerchantResponse(role))
	}

	return responseRoles
}

// mapMerchantResponseDeleteAt maps a *response.MerchantResponseDeleteAt to a *pb.MerchantResponseDeleteAt proto message.
//
// It is used to generate the response for the MerchantService.ListDeletedMerchant rpc method.
func (m *merchantQueryProtoMapper) mapMerchantResponseDeleteAt(merchant *response.MerchantResponseDeleteAt) *pb.MerchantResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if merchant.DeletedAt != nil {
		deletedAt = wrapperspb.String(*merchant.DeletedAt)
	}

	return &pb.MerchantResponseDeleteAt{
		Id:        int32(merchant.ID),
		Name:      merchant.Name,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

// mapMerchantResponsesDeleteAt maps a list of *response.MerchantResponseDeleteAt to a list of
// *pb.MerchantResponseDeleteAt proto responses.
//
// It iterates over each MerchantResponseDeleteAt in the input slice, converting
// them to their protobuf equivalent using the mapMerchantResponseDeleteAt function.
// This function is used to generate the response data for the MerchantService.ListDeletedMerchant rpc method.
func (s *merchantQueryProtoMapper) mapMerchantResponsesDeleteAt(roles []*response.MerchantResponseDeleteAt) []*pb.MerchantResponseDeleteAt {
	var responseRoles []*pb.MerchantResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMerchantResponseDeleteAt(role))
	}

	return responseRoles
}
