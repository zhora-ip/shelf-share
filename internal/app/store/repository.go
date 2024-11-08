package store

import "github.com/ZhoraIp/ShelfShare/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindAll() ([]*model.User, error)
}

type BookRepository interface {
	Create(*model.Book) error
	FindAll() ([]*model.Book, error)
}

type LibraryRepository interface {
	AddBook(UserId, BookId int) error
	FindAll(UserId int) ([]int, error)
}
