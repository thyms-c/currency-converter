package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/thyms-c/currency-converter/repositories"
	"github.com/thyms-c/currency-converter/services"
)

func main() {
	converterRepository := repositories.NewConverterRepository()
	converterService := services.NewConverterService(converterRepository)

	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)
		fmt.Printf("=> %s\n", converterService.ConvertDecimalToThaiBahtText(input))

	}
}
