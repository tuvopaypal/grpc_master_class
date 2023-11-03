package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tuvo1106/grpc_master_class/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Tu"},
		{FirstName: "Brenda"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreeet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
