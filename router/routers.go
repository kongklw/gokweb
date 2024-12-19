package router

import (
	"gokweb/controllers"
	"gokweb/middlewares"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogOutput() {
	f, _ := os.Create("gokweb.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Router() *gin.Engine {
	setupLogOutput()
	// r := gin.Default()

	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	r.GET("/hello", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"msg": "hello world"})
	})

	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.GET("/list", controllers.UserController{}.GetList)

		user.POST("/add", func(ctx *gin.Context) {
			// name := ctx.PostForm("name")
			// age := ctx.DefaultPostForm("age", "10")
			// fmt.Println("name is ", name, "and age are ", age)

			// json 接收方式
			param := make(map[string]interface{})
			err := ctx.BindJSON(&param)
			if err == nil {
				controllers.ReturnSuccess(ctx, 200, "ok", param, 1)
			}
			controllers.ReturnErr(ctx, 200, "something is wrong")

		})
	}

	sport := r.Group("/sport")
	{
		sport.GET("/list", controllers.SportController{}.GetSports)
		sport.GET("/sort", controllers.SportController{}.SportSort)
	}
	return r

}
