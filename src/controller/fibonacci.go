package controller

import (
	"net/http"
	"strconv"

	"fib_api/model"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type requestFibonacciNumber struct {
	number int
}

type fibonacci struct {
	fibonacci *decimal.Decimal
}

// GetFibonacci() return the Fibonacci number of the nth term
func GetFibonacci() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request requestFibonacciNumber
		var modelResponse fibonacci
		var err error

		// Recieved queryParam and validation
		requestString := ctx.Query("n")
		request.number, err = strconv.Atoi(requestString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "bad request",
			})
			return
		}
		if request.number < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "bad request",
			})
			return
		}

		// Calculate the Fibonacci number of the nth term
		modelResponse.fibonacci = model.CalclateFibonacci(&request.number)

		// Make response
		decimal.MarshalJSONWithoutQuotes = true
		response := make(map[string]*decimal.Decimal)
		response["result"] = modelResponse.fibonacci

		ctx.JSON(http.StatusOK, response)
	}
}
