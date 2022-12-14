package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	users "users/proto"
)

type Users struct{}

// Return a new handler
func New() *Users {
	return &Users{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Users) Call(ctx context.Context, req *users.Request, rsp *users.Response) error {
	log.Info("Received Users.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Users) Stream(ctx context.Context, req *users.StreamingRequest, stream users.Users_StreamStream) error {
	log.Infof("Received Users.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&users.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Users) PingPong(ctx context.Context, stream users.Users_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&users.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
