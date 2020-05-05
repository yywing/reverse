package main

import (
	"context"
	"flag"
	"fmt"
	"reverse/client"
	"reverse/service"
)

var (
	id           = flag.String("id", "test", "reverse client id")
	serveiceName = flag.String("service_name", "test", "service name")
	reversePort  = flag.Int("port", 10001, "self port")
	serverPort   = flag.Int("server_port", 10000, "server port")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	services := make(map[string]service.Service)
	services["test"] = service.NewBaseService(*serveiceName, fmt.Sprintf("localhost:%d", *serverPort))

	c := client.NewReverseClient(ctx, *id, fmt.Sprintf("localhost:%d", *reversePort), services)
	c.Run()
}
