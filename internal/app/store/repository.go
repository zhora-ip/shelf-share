package store

import (
	"github.com/zhora-ip/shelf-share/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindAll() ([]*model.User, error)
}

type BookRepository interface {
	Create(*model.Book) error
	FindByID(id int) (*model.Book, error)
	FindByAuthor(author string) ([]*model.Book, error)
	FindByTitle(title string) ([]*model.Book, error)
	FindByGenre(genre string) ([]*model.Book, error)
	FindAll() ([]*model.Book, error)
	UpdateFile(int, string) error
	UpdateGrade(int) error
}

type LibraryRepository interface {
	AddBook(UserId, BookId int) error
	FindAll(UserId int) ([]int, error)
}

type FeedbackRepository interface {
	Create(*model.Feedback) error
	FindByUser(id int) ([]*model.Feedback, error)
	FindByBook(id int) ([]*model.Feedback, error)
}

type DiscussionRepository interface {
	Create(*model.Discussion) error
	NewMessage(*model.Message) error
	FindByID(id int) (*model.Discussion, []*model.Message, error)
}
