package home

import (
	"os"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zeuxisoo/go-zenwords/pkg/tester"
)

var (
	engine *gin.Engine
)

func init() {
	engine = tester.CreateWebEngine()

	engine.GET("/", IndexGet)
	engine.GET("/robots.txt", RobotsTxtGet)
}

func TestHomeIndexGetOK(t *testing.T) {
	Convey("HomeIndexGet / should be OK", t, func() {
		responseRecorder := tester.PerformWebRequestGet(engine, "/")

		So(responseRecorder.Code, ShouldEqual, http.StatusOK)
		So(responseRecorder.Body.String(), ShouldEqual, "ZenWords")
	})
}

func TestHomeRobotsTxtGetOK(t *testing.T) {
	Convey("HomeRobotsTxtGet /robots.txt should be OK", t, func() {
		Convey("When robots.txt is not exists", func() {
			responseRecorder := tester.PerformWebRequestGet(engine, "/robots.txt")

			So(responseRecorder.Code, ShouldEqual, http.StatusOK)
			So(responseRecorder.Body.String(), ShouldContainSubstring, "User-agent")
			So(responseRecorder.Body.String(), ShouldContainSubstring, "Disallow")
		})

		Convey("When robots.txt is exists", func() {
			ioutil.WriteFile("robots.txt", []byte("Hello World\nrobots.txt"), 0644)

			responseRecorder := tester.PerformWebRequestGet(engine, "/robots.txt")

			So(responseRecorder.Code, ShouldEqual, http.StatusOK)
			So(responseRecorder.Body.String(), ShouldContainSubstring, "Hello World")
			So(responseRecorder.Body.String(), ShouldContainSubstring, "robots.txt")

			os.Remove("robots.txt")
		})
	})
}
