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
	Convey("When content is empty should be OK", t, func() {
		server  := ContentServiceServer{}
		request := &proto.ContentReplaceRequest{
			Content: "",
		}

		response, err := server.Replace(context.Background(), request)

		So(err, ShouldBeNil)
		So(response.Result, ShouldEqual, "")
	})
}
