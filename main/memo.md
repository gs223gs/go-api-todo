air install
```bash
cd /v1/rest
go install github.com/air-verse/air@latest
go mod init github.com/gs223gs/go-webapi-todo
go mod tidy
export PATH=$PATH:$(go env GOPATH)/bin
air init
```