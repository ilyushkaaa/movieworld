.PHONY: build
build:
	go build -o ./my_app_start ./cmd/app/main.go

.PHONY: lint
lint:
	chmod +x run_lint.sh
	./run_lint.sh

