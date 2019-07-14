package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}
}

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		client.Incr(path)
		log.Println("incr", path)
		c.Next()
	}
}
