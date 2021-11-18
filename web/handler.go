package web

import (
	"html/template"
	"net/http"

	"github.com/enjaku4/goreddit"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func NewHandler(store goreddit.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
		r.Get("/new", h.ThreadsCreate())
		r.Post("/", h.ThreadsStore())
		r.Post("/delete/{id}", h.ThreadsDelete())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store goreddit.Store
}

const threadsListHTML = `
<h1>Threads</h1>
<dl>
{{ range .Threads }}
	<dt><strong>{{ .Title }}</strong></dt>
	<dd>{{ .Description }}</dd>
	<dd>
		<form action="/threads/delete/{{.ID}}" method="POST">
			<button type="submit">Delete</button>
		</form>
	</dd>
{{ end }}
</dl>
<a href="/threads/new">Create thread</a>
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads []goreddit.Thread
	}

	tmpl := template.Must(template.New("").Parse(threadsListHTML))

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Threads: tt})
	}
}

const threadCreateHTML = `
<h1>New thread</h1>
<form action="/threads" method="POST">
	<table>
		<tr>
			<td>Title</td>
			<td><input type="text" name="title" /></td>
		</tr>
		<tr>
			<td>Description</td>
			<td><input type="text" name="description" /></td>
		</tr>
	</table>
	<button type="submit">Create thread</button>
</form>
`

func (h *Handler) ThreadsCreate() http.HandlerFunc {
	tmpl := template.Must(template.New("").Parse(threadCreateHTML))

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) ThreadsStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")

		err := h.store.CreateThread(&goreddit.Thread{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

func (h *Handler) ThreadsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}
