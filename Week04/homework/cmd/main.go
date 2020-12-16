package main

import (
	"Go-001/Week04/apis/homework/pb"
	"Go-001/Week04/homework/infra/gorm"
	"Go-001/Week04/homework/pkg/server"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	gctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(gctx)
	// grpc server
	g.Go(func() error {
		homesvc := InitHomeworkService(gorm.GormDB)
		svc := server.New(homesvc)
		listener, err := net.Listen("tcp", ":8888")
		if err != nil {
			panic(err)
		}
		gServer := grpc.NewServer()
		pb.RegisterHomeworkServer(gServer, svc)
		reflection.Register(gServer)
		go func() {
			<-ctx.Done()
			fmt.Println("grpc ctx done")
			gServer.GracefulStop()
		}()
		return gServer.Serve(listener)
	})
	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				fmt.Println("receive signal")
				cancel()
				return nil
			}
		}
	})

	err := g.Wait() // first error return
	fmt.Println(err)

}
