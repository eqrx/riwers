export BUILD_DATE=$(shell if hash gdate 2>/dev/null; then gdate --rfc-3339=seconds | sed 's/ /T/'; else date --rfc-3339=seconds | sed 's/ /T/'; fi)
export GITCOMMIT=$(shell git log -1 --pretty=format:"%H")
export GOLDFLAGS=-s -w -extldflags '-zrelro -znow' -X go.eqrx.net/mauzr.version=$(GITTAG) -X go.eqrx.net/mauzr.commit=$(GITCOMMIT) -X go.eqrx.net/mauzr.date=$(BUILD_DATE)
export GOFLAGS=-trimpath
export GITTAG=$(shell git describe --tags --always)
export IMAGE=docker.pkg.github.com/eqrx/mauzr/mauzr

.PHONY: all
all: build

.PHONY: dist/arm64/mauzr
dist/arm64/mauzr:
	GOARCH=arm64 go build -trimpath -ldflags "$(GOLDFLAGS)" -o $@ ./cmd/mauzr

.PHONY: dist/amd64/mauzr
dist/amd64/mauzr:
	GOARCH=amd64 go build -trimpath -ldflags "$(GOLDFLAGS)" -o $@ ./cmd/mauzr

.PHONY: build
build: dist/arm64/mauzr dist/amd64/mauzr

.PHONY: benchmark
benchmark:
	go test -bench=. ./...

.PHONY: test
test:
	go test -race ./...

.PHONY: vet
vet:
	golangci-lint run ./...

.PHONY: download
download:
	go mod download

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: build-image
build-image: build
	./build/buildah.sh

.PHONY: push-image
push-image:
	buildah manifest push --all $(IMAGE):$(GITTAG) docker://$(IMAGE):$(GITTAG)
