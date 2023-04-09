package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Method Not Allowed")
		if err != nil {
			panic(err)
		}
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		_, err := fmt.Fprint(writer, "POST Method")
		if err != nil {
			panic(err)
		}
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Method Not Allowed", string(body))
}
