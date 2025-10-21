package cardprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/card"
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
)

type cardQueryProtoMapper struct{}

func NewCardQueryProtoMapper() CardQueryProtoMapper {
	return &cardQueryProtoMapper{}
}

// ToProtoResponseCard maps a *response.CardResponse to a *pb.ApiResponseCard proto response.
//
// It is used to generate the response for the CardService.GetCard rpc method.
func (s *cardQueryProtoMapper) ToProtoResponseCard(status string, message string, card *response.CardResponse) *pb.ApiResponseCard {
	return &pb.ApiResponseCard{
		Status:  status,
		Message: message,
		Data:    s.mapCardResponse(card),
	}
}

// ToProtoResponsePaginationCard maps a pagination meta, status, message and a list of *response.CardResponse
// to a *pb.ApiResponsePaginationCard proto response.
//
// It is used to generate the response for the CardService.ListCard rpc method.
func (s *cardQueryProtoMapper) ToProtoResponsePaginationCard(pagination *pbhelpers.PaginationMeta, status string, message string, cards []*response.CardResponse) *pb.ApiResponsePaginationCard {
	return &pb.ApiResponsePaginationCard{
		Status:         status,
		Message:        message,
		Data:           s.mapCardResponses(cards),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationCardDeletedAt maps a pagination meta, status, message and a list of *response.CardResponseDeleteAt
// to a *pb.ApiResponsePaginationCardDeleteAt proto response.
//
// It is used to generate the response for the CardService.ListDeletedCard rpc method.
func (s *cardQueryProtoMapper) ToProtoResponsePaginationCardDeletedAt(pagination *pbhelpers.PaginationMeta, status string, message string, cards []*response.CardResponseDeleteAt) *pb.ApiResponsePaginationCardDeleteAt {
	return &pb.ApiResponsePaginationCardDeleteAt{
		Status:         status,
		Message:        message,
		Data:           s.mapCardResponsesDeleteAt(cards),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// mapCardResponse maps a *response.CardResponse to a *pb.CardResponse proto response.
//
// It is used to generate the response for the CardService.GetCard rpc method.
func (s *cardQueryProtoMapper) mapCardResponse(card *response.CardResponse) *pb.CardResponse {
	return &pb.CardResponse{
		Id:           int32(card.ID),
		UserId:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

// mapCardResponses maps a list of *response.CardResponse to a list of *pb.CardResponse
// proto response.
//
// It is used to generate the response for the CardService.ListCard rpc method.
func (s *cardQueryProtoMapper) mapCardResponses(roles []*response.CardResponse) []*pb.CardResponse {
	var responseRoles []*pb.CardResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapCardResponse(role))
	}

	return responseRoles
}

// mapCardResponseDeleteAt maps a *response.CardResponseDeleteAt to a *pb.CardResponseDeleteAt proto response.
//
// This function converts a CardResponseDeleteAt, which includes deletion information, into its protobuf
// representation. It handles the conversion of all fields including a nullable DeletedAt field.
func (s *cardQueryProtoMapper) mapCardResponseDeleteAt(card *response.CardResponseDeleteAt) *pb.CardResponseDeleteAt {
	res := &pb.CardResponseDeleteAt{
		Id:           int32(card.ID),
		UserId:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}

	if card.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(card.DeletedAt)
	}

	return res
}

// mapCardResponsesDeleteAt maps a list of *response.CardResponseDeleteAt to a list of
// *pb.CardResponseDeleteAt proto response.
//
// It is used to generate the response for the CardService.ListDeletedCard rpc method.
func (s *cardQueryProtoMapper) mapCardResponsesDeleteAt(roles []*response.CardResponseDeleteAt) []*pb.CardResponseDeleteAt {
	var responseRoles []*pb.CardResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapCardResponseDeleteAt(role))
	}

	return responseRoles
}
