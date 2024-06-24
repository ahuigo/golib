package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TcpServer(port string) {
	process := func(conn net.Conn) {
		defer conn.Close()
		for {
			var buf [128]byte
			n, err := conn.Read(buf[:])
			if err != nil {
				fmt.Println("Read from tcp server failed,err:", err)
				break
			} else {
				data := string(buf[:n])
				res := "hello, " + data
				fmt.Printf("Recived from client,data:%s\n", data)
				conn.Write([]byte(res))
			}
		}
	}
	// 监听TCP 服务端口
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("Listen tcp server failed,err:", err)
		return
	}

	for {
		// 建立socket连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listen.Accept failed,err:", err)
			continue
		}

		// 业务处理逻辑
		go process(conn)
	}
}

func TestTcpClient(t *testing.T) {
	port := "9090"
	go TcpServer(port)
	time.Sleep(time.Second)
	// internetSocket /opt/homebrew/Cellar/go/1.22.1/libexec/src/net/tcpsock_posix.go:85
	// /opt/homebrew/Cellar/go/1.22.1/libexec/src/syscall/syscall_bsd.go:161 sockaddr
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println("Connect to TCP server failed ,err:", err)
		return
	}

	// 响应服务端信息
	_, err = conn.Write([]byte("jackson"))
	if err != nil {
		fmt.Println("Write failed,err:", err)
	} else {
		// var b *bytes.Buffer = bytes.NewBuffer([]byte(""))
		// _, err = io.Copy(b, conn) // copy b to stdout
		b := make([]byte, 10)
		n, err := conn.Read(b) //non-block
		if err != nil {
			fmt.Println("io.Copy failed,err:", err)
		} else {
			fmt.Printf("response from server(%d):%s\n", n, string(b[:n]))
		}
	}

}
