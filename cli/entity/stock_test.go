package entity

import (
	"testing"

	"gotest.tools/assert"
)

func TestStockIsAllowedToSell(t *testing.T) {
	type TestCase struct {
		name     string
		quantity int
		stock    Stock

		want bool
	}

	tests := []TestCase{
		{
			name:     "should false when stock is not allowed to be sold",
			quantity: 999,
			stock: Stock{
				Quantity: 10,
			},

			want: false,
		},
		{
			name:     "should true when stock is allowed to be sold",
			quantity: 999,
			stock: Stock{
				Quantity: 999,
			},

			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.stock.IsAllowedToSell(test.quantity)
			assert.Equal(t, test.want, result)
		})
	}
}

func TestStockDeductQuantity(t *testing.T) {
	type TestCase struct {
		name     string
		quantity int
		stock    Stock

		want Stock
	}

	tests := []TestCase{
		{
			name:     "should not deduct quantity that is not allowed to be deducted",
			quantity: 999,
			stock: Stock{
				Quantity: 10,
			},

			want: Stock{
				Quantity: 10,
			},
		},
		{
			name:     "should deduct quantity that is allowed to be deducted",
			quantity: 999,
			stock: Stock{
				Quantity: 1000,
			},

			want: Stock{
				Quantity: 1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stock.DeductQuantity(test.quantity)
			assert.Equal(t, test.want, test.stock)
		})
	}
}

func TestStockUpdateBalance(t *testing.T) {
	type TestCase struct {
		name   string
		profit float32
		stock  Stock

		want Stock
	}

	tests := []TestCase{
		{
			name:   "should ignore balance update and set to 0 when stock quantity is 0",
			profit: 1000,
			stock: Stock{
				Quantity: 0,
				Balance:  1000,
			},

			want: Stock{
				Quantity: 0,
				Balance:  0,
			},
		},
		{
			name:   "should update balance when stock quantity is greter than 0",
			profit: 1000,
			stock: Stock{
				Quantity: 1,
				Balance:  1000,
			},

			want: Stock{
				Quantity: 1,
				Balance:  2000,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stock.UpdateBalance(test.profit)
			assert.Equal(t, test.want, test.stock)
		})
	}
}

func TestStockIsInLoss(t *testing.T) {
	type TestCase struct {
		name  string
		stock Stock

		want bool
	}

	tests := []TestCase{
		{
			name: "should true when stock is in loss",
			stock: Stock{
				Quantity: 10,
				Balance:  -1,
			},

			want: true,
		},
		{
			name: "should false when stock is not making either loss or profit",
			stock: Stock{
				Quantity: 10,
				Balance:  0,
			},

			want: false,
		},
		{
			name: "should false when stock is making profit",
			stock: Stock{
				Quantity: 999,
				Balance:  1,
			},

			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.stock.IsInLoss()
			assert.Equal(t, test.want, result)
		})
	}
}
