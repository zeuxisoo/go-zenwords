package servers

import (
	"context"

	"github.com/zeuxisoo/go-zenwords/rpc/protos"
	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// ContentServiceServer is implementing a protos.ContentServiceServer
type ContentServiceServer struct {
}

// Replace will return the filtered content
func (s *ContentServiceServer) Replace(context context.Context, request *protos.ContentReplaceRequest) (*protos.ContentReplaceResponse, error) {
	content := request.GetContent()
	result  := keywords.Filter(content)

	return &protos.ContentReplaceResponse{
		Result: result,
	}, nil
}
