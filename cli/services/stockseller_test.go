package services

import (
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"ganhocapital/services/result"
	"testing"

	"gotest.tools/assert"
)

func TestStockSellerImplSell(t *testing.T) {
	type TestCase struct {
		name  string
		stock entity.Stock
		order request.Order

		want    result.StockSold
		wantErr error
	}

	tests := []TestCase{
		{
			name: "should return ErrSellOrderQuantityInvalid when attemp to sell more stocks than available",
			stock: entity.Stock{
				Quantity: 99,
			},
			order: request.Order{
				UnitCost: 10,
				Quantity: 100,
			},

			want:    result.StockSold{},
			wantErr: ErrSellOrderQuantityInvalid,
		},
		{
			name: "should stock with quantity updated when stock is sold",
			stock: entity.Stock{
				AveragePrice: 20,
				Quantity:     200,
				Balance:      0,
			},
			order: request.Order{
				UnitCost: 20,
				Quantity: 100,
			},

			want: result.StockSold{
				Stock: entity.Stock{
					AveragePrice: 20,
					Quantity:     100,
					Balance:      0,
				},
			},
			wantErr: nil,
		},
		{
			name: "should return stock with updated balance when stock is sold with profit",
			stock: entity.Stock{
				AveragePrice: 20,
				Quantity:     200,
				Balance:      0,
			},
			order: request.Order{
				UnitCost: 40,
				Quantity: 100,
			},

			want: result.StockSold{
				Stock: entity.Stock{
					AveragePrice: 20,
					Quantity:     100,
					Balance:      2000,
				},
			},
			wantErr: nil,
		},
		{
			name: "should return stock with updated balance when stock is sold with loss",
			stock: entity.Stock{
				AveragePrice: 20,
				Quantity:     200,
				Balance:      0,
			},
			order: request.Order{
				UnitCost: 10,
				Quantity: 100,
			},

			want: result.StockSold{
				Stock: entity.Stock{
					AveragePrice: 20,
					Quantity:     100,
					Balance:      -1000,
				},
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			seller := NewStockSellerService()
			stockSold, err := seller.Sell(test.stock, test.order)

			assert.Equal(t, test.wantErr, err)
			assert.DeepEqual(t, test.want, stockSold)
		})
	}

}
