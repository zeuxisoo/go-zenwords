package servers

import (
	"testing"
	"context"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/zeuxisoo/go-zenwords/pkg/tester"
	"github.com/zeuxisoo/go-zenwords/rpc/proto"
)

func init() {
	tester.CreateRPC()
}

func TestContentReplaceMethodIsOK(t *testing.T) {
	Convey("Content replace should be OK", t, func() {
		server := ContentServiceServer{}

		Convey("When content is empty should be OK", func() {
			request := &proto.ContentReplaceRequest{
				Content: "",
			}

			response, err := server.Replace(context.Background(), request)

			So(err, ShouldBeNil)
			So(response.Result, ShouldEqual, "")
		})

		Convey("When content is not empty should be equals 'this is a *****'", func() {
			request := &proto.ContentReplaceRequest{
				Content: "this is a apple",
			}

			response, err := server.Replace(context.Background(), request)

			So(err, ShouldBeNil)
			So(response.Result, ShouldEqual, "this is a *****")
		})
	})
}
