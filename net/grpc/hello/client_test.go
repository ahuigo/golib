package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	//pb "github.com/gin-gonic/examples/grpc/pb"
	pb "demo/pb"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

func TestClient(t *testing.T) {
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	req := &pb.HelloRequest{Name: "ahuigo"}
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("output:%v\n", res.Message)
}
