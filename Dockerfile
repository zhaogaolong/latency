FROM golang:1.15.6-buster as builder
WORKDIR /workspace
ENV GOPROXY https://goproxy.io,direct
COPY . .
RUN go mod vendor -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o server main.go

FROM centos:7.9.2009 as runtime
WORKDIR /
RUN yum install iproute -y
COPY --from=builder /workspace/server /server
ENTRYPOINT ["bash", "-c", "/server"]
