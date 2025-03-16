package sqlstore

import (
	"database/sql"

	"github.com/zhora-ip/shelf-share/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db                   *sql.DB
	userRepository       *UserRepository
	bookRepository       *BookRepository
	libraryRepository    *LibraryRepository
	feedbackRepository   *FeedbackRepository
	discussionRepository *DiscussionRepository
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

func (s *Store) Feedback() store.FeedbackRepository {
	if s.feedbackRepository != nil {
		return s.feedbackRepository
	}

	s.feedbackRepository = &FeedbackRepository{s}
	return s.feedbackRepository
}

func (s *Store) Discussion() store.DiscussionRepository {
	if s.discussionRepository != nil {
		return s.discussionRepository
	}
	s.discussionRepository = &DiscussionRepository{s}
	return s.discussionRepository
}
