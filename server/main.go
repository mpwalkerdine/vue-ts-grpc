package main

import (
	"flag"
	"log"

	"gitlab.com/mpwalkerdine/vue-ts-grpc/server/pb"
	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "8081", "The port to bind gRPC to.")
	flag.Parse()

	gs := grpc.NewServer()

	s := &server{
		port:       ":" + *port,
		grpcServer: gs,
		examples:   map[string]*pb.Example{},
	}

	log.Fatal(s.Run())
}
