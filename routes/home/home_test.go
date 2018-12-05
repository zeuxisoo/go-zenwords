package home

import (
	"os"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zeuxisoo/go-zenwords/routes"
)

var (
	engine *gin.Engine
)

func init() {
	engine = routes.CreateEngine()

	engine.GET("/", IndexGet)
	engine.GET("/robots.txt", RobotsTxtGet)
}

func TestHomeIndexGetOK(t *testing.T) {
	Convey("HomeIndexGet / should be OK", t, func() {
		responseRecorder := routes.PerformRequest(engine, "GET", "/")

		So(responseRecorder.Code, ShouldEqual, http.StatusOK)
		So(responseRecorder.Body.String(), ShouldEqual, "ZenWords")
	})
}

func TestHomeRobotsTxtGetOK(t *testing.T) {
	Convey("HomeRobotsTxtGet /robots.txt should be OK", t, func() {
		Convey("When robots.txt is not exists", func() {
			responseRecorder := routes.PerformRequest(engine, "GET", "/robots.txt")

			So(responseRecorder.Code, ShouldEqual, http.StatusOK)
			So(responseRecorder.Body.String(), ShouldContainSubstring, "User-agent")
			So(responseRecorder.Body.String(), ShouldContainSubstring, "Disallow")
		})

		Convey("When robots.txt is exists", func() {
			ioutil.WriteFile("robots.txt", []byte("Hello World\nrobots.txt"), 0644)

			responseRecorder := routes.PerformRequest(engine, "GET", "/robots.txt")

			So(responseRecorder.Code, ShouldEqual, http.StatusOK)
			So(responseRecorder.Body.String(), ShouldContainSubstring, "Hello World")
			So(responseRecorder.Body.String(), ShouldContainSubstring, "robots.txt")

			os.Remove("robots.txt")
		})
	})
}