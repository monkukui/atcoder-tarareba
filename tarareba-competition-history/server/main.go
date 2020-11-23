package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/pb"
	service "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/service"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("run..")

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	tararebaService := service.NewTararebaService()

	// 実行したい実処理を server に登録する
	pb.RegisterTararebaServiceServer(server, tararebaService)
	server.Serve(listenPort)
}
