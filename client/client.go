package client

import (
	"context"
	"io"
	pb "reverse/proto"
	"reverse/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var logger = logrus.New()

const (
	MAX_SEND_MSG_SIZE    = 12 * 8 * 1024 * 1024
	MAX_RECEIVE_MSG_SIZE = 12 * 8 * 1024 * 1024
)

type ReverseClient struct {
	ctx        context.Context
	ID         string
	ServerAddr string
	Services   map[string]service.Service
}

func NewReverseClient(ctx context.Context, id string, serverAddr string, services map[string]service.Service) *ReverseClient {
	return &ReverseClient{
		ctx:        ctx,
		ID:         id,
		ServerAddr: serverAddr,
		Services:   services,
	}
}

func (c *ReverseClient) Run() {
	for {
		err := c.connect()
		logger.Errorf("reverse client connect error: %s", err)
	}

}

func (c *ReverseClient) handle(request *pb.ConnectRequest) {
	s, ok := c.Services[request.ServiceName]
	if !ok {
		logger.Errorf("reverse client handler not found: %s", request.ServiceName)
	}
	s.Handle(request)
}

func (c *ReverseClient) connect() error {
	// 建立连接
	conn, err := grpc.Dial(c.ServerAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	grpcClient := pb.NewReverseClient(conn)

	msg := &pb.RegisterMessage{
		ID: c.ID,
	}

	stream, err := grpcClient.ServerConnect(c.ctx, msg, grpc.MaxCallSendMsgSize(MAX_SEND_MSG_SIZE), grpc.MaxCallRecvMsgSize(MAX_RECEIVE_MSG_SIZE))

	if err != nil {
		return err
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil
			}
		}
		go c.handle(request)
	}

	return nil
}
