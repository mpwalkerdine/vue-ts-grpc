package main

import (
	"flag"
	"log"

	"gitlab.com/mpwalkerdine/vue-ts-grpc/server/pb"
	"google.golang.org/grpc"
)

func main() {
	webport := flag.String("webport", "8081", "The port to bind gRPC-web to.")
	port := flag.String("port", "8082", "The port to bind gRPC to.")
	flag.Parse()

	gs := grpc.NewServer()

	s := &server{
		port:       ":" + *port,
		webPort:    ":" + *webport,
		grpcServer: gs,
		examples:   map[string]*pb.Example{},
	}

	log.Fatal(s.Run())
}
