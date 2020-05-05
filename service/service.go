package service

import (
	"reverse/handler"
	pb "reverse/proto"
	"reverse/utils"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type Service interface {
	// service name is unique on one reverse client.
	GetName() string
	// Handle is how to proxy request from client to server
	Handle(*pb.ConnectRequest)
}

type BaseService struct {
	name       string
	serverAddr string
}

func NewBaseService(name, serverAddr string) *BaseService {
	return &BaseService{
		name:       name,
		serverAddr: serverAddr,
	}
}

func (s *BaseService) GetName() string {
	return s.name
}

func (s *BaseService) Handle(request *pb.ConnectRequest) {
	from := utils.IPAndPort(request.IP, uint16(request.Port))
	err := handler.Handle(from, s.serverAddr, strings.ToLower(request.Protocol.String()), request.Timeout)
	logger.Errorf("%v handle error: %v", s.name, err)
}
