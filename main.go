package main

import (
	"gokweb/router"
)

func main() {

	r := router.Router()

	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000") // set self define port

}
