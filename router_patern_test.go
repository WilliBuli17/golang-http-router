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

func TestNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/item/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		productId := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + productId + " Item " + itemId
		_, err := fmt.Fprint(writer, text)
		if err != nil {
			panic(err)
		}
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/2/item/4", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Product 2 Item 4", string(body))
}

func TestCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		_, err := fmt.Fprint(writer, text)
		if err != nil {
			panic(err)
		}
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Image : /small/profile.png", string(body))
}
