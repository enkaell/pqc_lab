package main

import (
	"log"
	"net"

	"github.com/enkaell/pqc_lab/internals/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	serverRegister := grpc.NewServer()
	server.RegisterAlgorithmServiceServer(serverRegister)

}
