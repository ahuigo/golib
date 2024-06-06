package main

import (
	"log"
	"net"
	"testing"

	//pb "github.com/gin-gonic/examples/grpc/pb"
	pb "demo/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func TestHelloServerWithProxy(t *testing.T) {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(authenticateClient),
	}

	log.Println("listen ", ":50051")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(opts...)

	// Register reflection service on gRPC server.
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestHelloServer(t *testing.T) {
	log.Println("listen ", ":50051")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register reflection service on gRPC server.
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// authenticateClient is a unary interceptor function to handle auth
func authenticateClient(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC req method: %s, %v", info.FullMethod, req)
	checkPasswd := false
	if checkPasswd {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
		}

		// Extract user and password
		user := md["user"]
		password := md["password"]

		// Validate user and password
		if len(user) == 0 || user[0] == password[0] {
			return nil, grpc.Errorf(codes.Unauthenticated, "invalid user or password")
		}
	}
	return handler(ctx, req)
}
