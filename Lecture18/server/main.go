package main

import (
	"justCode/Lecture18/server/grpc"
	"log"
	"net"

	gogrpc "google.golang.org/grpc"

	protoport "justCode/Lecture18/proto/gw"
)

func main() {
	log.Println("server")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := gogrpc.NewServer()
	protoServer := grpc.Port{}

	protoport.RegisterPortServiceServer(server, protoServer)

	log.Fatalln(server.Serve(listener))
}
