VERSION ?= $(cat VERSION)

win_build:
	go build -o ./bin/server.exe ./cmd/server.go

linux_build:
	go build -o ./bin/server ./cmd/server.go

docker-build:
	docker build -t harbor.devops/cd/backend:$(VERSION) .

release:
	docker build -t harbor.devops/cd/backend:$(VERSION) .
	docker push harbor.devops/cd/backend:$(VERSION)