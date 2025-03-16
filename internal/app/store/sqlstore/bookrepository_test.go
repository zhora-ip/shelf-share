package sqlstore_test

import (
	"testing"

	"github.com/zhora-ip/shelf-share/internal/app/model"
	"github.com/zhora-ip/shelf-share/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestBookRepository_Create(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books")

	s := sqlstore.New(db)
	b := model.TestBook(t)
	u := model.TestUser(t)

	s.User().Create(u)
	b.CreatedBy = u.ID

	err := s.Book().Create(b)
	assert.NoError(t, err)
	assert.NotNil(t, b)
}

func TestBookRepository_FindByID(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	b1 := model.TestBook(t)

	s.User().Create(u)
	b1.CreatedBy = u.ID
	s.Book().Create(b1)

	b2, err := s.Book().FindByID(b1.ID)
	assert.NoError(t, err)
	assert.Equal(t, b1, b2)
}
func TestBookRepository_FindByAuthor(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	bs := model.TestBooks(t)

	s.User().Create(u)

	for i := range bs {
		bs[i].CreatedBy = u.ID
		s.Book().Create(bs[i])
	}

	for _, v := range bs {
		b, err := s.Book().FindByAuthor(v.Author)
		assert.NoError(t, err)
		assert.Equal(t, []*model.Book{v}, b)
	}
}
func TestBookRepository_FindByTitle(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	bs := model.TestBooks(t)

	s.User().Create(u)

	for i := range bs {
		bs[i].CreatedBy = u.ID
		s.Book().Create(bs[i])
	}

	for _, v := range bs {
		b, err := s.Book().FindByTitle(v.Title)
		assert.NoError(t, err)
		assert.Equal(t, []*model.Book{v}, b)
	}
}
func TestBookRepository_FindByGenre(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	bs := model.TestBooks(t)

	s.User().Create(u)

	for i := range bs {
		bs[i].CreatedBy = u.ID
		s.Book().Create(bs[i])
	}

	for _, v := range bs {
		b, err := s.Book().FindByGenre(v.Genre)
		assert.NoError(t, err)
		assert.Equal(t, []*model.Book{v}, b)
	}
}
func TestBookRepository_FindAll(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	bs := model.TestBooks(t)

	s.User().Create(u)

	for i, v := range bs {
		bs[i].CreatedBy = u.ID
		s.Book().Create(v)
	}

	b, err := s.Book().FindAll()
	assert.NoError(t, err)
	assert.Equal(t, bs, b)
}

func TestBookRepository_UpdateFile(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	defer close("users", "books")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	b := model.TestBook(t)

	s.User().Create(u)
	b.CreatedBy = u.ID
	s.Book().Create(b)

	err := s.Book().UpdateFile(b.ID, "pdf")
	bn, _ := s.Book().FindByID(b.ID)

	assert.NoError(t, err)
	assert.Equal(t, bn.Format, "pdf")
}
