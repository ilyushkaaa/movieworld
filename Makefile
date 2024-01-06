.PHONY: build
build:
	go build -o ./service_review_start ./service_review/cmd/service_review/main.go
	go build -o ./service_auth_start ./service_auth/cmd/service_auth/main.go
	go build -o ./app_start .app/cmd/app/main.go

.PHONY: lint
lint:
	chmod +x run_lint.sh
	./run_lint.sh

