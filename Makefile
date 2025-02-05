.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: verify
verify: test lint

.PHONY: build
build:
	docker build .

.PHONY: docs
docs:
	go run main.go docs/icons/go.svg --color=blue
	go run main.go docs/icons/go-badge.svg --object=github.com/essentialkaos/go-badge --short-object-name
	go run main.go docs/icons/depexplorer.svg --object=github.com/artarts36/depexplorer --short-object-name
