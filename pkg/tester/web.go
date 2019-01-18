package tester

import (
	"io"
	"strings"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// CreateWebEngine return the gin engine for test case
func CreateWebEngine() *gin.Engine {
	keywords.NewKeywords("../../words.txt")

	engine := gin.New()
	engine.Use(gin.Recovery())

	return engine
}

// PerformWebRequest return the response detail for related method and path
func PerformWebRequest(handler http.Handler, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, body)

	if strings.ToUpper(method) == "POST" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	return responseRecorder
}

// PerformWebRequestGet return the response detail for GET method and path
func PerformWebRequestGet(handler http.Handler, path string) *httptest.ResponseRecorder {
	return PerformWebRequest(handler, "GET", path, nil)
}

// PerformWebRequestPost return the response detail for POST method and path
func PerformWebRequestPost(handler http.Handler, path string) *httptest.ResponseRecorder {
	return PerformWebRequest(handler, "POST", path, nil)
}

// PerformWebRequestPostWithBody return the response detail for POST method, path and body
func PerformWebRequestPostWithBody(handler http.Handler, path string, body string) *httptest.ResponseRecorder {
	return PerformWebRequest(handler, "POST", path, strings.NewReader(body))
}
