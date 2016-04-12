package gocraftsplit

import (
	"github.com/gocraft/web"
	"encoding/json"
)

type StatusHandler struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func (c *StatusHandler) Get(rw web.ResponseWriter, req *web.Request) {
	c.Status = "success"

	json.NewEncoder(rw).Encode(c)
}
