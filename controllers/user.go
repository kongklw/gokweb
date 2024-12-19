package controllers

import "github.com/gin-gonic/gin"

// 如果controllers 里面有很多文件。同时想使用GetList作为函数名称 。那么就会导致别的文件中不能定义该函数。
// 解决方案。 定义一个struct （类似python class）

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")

	ReturnSuccess(c, 200, "success", "xiaoming class1 "+id, 1)
}

func (u UserController) GetList(c *gin.Context) {

	ReturnErr(c, 404, "not found about this one")
}
