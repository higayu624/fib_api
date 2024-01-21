package model_test

import (
	"testing"

	"fib_api/model"

	"github.com/shopspring/decimal"
)

func TestCalclateFibonacci(t *testing.T) {
	t.Parallel()

	fibonacciOutputOK, _ := decimal.NewFromString("218922995834555169026")

	type input struct {
		number int
	}
	type want struct {
		fibonacci *decimal.Decimal
	}

	tests := map[string]struct {
		input input
		want  want
	}{
		"OK": {
			input{
				number: 99,
			},
			want{
				fibonacci: &fibonacciOutputOK,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			response := model.CalclateFibonacci(tt.input.number)
			if *response != *tt.want.fibonacci {
				t.Errorf("test case [%s]: CalclateFibonacci response is not equal want.fibonacci", name)
				t.Errorf("want %v", tt.want.fibonacci)
				t.Errorf("response %v", response)
			}
		})
	}
}
