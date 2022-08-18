package main

import (
	"flag"
	"ginapp/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("p", "4500", "Public Server Port")
	dir := flag.String("d", "", "change directory")
	staticFS := flag.Bool("s", false, "static fs")
	staticBasePath := flag.String("sb", "", "static base path, like /a")
	flag.Parse()

	// chang directory
	if *dir != "" {
		//home, _ := os.UserHomeDir()
		if err := os.Chdir(*dir); err != nil {
			panic(err)
		}
	}

	// gin+port
	r := gin.New()
	//  https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	r.SetTrustedProxies([]string{"127.0.0.1"})
	server.Register(r, *staticFS, *staticBasePath)

	//http.Handle("/", r)
	r.Run(":" + *port)
}
