.PHONY: postgres migrate rollback web

postgres:
	docker run --rm -it -v goreddit_pg:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres:12

migrate:
	migrate -source file://migrations -database postgres://postgres:secret@localhost/postgres?sslmode=disable up

rollback:
	migrate -source file://migrations -database postgres://postgres:secret@localhost/postgres?sslmode=disable down

web:
	reflex -s go run ./cmd/goreddit/main.go
