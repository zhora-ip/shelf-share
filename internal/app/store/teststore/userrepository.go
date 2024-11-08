package teststore

import (
	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/ZhoraIp/ShelfShare/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.users[u.Id] = u
	u.Id = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	if len(r.users) == 0 {
		return nil, store.ErrRecordNotFound
	}

	users := []*model.User{}

	for _, v := range r.users {
		users = append(users, v)
	}

	return users, nil
}
