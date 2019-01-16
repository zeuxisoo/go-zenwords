package servers

import (
	"context"

	"github.com/zeuxisoo/go-zenwords/rpc/proto"
	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// ContentServiceServer is implementing a proto.ContentServiceServer
type ContentServiceServer struct {
}

// Replace will return the filtered content
func (s *ContentServiceServer) Replace(context context.Context, request *proto.ContentReplaceRequest) (*proto.ContentReplaceResponse, error) {
	content := request.GetContent()
	result  := keywords.Filter(content)

	return &proto.ContentReplaceResponse{
		Result: result,
	}, nil
}
