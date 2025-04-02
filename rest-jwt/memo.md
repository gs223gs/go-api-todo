air install
```bash
cd /v1/rest

go install github.com/air-verse/air@latest

go mod tidy

export PATH=$PATH:$(go env GOPATH)/bin

air init
```