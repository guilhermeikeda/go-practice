package services

import (
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"testing"

	"gotest.tools/assert"
)

func TestSellingTaxCalculator_Calculate(t *testing.T) {
	type TestCase struct {
		name  string
		stock entity.Stock
		order request.Order

		want float32
	}

	tests := []TestCase{
		{
			name: "should return 0 when order total cost is lower than 20.000",
			order: request.Order{
				Quantity: 1,
				UnitCost: 19999,
			},

			want: 0,
		},
		{
			name: "should return 0 when selling order result in loss",
			order: request.Order{
				Quantity: 1,
				UnitCost: 30000,
			},
			stock: entity.Stock{
				Quantity:     1,
				AveragePrice: 40000,
			},

			want: 0,
		},
		{
			name: "should return 0 when order profit equals stock losses",
			order: request.Order{
				Quantity: 1,
				UnitCost: 80000,
			},
			stock: entity.Stock{
				Quantity:     1,
				AveragePrice: 40000,
				Balance:      -40000,
			},

			want: 0,
		},
		{
			name: "should return 20% of order profit when stock does not have previous losses",
			order: request.Order{
				Quantity: 1,
				UnitCost: 80000,
			},
			stock: entity.Stock{
				Quantity:     1,
				AveragePrice: 40000,
				Balance:      0,
			},

			want: 8000,
		},
		{
			name: "should return 20% of order profit discounting stock previous losses",
			order: request.Order{
				Quantity: 1,
				UnitCost: 80000,
			},
			stock: entity.Stock{
				Quantity:     1,
				AveragePrice: 40000,
				Balance:      -997,
			},

			want: 7800.60,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calculator := NewSellingTaxCalculator()
			tax := calculator.Calculate(test.stock, test.order)

			assert.Equal(t, test.want, tax)
		})
	}
}
