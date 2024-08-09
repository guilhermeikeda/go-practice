package services

import (
	"errors"
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"ganhocapital/services/result"
)

var (
	ErrSellOrderQuantityInvalid = errors.New("invalid sell order quantity")
)

type stockSeller interface {
	Sell(entity.Stock, request.Order) (result.StockSold, error)
}

type stockSellerImpl struct {
}

func NewStockSellerService() stockSeller {
	return stockSellerImpl{}
}

func (s stockSellerImpl) Sell(stock entity.Stock, order request.Order) (result.StockSold, error) {
	var (
		orderQuantity = order.Quantity
	)

	if !stock.IsAllowedToSell(orderQuantity) {
		return result.StockSold{}, ErrSellOrderQuantityInvalid
	}

	stock.DeductQuantity(orderQuantity)

	profit := calculateProfit(stock, order)
	stock.UpdateBalance(profit)

	result := result.StockSold{
		Stock: stock,
	}
	return result, nil
}

func calculateProfit(stock entity.Stock, order request.Order) float32 {
	var (
		orderQuantity = order.Quantity
		orderUnitCost = order.UnitCost
	)

	totalSellPrice := float32(orderQuantity) * orderUnitCost
	stockAveragePriceForOder := float32(orderQuantity) * stock.AveragePrice
	profit := totalSellPrice - stockAveragePriceForOder

	return profit
}
