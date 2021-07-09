export BUILD_INFO_PKG="github.com/RashadAnsari/gola/pkg/version"

export LDFLAGS="-w -s -X '$(BUILD_INFO_PKG).AppVersion=latest' -X '$(BUILD_INFO_PKG).Date=$$(date)' -X '$(BUILD_INFO_PKG).BuildVersion=$$(git rev-parse HEAD | cut -c 1-8)' -X '$(BUILD_INFO_PKG).VCSRef=$$(git rev-parse --abbrev-ref HEAD)'"

build:
	go build -ldflags $(LDFLAGS) ./cmd/gola

install:
	go install -ldflags $(LDFLAGS) ./cmd/gola

format:
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -s -w {} +
	find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w  -local github.com/RashadAnsari {} +

lint:
	golangci-lint -c build/ci/.golangci.yml run ./...

test:
	go test -ldflags $(LDFLAGS) -v -race -p 1 `go list ./...

ci-test:
	go test -ldflags $(LDFLAGS) -v -race -p 1 -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func coverage.txt