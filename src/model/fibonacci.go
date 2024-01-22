package model

import "github.com/shopspring/decimal"

func CalclateFibonacci(number *int) *decimal.Decimal {
	var twoPrevious, Previous decimal.Decimal // 2の31乗の桁まで計算可能な型

	twoPrevious, Previous = decimal.NewFromInt(0), decimal.NewFromInt(1)
	for i := 0; i < *number; i++ {
		twoPrevious, Previous = Previous, twoPrevious.Add(Previous)
	}
	return &twoPrevious
}
