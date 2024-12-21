package sqlstore_test

import (
	"testing"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/ZhoraIp/ShelfShare/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestFeedbackRepository_Create(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books", "feedback")
	s := sqlstore.New(db)

	u := model.TestUser(t)
	b := model.TestBook(t)
	f := model.TestFeedback(t)

	s.User().Create(u)
	b.CreatedBy = u.ID
	s.Book().Create(b)
	f.UserID = u.ID
	f.BookID = b.ID

	err := s.Feedback().Create(f)
	assert.NoError(t, err)
}

func TestFeedbackRepository_FindByUser(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books", "feedback")
	s := sqlstore.New(db)

	u := model.TestUser(t)
	b := model.TestBook(t)
	f1 := model.TestFeedback(t)

	s.User().Create(u)
	b.CreatedBy = u.ID
	s.Book().Create(b)
	f1.UserID = u.ID
	f1.BookID = b.ID

	s.Feedback().Create(f1)
	f2, err := s.Feedback().FindByUser(u.ID)

	assert.Equal(t, []*model.Feedback{f1}, f2)
	assert.NoError(t, err)
}

func TestFeedbackRepository_FindByBook(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books", "feedback")
	s := sqlstore.New(db)

	u := model.TestUser(t)
	b := model.TestBook(t)
	f1 := model.TestFeedback(t)

	s.User().Create(u)
	b.CreatedBy = u.ID
	s.Book().Create(b)
	f1.UserID = u.ID
	f1.BookID = b.ID

	s.Feedback().Create(f1)
	f2, err := s.Feedback().FindByBook(b.ID)

	assert.Equal(t, []*model.Feedback{f1}, f2)
	assert.NoError(t, err)
}
