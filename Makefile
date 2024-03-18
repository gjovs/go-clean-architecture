run-app: 
	go run ./cmd/app/
build: 
	env GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64 github.com/gjovs/clean/
generate-grpc:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/proto/*.proto