package main

import (
	"net/http"
	"time"

	"fib_api/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var fibAPIRoot = "fib"

func initRouter() *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "404 page not found",
		})
	})

	router.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"GET",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowOrigins: []string{
			"http://localhost",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	router.GET(fibAPIRoot, controller.GetFibonacci())

	return router
}
