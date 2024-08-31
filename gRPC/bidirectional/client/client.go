package main

import (
	proto "bidirectionalGrpc/protoc"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	//connection to internal grpc server
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)

	// req := &proto.HelloRequest{Req: "boht hogai padhai"}
	// client.ServerReply(context.TODO(), req)

	//implement Rest API
	r := gin.Default()
	r.GET("/sent", clientServerConnection)
	r.Run(":8000")
}

func clientServerConnection(c *gin.Context) {

	requests := []*proto.HelloRequest{
		{Req: "first"},
		{Req: "second"},
		{Req: "third"},
		{Req: "fourth"},
		{Req: "fifth"},
		{Req: "sixth"},
	}

	grpcStream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("something error")
		return
	}

	for _, req := range requests {
		err = grpcStream.Send(req)
		if err != nil {
			fmt.Println("request is not fulfilled")
			return
		}
	}

	res, err := grpcStream.CloseAndRecv()
	if err != nil {
		fmt.Println("ther is some error in closing ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message_count": res,
	})
}
