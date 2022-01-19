set GOOS=linux
set GOARCH=amd64
CGO_ENABLE=0 go build -o  rename main.go