package sqlstore

import (
	"database/sql"

	"github.com/ZhoraIp/ShelfShare/internal/app/store"
)

type LibraryRepository struct {
	store *Store
}

func (r *LibraryRepository) AddBook(UserId, BookId int) error {
	stmt, err := r.store.db.Prepare("INSERT INTO library (user_id, book_id) VALUES($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(UserId, BookId)
	if err != nil {
		return err
	}
	return nil

}

func (r *LibraryRepository) FindAll(UserId int) ([]int, error) {

	library := []int{}

	rows, err := r.store.db.Query("SELECT book_id FROM library WHERE user_id = $1", UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var l int

		if err := rows.Scan(
			&l,
		); err != nil {
			return nil, err
		}

		library = append(library, l)
	}

	return library, nil
}
