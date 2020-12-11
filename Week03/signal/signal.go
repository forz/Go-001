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
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go func() {
		var sig os.Signal
		sig = <-sigChan
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //nolint
		defer cancel()
		log.Printf("接收信号: %s ,优雅关闭", sig)
		// 不接error,因为函数调用后,main直接返回,error打印不出来.
		_ = c.GracefulStop(ctx) //nolint
	}()
}
