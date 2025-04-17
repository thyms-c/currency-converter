package repositories

import (
	"strings"
)

type ConverterRepository interface {
	ConvertNumberToThaiText(number int64) string
	SplitNumberToSegments(number int64) []int64
	ConvertSegmentToThaiText(number int64) string
}

type converterRepositoryImpl struct {
	thaiDigits []string
	thaiUnits  []string
}

func NewConverterRepository() ConverterRepository {
	return &converterRepositoryImpl{
		thaiDigits: []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"},
		thaiUnits:  []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"},
	}
}

func (c *converterRepositoryImpl) ConvertNumberToThaiText(number int64) string {
	if number == 0 {
		return "ศูนย์"
	}

	var result strings.Builder
	segments := c.SplitNumberToSegments(number)

	for i := len(segments) - 1; i >= 0; i-- {
		segment := segments[i]
		if segment == 0 {
			continue
		}
		text := c.ConvertSegmentToThaiText(segment)
		if i > 0 {
			text += "ล้าน"
		}
		result.WriteString(text)
	}

	return result.String()
}

func (c *converterRepositoryImpl) ConvertSegmentToThaiText(number int64) string {
	var result strings.Builder
	digits := []int{}

	for number > 0 {
		digits = append(digits, int(number%10))
		number /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]
		unit := c.thaiUnits[i]

		if d == 0 {
			continue
		}

		switch {
		case i == 0 && d == 1 && len(digits) > 1:
			result.WriteString("เอ็ด")
		case i == 1 && d == 2:
			result.WriteString("ยี่" + unit)
		case i == 1 && d == 1:
			result.WriteString(unit)
		default:
			result.WriteString(c.thaiDigits[d] + unit)
		}
	}
	return result.String()
}

func (c *converterRepositoryImpl) SplitNumberToSegments(number int64) []int64 {
	var segments []int64

	for number > 0 {
		segments = append(segments, number%1_000_000)
		number /= 1_000_000
	}
	return segments
}
