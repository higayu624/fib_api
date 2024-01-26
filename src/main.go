package main

import "log"

func main() {
	log.Printf("Server activation")
	log.Printf("waiting http access")

	router := initRouter()
	router.Run(":5000")
}
