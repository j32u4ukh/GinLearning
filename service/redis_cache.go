package service

import (
	"GinLearning/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"
)

func CacheOneUserDecorator(h gin.HandlerFunc, porm string, readKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyId := c.Param(porm)
		redisKey := fmt.Sprintf(readKeyPattern, keyId)
		// conn := database.RedisDefaultPool.Get()
		conn := database.GetRedisConn()
		defer conn.Close()
		// data, err := redis.Bytes(conn.Do("GET", redisKey))
		data, err := database.GetRedis(conn, redisKey)
		if err != nil {
			h(c)
			dbResult, exists := c.Get("dbResult")
			if !exists {
				dbResult = empty
			}
			redisData, _ := ffjson.Marshal(dbResult)
			conn.Do("SETEX", redisKey, 30, redisData)
			c.JSON(http.StatusOK, gin.H{
				"message": "From DB",
				"data":    dbResult,
			})
			return
		}

		ffjson.Unmarshal(data, &empty)
		c.JSON(http.StatusOK, gin.H{
			"message": "From Redis",
			"data":    empty,
		})
	}
}

func CacheAllUserDecorator(h gin.HandlerFunc, redisKey string, empty interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := database.GetRedisConn()
		defer conn.Close()
		data, err := database.GetRedis(conn, redisKey)
		if err != nil {
			h(c)
			dbAllUser, exists := c.Get("dbAllUser")
			if !exists {
				dbAllUser = empty
			}
			redisData, _ := ffjson.Marshal(dbAllUser)
			conn.Do("SETEX", redisKey, 30, redisData)
			c.JSON(http.StatusOK, gin.H{
				"message": "From DB",
				"data":    dbAllUser,
			})
			return
		}
		ffjson.Unmarshal(data, &empty)
		c.JSON(http.StatusOK, gin.H{
			"message": "From Redis",
			"data":    empty,
		})
	}
}
