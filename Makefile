run:
	go run cmd/main.go

mock:
	mockgen -source=internal/port/db.go -destination=internal/database/mocks/db_mock.go -package=mocks

test:
	go test ./...

up:
	docker-compose up

down:
	docker-compose down


