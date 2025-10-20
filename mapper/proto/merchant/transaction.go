package merchantprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
)

type merchantTransactionProtoMapper struct {
}

func NewMerchantTransactionProtoMapper() MerchantTransactionProtoMapper {
	return &merchantTransactionProtoMapper{}
}

// ToProtoResponsePaginationMerchantTransaction maps a pagination meta, status, message and a list of *response.MerchantTransactionResponse
// to a *pb.ApiResponsePaginationMerchantTransaction proto response.
//
// It is used to generate the response for the MerchantService.ListMerchantTransaction rpc method.
func (m *merchantTransactionProtoMapper) ToProtoResponsePaginationMerchantTransaction(pagination *pbhelpers.PaginationMeta, status string, message string, merchants []*response.MerchantTransactionResponse) *pb.ApiResponsePaginationMerchantTransaction {

	return &pb.ApiResponsePaginationMerchantTransaction{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantTransactionResponses(merchants),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// mapMerchantTransactionResponse maps a *response.MerchantTransactionResponse to a *pb.MerchantTransactionResponse proto message.
//
// It is used to generate the response for the MerchantService.ListMerchantTransaction rpc method.
func (m *merchantTransactionProtoMapper) mapMerchantTransactionResponse(merchant *response.MerchantTransactionResponse) *pb.MerchantTransactionResponse {
	return &pb.MerchantTransactionResponse{
		Id:              int32(merchant.ID),
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantId:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

// mapMerchantTransactionResponses maps a list of *response.MerchantTransactionResponse to a list of
// *pb.MerchantTransactionResponse proto responses.
//
// It iterates over each MerchantTransactionResponse in the input slice, converting
// them to their protobuf equivalent using the mapMerchantTransactionResponse function.
// This function is used to generate the response data for merchant transaction-related RPC methods.
func (s *merchantTransactionProtoMapper) mapMerchantTransactionResponses(roles []*response.MerchantTransactionResponse) []*pb.MerchantTransactionResponse {
	var responseRoles []*pb.MerchantTransactionResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMerchantTransactionResponse(role))
	}

	return responseRoles
}
