package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewTararebaServiceClient(conn)

	var userId string
	fmt.Print("AtCoder ID >> ")
	fmt.Scan(&userId)
	message := &pb.GetCompetitionHistoryRequest{UserId: userId}
	res, err := client.GetCompetitionHistory(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")
	fmt.Println("")

	if err == nil {
		fmt.Println(res)
	}
}
