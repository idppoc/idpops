CT_DOCKER_REPO := docker.io

DOCKER_IMAGE_NAME := idpapiserver
DOCKER_IMAGE_VERSION := ${CT_DOCKER_REPO}/madhukirans/${DOCKER_IMAGE_NAME}:$(shell date -u +%Y%m%d%H%M%S)

all: build image push

build:
	cd go/apiserver
	GOOS=linux GOARCH=amd64 go build -o vault-sidecar -ldflags "-s -w"
image:
	cd go && docker build -t ${DOCKER_IMAGE_VERSION}  .
push:
	docker push ${DOCKER_IMAGE_VERSION}
