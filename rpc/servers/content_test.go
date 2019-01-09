package servers

import (
	"testing"
	"context"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/zeuxisoo/go-zenwords/rpc/protos"
)

func TestContentServiceServerReplaceIsOK(t *testing.T) {
	Convey("ContentServiceServer Replace method is OK", t, func() {
		server  := ContentServiceServer{}
		request := &protos.ContentReplaceRequest{}

		response, err := server.Replace(context.Background(), request)

		So(err, ShouldBeNil)
		So(response.Result, ShouldEqual, "")
	})
}
