package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Genre       string  `json:"genre"`
	Description string  `json:"description"`
	AvgGrade    float32 `json:"avg_grade"`
	Format      string  `json:"format"`
	S3ID        int     `json:"-"`
	CreatedBy   int     `json:"created_by"`
}

func (b *Book) Validate() error {
	return validation.ValidateStruct(
		b,
		validation.Field(&b.Title, validation.Required),
		validation.Field(&b.Author, validation.Required),
		validation.Field(&b.Genre, validation.Required),
		validation.Field(&b.Description, validation.Required),
	)
}
