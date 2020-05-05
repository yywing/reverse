package rpc

import (
	"context"
	"sync"

	pb "reverse/proto"

	"github.com/sirupsen/logrus"
	errors "golang.org/x/xerrors"
	"google.golang.org/grpc"
)

var logger = logrus.New()

type ReverseService struct {
	mutex sync.Mutex
	chs   map[string](chan *pb.ConnectRequest)
}

func NewReverseService() *ReverseService {
	return &ReverseService{
		mutex: sync.Mutex{},
		chs:   make(map[string](chan *pb.ConnectRequest)),
	}
}

func (s *ReverseService) Register(srv *grpc.Server) {
	pb.RegisterReverseServer(srv, s)
}

func (s *ReverseService) ClientConnect(ctx context.Context, request *pb.ClientConnectRequest) (*pb.ConnectResponse, error) {
	ch, ok := s.chs[request.ID]
	if !ok {
		return nil, errors.Errorf("server connect stream load error")
	}
	ch <- request.Request
	return &pb.ConnectResponse{}, nil
}

func (s *ReverseService) ServerConnect(request *pb.RegisterMessage, stream pb.Reverse_ServerConnectServer) error {
	ch := make(chan *pb.ConnectRequest)
	s.chs[request.ID] = ch

	defer func() {
		s.mutex.Lock()
		delete(s.chs, request.ID)
		close(ch)
		s.mutex.Unlock()
	}()

	for msg := range ch {
		err := stream.Send(msg)
		if err != nil {
			logger.Errorf("server connect send msg to %s error: %s", request.ID, err)
		}
	}

	return nil
}
