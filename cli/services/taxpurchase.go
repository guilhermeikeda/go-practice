package services

import (
	"ganhocapital/entity"
	"ganhocapital/services/request"
)

type purchaseTaxCalculator struct {
}

func NewPurchaseTaxCalculator() TaxCalculator {
	return purchaseTaxCalculator{}
}

func (calculator purchaseTaxCalculator) Calculate(stock entity.Stock, order request.Order) float32 {
	return 0
}
