package server

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type Control struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

func ListenControl(addr string) (*Control, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer()
	return &Control{
		listener:   lis,
		grpcServer: grpcServer,
	}, nil
}

func (c *Control) Start() error {
	slog.Info("starting control server", "addr", c.listener.Addr())
	return c.grpcServer.Serve(c.listener)
}

func (c *Control) GracefulStop() {
	slog.Info("gracefully stopping control server")
	c.grpcServer.GracefulStop()
}

func (c *Control) Stop() {
	slog.Info("stopping control server")
	c.grpcServer.Stop()
}
