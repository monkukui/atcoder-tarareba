package main

import (
	"fmt"
	"log"
	"net"

	service "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/service"
	pb "github.com/monkukui/atcoder-tarareba/tarareba-competition-history/tarareba_competition_history_pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("tarareba-competition-history run..")

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
