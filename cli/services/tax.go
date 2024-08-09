package services

import (
	"ganhocapital/entity"
	"ganhocapital/services/request"
	"math"
)

type TaxCalculator interface {
	Calculate(entity.Stock, request.Order) float32
}

func round(value float32) float32 {
	shift := float32(100)

	return (float32(math.Round(float64(value*shift))) / shift)
}
