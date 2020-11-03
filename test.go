package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/caramis/grpc-client-connection-test/services"
	"google.golang.org/grpc"
)

func main() {
	wg := new(sync.WaitGroup)

	servers := make([]*grpc.Server, 0)

	const DELAY = 1 * time.Second
	const WAIT = 3 * time.Second

	// ServiceServer1
	{
		// Start ServiceServer1
		wg.Add(1)
		go func() {
			defer wg.Done()

			serviceServer := services.NewServiceServer1()
			servers = append(servers, serviceServer.Server)

			err := services.Serve(serviceServer.Server, 9091)
			if err != nil {
				panic(err)
			}
		}()

		// Dial and RPC to ServiceServer2
		wg.Add(1)
		go func() {
			defer wg.Done()

			conn, err := services.Dial(9092)
			if err != nil {
				panic(err)
			}

			result, err := services.Request(conn)
			if err != nil {
				panic(err)
			}

			fmt.Printf("[Test 1] From ServiceServer2: %d\n", result)

			services.CloseConn(conn)
		}()
	}

	// For dialing test
	time.Sleep(DELAY)

	// ServiceServer2
	{
		// Start ServiceServer2
		wg.Add(1)
		go func() {
			defer wg.Done()

			serviceServer := services.NewServiceServer2()
			servers = append(servers, serviceServer.Server)

			err := services.Serve(serviceServer.Server, 9092)
			if err != nil {
				panic(err)
			}
		}()

		// Dial and RPC to ServiceServer1
		wg.Add(1)
		go func() {
			defer wg.Done()

			conn, err := services.Dial(9091)
			if err != nil {
				panic(err)
			}

			result, err := services.Request(conn)
			if err != nil {
				panic(err)
			}

			fmt.Printf("[Test 2] From ServiceServer1: %d\n", result)

			services.CloseConn(conn)
		}()
	}

	go func() {
		time.Sleep(WAIT)
		for _, server := range servers {
			server.GracefulStop()
		}
	}()

	wg.Wait()
	fmt.Println("Test done !!!")
}
