package client

import (
	"context"
	"fmt"
	pb "reverse/proto"
	"reverse/service"
	"time"

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
	KeepAlive  int
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
		logger.Info("reverse client start connet ...")
		err := c.connect()
		logger.Errorf("reverse client connect error: %s", err)
		time.Sleep(1 * time.Second)
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

	ctx, cancel := context.WithCancel(c.ctx)
	defer cancel()

	stream, err := grpcClient.ServerConnect(ctx, grpc.MaxCallSendMsgSize(MAX_SEND_MSG_SIZE), grpc.MaxCallRecvMsgSize(MAX_RECEIVE_MSG_SIZE))

	if err != nil {
		return err
	}

	msg := &pb.ServerRequest{
		Message: &pb.ServerRequest_Register{
			Register: &pb.RegisterMessage{
				ID: c.ID,
			},
		},
	}
	err = stream.Send(msg)

	if err != nil {
		return err
	}

	go func() {
		t := time.NewTicker(time.Second * time.Duration(c.KeepAlive))
		defer t.Stop()
		for range t.C {
			msg := &pb.ServerRequest{
				Message: &pb.ServerRequest_Heartbeat{
					Heartbeat: &pb.HeartbeatMessage{},
				},
			}
			err = stream.Send(msg)
			if err != nil {
				logger.Errorf("stream send err: %v", err)
				cancel()
				return
			}
		}
	}()

	for {
		request, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("stream receive err: %v", err)
		}
		go c.handle(request)
	}
}
