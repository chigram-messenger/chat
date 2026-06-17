build:
	go build -o bin/chat cmd/chat/main.go
run: build

	./bin/chat --config=./configs/config.yaml

run-postgres:
	docker compose down -v
	docker compose up -d