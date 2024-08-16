package fxdemo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"go.uber.org/fx"
)

func NewLogger(lc fx.Lifecycle) *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Start logger.")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stop logger.")
			return nil
		},
	})
	return logger
}

func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger.Print("Got a request.")
		fmt.Fprintln(w, "Hello Robby")
	}), nil
}

// Here, NewMux makes an HTTP mux available to other functions. Since
// constructors are called lazily, we know that NewMux won't be called unless
// some other function wants to register a handler. This makes it easy to use
// Fx's Lifecycle to start an HTTP server only if we have handlers registered.
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	// First, we construct the mux and server. We don't want to start the server until all handlers are registered.
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// Hooks are executed in dependency order. At startup, NewLogger's hooks run
	// before NewMux's. On shutdown, the order is reversed.
	//
	// Returning an error from OnStart hooks interrupts application startup. Fx
	// immediately runs the OnStop portions of any successfully-executed OnStart hooks (so that types which started cleanly can also shut down cleanly), then exits.
	//
	// Returning an error from OnStop hooks logs a warning, but Fx continues to run the remaining hooks.
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks.
		// By default, hooks have a total of 15 seconds to complete. Timeouts are passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// Register mounts our HTTP handler on the mux.
// Unlike constructors, invocations are called eagerly. See the main function below for details.
func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

func TestMux2(t *testing.T) {
	app := fx.New(
		// Remember that constructors are called lazily, so this block doesn't do much on its own.
		fx.Provide(
			NewLogger,
			NewHandler,
			NewMux,
		),
		// Since constructors are called lazily, we need some invocations to kick-start our application.
		fx.Invoke(Register),

		// This is optional. With this, you can control where Fx logs
		// its events. In this case, we're using a NopLogger to keep
		// our test silent. Normally, you'll want to use an
		// fxevent.ZapLogger or an fxevent.ConsoleLogger.
		// fx.WithLogger(
		// 	func() fxevent.Logger {
		// 		// return &fxevent.ZapLogger{glogger.GetLogger("proj", glogger.DebugLevel).Desugar()} //disable logger
		// 		return fxevent.NopLogger //disable logger
		// 	},
		// ),
	)
	fmt.Println("3. app.Start...")

	// In a typical application, we could just use app.Run() here.
	// Since we don't want this example to run forever, we'll use the more-explicit Start
	// and Stop.
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// app.Run() //Run() 阻塞, Start() 不阻塞
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	// Normally, we'd block here with <-app.Done(). Instead, we'll make an HTTP
	// request to demonstrate that our server is running.
	if resp, err := http.Get("http://localhost:8080"); err != nil {
		log.Fatal(err)
	} else {
		bytes, _ := ioutil.ReadAll(resp.Body)
		log.Printf("body:%#v\n", string(bytes))
	}

	fmt.Println("4. app.Stop...")
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}

}
