package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Nickname: "Vasiliy",
		Password: "1231231",
		Email:    "vasyliy@mail.ru",
	}
}

func TestUsers(t *testing.T) []*User {
	t.Helper()
	return []*User{
		{
			Nickname: "Vasiliy",
			Password: "password",
			Email:    "vasyliy@mail.ru",
		},
		{
			Nickname: "Andrey",
			Password: "password",
			Email:    "Andrey@mail.ru",
		},
	}
}

func TestBook(t *testing.T) *Book {
	t.Helper()
	return &Book{
		ID:          1,
		Title:       "Title",
		Author:      "Author",
		Genre:       "Genre",
		Description: "Description",
		AvgGrade:    0,
		Format:      "",
		S3ID:        0,
		CreatedBy:   1,
	}
}

func TestBooks(t *testing.T) []*Book {
	t.Helper()
	return []*Book{
		{
			ID:          1,
			Title:       "Title1",
			Author:      "Author1",
			Genre:       "Genre1",
			Description: "Description1",
			AvgGrade:    0,
			Format:      "",
			S3ID:        0,
			CreatedBy:   1,
		},
		{
			ID:          2,
			Title:       "Title2",
			Author:      "Author2",
			Genre:       "Genre2",
			Description: "Description2",
			AvgGrade:    0,
			Format:      "",
			S3ID:        0,
			CreatedBy:   1,
		},
		{
			ID:          3,
			Title:       "Title3",
			Author:      "Author3",
			Genre:       "Genre3",
			Description: "Description3",
			AvgGrade:    0,
			Format:      "",
			S3ID:        0,
			CreatedBy:   1,
		},
	}
}

func TestLibrary(t *testing.T) []*Library {
	return []*Library{
		{
			UserID: 1,
			BookID: 1,
		},
		{
			UserID: 1,
			BookID: 2,
		},
	}
}

func TestFeedback(t *testing.T) *Feedback {
	return &Feedback{
		ID:       1,
		UserID:   1,
		BookID:   1,
		Feedback: "feedback",
		Grade:    5,
	}
}
