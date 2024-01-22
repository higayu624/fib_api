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
		want want
	}{
		"OK": {
			want: want{
				status: http.StatusOK,
				body:   `{"result":218922995834555169026}`,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			number := 99
			fibAPIRoot := fmt.Sprintf("/fib?n=%d", number)

			response := httptest.NewRecorder()

			engine := gin.New()
			engine.Use(func(ctx *gin.Context) {
				c, _ := gin.CreateTestContext(response)

				c.Request, _ = http.NewRequest(
					http.MethodGet,
					fmt.Sprintf("/fib?n=%d", number),
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
