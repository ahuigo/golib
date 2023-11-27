package server

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// curl -H 'Content-Type: application/json' 'm:4500/bind/f' -d '{"extra":{"name":"ahui", "age":1}}'
func BindServer(c *gin.Context) {
	//backup
	buf, _ := io.ReadAll(c.Request.Body)
	// revert main
	c.Request.Body = io.NopCloser(bytes.NewBuffer(buf)) // important!!

	// bind
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("bind error:", err)
	}
	if user.Name == "" {
		user.Name = c.Query("name")
	}
	fmt.Printf("user:%#v \n", user)
	fmt.Printf("user.time:%v \n", user.Time)

	// c.String(http.StatusOK, res)
	c.JSON(http.StatusOK, user)

}

// curl -X PUT 'm:4500/bind/f' -F file1=@go.mod
func BindFileServer(c *gin.Context) {
	file := struct {
		File1 *multipart.FileHeader `form:"file1" `
	}{}
	if err := c.ShouldBind(&file); err != nil {
		fmt.Println("bind error:", err)
	}
	if file.File1 == nil {
		file.File1, _ = c.FormFile("file1")
	}
	fmt.Println("bind file1:", file.File1)
	c.JSON(http.StatusOK, file)

}

func createMultipartFileHeader(filePath string) (*multipart.FileHeader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("fileField", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	_, params, err := mime.ParseMediaType(writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	boundary, ok := params["boundary"]
	if !ok {
		return nil, errors.New("no boundary")
	}

	reader := multipart.NewReader(body, boundary)
	mf, _ := reader.ReadForm(1 << 8)
	fileHeader := mf.File["file"][0]

	return fileHeader, nil
}
