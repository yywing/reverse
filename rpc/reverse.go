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

func (s *ReverseService) PrintServices() {
	logger.Infof("register services %+v", s.chs)
}

func (s *ReverseService) Register(srv *grpc.Server) {
	pb.RegisterReverseServer(srv, s)
}

func (s *ReverseService) ClientConnect(ctx context.Context, request *pb.ClientConnectRequest) (*pb.ConnectResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	ch, ok := s.chs[request.ID]
	if !ok {
		return nil, errors.Errorf("server connect stream load error with id: %v", request.ID)
	}
	ch <- request.Request
	return &pb.ConnectResponse{}, nil
}

func (s *ReverseService) ServerConnect(stream pb.Reverse_ServerConnectServer) error {
	first, err := stream.Recv()
	if err != nil {
		logger.Errorf("server connect get firest msg err: %s", err)
	}
	request, ok := first.Message.(*pb.ServerRequest_Register)
	if !ok {
		logger.Errorf("server connect first msg type error")
	}
	ch := make(chan *pb.ConnectRequest)
	s.chs[request.Register.ID] = ch

	defer func() {
		s.mutex.Lock()
		delete(s.chs, request.Register.ID)
		close(ch)
		s.mutex.Unlock()
		logger.Infof("server connect %v exit!", request.Register.ID)
	}()

	ctx, cancel := context.WithCancel(stream.Context())

	go func() {
		for {
			_, err := stream.Recv()
			if err != nil {
				logger.Errorf("server connect heartbeat receive err: %v", err)
				cancel()
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			logger.Infof("server connect %v send msg exit!", request.Register.ID)
			return nil
		case msg := <-ch:
			err := stream.Send(msg)
			if err != nil {
				logger.Errorf("server connect send msg to %s error: %s", request.Register.ID, err)
			}
		}
	}
}
