package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"msg": "hello world"})
	})

	user := r.Group("/user")
	{
		user.GET("/list", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"data": "xiaoming"})
		})

		user.POST("/add", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "add successful"})
		})
	}
	return r

}
