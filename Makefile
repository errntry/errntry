.PHONY: ;
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

export GO111MODULE=on

GOIMPORTS=goimports
BINARY_NAME=errntry
DOCKER_TAG=`date -u +%Y%m%d`
GOTEST=go test

all: dep
dep: dep-js dep-go
test: test-js test-go
test-go: run-test-go
test-short: run-test-short
test-all: test
fmt: fmt-js fmt-go
upgrade: dep-upgrade-js dep-upgrade
generate:
	go generate ./pkg/... ./internal/...
	go fmt ./pkg/... ./internal/...
dep-list:
	go list -u -m -json all | go-mod-outdated -direct
dep-js:
	(cd frontend &&	npm install --silent --legacy-peer-deps && npm audit fix)
dep-go:
	go build -v ./...
dep-upgrade:
	go get -u -t ./...
dep-upgrade-js:
	(cd frontend &&	npm --depth 3 update --legacy-peer-deps)
build-js:
	(cd frontend &&	env NODE_ENV=production npm run build)
build-race:
	rm -f $(BINARY_NAME)
	scripts/build.sh race $(BINARY_NAME)
watch-js:
	(cd frontend &&	env NODE_ENV=development npm run watch)
test-js:
	$(info Running JS unit tests...)
	(cd frontend &&	env NODE_ENV=development BABEL_ENV=test npm run test)
run-test-short:
	$(info Running short Go unit tests in parallel mode...)
	$(GOTEST) -parallel 2 -count 1 -cpu 2 -short -timeout 5m ./pkg/... ./internal/...
run-test-go:
	$(info Running all Go unit tests...)
	$(GOTEST) -parallel 1 -count 1 -cpu 1 -tags slow -timeout 20m ./pkg/... ./internal/...
test-parallel:
	$(info Running all Go unit tests in parallel mode...)
	$(GOTEST) -parallel 2 -count 1 -cpu 2 -tags slow -timeout 20m ./pkg/... ./internal/...
test-verbose:
	$(info Running all Go unit tests in verbose mode...)
	$(GOTEST) -parallel 1 -count 1 -cpu 1 -tags slow -timeout 20m -v ./pkg/... ./internal/...
test-race:
	$(info Running all Go unit tests with race detection in verbose mode...)
	$(GOTEST) -tags slow -race -timeout 60m -v ./pkg/... ./internal/...
test-codecov:
	$(info Running all Go unit tests with code coverage report for codecov...)
	go test -parallel 1 -count 1 -cpu 1 -failfast -tags slow -timeout 30m -coverprofile coverage.txt -covermode atomic ./pkg/... ./internal/...
	scripts/codecov.sh -t $(CODECOV_TOKEN)
test-coverage:
	$(info Running all Go unit tests with code coverage report...)
	go test -parallel 1 -count 1 -cpu 1 -failfast -tags slow -timeout 30m -coverprofile coverage.txt -covermode atomic ./pkg/... ./internal/...
	go tool cover -html=coverage.txt -o coverage.html
clean:
	rm -f $(BINARY_NAME)
	rm -f *.log
	rm -rf node_modules
	rm -rf storage/testdata
	rm -rf storage/backups
	rm -rf storage/cache
	rm -rf frontend/node_modules
lint-js:
	(cd frontend &&	npm run lint)
fmt-js:
	(cd frontend &&	npm run fmt)
fmt-imports:
	goimports -w pkg internal
fmt-go:
	go fmt ./pkg/... ./internal/...
tidy:
	go mod tidy