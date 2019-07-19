package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "sdkeji/person/pkg/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("dail failed. err:[%v]\n", err)
		return
	}
	client := pb.NewPersonsClient(conn)
	res, err := client.CreatePerson(context.Background(), &pb.CreatePersonRequest{
		Name:   "spider2001",
		Age:    21,
		Gender: 2,
	})
	if err != nil {
		fmt.Errorf("client person failed.err: [%v]", err)
		return
	}
	fmt.Printf("message from server: %v\n", res.GetId())
}
