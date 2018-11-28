generate: server/pb/example.pb.go

server/pb/example.pb.go: server/pb/example.proto
	@cd server; go generate