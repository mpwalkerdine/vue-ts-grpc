package main

import (
	"context"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"gitlab.com/mpwalkerdine/vue-ts-grpc/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

//go:generate protoc --go_out=plugins=grpc,paths=source_relative:. pb/example.proto

type server struct {
	port       string
	grpcServer *grpc.Server
	examples   map[string]*pb.Example
}

func (s *server) Run() error {
	options := []grpcweb.Option{
		// gRPC-Web compatibility layer with CORS configured to accept on every request
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(string) bool {
			return true
		}),
	}
	wrappedGrpc := grpcweb.WrapServer(s.grpcServer, options...)
	pb.RegisterExampleServiceServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		wrappedGrpc.ServeHTTP(resp, req)
	})
	httpServer := &http.Server{
		Addr:           s.port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	defer httpServer.Shutdown(nil)

	return httpServer.ListenAndServe()
}

func (s *server) CreateExample(ctx context.Context, in *pb.CreateExampleRequest) (*pb.Example, error) {
	example := in.Example
	name := example.Name

	if _, duplicate := s.examples[name]; duplicate {
		return nil, status.Errorf(codes.AlreadyExists, "item with name %q already exists", name)
	}
	s.examples[name] = example

	return example, nil
}

func (s *server) DeleteExample(ctx context.Context, in *pb.DeleteExampleRequest) (*pb.Empty, error) {
	if _, found := s.examples[in.Name]; !found {
		return nil, status.Errorf(codes.NotFound, "item with name %q not found", in.Name)
	}

	delete(s.examples, in.Name)

	return &pb.Empty{}, nil
}

func (s *server) ListExamples(ctx context.Context, in *pb.ListExamplesRequest) (*pb.ListExamplesResponse, error) {
	examples := []*pb.Example{}

	for _, v := range s.examples {
		examples = append(examples, v)
	}

	return &pb.ListExamplesResponse{
		Examples: examples,
	}, nil
}
