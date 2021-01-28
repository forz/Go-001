package chann

import (
	"Go-001/Week09/protocol"
	"bufio"
	"fmt"
	"net"
	"sync"
)

// Channel used by message pusher send msg to write goroutine.
type Channel struct {
	Writer *bufio.Writer
	Reader *bufio.Reader
	signal chan *protocol.Proto
	mutex  sync.RWMutex
}

func (ch *Channel) Read() {
	for {
		p := &protocol.Proto{}
		err := p.Read(ch.Reader)
		if err != nil {
			fmt.Println("读请求失败", err)
			return
		}
		fmt.Println("收到消息:", string(p.Body))
		ch.signal <- p
	}
}

func (ch *Channel) Write() {
	for p := range ch.signal {
		ch.mutex.Lock()
		p.Body = []byte(fmt.Sprintf("server:返回消息:%d", p.Seq))
		err := p.Write(ch.Writer)
		if err != nil {
			fmt.Println(err)
			return
		}
		ch.mutex.Unlock()
	}
}

func New(conn net.Conn) *Channel {
	c := new(Channel)
	c.Writer = bufio.NewWriter(conn)
	c.Reader = bufio.NewReader(conn)
	c.signal = make(chan *protocol.Proto, 5)
	return c
}
