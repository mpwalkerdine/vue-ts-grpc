default: build

generate: server/pb/example.pb.go

server/pb/example.pb.go: server/pb/example.proto
	@cd server; go generate

build: generate bin/example-server

bin/example-server: $(shell find server -type f)
	@cd server; go build -o ../bin/example-server

run: build
	@bin/example-server