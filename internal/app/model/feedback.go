package model

import validation "github.com/go-ozzo/ozzo-validation"

type Feedback struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	BookID   int    `json:"book_id"`
	Feedback string `json:"feedback"`
	Grade    int    `json:"grade"`
}

func (f *Feedback) Validate() error {
	return validation.ValidateStruct(
		f,
		validation.Field(&f.UserID, validation.Required),
		validation.Field(&f.BookID, validation.Required),
		validation.Field(&f.Feedback, validation.Required),
		validation.Field(&f.Grade, validation.Required),
	)
}
