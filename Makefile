postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Maliborh521908 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root yatzes_db

dropdb:
	docker exec -it postgres16 dropdb yatzes_db

migrateup:
	migrate -path db/migration -database "postgresql://root:Maliborh521908@localhost:5432/yatzes_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:Maliborh521908@localhost:5432/yatzes_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  proto/*.proto

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test proto