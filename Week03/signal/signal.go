package signal

import (
	"Go-000/Week03/cycle"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HookSignals(c *cycle.Cycle) {
	sigChan := make(chan os.Signal)
	signal.Notify(
		sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		var sig os.Signal
		for {
			sig = <-sigChan
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Printf("接收信号: %s ,优雅关闭", sig)
				// 不接error,因为函数调用后,main直接返回,error打印不出来.
				_ = c.GracefulStop(context.TODO()) //nolint
			}
			// 3s后可接收新信号,防止信号大量传入.
			time.Sleep(time.Second * 3) //nolint
		}
	}()
}
