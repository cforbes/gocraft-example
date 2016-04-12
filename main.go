package main

import (
	"net/http"
	"github.com/cforbes/examples/gocraftSplit/service"
)

func main() {
	http.ListenAndServe(`localhost:8080`, gocraftsplit.GetRoutes())
}
