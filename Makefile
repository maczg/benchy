BINARY_NAME=benchy
VERSION ?= 0.0.1
IMAGE_TAG_BASE ?= quay.io/massigollo/benchy
IMG ?= $(IMAGE_TAG_BASE):$(VERSION)

.PHONY: build
build:
	@go build -o bin/$(BINARY_NAME) cmd/main.go

.PHONY: build
run: test build
	@bin/$(BINARY_NAME)

.PHONY: test
test:
	@go test -v ./...

.PHONY: clean
clean:
	@rm -rf bin/$(BINARY_NAME)

.PHONY: docker-build
docker-build: test
	@docker build -t ${IMG} .

.PHONY: docker-push
docker-push: docker-build
	@docker push ${IMG}

.PHONY: docker-build-load
docker-build-load:
	@docker build -t $(IMAGE_TAG_BASE):load load-gen

.PHONY: docker-push-load
docker-push-load: docker-build-load
	@docker push $(IMAGE_TAG_BASE):load


deploy:
	@kubectl cluster-info | head -n -2
	@echo "Current ns: $$(kubectl config get-contexts | grep -e "^\*" | awk '{print $$5}')"
	@kubectl apply -f release

destroy:
	kubectl delete -f release
