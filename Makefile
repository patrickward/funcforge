# ==================================================================================== #
## :
## HELPERS:
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
## :
## QUALITY CONTROL:
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

## lint: run linters
.PHONY: lint
lint:
	golangci-lint run

# ==================================================================================== #
## :
## DEVELOPMENT:
# ==================================================================================== #

## docs: generate the godoc documentation and serve it on localhost:6060
.PHONY: docs
docs:
	@godoc -http=:6060 -notes="BUG|TODO|NOTE|IMPORTANT" -play

## test: run all tests for the project
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/unit pkg=$1: run all unit tests for the given package
.PHONY: test/unit
test/unit:
	@if [ -z "${pkg}" ]; then echo "pkg is required. It should the path to the package to test"; exit 1; fi
	go test -v -race -buildvcs ${pkg}/...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## test/bench pkg=$1: run all benchmarks for the given package
.PHONY: test/bench
test/bench:
	@if [ -z "${pkg}" ]; then echo "pkg is required. It should the path to the package to test"; exit 1; fi
	go test -v ${pkg} -bench=. -benchmem -run ^$ #gosetup
