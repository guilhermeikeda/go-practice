package services

import (
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"ganhocapital/services/result"
	"testing"

	"gotest.tools/assert"
)

func TestStockPurchaserPurchase(t *testing.T) {
	type TestCase struct {
		name  string
		stock entity.Stock
		order request.Order

		want    result.StockPurchased
		wantErr error
	}

	tests := []TestCase{
		{
			name: "should return stock with quantity updated when stock is purchased for the same price",
			stock: entity.Stock{
				AveragePrice: 10,
				Balance:      0,
				Quantity:     100,
			},
			order: request.Order{
				UnitCost: 10,
				Quantity: 100,
			},

			want: result.StockPurchased{
				Stock: entity.Stock{
					Quantity:     200,
					AveragePrice: 10,
					Balance:      0,
				},
			},
			wantErr: nil,
		},
		{
			name: "should return stock with average price updated when stock is purchased for a different price",
			stock: entity.Stock{
				AveragePrice: 10,
				Balance:      0,
				Quantity:     100,
			},
			order: request.Order{
				UnitCost: 20,
				Quantity: 100,
			},

			want: result.StockPurchased{
				Stock: entity.Stock{
					Quantity:     200,
					AveragePrice: 15,
					Balance:      0,
				},
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			purchaser := NewStockPurchaseService()
			stockPurchased, err := purchaser.Purchase(test.stock, test.order)

			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, stockPurchased)
		})
	}

}
