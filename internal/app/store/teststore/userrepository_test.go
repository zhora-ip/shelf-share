package teststore_test

import (
	"testing"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/ZhoraIp/ShelfShare/internal/app/store"
	"github.com/ZhoraIp/ShelfShare/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserrepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	_, err := s.User().FindByEmail("alex@mail.ru")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail("vasyliy@mail.ru")
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
