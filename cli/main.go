package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"ganhocapital/entity"
	"ganhocapital/input"
	"ganhocapital/output"
	"ganhocapital/services"
	"ganhocapital/services/request"
	"ganhocapital/services/result"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		rawInput, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(fmt.Sprintf("failed to read input: %s", err.Error()))
		}

		operations := []input.Operation{}
		err = json.Unmarshal([]byte(rawInput), &operations)
		if err != nil {
			panic(fmt.Sprintf("invalid simulation %s", rawInput))
		}

		result := performSimulation(operations)
		printResult(result)
	}
}

func performSimulation(operations []input.Operation) output.Simulation {
	var simulationOutput output.Simulation
	stock := entity.Stock{
		Quantity:     0,
		AveragePrice: 0,
	}

	for _, operation := range operations {
		if !operation.IsValid() {
			panic("invalid operation")
		}

		var orderResult result.Order
		if operation.IsBuyOperation() {
			orderResult = performBuyOperation(stock, operation)
		} else {
			orderResult = performSellOperation(stock, operation)
		}

		stock = orderResult.Stock
		simulationOutput = append(simulationOutput, output.Charge{
			Tax: orderResult.Charges.Tax,
		})
	}

	return simulationOutput
}

func performBuyOperation(stock entity.Stock, operation input.Operation) result.Order {
	var (
		stockService = services.NewStockService()

		orderRequest = request.Order{
			UnitCost: operation.UnitCost,
			Quantity: operation.Quantity,
		}
	)

	orderResult, err := stockService.Buy(stock, orderRequest)
	if err != nil {
		panic("failed to perform buy operation")
	}

	return orderResult
}

func performSellOperation(stock entity.Stock, operation input.Operation) result.Order {
	var (
		stockService = services.NewStockService()

		orderRequest = request.Order{
			UnitCost: operation.UnitCost,
			Quantity: operation.Quantity,
		}
	)

	orderResult, err := stockService.Sell(stock, orderRequest)
	if err != nil {
		panic("failed to perform sell operation")
	}

	return orderResult
}

func printResult(simulationOutput output.Simulation) {
	simulationOutputAsJson, err := json.Marshal(simulationOutput)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal simulation output to json: %s", err.Error()))
	}

	fmt.Println(string(simulationOutputAsJson))
}
