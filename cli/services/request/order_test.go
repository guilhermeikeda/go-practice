package request

import (
	"testing"

	"gotest.tools/assert"
)

func TestOrderRequest(t *testing.T) {
	type TestCase struct {
		name  string
		order Order

		want float32
	}

	tests := []TestCase{
		{
			name: "should return order total cost",
			order: Order{
				UnitCost: 10,
				Quantity: 100,
			},
			want: 1000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, test.order.TotalCost())
		})
	}
}
