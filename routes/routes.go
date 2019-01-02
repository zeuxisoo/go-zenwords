package routes

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// CreateEngine return the gin engine for test case
func CreateEngine() *gin.Engine {
	keywords.NewKeywords("../../words.txt")

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

// PerformRequestGet return the response detail for GET method and path
func PerformRequestGet(handler http.Handler, path string) *httptest.ResponseRecorder {
	return PerformRequest(handler, "GET", path)
}
