package main

import (
	"bytes"
	_ "fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var url = "http://localhost:9002/upload"

// Transfer-Encoding: chunked
func uploadChunked(filename string) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()
		part, err := m.CreateFormFile("filename", filename)
		if err != nil {
			return
		}
		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()
	http.Post(url, m.FormDataContentType(), r)
}

//Content-Length
func uploadFile(filename string) (err error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	defer writer.Close()
	part, err := writer.CreateFormFile("myFile", filename)
	if err != nil {
		return err
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = io.Copy(part, file); err != nil {
		return err
	}
	http.Post(url, writer.FormDataContentType(), buf)
	return
}

func main() {
	uploadChunked("README.md")
}
