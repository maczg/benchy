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

PLATFORMS ?= linux/arm64,linux/amd64
.PHONY: docker-buildx
docker-buildx: test ##
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile > Dockerfile.cross
	- docker buildx create --name benchy-builder
	docker buildx use benchy-builder
	- docker buildx build --push --platform=$(PLATFORMS) --tag ${IMG} -f Dockerfile.cross .
	- docker buildx rm benchy-builder
	#rm Dockerfile.cross

.PHONY: docker-push
docker-push: docker-build
	@docker push ${IMG}

.PHONY: docker-build-load
docker-build-load:
	@docker build -t $(IMAGE_TAG_BASE):load load-gen


LOAD_PLATFORMS ?= linux/amd64
.PHONY: docker-buildx-load
docker-buildx-load:
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' load-gen/Dockerfile > load-gen/Dockerfile.cross
	- docker buildx create --name load-benchy-builder
	docker buildx use load-benchy-builder
	- docker buildx build --push --platform=$(LOAD_PLATFORMS) --tag $(IMAGE_TAG_BASE):load -f load-gen/Dockerfile.cross load-gen
	- docker buildx rm load-benchy-builder


.PHONY: docker-push-load
docker-push-load: docker-build-load
	@docker push $(IMAGE_TAG_BASE):load


deploy:
	@kubectl cluster-info | head -n -2
	@echo "Current ns: $$(kubectl config get-contexts | grep -e "^\*" | awk '{print $$5}')"
	@kubectl apply -f release

destroy:
	kubectl delete -f release
