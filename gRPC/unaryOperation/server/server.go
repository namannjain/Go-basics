package main

// import "google.golang.org/grpc/profiling/proto"
import (
	"context"
	"errors"
	"fmt"
	"net"
	proto "unaryOperationGrpc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	grpcServer := grpc.NewServer() //engine
	proto.RegisterExampleServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) ServerReply(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("receive request from client", req.Req)
	fmt.Println("hello from server")
	return &proto.HelloResponse{}, errors.New("")
}
