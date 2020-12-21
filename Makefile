dev:
	mkdir -p bin
	go build -o bin/server main.go
	./bin/server

docker-build:
	docker build -f Dockerfile . -t docker.pkg.github.com/zhaogaolong/latency/server:latest --no-cache
