mount:
	[ -d ./bin ] && rm -r bin
	mkdir bin
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/filterList src/handlers/filterList/filterList.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getFullList src/handlers/getFullList/getFullList.go

local:
	sls offline --useDocker --host 0.0.0.0

push:
	sls deploy
