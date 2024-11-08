package sqlstore

import (
	"database/sql"

	"github.com/ZhoraIp/ShelfShare/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db                *sql.DB
	userRepository    *UserRepository
	bookRepository    *BookRepository
	libraryRepository *LibraryRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{s}
	return s.userRepository
}

func (s *Store) Book() store.BookRepository {
	if s.bookRepository != nil {
		return s.bookRepository
	}

	s.bookRepository = &BookRepository{s}
	return s.bookRepository
}

func (s *Store) Library() store.LibraryRepository {
	if s.libraryRepository != nil {
		return s.libraryRepository
	}

	s.libraryRepository = &LibraryRepository{s}
	return s.libraryRepository
}
