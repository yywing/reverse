package service

import (
	"io"
	"net"
	pb "reverse/proto"
	"reverse/utils"
	"sync"
	"time"

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
	protocol   string
	conn       net.Conn
	Keepalive  int
}

func NewBaseService(name, serverAddr string) *BaseService {
	return &BaseService{
		name:       name,
		serverAddr: serverAddr,
	}
}

func (s *BaseService) Check() {
	logger.Infof("%v", s.serverAddr)

	if s.conn != nil {
		s.conn.
	}

	toConnection, err := net.DialTimeout(s.protocol, s.serverAddr, time.Duration(s.Keepalive)*time.Second)
	if err != nil {
		return err
	}
	defer toConnection.Close()
}

func (s *BaseService) GetName() string {
	return s.name
}

func (s *BaseService) Handle(request *pb.ConnectRequest) {
	from := utils.IPAndPort(request.IP, uint16(request.Port))
	logger.Infof("%v", from)
	conn, err := net.DialTimeout(s.protocol, from, time.Duration(request.Timeout)*time.Second)
	if err != nil {
		logger.Errorf("%s connect to client error: %v", err)
		return
	}
	defer conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := io.Copy(conn, s.conn)
		if err != nil {
			logger.Errorf("from %v to %v send data error: %v\n", conn.RemoteAddr(), s.conn.RemoteAddr(), err)
		}
	}()

	go func() {
		defer wg.Done()
		_, err := io.Copy(s.conn, conn)
		if err != nil {
			logger.Errorf("from %v to %v send data error: %v\n", s.conn.RemoteAddr(), conn.RemoteAddr(), err)
		}
	}()

	// TODO: 当一方连接断了之后另一方会断么？
	wg.Wait()
	logger.Infof("from %v to %v handler exit.", conn.RemoteAddr(), s.conn.RemoteAddr())
	return
}
