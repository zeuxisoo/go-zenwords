package servers

import (
	"context"

	"github.com/zeuxisoo/go-zenwords/rpc/protos"
)

// ContentServiceServer is implementing a protos.ContentServiceServer
type ContentServiceServer struct {
}

// Replace will return the filtered content
func (s *ContentServiceServer) Replace(context context.Context, request *protos.ContentReplaceRequest) (*protos.ContentReplaceResponse, error) {
	return &protos.ContentReplaceResponse{
		// TODO: filter content action
		Result: "",
	}, nil
}
