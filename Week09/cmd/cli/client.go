package main

import (
	"Go-001/Week09/protocol"
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	wr := bufio.NewWriter(conn)
	rd := bufio.NewReader(conn)
	go func() {
		for i := 0; i < 10; i++ {
			proto := protocol.Proto{
				Ver:       1,
				Operation: 1,
				Seq:       int32(i),
				Body:      []byte(fmt.Sprintf("client:发送消息,当前为:%d", i)),
			}

			if err = proto.Write(wr); err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
	for {
		var p protocol.Proto
		if err = p.Read(rd); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("收到消息:", string(p.Body))
	}
}
