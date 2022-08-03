package main

import (
	"crypto/tls"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 500)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	println("read:", string(buf[:n]))
}
