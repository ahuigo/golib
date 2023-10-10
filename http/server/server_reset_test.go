package main

import (
	"log"
	"net"
	"net/http"
	"testing"
)

func handler(w http.ResponseWriter, req *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		log.Fatalf("Hijacking failed: %v", err)
		return
	}
	conn.Close() // 这将模拟 "connection reset by peer" 错误
}

// 当client发起海量请求时，client端会出现大量的connection reset by peer错误
func TestHttpServerHello(t *testing.T) {
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":7777", nil))

	var content = []byte(`hello world`)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(content)
	})
	log.Fatal(http.ListenAndServe(":7777", nil))

}

// 当client发起大量请求时，client端会出现大量的connection reset by peer错误
func TestTcpNull(t *testing.T) {
	handleConnection := func(conn net.Conn) {
		defer conn.Close()
		buffer := make([]byte, 256)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		conn.Close() // 在数据传输过程中关闭连接
	}

	listener, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatal(err)
	}
	n := 0
	for {
		n += 1
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		t.Log("accept connection:", conn.RemoteAddr().String(), n)
		go handleConnection(conn) // 处理连接在协程中
		if n > 1000 {
			break
		}
	}
}
