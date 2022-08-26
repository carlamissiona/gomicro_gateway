package main

import (
	"bookgateway/handler"
	pb "bookgateway/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/micro/go-micro/v2/web"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("bookgateway"),
	)

	// Register handler
	pb.RegisterBookgatewayHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}



	websv := web.NewService(
        web.Name("echo.micro.api.gateway"),
        web.Version("latest"),
    )

    // Initialise service
    websv.Init()

    // Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    g := e.Group("/api")
    g.GET("/list", handler.List)

    // Register Handler
    service.Handle("/", e)

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
