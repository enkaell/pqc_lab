package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/enkaell/pqc_lab/internals/repository"
	"github.com/enkaell/pqc_lab/internals/server"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type myAlgorithmServiceServer struct {
	server.UnimplementedAlgorithmServiceServer
}

type structError struct{}

func (m *structError) Error() string {
	return "Struct init error"
}

// Main RPC's signatures

func (s *myAlgorithmServiceServer) GetAllAlgorithms(context.Context, *emptypb.Empty) (*server.AlgorithmList, error) {
	algorithms, err := repository.DBGetAllAlgorithms()
	if err != nil {
		log.Fatal("Database fetching error: ", err)

	}
	response := &server.AlgorithmList{Algorithms: algorithms}
	return response, nil
}

func (s myAlgorithmServiceServer) GetAllAlgVersionByAlgName(ctx context.Context, req *server.Request) (*server.AlgorithmVersionList, error) {
	fmt.Printf("Res information: %v: \n", req)
	versions, err := repository.DBGetGetAllAlgVersionByAlgName(req.GetName())
	if err != nil {
		log.Fatal("Database fetching error: ", err)
	}
	response := &server.AlgorithmVersionList{Versions: versions}
	return response, nil
}

func (s myAlgorithmServiceServer) GetAllAlgRunsByVersion(ctx context.Context, req *server.Request) (*server.AlgorithmVersionRunList, error) {
	fmt.Printf("Res information: %v: \n", req)
	runs, err := repository.DBGetAllAlgRunsByVersion(req.GetName())
	if err != nil {
		log.Fatal("Database fetching error: ", err)
	}
	response := &server.AlgorithmVersionRunList{Runs: runs}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	database, err := repository.DBInitDatabaseConn()
	if err != nil {
		log.Fatalf("Cannot connect to database: %s", err)
	}
	defer database.Close()
	grpcServer := grpc.NewServer()
	service := &myAlgorithmServiceServer{}
	server.RegisterAlgorithmServiceServer(grpcServer, service)
	fmt.Println("Server just started")
	grpcServer.Serve(lis)

}
