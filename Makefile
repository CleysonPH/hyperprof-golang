dsn=mysql://root:root@tcp(localhost:3306)/hyperprof

.PHONY: run
run:
	go run cmd/app/app.go

.PHONY: dev
dev:
	air

.PHONY: build
build:
	go build -o bin/app cmd/app/app.go

.PHONY: makemigration
makemigration:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrateup
migrateup:
	migrate -path migrations -database $(dsn) -verbose up

.PHONY: migratedown
migratedown:
	migrate -path migrations -database $(dsn) -verbose down

.PHONY: build
build:
	go build -o bin/app cmd/app/app.go
