package main

import (
	"fmt"
	"github.com/cforbes/examples/gocraftSplit/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedisHandler(t *testing.T) {
	router := gocraftsplit.GetRoutes()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/status", nil)

	router.ServeHTTP(recorder, request)
	// res, err := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(recorder.Body.Bytes()))
}
