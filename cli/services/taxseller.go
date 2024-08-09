package services

import (
	"ganhocapital/config"
	"ganhocapital/entity"
	"ganhocapital/services/request"
)

type sellingTaxCalculator struct {
}

func NewSellingTaxCalculator() TaxCalculator {
	return sellingTaxCalculator{}
}

func (calculator sellingTaxCalculator) Calculate(stock entity.Stock, order request.Order) float32 {
	if order.TotalCost() <= config.TaxFreeThreshold {
		return 0
	}

	if calculator.orderResultInLoss(stock, order) {
		return 0
	}

	orderProfit := calculateProfit(stock, order)
	if stock.IsInLoss() {
		newBalance := orderProfit + stock.Balance
		if newBalance <= 0 {
			return 0
		}
		return round(newBalance * config.TaxPercentage)
	}

	return round(orderProfit * config.TaxPercentage)
}

func (calculator sellingTaxCalculator) orderResultInLoss(stock entity.Stock, order request.Order) bool {
	profit := calculateProfit(stock, order)
	return profit < 0
}
