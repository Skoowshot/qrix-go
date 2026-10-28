package server

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type Network struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

func ListenNetwork(addr string) (*Network, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer()
	return &Network{
		listener:   lis,
		grpcServer: grpcServer,
	}, nil
}

func (c *Network) Start() error {
	slog.Info("starting network server", "addr", c.listener.Addr())
	return c.grpcServer.Serve(c.listener)
}

func (c *Network) GracefulStop() {
	slog.Info("gracefully stopping network server")
	c.grpcServer.GracefulStop()
}

func (c *Network) Stop() {
	slog.Info("stopping network server")
	c.grpcServer.Stop()
}
