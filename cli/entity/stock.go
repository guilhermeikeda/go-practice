package entity

type Stock struct {
	Quantity     int
	AveragePrice float32

	Balance float32
}

func (s Stock) QuantityAsFloat() float32 {
	return float32(s.Quantity)
}

func (s Stock) IsAllowedToSell(quantity int) bool {
	return s.Quantity >= quantity
}

func (stock *Stock) DeductQuantity(quantity int) {
	resultQuantity := stock.Quantity - quantity
	if resultQuantity < 0 {
		return
	}

	stock.Quantity = resultQuantity
}

func (stock *Stock) UpdateBalance(profit float32) {
	if stock.Quantity == 0 {
		stock.Balance = 0
		return
	}

	stock.Balance += profit
}

func (stock *Stock) IsInLoss() bool {
	return stock.Balance < 0
}
