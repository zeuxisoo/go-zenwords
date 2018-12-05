package routes

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// CreateEngine return the gin engine for test case
func CreateEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())

	return engine
}

// PerformRequest return the response detail for related method and path
func PerformRequest(handler http.Handler, method string, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	return responseRecorder
}
