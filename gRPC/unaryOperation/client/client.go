package main

import (
	"context"
	"net/http"
	proto "unaryOperationGrpc/protoc"

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
	r.GET("/send-message-to-server/:message", clientServerConnection)
	r.Run(":8000")
}

func clientServerConnection(c *gin.Context) {
	newMessage := c.Param("message")

	req := &proto.HelloRequest{Req: newMessage}
	client.ServerReply(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully to server" + newMessage,
	})
}
