package gocraftsplit

import (
	"fmt"
	"net/http"
	"github.com/gocraft/web"
	"gopkg.in/redis.v3"
	"encoding/json"
)

type RedisContext struct {
	*StatusHandler
	Status string `json:"ping"`
	Error  error  `json:"error"`
}

func (c *RedisContext) Ping(rw web.ResponseWriter, req *web.Request) {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// defer redis.Close()

	c.Status, c.Error = redis.Ping().Result()

	if c.Error != nil {
		fmt.Println(`error`)
		rw.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(c)
}