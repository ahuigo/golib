package server

import (
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var _path404 = ""

func staticFsHandler(r *gin.Engine, path404 string) {
	_path404 = path404
	// redirect /index.html to /
	// r.Static("/", "./")

	staticFS(&r.RouterGroup, "/", gin.Dir("./", false))
}

func staticFS(group *gin.RouterGroup, relativePath string, fs http.FileSystem) {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := createStaticHandler(group, relativePath, fs)
	urlPattern := path.Join(relativePath, "/*filepath")

	// Register GET and HEAD handlers
	group.GET(urlPattern, handler)
	group.HEAD(urlPattern, handler)
}

func createStaticHandler(group *gin.RouterGroup, relativePath string, fs http.FileSystem) gin.HandlerFunc {
	absolutePath := joinPaths(group.BasePath(), relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *gin.Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open(file)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			if(_path404==""){
				return
			}
			c.Request.URL.Path = _path404
			f, err := fs.Open(_path404)
			if err != nil {
				return
			}
			defer f.Close()
			io.Copy(c.Writer, f)
			return

			// c.Request.URL.Path =  "/a/404.html"
		} else {
			f.Close()
		}

		// Replace `/index.html` with `/` to stop 301 redirect
		if strings.HasSuffix(c.Request.URL.Path, "/index.html") {
			c.Request.URL.Path = strings.TrimSuffix(c.Request.URL.Path, "index.html")
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
