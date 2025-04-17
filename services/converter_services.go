package services

import (
	"github.com/shopspring/decimal"
	"github.com/thyms-c/currency-converter/repositories"
)

type ConverterService interface {
	ConvertDecimalToThaiBahtText(amount decimal.Decimal) string
}
type converterServiceImpl struct {
	coverterRepository repositories.ConverterRepository
}

func NewConverterService(converterRepository repositories.ConverterRepository) ConverterService {
	return &converterServiceImpl{
		coverterRepository: converterRepository,
	}
}

func (c *converterServiceImpl) ConvertDecimalToThaiBahtText(amount decimal.Decimal) string {
	intPart := amount.Truncate(0)
	fracPart := amount.Sub(intPart).Mul(decimal.NewFromInt(100)).Round(0)

	result := c.coverterRepository.ConvertNumberToThaiText(intPart.IntPart()) + "บาท"
	if fracPart.IsZero() {
		result += "ถ้วน"
	} else {
		result += c.coverterRepository.ConvertNumberToThaiText(fracPart.IntPart()) + "สตางค์"
	}
	return result
}
