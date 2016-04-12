package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"net/http"
	"github.com/cforbes/examples/gocraftSplit/service"
)

func TestRedisHandler (t *testing.T) {
	router := gocraftsplit.GetRoutes()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/status", nil)

	router.ServeHTTP(recorder, request)
	// res, err := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(recorder.Body.Bytes()))
}