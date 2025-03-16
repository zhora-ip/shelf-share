VERSION=
USER=
PASSWORD=
NAME=
STEP=
HOST=172.31.122.239
PORT=5433
DB=shelfshare


.PHONY: build
build:
	go build -v ./cmd/apishelfshare

.PHONY: test
test:
	go test -v -race -timeout 30s ./...


.PHONY: migration_create
migration_create:
	migrate create -ext sql -dir database/migration/ -seq ${NAME}

.PHONY: migration_up
migration_up:
	migrate -path database/migration/ -database "postgresql://${USER}:${PASSWORD}@${HOST}:${PORT}/${DB}?sslmode=disable" -verbose up 

.PHONY: migration_down
migration_down:
	migrate -path database/migration/ -database "postgresql://${USER}:${PASSWORD}@${HOST}:${PORT}/${DB}?sslmode=disable" -verbose down ${STEP}

.PHONY: migration_force
migration_force:
	migrate -path database/migration/ -database "postgresql://${USER}:${PASSWORD}@${HOST}:${PORT}/${DB}?sslmode=disable" force ${VERSION}

.DEFAULT_GOAL := build

