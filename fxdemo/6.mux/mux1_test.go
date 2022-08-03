package fx

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"go.uber.org/fx"
)

// mystruct 自定义类型
type mystruct struct{}

// NewMyConstruct 返回mystruct结构体实例, 会初始化日志
func NewMyConstruct(logger *log.Logger) mystruct {
	logger.Println("Executing NewMyConstruct.")
	return mystruct{}
}

// NewMux 启动http server
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		// app.start调用
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Print(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// NewHandler 获取handler
func NewHandler(lc fx.Lifecycle, logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	lc.Append(fx.Hook{
		// app.start调用
		OnStart: func(i context.Context) error {
			logger.Println("handler onstart..")
			return nil
		},
		OnStop: func(i context.Context) error {
			logger.Println("handler onstop..")
			return nil
		},
	})

	sayHello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Robby")
	}

	return http.HandlerFunc(sayHello), nil
}

// NewLogger : fx 懒惰的调用 构造器， 仅有其它的函数需要 logger时， 才会调用 Newlogger. 一旦 被初始化，loggere 将会缓存并重用，因此在整个程序内， 其将是一个单实例
func NewLogger(lc fx.Lifecycle) *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")

	lc.Append(fx.Hook{
		// app.start调用
		OnStart: func(i context.Context) error {
			logger.Println("logger onstart..")
			return nil
		},
		// app.stop调用
		OnStop: func(i context.Context) error {
			logger.Println("logger onstop..")
			return nil
		},
	})
	return logger
}

// invokeNothingUse 什么都不做
func invokeNothingUse() {
	fmt.Println("2.1 invokeNothingUse...")
}

// invokeRegister (最重要)将路由配置与handler匹配，初始化ServeMux、HandlerLogger
func invokeRegister(mux *http.ServeMux, logger *log.Logger, h http.Handler) {
	logger.Println("2.2 invokeRegiste...")
	mux.Handle("/", h)
}

func TestMux(t *testing.T) {
	app := fx.New(
		fx.Provide(
			NewMyConstruct,
			NewMux,
			NewLogger,
			NewHandler,
		),

		// 一般来说invokeRegister 按参数顺序初始化。Run() 后也越早触发onStart,  也越晚触发onStop
		// 不过其中参数 mux *http.ServeMux 依赖logger, 最早初始化的变是NewLogger
		fx.Invoke(invokeNothingUse, invokeRegister),
	)

	fmt.Println("3! app.Start()...")
	// 启动container
	// 在典型的应用程序中， 我们仅需要运行 app.Run 。应为我们想要程序是一个http server， 因此我们需要更加复杂的 Start 和 Stop
	app.Run()

	/*
	   如果使用start和stop启动app, 那么程序会停止:http server不会在后台运行
	*/
	//startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()
	//if err := app.Start(startCtx); err != nil {
	//  log.Fatal(err)
	//}

	//http.Get("http://localhost:8080/")

	//// 停止container
	//stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()
	//if err := app.Stop(stopCtx); err != nil {
	//  log.Fatal(err)
	//}

}
