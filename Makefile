VERSION ?= 0.0.1
#VERSION ?= $(git rev-parse --short HEAD)

IMAGE_TAG_BASE ?= quay.io/massigollo/benchy
PLATFORMS ?= linux/amd64
IMG ?= $(IMAGE_TAG_BASE):$(VERSION)

COMMAND ?= simple

.PHONY: test
test:
	@go test -v ./...

.PHONY: run
run:
	@go build -o bin/benchy main.go
	@./bin/benchy ${COMMAND}

.PHONY: build
build: test
	@docker buildx build --platform=${PLATFORMS} --tag ${IMG} .

.PHONY: push
push: build
	@docker push ${IMG}

.PHONY: build-load
build-load:
	@docker build -t $(IMAGE_TAG_BASE):locust-${VERSION} tools/load-gen

.PHONY: push-load
push-load: build-load
	@docker push $(IMAGE_TAG_BASE):locust-${VERSION}

deploy:
	@kubectl cluster-info | head -n -2
	@echo "Current ns: $$(kubectl config get-contexts | grep -e "^\*" | awk '{print $$5}')"
	@kubectl apply -f release/kubernetes

destroy:
	kubectl delete -f release/kubernetes
