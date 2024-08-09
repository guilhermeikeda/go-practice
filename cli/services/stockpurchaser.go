package services

import (
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"ganhocapital/services/result"
)

type stockPurchaser interface {
	Purchase(entity.Stock, request.Order) (result.StockPurchased, error)
}

type stockPurchaserImpl struct {
}

func NewStockPurchaseService() stockPurchaser {
	return stockPurchaserImpl{}
}

func (s stockPurchaserImpl) Purchase(stock entity.Stock, order request.Order) (result.StockPurchased, error) {
	var (
		operationQuantity = float32(order.Quantity)
		currentQuantity   = float32(stock.Quantity)

		orderUnitCost = order.UnitCost
		orderQuantity = order.Quantity
	)

	stock.AveragePrice = ((stock.AveragePrice * stock.QuantityAsFloat()) + (orderUnitCost * operationQuantity)) / (currentQuantity + operationQuantity)
	stock.Quantity += orderQuantity

	result := result.StockPurchased{
		Stock: stock,
	}
	return result, nil
}
