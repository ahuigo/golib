package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// fileReadHandler _
func fileReadHandler(c *gin.Context) {
	// filepath := c.DefaultQuery("path", "tmp/a.txt")
	filepath := c.DefaultQuery("path", "tmp/new.json")
	if strings.HasSuffix(filepath, ".json") {
		c.Writer.Header().Set("Content-type", "application/json")
	} else {
		c.Writer.Header().Set("Content-type", "text/html; charset=utf-8")
	}

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	io.Copy(c.Writer, f)
}

func fileWriteHandler(c *gin.Context) {
	name := c.PostForm("name")
	file, err := c.FormFile("file1")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]
	filelist := []string{}
	for _, file := range files {
		filelist = append(filelist, file.Filename)
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "./tmp/upload.txt"); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     gin.H{"name": name},
		"file1":    filename,
		"filelist": filelist,
		"msg":      "upload to ./tmp/upload.txt",
	})
}
