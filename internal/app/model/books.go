package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Genre       string  `json:"genre"`
	Description string  `json:"description"`
	AvgGrade    float32 `json:"-"`
	Format      string  `json:"-"`
	S3Id        int     `json:"-"`
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
