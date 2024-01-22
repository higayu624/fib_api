package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"fib_api/controller"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

func TestGetFibonacci(t *testing.T) {
	t.Parallel()

	type want struct {
		status int
		body   string
	}

	tests := map[string]struct {
		number any
		want   want
	}{
		"OK": {
			number: 99,
			want: want{
				status: http.StatusOK,
				body:   `{"result":218922995834555169026}`,
			},
		},
		"NG when the number type is not int": {
			number: "d",
			want: want{
				status: http.StatusBadRequest,
				body:   `{"message":"bad request","status":400}`,
			},
		},
		"NG when the number is negative": {
			number: -1,
			want: want{
				status: http.StatusBadRequest,
				body:   `{"message":"bad request","status":400}`,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			fibAPIRoot := fmt.Sprintf("/fib?n=%d", tt.number)

			response := httptest.NewRecorder()

			engine := gin.New()
			engine.Use(func(ctx *gin.Context) {
				c, _ := gin.CreateTestContext(response)
				c.Request, _ = http.NewRequest(
					http.MethodGet,
					fmt.Sprintf("/fib?n=%d", tt.number),
					nil,
				)
				c.Request.Header.Set("content-Type", "application/json")
			})

			engine.GET("/fib", controller.GetFibonacci())

			req := httptest.NewRequest("GET", fibAPIRoot, nil)
			engine.ServeHTTP(response, req)

			if response.Code != tt.want.status {
				t.Errorf("status code got %d, want %d", response.Code, tt.want.status)
			}
			if diff := cmp.Diff(response.Body.String(), tt.want.body); diff != "" {
				t.Errorf("GetFibonacci returned unexpected diff %s", diff)
			}
		})
	}
}
