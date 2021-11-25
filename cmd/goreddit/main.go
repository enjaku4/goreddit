package main

import (
	"log"
	"net/http"

	"github.com/enjaku4/goreddit/postgres"
	"github.com/enjaku4/goreddit/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	csrfKey := []byte("01234567890123456789012345678901")
	h := web.NewHandler(store, csrfKey)
	http.ListenAndServe(":3000", h)
}
