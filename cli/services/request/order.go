package request

type Order struct {
	UnitCost float32
	Quantity int
}

func (order Order) TotalCost() float32 {
	return float32(order.Quantity) * order.UnitCost
}
