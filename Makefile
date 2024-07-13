# CURRENT_DIR=$(shell pwd)
# DB_URL=postgres://postgres:root@localhost:5432/restaurant_reservation_users?sslmode=disable

# tidy:
# 	go mod tidy
# gen-proto:
# 	./scripts/gen-proto.sh ${CURRENT_DIR}
# exp:
# 	export POSTGRESQL_URL=${DB_URL}

# mig-create:
# 	@migrate create -ext sql -dir migrations -seq ${1}_${2}_${3}

# mig-up:
# 	@migrate -database ${DB_URL} -path migrations -verbose up

# mig-down:
# 	@migrate -database ${DB_URL} -path migrations -verbose down

# mig-force:
# 	@migrate -database ${DB_URL} -path migrations -verbose force ${1}

CURRENT_DIR=$(shell pwd)
DB_URL=postgres://postgres:root@localhost:5432/reservation_service?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}
# proto-gen:
# 	./scripts/gen-proto.sh /home/abbos/go/src/github.com/exem/UserService


run:
	go run cmd/main.go

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