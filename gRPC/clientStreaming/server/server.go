package main

// import "google.golang.org/grpc/profiling/proto"
import (
	proto "bidirectionalGrpc/protoc"
	"fmt"
	"io"
	"net"
	"strconv"

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

func (s *server) ServerReply(grpcStream proto.Example_ServerReplyServer) error {
	total := 0
	for {
		req, err := grpcStream.Recv()
		if err == io.EOF {
			return grpcStream.SendAndClose(&proto.HelloResponse{
				Res: strconv.Itoa(total),
			})
		}

		if err != nil {
			return err
		}

		total += 1
		fmt.Println(req)
	}
}
