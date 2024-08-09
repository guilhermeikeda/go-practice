package input

type OperationFixture struct {
	operations []Operation
}

func NewOperationFixture() OperationFixture {
	return OperationFixture{
		operations: []Operation{},
	}
}

func (fixture OperationFixture) AddBuyOperation(quantity int, unitCost float32) OperationFixture {
	fixture.operations = append(fixture.operations, Operation{
		Type:     "buy",
		UnitCost: unitCost,
		Quantity: quantity,
	})

	return fixture
}

func (fixture OperationFixture) AddSellOperation(quantity int, unitCost float32) OperationFixture {
	fixture.operations = append(fixture.operations, Operation{
		Type:     "sell",
		UnitCost: unitCost,
		Quantity: quantity,
	})

	return fixture
}

func (fixture OperationFixture) Build() []Operation {
	return fixture.operations
}
