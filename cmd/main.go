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
	response := &server.AlgorithmList{Algorithms: []*server.Algorithm{{Id: 1, Name: "Algone"}, {Id: 2, Name: "Algtwo"}, {Id: 3, Name: "Algthree"}}}
	return response, nil
}

func (s myAlgorithmServiceServer) GetAllAlgVersionByAlgName(ctx context.Context, req *server.Request) (*server.AlgorithmVersionList, error) {
	fmt.Printf("Res information: %v: \n", req)
	res := &server.AlgorithmVersionList{Versions: []*server.Version{}}
	versions := []*server.Version{
		{Id: 1, AlgorithmId: 1, AlgorithmName: "Algone", Name: "AlgoneVerone"},
		{Id: 2, AlgorithmId: 1, AlgorithmName: "Algone", Name: "Algonevertwo"},
		{Id: 3, AlgorithmId: 1, AlgorithmName: "Algone", Name: "Algoneverthree"},
	}
	for _, ver := range versions {
		if req.Name == ver.AlgorithmName {
			res.Versions = append(res.Versions, ver)
		}

	}
	return res, nil
}

func (s myAlgorithmServiceServer) GetAllAlgRunsByVersion(ctx context.Context, req *server.Request) (*server.AlgorithmVersionRunList, error) {
	fmt.Printf("Res information: %v: \n", req)
	response := &server.AlgorithmVersionRunList{}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	var repo repository.Database = repository.DB{}
	repo, err = repo.initDatabaseConn()

	grpcServer := grpc.NewServer()
	service := &myAlgorithmServiceServer{}
	server.RegisterAlgorithmServiceServer(grpcServer, service)
	fmt.Println("Server just started")
	grpcServer.Serve(lis)

}
