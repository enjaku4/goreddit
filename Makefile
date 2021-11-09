.PHONY: postgres migrate rollback

postgres:
	docker run --rm -ti -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres:12

migrate:
	migrate -source file://migrations -database postgres://postgres:secret@localhost/postgres?sslmode=disable up

rollback:
	migrate -source file://migrations -database postgres://postgres:secret@localhost/postgres?sslmode=disable down
