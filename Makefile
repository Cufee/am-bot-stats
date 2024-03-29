SERVICE := am-bot-stats
NAMESPACE := aftermath-bot
REGISTRY := ghcr.io/byvko-dev
# 
VERSION = $(shell git rev-parse --short HEAD)
TAG := ${REGISTRY}/${SERVICE}

echo:
	@echo "Tag:" ${TAG}

pull:
	git pull

build:
	go mod tidy
	go mod vendor
	docker build -t ${TAG}:${VERSION} -t ${TAG}:latest .
	docker image prune -f

push:
	docker push ${TAG}:latest

restart:
	kubectl rollout restart statefulset/${SERVICE} -n ${NAMESPACE}