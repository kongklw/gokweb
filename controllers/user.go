package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 200, "success", "xiaoming class1", 1)
}

func GetList(c *gin.Context) {

	ReturnErr(c, 404, "not found about this one")
}
