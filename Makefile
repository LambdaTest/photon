GOBIN=$(GOPATH)/bin
GOSRC=$(GOPATH)/src
u := $(if $(update),-u)

$(GOBIN)/wire:
	GO111MODULE=off go get github.com/google/wire/cmd/wire

.PHONY: wire
wire: $(GOBIN)/wire
	wire gen ./cmd/...

.PHONY: deps
deps:
	go get ${u} -d
	go mod tidy
	go mod vendor

$(GOBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.42.1

.PHONY: lint
lint: $(GOBIN)/golangci-lint
	$(GOBIN)/golangci-lint run --timeout 5m

$(GOBIN)/goreportcard-cli:
	git clone https://github.com/gojp/goreportcard.git
	cd goreportcard
	make install
	go install ./cmd/goreportcard-cli
	cd..
	rm -rf goreportcard

.PHONY: report-card
report-card: $(GOBIN)/goreportcard-cli
	$(GOBIN)/goreportcard-cli -v

.PHONY: build
build:
	./build.sh
