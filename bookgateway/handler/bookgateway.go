package handler

import (
	"context"
    "fmt"
	log "github.com/micro/micro/v3/service/logger"

	bookgateway "bookgateway/proto"
)

type Bookgateway struct{}

// Return a new handler
func New() *Bookgateway {
	return &Bookgateway{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Bookgateway) Call(ctx context.Context, req *bookgateway.Request, rsp *bookgateway.Response) error {
	log.Info("Received Bookgateway.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}



// Call is a single request handler called via client.Call or the generated client code
func (e *Bookgateway) List(ctx context.Context, req *bookgateway.ListRequest, rsp *bookgateway.Response) error {
	log.Info("Received Bookgateway.Call request")
	 strmsg := fmt.Println("API Client : %v" , req.Name)
	 strmsg = strmsg + fmt.Println("Here are the list of services /books /providers /users")
	rsp.Msg = "List services books book providers " + req.Name
	return nil
}
