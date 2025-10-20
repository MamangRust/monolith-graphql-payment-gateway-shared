package merchantprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
)

type merchantCommandProtoMapper struct{}

func NewMerchantCommandProtoMapper() MerchantCommandProtoMapper {
	return &merchantCommandProtoMapper{}
}

// ToProtoResponseMerchant maps a status, message and a *response.MerchantResponse
// to a *pb.ApiResponseMerchant proto response.
//
// It is used to generate the response for the MerchantService.GetMerchant rpc method.
func (m *merchantCommandProtoMapper) ToProtoResponseMerchant(status string, message string, res *response.MerchantResponse) *pb.ApiResponseMerchant {
	return &pb.ApiResponseMerchant{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponse(res),
	}
}

func (m *merchantCommandProtoMapper) ToProtoResponseMerchantDeleteAt(status string, message string, res *response.MerchantResponseDeleteAt) *pb.ApiResponseMerchantDeleteAt {
	return &pb.ApiResponseMerchantDeleteAt{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponseDeleteAt(res),
	}
}

// ToProtoResponseMerchantAll maps a status and message to its protobuf representation
// specifically for bulk merchant operations.
// It includes the status and message of the API response.
// Parameters:
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//
// Returns:
//
//	A pointer to ApiResponseMerchantAll containing the status and message.
func (m *merchantCommandProtoMapper) ToProtoResponseMerchantAll(status string, message string) *pb.ApiResponseMerchantAll {
	return &pb.ApiResponseMerchantAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseMerchantDelete converts a status and message to its protobuf representation
// specifically for merchant deletion operations.
// It includes the status and message of the API response.
// Parameters:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseMerchantDelete containing the status and message.
func (m *merchantCommandProtoMapper) ToProtoResponseMerchantDelete(status string, message string) *pb.ApiResponseMerchantDelete {
	return &pb.ApiResponseMerchantDelete{
		Status:  status,
		Message: message,
	}
}

// mapMerchantResponse maps a *response.MerchantResponse to a *pb.MerchantResponse proto message.
//
// It is used to generate the response for the MerchantService.GetMerchant rpc method.
func (m *merchantCommandProtoMapper) mapMerchantResponse(merchant *response.MerchantResponse) *pb.MerchantResponse {
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

func (m *merchantCommandProtoMapper) mapMerchantResponseDeleteAt(merchant *response.MerchantResponseDeleteAt) *pb.MerchantResponseDeleteAt {
	res := &pb.MerchantResponseDeleteAt{
		Id:        int32(merchant.ID),
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}

	if merchant.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(merchant.DeletedAt)
	}

	return res
}
