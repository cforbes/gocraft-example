package main

import (
	"github.com/cforbes/examples/gocraftSplit/service"
	"net/http"
)

func main() {
	http.ListenAndServe(`localhost:8080`, gocraftsplit.GetRoutes())
}
