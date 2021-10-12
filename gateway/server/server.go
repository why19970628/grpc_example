package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	echopb "github/why19970628/grpc_example/gateway/proto/echo"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type server struct{
	echopb.UnimplementedEchoServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Echo(ctx context.Context, in *echopb.StringMessage) (*echopb.StringMessage, error) {
	return &echopb.StringMessage{Value: in.Value}, nil
}

func main() {
	registerGRPC()
	registerGateWay()
}

func registerGRPC()  {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}


	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	echopb.RegisterEchoServiceServer(s,  NewServer())
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
}


func registerGateWay()  {

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = echopb.RegisterEchoServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}


	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	log.Fatalln(gwServer.ListenAndServe())
}