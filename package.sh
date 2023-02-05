mkdir bin
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/filterList src/handlers/filterList/filterList.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getFullList src/handlers/getFullList/getFullList.go