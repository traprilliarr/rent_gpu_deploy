package converter

import (
	"rent_gpu_be/internal/entity"
	"rent_gpu_be/internal/model"
)

func OrderToResponses(order *[]entity.Order) []*model.OrderResponse {
	orderResponses := make([]*model.OrderResponse, 0, len(*order))
	for _, order := range *order {
		orderResponse := &model.OrderResponse{
			ID: order.ID,
		}
		orderResponses = append(orderResponses, orderResponse)
	}
	return orderResponses
}

func OrderToResponse(order *entity.Order) *model.OrderResponse {
	response := model.OrderResponse{
		ID:                 order.ID,
		OrderName:          order.ID,
		PaymentAddress:     order.PaymentAddress,
		PaymentTransaction: order.Hash,
		Value:              order.Value,
		Status:             order.Status,
	}
	return &response
}
