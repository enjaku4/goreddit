package web

import (
	"encoding/gob"
	"strings"
)

func init() {
	gob.Register(CreateCommentForm{})
	gob.Register(CreatePostForm{})
	gob.Register(CreateThreadForm{})
	gob.Register(FormErrors{})
	gob.Register(RegisterForm{})
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

type RegisterForm struct {
	Username      string
	Password      string
	UsernameTaken bool

	Errors FormErrors
}

func (f *RegisterForm) Validate() bool {
	f.Errors = FormErrors{}

	f.Username = strings.TrimSpace(f.Username)
	f.Password = strings.TrimSpace(f.Password)

	if f.Username == "" {
		f.Errors["Username"] = "Please enter a username."
	} else if f.UsernameTaken {
		f.Errors["Username"] = "This username is already taken."
	}

	if f.Password == "" {
		f.Errors["Password"] = "Please enter a password."
	} else if len(f.Password) < 8 {
		f.Errors["Password"] = "Your password must be at least 8 characters long"
	}

	return len(f.Errors) == 0
}
