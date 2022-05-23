build-rest:
	go build -o ./bin/cmd-rest/ ./cmd/rest/main.go
	cp ./configs/config.yaml ./bin/cmd-rest/

build-gql:
	go build -o ./bin/cmd-gql/ ./cmd/gql/server.go
	cp ./configs/config.yaml ./bin/cmd-gql/

build-all: build-rest build-gql

clean:
	rm -rf bin/
