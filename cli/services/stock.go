package services

import (
	"errors"
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"ganhocapital/services/result"
)

type StockService interface {
	Buy(stock entity.Stock, order request.Order) (result.Order, error)
	Sell(stock entity.Stock, order request.Order) (result.Order, error)
}

type stockService struct {
	stockPurchaser stockPurchaser
	stockSeller    stockSeller

	purchaseTaxCalculator TaxCalculator
	sellingTaxCalculator  TaxCalculator
}

func NewStockService() StockService {
	return &stockService{
		stockPurchaser: NewStockPurchaseService(),
		stockSeller:    NewStockSellerService(),

		purchaseTaxCalculator: NewPurchaseTaxCalculator(),
		sellingTaxCalculator:  NewSellingTaxCalculator(),
	}
}

func (service stockService) Buy(stock entity.Stock, order request.Order) (result.Order, error) {
	tax := service.purchaseTaxCalculator.Calculate(stock, order)

	purchaseResult, err := service.stockPurchaser.Purchase(stock, order)
	if err != nil {
		return result.Order{}, errors.New("failed to purchase stock")
	}

	result := result.Order{
		Stock: purchaseResult.Stock,
		Charges: result.OrderCharge{
			Tax: tax,
		},
	}
	return result, nil
}

func (service stockService) Sell(stock entity.Stock, order request.Order) (result.Order, error) {
	tax := service.sellingTaxCalculator.Calculate(stock, order)

	purchaseResult, err := service.stockSeller.Sell(stock, order)
	if err != nil {
		return result.Order{}, errors.New("failed to sell stock")
	}

	result := result.Order{
		Stock: purchaseResult.Stock,
		Charges: result.OrderCharge{
			Tax: tax,
		},
	}
	return result, nil
}
