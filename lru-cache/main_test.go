package main

import (
	"strconv"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"gotest.tools/assert"
)

func TestLRU(t *testing.T) {
	type Params struct {
		key            int
		value          *int
		expectedOutput *int
	}
	type TestCase struct {
		capacity int
		input    []map[string]Params
	}

	tests := []TestCase{
		{
			input: []map[string]Params{
				{"put": Params{key: 1, value: ptr.Int(1)}},
				{"put": Params{key: 2, value: ptr.Int(2)}},
				{"get": Params{key: 1, expectedOutput: ptr.Int(1)}},
				{"put": Params{key: 3, value: ptr.Int(3)}},
				{"get": Params{key: 2, expectedOutput: ptr.Int(-1)}},
			},
			capacity: 2,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			lruCache := Constructor(test.capacity)

			for _, operations := range test.input {
				for operation, params := range operations {
					switch operation {
					case "put":
						lruCache.Put(params.key, *params.value)
					case "get":
						assert.Equal(t, *params.expectedOutput, lruCache.Get(params.key))
					}
				}
			}
		})
	}
}
