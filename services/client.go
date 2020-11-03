package services

import (
	"context"
	"strconv"
	"strings"

	"github.com/caramis/grpc-client-connection-test/test_supplements/gen_src/test"
	"google.golang.org/grpc"
)

// Dial ...
func Dial(port int) (*grpc.ClientConn, error) {
	address := strings.Join([]string{":", strconv.Itoa(port)}, "")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Request ...
func Request(conn *grpc.ClientConn) (int, error) {
	client := test.NewTestServiceClient(conn)
	// It returns static result
	response, err := client.Add(context.Background(), &test.Request{Number1: 1, Number2: 2})
	if err != nil {
		return 0, err
	}
	return int(response.Result), nil
}

// CloseConn ...
func CloseConn(conn *grpc.ClientConn) error {
	return conn.Close()
}
