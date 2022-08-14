package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	//pb "github.com/gin-gonic/examples/grpc/pb"
	pb "demo/pb"

	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
    // Refer to: https://github.com/grpc/grpc-go/examples/helloworld
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
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


// curl -v 'http://localhost:8052/rest/n/thinkerou'
func TestGonic(t *testing.T) {
    // Refer to: https://github.com/grpc/grpc-go/examples/helloworld
    // conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// Set up a http server.
	r := gin.Default()
	r.GET("/rest/n/:name", func(c *gin.Context) {
		name := c.Param("name")

		// Contact the server and print out its response.
		req := &pb.HelloRequest{Name: name}
		res, err := client.SayHello(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
