package services

import (
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"testing"

	"gotest.tools/assert"
)

func TestPurchaseTaxCalculator_Calculate(t *testing.T) {
	type TestCase struct {
		name  string
		stock entity.Stock
		order request.Order

		want float32
	}

	tests := []TestCase{
		{
			name: "should return 0 for purchases",
			order: request.Order{
				Quantity: 1,
				UnitCost: 19999,
			},

			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calculator := NewPurchaseTaxCalculator()
			tax := calculator.Calculate(test.stock, test.order)

			assert.Equal(t, test.want, tax)
		})
	}
}
