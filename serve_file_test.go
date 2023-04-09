package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var fileResources embed.FS

func TestServeFile(t *testing.T) {
	dir, err := fs.Sub(fileResources, "resources")
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello i am txt", string(body))
}
