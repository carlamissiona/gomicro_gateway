package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	books "books/proto"
)

type Books struct{}

// Return a new handler
func New() *Books {
	return &Books{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Books) List(ctx context.Context, req *books.ListRequest, rsp *books.Response) error {
	log.Info("Received Books.Call request")
	rsp.Msg = "Hello List function"
	return nil
}

func (e *Books) Get(ctx context.Context, req *books.Request, rsp *books.Response) error {
	log.Info("Received Books.Call request")
	rsp.Msg = "Hello Get function"
	return nil
}