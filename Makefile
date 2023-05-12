build:
	docker-compose build todo

run:
	docker-compose up

migrate:
	migrate -path ./migration -database 'postgres://postgres:postgres@0.0.0.0:5432/todo?sslmode=disable' up