FROM golang:1.19-alpine3.16 AS builder
WORKDIR /cd
ENV GOPROXY https://goproxy.cn,direct

COPY go.mod .
COPY go.sum .
RUN GOPROXY=https://goproxy.cn go mod download
COPY . .

RUN go build -o ./server ./cmd/server.go

FROM alpine
COPY --from=builder /cd/server .
COPY --from=builder /cd/conf/kubecfg.yaml /conf/kubecfg.yaml
EXPOSE 8080
CMD ["/bin/bash", "-c", "./server"]