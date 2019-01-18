package tester

import (
	"io"
	"strings"
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
func PerformRequest(handler http.Handler, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, body)

	if strings.ToUpper(method) == "POST" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	return responseRecorder
}

// PerformRequestGet return the response detail for GET method and path
func PerformRequestGet(handler http.Handler, path string) *httptest.ResponseRecorder {
	return PerformRequest(handler, "GET", path, nil)
}

// PerformRequestPost return the response detail for POST method and path
func PerformRequestPost(handler http.Handler, path string) *httptest.ResponseRecorder {
	return PerformRequest(handler, "POST", path, nil)
}

// PerformRequestPostWithBody return the response detail for POST method, path and body
func PerformRequestPostWithBody(handler http.Handler, path string, body string) *httptest.ResponseRecorder {
	return PerformRequest(handler, "POST", path, strings.NewReader(body))
}
