package routes

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	engine = createEngine()
)

func TestHomeIndexGetOK(t *testing.T) {
	responseRecorder := performRequest(engine, "GET", "/")

	Convey("HomeIndexGet / should be OK", t, func() {
		So(responseRecorder.Code, ShouldEqual, http.StatusOK)
		So(responseRecorder.Body.String(), ShouldEqual, "ZenWords")
	})
}

func TestHomeRobotsTxtGetOK(t *testing.T) {
	responseRecorder := performRequest(engine, "GET", "/robots.txt")

	Convey("HomeRobotsTxtGet /robots.txt should be OK", t, func() {
		So(responseRecorder.Code, ShouldEqual, http.StatusOK)
		So(responseRecorder.Body.String(), ShouldContainSubstring, "User-agent")
		So(responseRecorder.Body.String(), ShouldContainSubstring, "Disallow")
	})
}
