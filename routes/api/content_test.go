package api

import (
	"net/http"
	"testing"
	"encoding/json"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/zeuxisoo/go-zenwords/routes"
	"github.com/zeuxisoo/go-zenwords/pkg/contracts"
)

var (
	engine *gin.Engine
)

func init() {
	engine = routes.CreateEngine()

	engine.
		Group("/api").
		Group("/content").
			POST("/replace", ContentReplacePost)
}

func TestContentReplacePostOK(t *testing.T) {
	Convey("ContentReplacePost /api/content/replace should be OK", t, func() {
		responseRecorder := routes.PerformRequest(engine, "POST", "/api/content/replace")

		Convey("Http status is 200", func() {
			So(responseRecorder.Code, ShouldEqual, http.StatusOK)
		})

		Convey("JSON should not be nil", func() {
			data := contracts.APIContentReplacePost{}

			So(json.Unmarshal(responseRecorder.Body.Bytes(), &data), ShouldBeNil)
			So(data, ShouldResemble, contracts.APIContentReplacePost{})
		})

		Convey("Result key should be empty", func() {
			data := contracts.APIContentReplacePost{}

			json.Unmarshal(responseRecorder.Body.Bytes(), &data)

			So(data.Result, ShouldEqual, "")
		})
	})
}
