package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "my-otel-demo/proto" // 替换为你的模块名

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

// initTracerProvider 初始化并注册 OpenTelemetry Tracer Provider
// 它负责创建和导出 trace 数据
func initTracerProvider(ctx context.Context) (func(context.Context) error, error) {
	// OTLP Collector 的地址
	// 对于 HTTP Exporter，使用基础 URL，不要包含 /v1/traces 路径
	// HTTP exporter 会自动添加 /v1/traces 路径
	const otelCollectorEndpoint = "otel-collector.local"
	// 创建一个 resource 来标识我们的应用
	// service.name 是必须的，它会显示在 Jaeger/Grafana 等后端 UI 中
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("app-name"), // 设置服务名
			semconv.ServiceVersion("v1.0.0"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// 创建 OTLP/HTTP Exporter（因为服务器支持 HTTPS/JSON）
	// 从 curl 测试可以看出，服务器期望 HTTP 请求而不是 gRPC
	traceExporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(otelCollectorEndpoint),
		// HTTP exporter 已经自动使用 HTTPS，不需要 WithInsecure
		otlptracehttp.WithTimeout(10*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// 创建 Tracer Provider
	// 我们使用 BatchSpanProcessor 将 Span 批量发送，提高性能
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(res),
	)

	// 将我们创建的 Tracer Provider 设置为全局的
	otel.SetTracerProvider(tracerProvider)

	// 设置全局的 Propagator，用于跨服务上下文传递
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// 返回一个 shutdown 函数，用于在应用退出时优雅地关闭 Provider
	return tracerProvider.Shutdown, nil
}

// server 结构体实现了我们 proto 中定义的 GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
	tracer trace.Tracer // 用于手动创建 Span
}

// SayHello 是 gRPC 服务的实现
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	// ================= 手动插桩示例 =================
	// 自动插桩（interceptor）已经为我们创建了一个父 Span
	// 我们可以在这个函数内部创建子 Span 来追踪更具体的操作
	_, childSpan := s.tracer.Start(ctx, "internal-processing")
	defer childSpan.End()

	// 给 Span 添加属性（Attributes），方便查询和分析
	childSpan.SetAttributes(
		attribute.String("request.name", in.GetName()),
		attribute.Int("processing.step", 1),
	)

	// 模拟一些耗时操作
	time.Sleep(50 * time.Millisecond)
	childSpan.AddEvent("Finished processing step 1") // 在 Span 中添加一个事件

	// ===============================================

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	ctx := context.Background()

	// 初始化 Tracer Provider
	shutdown, err := initTracerProvider(ctx)
	if err != nil {
		log.Fatalf("failed to initialize tracer provider: %v", err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown tracer provider: %v", err)
		}
	}()

	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器，并添加 OTEL 拦截器
	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		// grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()), // 自动追踪
	)

	/**
	有反射功能时，可以这样调用：
		grpcurl -plaintext localhost:50051 list
	# 调用具体方法：
	grpcurl -plaintext -d '{"name": "world"}' localhost:50051 greeter.Greeter/SayHello
	*/
	reflection.Register(s)

	// 注册我们的服务实现
	pb.RegisterGreeterServer(s, &server{
		tracer: otel.Tracer("my-grpc-server"), // 获取一个 Tracer 实例
	})

	log.Println("gRPC server listening at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
