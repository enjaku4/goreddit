package web

import (
	"encoding/gob"
	"strings"
)

func init() {
	gob.Register(CreatePostForm{})
	gob.Register(CreateThreadForm{})
	gob.Register(FormErrors{})
	gob.Register(CreateCommentForm{})
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

type CreateThreadForm struct {
	Title       string
	Description string

	Errors FormErrors
}

func (f *CreateThreadForm) Validate() bool {
	f.Errors = FormErrors{}

	f.Title = strings.TrimSpace(f.Title)
	f.Description = strings.TrimSpace(f.Description)

	if f.Title == "" {
		f.Errors["Title"] = "Please enter a title."
	}

	if f.Description == "" {
		f.Errors["Description"] = "Please enter a description."
	}

	return len(f.Errors) == 0
}

type CreateCommentForm struct {
	Content string

	Errors FormErrors
}

func (f *CreateCommentForm) Validate() bool {
	f.Errors = FormErrors{}

	f.Content = strings.TrimSpace(f.Content)

	if f.Content == "" {
		f.Errors["Content"] = "Please enter a comment."
	}

	return len(f.Errors) == 0
}
