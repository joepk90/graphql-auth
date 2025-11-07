mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
base_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))


DOCKER_REGISTRY?=jparkkennaby
SERVICE=graphql-auth
DOCKER_REPOSITORY=$(DOCKER_REGISTRY)/$(SERVICE)

BUILDENV :=
BUILDENV += CGO_ENABLED=0

GIT_HASH := $(GITHUB_SHA)
ifeq ($(GIT_HASH),)
  GIT_HASH := $(shell git rev-parse HEAD)
endif
LINKFLAGS :=-s -X main.gitHash=$(GIT_HASH)
TESTFLAGS := -v -cover
LINT_EXCLUDE=pb.go|pb.gw.go
LINT_FLAGS :=--disable  errcheck --disable staticcheck --timeout=2m

LINTER_EXE := golangci-lint
LINTER := $(GOPATH)/bin/$(LINTER_EXE)

EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
join-with = $(subst $(SPACE),$1,$(strip $2))


LEXC :=
ifdef LINT_EXCLUDE
	LEXC := $(call join-with,|,$(LINT_EXCLUDE))
endif

$(LINTER):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(GOPATH)/bin

.PHONY: lint
lint: $(LINTER)
ifdef LEXC
	$(LINTER) --exclude '$(LEXC)' run $(LINT_FLAGS) ./...
else
	$(LINTER) run $(LINT_FLAGS) ./...
endif

.PHONY: install
install:
	go get -v -t -d ./... 2>&1 | sed -e "s/[[:alnum:]]*:x-oauth-basic/redacted/"

.PHONY: clean
clean:
	rm -f $(SERVICE)

# builds our binary
$(SERVICE): clean
	$(BUILDENV) go build -o $(SERVICE) -a -ldflags '$(LINKFLAGS)' ./cmd/$(SERVICE)

dev:
	. ./.env && go run ./cmd/graphql-auth/main.go

.PHONY: test
test:
	$(BUILDENV) go test $(TESTFLAGS) ./...

.PHONY: all
all: clean $(LINTER) lint test build

docker-build: $(SERVICE)

mocks:
	go get go.uber.org/mock/mockgen
	go install go.uber.org/mock/mockgen
	go generate ./...
