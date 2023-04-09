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

type LogMiddleware struct {
	Handler http.Handler
}

func (l *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Recive Request")
	l.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		_, err := fmt.Fprint(writer, "Middleware")
		if err != nil {
			panic(err)
		}
	})

	middleware := LogMiddleware{ // saat ini dibuat tidak menggunakan pointer karena aplikasi tidak memerlukan modifikasi pada objek middleware selama penggunaannya
		Handler: router,
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Middleware", string(body))
}
