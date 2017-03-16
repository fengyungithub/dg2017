GOPATH := ${PWD}/vendor:${PWD}
export GOPATH

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o api github.com/nickschuch/api
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o cli github.com/nickschuch/cli
