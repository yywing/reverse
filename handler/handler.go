package handler

import (
	"io"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Handle(from, to, protocol string, timeout int64) error {
	logger.Infof("%v", from)
	fromConnection, err := net.DialTimeout(protocol, from, time.Duration(timeout*time.Second))
	if err != nil {
		return err
	}
	defer fromConnection.Close()

	logger.Infof("%v", to)
	toConnection, err := net.DialTimeout(protocol, to, time.Duration(timeout*time.Second))
	if err != nil {
		return err
	}
	defer toConnection.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := io.Copy(fromConnection, toConnection)
		logger.Errorf("from %v to %v send data error: %v\n", fromConnection.RemoteAddr(), toConnection.RemoteAddr(), err)

	}()

	go func() {
		defer wg.Done()
		_, err := io.Copy(toConnection, fromConnection)
		logger.Errorf("from %v to %v send data error: %v\n", toConnection.RemoteAddr(), fromConnection.RemoteAddr(), err)
	}()

	// TODO: 当一方连接断了之后另一方会断么？
	wg.Wait()
	logger.Infof("from %v to %v handler exit.", fromConnection.RemoteAddr(), toConnection.RemoteAddr())
	return nil
}
