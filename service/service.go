package service

import (
	pb "reverse/proto"
)

type Service interface {
	// service name is unique on one reverse client.
	GetName() string
	// Handle is how to proxy request from client to server
	Handle(*pb.ConnectRequest)
}
