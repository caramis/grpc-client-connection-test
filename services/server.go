package services

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/caramis/grpc-client-connection-test/test_supplements/gen_src/test"
	"google.golang.org/grpc"
)

// ServiceServer1 is server for testing
type ServiceServer1 struct {
	Server *grpc.Server
	test.UnimplementedTestServiceServer
}

// Add for Server1 returns "10"
func (s ServiceServer1) Add(context.Context, *test.Request) (*test.Response, error) {
	return &test.Response{Code: 0, Result: 10}, nil
}

// NewServiceServer1 returns new ServiceServer1
func NewServiceServer1() ServiceServer1 {
	serviceServer := ServiceServer1{
		Server: grpc.NewServer(),
	}
	test.RegisterTestServiceServer(serviceServer.Server, serviceServer)

	return serviceServer
}

// ServiceServer2 is server for testing
type ServiceServer2 struct {
	Server *grpc.Server
	test.UnimplementedTestServiceServer
}

// Add for Server2 returns "20"
func (s ServiceServer2) Add(context.Context, *test.Request) (*test.Response, error) {
	return &test.Response{Code: 0, Result: 20}, nil
}

// NewServiceServer2 returns new ServiceServer1
func NewServiceServer2() ServiceServer2 {
	serviceServer := ServiceServer2{
		Server: grpc.NewServer(),
	}
	test.RegisterTestServiceServer(serviceServer.Server, serviceServer)

	return serviceServer
}

// Serve serves TestService
func Serve(grpcServer *grpc.Server, port int) error {
	// Listen
	address := strings.Join([]string{":", strconv.Itoa(port)}, "")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// Serve
	{
		err := grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
