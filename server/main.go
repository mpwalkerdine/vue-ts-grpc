package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "8081", "The port to bind gRPC to.")
	flag.Parse()

	gs := grpc.NewServer()

	s := &server{
		port:       ":" + *port,
		grpcServer: gs,
	}

	log.Fatal(s.Run())
}
