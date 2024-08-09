package input

type Operation struct {
	Type     string  `json:"operation"`
	UnitCost float32 `json:"unit-cost"`
	Quantity int     `json:"quantity"`
}

func (operation Operation) IsBuyOperation() bool {
	return operation.Type == "buy"
}

func (operation Operation) IsSellOperation() bool {
	return operation.Type == "sell"
}

func (operation Operation) IsTypeValid() bool {
	return (operation.IsBuyOperation() || operation.IsSellOperation())
}

func (operation Operation) IsUnitCostValid() bool {
	return operation.UnitCost > 0
}

func (operation Operation) IsQuantityValid() bool {
	return operation.Quantity > 0
}

func (operation Operation) IsValid() bool {
	return operation.IsTypeValid() &&
		operation.IsUnitCostValid() &&
		operation.IsQuantityValid()
}
