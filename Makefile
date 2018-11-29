default: build

build: client/node_modules generate bin/example-server

generate: server/pb/example.pb.go client/src/_proto/example.pb.js

server/pb/example.pb.go: server/pb/example.proto
	@cd server; go generate

client/src/_proto/example.pb.js: server/pb/example.proto client/node_modules
	@mkdir -p client/src/_proto
	@protoc \
		-I server/pb \
		--plugin="protoc-gen-ts=$(CURDIR)/client/node_modules/.bin/protoc-gen-ts" \
		--js_out="import_style=commonjs,binary:$(CURDIR)/client/src/_proto" \
		--ts_out="service=true:$(CURDIR)/client/src/_proto" \
		example.proto

client/node_modules: # only runs on first build because then it exists
	@cd client; npm install

bin/example-server: $(shell find server -type f)
	@cd server; go build -o ../bin/example-server

run: build
	@sh -c "'$(CURDIR)/scripts/run.sh'"

clean:
	@rm -rf bin client/node_modules