package main

import (
	"ganhocapital/input"
	"ganhocapital/output"
	"testing"

	"gotest.tools/assert"
)

type TestCase struct {
	name       string
	operations []input.Operation

	want output.Simulation
}

func TestPerformSimulation(t *testing.T) {

	tests := []TestCase{
		{
			name:       "case #1",
			operations: testCase1.operations,
			want:       testCase1.want,
		},
		{
			name:       "case #2",
			operations: testCase2.operations,
			want:       testCase2.want,
		},
		{
			name:       "case #3",
			operations: testCase3.operations,
			want:       testCase3.want,
		},
		{
			name:       "case #4",
			operations: testCase4.operations,
			want:       testCase4.want,
		},
		{
			name:       "case #5",
			operations: testCase5.operations,
			want:       testCase5.want,
		},
		{
			name:       "case #6",
			operations: testCase6.operations,
			want:       testCase6.want,
		},
		testCase7,
		{
			name:       "case #8",
			operations: testCase8.operations,
			want:       testCase8.want,
		},
		{
			name:       "custom case #1",
			operations: customTestCase1.operations,
			want:       customTestCase1.want,
		},
		{
			name:       "custom case #1",
			operations: customTestCase2.operations,
			want:       customTestCase2.want,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			result := performSimulation(test.operations)

			assert.DeepEqual(t, test.want, result)
		})
	}
}

func TestPerformSimulationForSpecificCase(t *testing.T) {
	testCase := customTestCase1
	operations := testCase.operations
	want := testCase.want

	result := performSimulation(operations)
	assert.DeepEqual(t, want, result)
}

var (
	fixture = input.NewOperationFixture()
)

var testCase1 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(100, 10.00).
		AddSellOperation(50, 15.00).
		AddSellOperation(50, 15.00).
		Build(),

	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
	},
}

var testCase2 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddSellOperation(5000, 20.00).
		AddSellOperation(5000, 5.00).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 10000,
		},
		{
			Tax: 0,
		},
	},
}

var testCase3 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddSellOperation(5000, 5.00).
		AddSellOperation(3000, 20.00).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 1000,
		},
	},
}

var testCase4 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddBuyOperation(5000, 25.00).
		AddSellOperation(10000, 15.00).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
	},
}

var testCase5 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddBuyOperation(5000, 25.00).
		AddSellOperation(10000, 15.00).
		AddSellOperation(5000, 25.00).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 10000,
		},
	},
}

var testCase6 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddSellOperation(5000, 2.00).
		AddSellOperation(2000, 20.00).
		AddSellOperation(2000, 20.00).
		AddSellOperation(1000, 25.00).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 3000,
		},
	},
}

var testCase7 TestCase = TestCase{
	name: "case #7",
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddSellOperation(5000, 2.00).
		AddSellOperation(2000, 20.00).
		AddSellOperation(2000, 20.00).
		AddSellOperation(1000, 25.00).
		AddBuyOperation(10000, 20.00).
		AddSellOperation(5000, 15.00).
		AddSellOperation(4350, 30.00).
		AddSellOperation(650, 30.00).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 3000,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 3700,
		},
		{
			Tax: 0,
		},
	},
}

var testCase8 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddSellOperation(10000, 50.00).
		AddBuyOperation(10000, 20.00).
		AddSellOperation(10000, 50.00).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 80000,
		},
		{
			Tax: 0,
		},
		{
			Tax: 60000,
		},
	},
}

var customTestCase1 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 3.98).
		AddBuyOperation(5000, 2.37).
		AddSellOperation(2000, 10.99).
		AddSellOperation(1000, 0.01).
		AddSellOperation(12000, 100.33).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 3018.67,
		},
		{
			Tax: 0,
		},
		{
			Tax: 232528,
		},
	},
}

var customTestCase2 TestCase = TestCase{
	operations: fixture.
		AddBuyOperation(10000, 10.00).
		AddSellOperation(2000, 5.00).
		AddSellOperation(2000, 9.99).
		AddSellOperation(2000, 8.78).
		AddSellOperation(2000, 8.78).
		AddSellOperation(2000, 4.29).
		Build(),
	want: output.Simulation{
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
		{
			Tax: 0,
		},
	},
}
