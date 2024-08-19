postgres:
	docker run --rm -p 5433:5432 -d --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:16-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres dropdb simple_bank

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb postgres createdb migrateup migratedown dropdb sqlc test

#start: postgres createdb migrateup


