package sqlstore_test

import (
	"testing"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/ZhoraIp/ShelfShare/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestLibraryRepository_AddBook(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books", "library")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	b := model.TestBook(t)

	s.User().Create(u)
	b.CreatedBy = u.ID
	s.Book().Create(b)

	err := s.Library().AddBook(u.ID, b.ID)
	assert.NoError(t, err)
}

func TestLibraryRepository_FindAll(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books", "library")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	bs := model.TestBooks(t)

	s.User().Create(u)

	for i := range bs {
		bs[i].CreatedBy = u.ID
		s.Book().Create(bs[i])
		s.Library().AddBook(u.ID, bs[i].ID)
	}

	l, err := s.Library().FindAll(u.ID)

	assert.NoError(t, err)
	for i, v := range l {
		assert.Equal(t, bs[i].ID, v)
	}
}
