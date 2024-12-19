package controllers

import (
	"fmt"
	"gokweb/cache"
	"gokweb/models"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type SportController struct {
}

func (s SportController) GetSports(c *gin.Context) {

	sport, err := models.GetSportList()
	if err != nil {
		fmt.Println("some thing is wrong", err.Error())
	}

	ReturnSuccess(c, 200, "successful", sport, 2)

}

// 进行sport点赞排序
func (s SportController) SportSort(c *gin.Context) {
	err := cache.Rdb.Set(cache.Rctx, "football", "10", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := cache.Rdb.Get(cache.Rctx, "football").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("football value is ", val)

	val2, err := cache.Rdb.Get(cache.Rctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	ReturnSuccess(c, 200, "hello", "world", 0)

}
