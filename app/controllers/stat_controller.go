package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	xerrors "quarxlab/lib/errors"
)

type statController int

const StatController = statController(0)

func (this statController) Count(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		errJson := xerrors.NewError(9900)
		panic(errJson)
	}
	keys := client.Keys("*").Val()
	kv := make(map[string]string, len(keys))
	for _, k := range keys {
		val := client.Get(k).Val()
		kv[k] = val
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": kv})
}
