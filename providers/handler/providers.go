package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	providers "providers/proto"
)

type Providers struct{}

// Return a new handler
func New() *Providers {
	return &Providers{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Providers) Call(ctx context.Context, req *providers.Request, rsp *providers.Response) error {
	log.Info("Received Providers.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Providers) Stream(ctx context.Context, req *providers.StreamingRequest, stream providers.Providers_StreamStream) error {
	log.Infof("Received Providers.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&providers.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Providers) PingPong(ctx context.Context, stream providers.Providers_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&providers.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
