COVERAGE_OUT := coverage.out
COVERAGE_HTML := coverage.html

.PHONY: all build fmt fmt-check vet test test-verbose test-race \
	cover cover-html cover-func doc lint tidy verify clean

all: fmt vet test

## Build the package.
build:
	go build ./...

## Reformat source files in place.
fmt:
	go fmt ./...

## Fail if any file is not gofmt-formatted, without modifying files.
fmt-check:
	@unformatted="$$(gofmt -l .)"; \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not gofmt-formatted:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

## Run go vet static analysis.
vet:
	go vet ./...

## Run tests.
test:
	go test ./...

## Run tests with verbose output.
test-verbose:
	go test -v ./...

## Run tests with the race detector.
test-race:
	go test -race ./...

## Run tests with coverage instrumentation and print a summary.
cover:
	go test ./... -coverprofile=$(COVERAGE_OUT)
	go tool cover -func=$(COVERAGE_OUT)

## Run tests with coverage and print per-function coverage only (no test output).
cover-func: cover

## Run tests with coverage and open an HTML report.
cover-html: cover
	go tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)
	@echo "Coverage report written to $(COVERAGE_HTML)"

## Show package documentation as rendered by go doc.
doc:
	go doc -all .

## Run staticcheck if it is installed, otherwise skip.
lint:
	@if command -v staticcheck >/dev/null 2>&1; then \
		staticcheck ./...; \
	else \
		echo "staticcheck not installed, skipping"; \
	fi

## Tidy go.mod/go.sum.
tidy:
	go mod tidy

## Run the full verification suite used before considering the repo done.
verify: fmt-check vet test test-race cover lint

## Remove build/test artifacts.
clean:
	go clean ./...
	rm -f $(COVERAGE_OUT) $(COVERAGE_HTML)
