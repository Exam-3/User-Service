CURRENT_DIR=$(shell pwd)
DB_URL=postgres://postgres:root@localhost:5432/users?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

run:
	go run cmd/service/main.go

migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down

migrate_force:
	migrate -path migrations -database ${DB_URL} -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq insert_to_tables

test:
	go test ./storage/postgres

tidy:
	go mod tidy
	go mod vendor

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs
