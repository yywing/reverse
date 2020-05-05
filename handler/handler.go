package handler

import (
	"io"
	"net"
	"sync"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Handler(from, to, protocol string, timeout int32) error {
	fromConnection, err := net.Dial(protocol, from)
	if err != nil {
		return err
	}
	defer fromConnection.Close()

	toConnection, err := net.Dial(protocol, to)
	if err != nil {
		return err
	}
	defer toConnection.Close()

	wg := sync.WaitGroup{}
	go proxyRequest(fromConnection, toConnection)
	go proxyRequest(toConnection, fromConnection)
	// TODO: 当一方连接断了之后另一方会断么？
	wg.Wait()
	logger.Infof("from %v to %v handler exit.", fromConnection.RemoteAddr(), toConnection.RemoteAddr())
}

// Forward all requests from r to w
func proxyRequest(r net.Conn, w net.Conn, wg sync.WaitGroup) {
	_, err := io.Copy(w, r)
	logger.Errorf("from %v to %v send data error: %v\n", r.RemoteAddr(), w.RemoteAddr(), err)
	wg.Done()
}
