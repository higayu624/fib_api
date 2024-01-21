package main

import (
	"time"

	"fib_api/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var fibAPIRoot = "fib"

func initRouter() *gin.Engine {
	router := gin.Default()
	// router.ContextWithFallback = true

	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"GET",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Content-Type",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost",
		},
		// // cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	router.GET(fibAPIRoot, controller.GetFibonacci())

	return router
}
