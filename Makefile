build:
	meta-g && wire ./app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o meta main.go
run:
	go run main.go