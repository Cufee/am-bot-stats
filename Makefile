SERVICE := am-bot-stats
NAMESPACE := aftermath
REGISTRY := docker.io/vkouzin
# 
VERSION = $(shell git rev-parse --short HEAD)
TAG := ${REGISTRY}/${SERVICE}

echo:
	@echo "Tag:" ${TAG}

pull:
	git pull

build:
	docker build -t ${TAG}:${VERSION} .
	docker tag ${TAG}:${VERSION} ${TAG}:latest
	docker image prune -f

push:
	docker push ${TAG}:latest

apply:
	kubectl apply -f ./_kube-yml

restart:
	kubectl rollout restart statefulset/${SERVICE} -n ${NAMESPACE}

ctx:
	kubectl config set-context --current --namespace=${NAMESPACE}