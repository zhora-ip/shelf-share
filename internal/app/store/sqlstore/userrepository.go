package sqlstore

import (
	"database/sql"

	"github.com/zhora-ip/shelf-share/internal/app/model"
	"github.com/zhora-ip/shelf-share/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO users (nickname, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		u.Nickname,
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, nickname, email, encrypted_password from users where email = $1",
		email,
	).Scan(&u.ID,
		&u.Nickname,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, nickname, email, encrypted_password from users where id = $1",
		id,
	).Scan(&u.ID,
		&u.Nickname,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	users := []*model.User{}

	rows, err := r.store.db.Query("SELECT id, nickname, email, encrypted_password from users")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(
			&u.ID,
			&u.Nickname,
			&u.Email,
			&u.EncryptedPassword,
		); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
