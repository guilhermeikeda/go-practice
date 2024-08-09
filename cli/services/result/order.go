package result

import "ganhocapital/entity"

type Order struct {
	Stock   entity.Stock
	Charges OrderCharge
}

type OrderCharge struct {
	Tax float32
}
