package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main5() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //例如链接中止
			continue
		}
		go handleConn5(conn) //一次处理一个连接
	}
}

func handleConn5(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:05:05\n"))
		if err != nil {
			return //例如链接断开
		}
		time.Sleep(1 * time.Second)
	}
}
