package main

import (
	"flag"
	"fmt"
	_ "ginapp/conf"
	"ginapp/fslib"
	"ginapp/server"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("p", "4500", "Public Server Port")
	dir := flag.String("d", "", "change directory")
	staticFS := flag.Bool("s", false, "static fs")
	staticBasePath := flag.String("s4", "", "static 404 path, e.g.: /a/404.html")
	flag.Parse()

	// chang directory
	if *dir != "" {
		//home, _ := os.UserHomeDir()
		if err := fslib.Chdir(*dir); err != nil {
			panic(err)
		}
	}

	// gin+port
	engine := gin.New()
	//  https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	server.Register(engine, *staticFS, *staticBasePath)

	//http.Handle("/", r)
	if false {
		engine.Run(":" + *port)
	} else {
		// set timeout
		publicServer := &http.Server{
			Addr: fmt.Sprintf(":%s", *port),
            // https://ieftimov.com/posts/testing-in-go-test-doubles-by-example/
			// Handler: engine,
			// Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "Timeout!\n"),
			Handler: http.TimeoutHandler(engine, 10*time.Second, "Timeout!\n"),

			// ReadTimeout: the maximum duration for reading the entire request, including the body
			ReadTimeout: 2 * time.Second,
			// WriteTimeout: the maximum duration before timing out writes of the response
			WriteTimeout: 2 * time.Second,
			// IdleTimetout: the maximum amount of time to wait for the next request when keep-alive is enabled
			IdleTimeout:       30 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
			// TLSConfig:         tlsConfig,
			MaxHeaderBytes: 1 << 20,
		}
		// service connections
		if err := publicServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}

}
