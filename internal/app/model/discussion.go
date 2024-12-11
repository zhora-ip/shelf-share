package model

import validation "github.com/go-ozzo/ozzo-validation"

type Discussion struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Message struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	DiscussionID int    `json:"discussion_id"`
	Message      string `json:"message"`
}

func (d *Discussion) Validate() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.UserID, validation.Required),
		validation.Field(&d.Title, validation.Required),
		validation.Field(&d.Description, validation.Required),
	)
}

func (m *Message) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.UserID, validation.Required),
		validation.Field(&m.DiscussionID, validation.Required),
		validation.Field(&m.Message, validation.Required),
	)
}
