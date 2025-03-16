package sqlstore_test

import (
	"testing"

	"github.com/zhora-ip/shelf-share/internal/app/model"
	"github.com/zhora-ip/shelf-share/internal/app/store"
	"github.com/zhora-ip/shelf-share/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserrepository_Create(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	s := sqlstore.New(db)
	defer close("users")

	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	s := sqlstore.New(db)
	defer close("users")

	_, err := s.User().FindByEmail("alex@mail.ru")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail("vasyliy@mail.ru")
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindAll(t *testing.T) {
	db, close := sqlstore.TestDB(t, databaseURL)
	s := sqlstore.New(db)
	defer close("users")

	users := model.TestUsers(t)

	for _, u := range users {
		s.User().Create(u)
	}

	us, err := s.User().FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, us)
}
