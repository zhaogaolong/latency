dev:
	mkdir -p bin
	go build -o bin/server main.go
	./bin/server

docker-image-build:
	docker build -f Dockerfile . -t docker.pkg.github.com/zhaogaolong/latency/server:2020-12-20-1 --no-cache --squash
