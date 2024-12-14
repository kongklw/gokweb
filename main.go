package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"msg": "hello world"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000") // set self define port

}
