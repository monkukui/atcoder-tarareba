package main

import (
	"fmt"
	"log"
	"net"

	service "github.com/monkukui/atcoder-tarareba/tarareba-algorithms/service"
	pb "github.com/monkukui/atcoder-tarareba/tarareba-algorithms/tarareba_algorithms_pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("run..")

	// listenPort, err := net.Listen("tcp", ":19003")
	listenPort, err := net.Listen("tcp", ":19004") // これ、19004 にしました
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	tararebaService := service.NewTararebaService()

	// 実行したい実処理を server に登録する
	pb.RegisterTararebaServiceServer(server, tararebaService)
	server.Serve(listenPort)
}
