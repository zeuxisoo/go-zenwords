package routes

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func createEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.GET("/", HomeIndexGet)
	engine.GET("/robots.txt", HomeRobotsTxtGet)

	return engine
}

func performRequest(handler http.Handler, method string, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	return responseRecorder
}
