.PHONY: postgres migrate rollback web

postgres:
	docker run --rm -it -v goreddit_pg:/var/lib/postgresql/data -p 5433:5432 -e POSTGRES_PASSWORD=secret postgres:12

migrate:
	migrate -source file://migrations -database postgres://postgres:secret@localhost:5433/postgres?sslmode=disable up

rollback:
	migrate -source file://migrations -database postgres://postgres:secret@localhost:5433/postgres?sslmode=disable down

web:
	reflex -s go run ./cmd/goreddit/main.go
