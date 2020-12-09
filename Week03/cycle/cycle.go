package cycle

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type Cycle struct {
	servers []*http.Server
}

func (c *Cycle) AddServer(s *http.Server) {
	c.servers = append(c.servers, s)
}
func (c *Cycle) GracefulStop(ctx context.Context) (err error) {
	var eg errgroup.Group
	for _, s := range c.servers {
		ts := s
		eg.Go(func() error {
			return ts.Shutdown(ctx)
		})
	}
	return eg.Wait()
}

func (c *Cycle) runServers() error {
	var eg errgroup.Group
	// start multi servers
	for _, s := range c.servers {
		s := s
		eg.Go(func() (err error) {
			log.Println("服务启动:", s.Addr)
			defer func() {
				log.Println("服务关闭:", s.Addr, "err:", err)
			}()
			return s.ListenAndServe()
		})
	}
	return eg.Wait()
}

func (c *Cycle) Stop() (err error) {
	var eg errgroup.Group
	for _, s := range c.servers {
		s := s
		eg.Go(s.Close)
	}
	return eg.Wait()
}
func (c *Cycle) Run() error {
	var eg errgroup.Group
	eg.Go(c.runServers)
	return eg.Wait()
}
