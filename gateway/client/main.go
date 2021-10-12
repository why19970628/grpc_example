package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "github/why19970628/grpc_example/gateway/proto/echo"
	"log"
	"time"
)

func main() {
	addr := "0.0.0.0:8090"
	ctx, cel := context.WithTimeout(context.Background(), time.Second*2)
	defer cel()
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
	}
	//conn, err := grpc.Dial(":8090")
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewEchoServiceClient(conn)
	context := context.Background()
	body := &pb.StringMessage{
		Value : "Grpc",
	}

	r, err := c.Echo(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Value)
}