package web

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/enjaku4/goreddit"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CommentHandler struct {
	store    goreddit.Store
	sessions *scs.SessionManager
}

func (h *CommentHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := r.FormValue("content")

		idStr := chi.URLParam(r, "postID")

		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.CreateComment(&goreddit.Comment{
			ID:      uuid.New(),
			PostID:  id,
			Content: content,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.sessions.Put(r.Context(), "flash", "Your comment has been submitted")

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}

func (h *CommentHandler) Vote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c, err := h.store.Comment(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		dir := r.URL.Query().Get("dir")

		if dir == "up" {
			c.Votes++
		} else if dir == "down" {
			c.Votes--
		}

		if err := h.store.UpdateComment(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
