package web

import (
	"encoding/gob"
	"strings"
)

func init() {
	gob.Register(CreatePostForm{})
	gob.Register(FormErrors{})
}

type FormErrors map[string]string

type CreatePostForm struct {
	Title   string
	Content string

	Errors FormErrors
}

func (f *CreatePostForm) Validate() bool {
	f.Errors = FormErrors{}

	f.Title = strings.TrimSpace(f.Title)
	f.Content = strings.TrimSpace(f.Content)

	if f.Title == "" {
		f.Errors["Title"] = "Please enter a title."
	}

	if f.Content == "" {
		f.Errors["Content"] = "Please enter a text."
	}

	return len(f.Errors) == 0
}
