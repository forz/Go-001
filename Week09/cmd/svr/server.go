package main

import (
	"Go-001/Week09/chann"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	gw, _ := errgroup.WithContext(ctx)

	gw.Go(func() error {
		return tcpServer(ctx)
	})

	gw.Go(func() error {
		sigs := []os.Signal{os.Interrupt}

		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, sigs...)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-sigch:
				cancelFn()
				return fmt.Errorf("receive signal:%s", sig)
			}
		}
	})

	if err := gw.Wait(); err != nil {
		fmt.Println("exit:", err)
	}
}

func tcpServer(ctx context.Context) error {
	address := ":8888"
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		listen.Close()
	}()
	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		go handConn(conn)
	}
}

func handConn(conn net.Conn) {
	ch := chann.New(conn)
	go ch.Write()
	go ch.Read()
}
