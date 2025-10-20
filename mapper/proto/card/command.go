package cardprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
)

type cardCommandProtoMapper struct{}

func NewCardCommandProtoMapper() CardCommandProtoMapper {
	return &cardCommandProtoMapper{}
}

// ToProtoResponseCard maps a *response.CardResponse to a *pb.ApiResponseCard proto response.
//
// It is used to generate the response for the CardService.GetCard rpc method.
func (s *cardCommandProtoMapper) ToProtoResponseCard(status string, message string, card *response.CardResponse) *pb.ApiResponseCard {
	return &pb.ApiResponseCard{
		Status:  status,
		Message: message,
		Data:    s.mapCardResponse(card),
	}
}

func (s *cardCommandProtoMapper) ToProtoResponseCardDeleteAt(status string, message string, card *response.CardResponseDeleteAt) *pb.ApiResponseCardDeleteAt {
	return &pb.ApiResponseCardDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapCardResponseDeleteAt(card),
	}
}

// ToProtoResponseCardDeleteAt maps a status and message to a *pb.ApiResponseCardDelete proto response.
//
// It is used to generate the response for the CardService.DeleteCard rpc method.
func (s *cardCommandProtoMapper) ToProtoResponseCardDelete(status string, message string) *pb.ApiResponseCardDelete {
	return &pb.ApiResponseCardDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseCardAll maps a status and message to a *pb.ApiResponseCardAll proto response.
//
// It is used to generate the response for the CardService.GetAllCard rpc method.
func (s *cardCommandProtoMapper) ToProtoResponseCardAll(status string, message string) *pb.ApiResponseCardAll {
	return &pb.ApiResponseCardAll{
		Status:  status,
		Message: message,
	}
}

// mapCardResponse maps a *response.CardResponse to a *pb.CardResponse proto response.
//
// It is used to generate the response for the CardService.GetCard rpc method.
func (s *cardCommandProtoMapper) mapCardResponse(card *response.CardResponse) *pb.CardResponse {
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

func (s *cardCommandProtoMapper) mapCardResponseDeleteAt(card *response.CardResponseDeleteAt) *pb.CardResponseDeleteAt {
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
