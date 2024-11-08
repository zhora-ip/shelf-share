package sqlstore

import (
	"database/sql"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/ZhoraIp/ShelfShare/internal/app/store"
)

type BookRepository struct {
	store *Store
}

func (r *BookRepository) Create(b *model.Book) error {

	/* validations + before create */

	if err := b.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO books (author, title, genre, description) VALUES ($1, $2, $3, $4) RETURNING id",
		b.Author,
		b.Title,
		b.Genre,
		b.Description,
	).Scan(&b.Id)

}

func (r *BookRepository) FindAll() ([]*model.Book, error) {
	books := []*model.Book{}

	rows, err := r.store.db.Query("SELECT id, author, title, genre, description, COALESCE(avg_grade, 0.0), COALESCE(format, '') from books")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		b := model.Book{}

		if err := rows.Scan(
			&b.Id,
			&b.Author,
			&b.Title,
			&b.Genre,
			&b.Description,
			&b.AvgGrade,
			&b.Format,
		); err != nil {
			return nil, err
		}
		books = append(books, &b)
	}

	return books, nil

}

/*type Book struct {
	Id          int     `json:"id"`
	Author      string  `json:"author"`
	Genre       string  `json:"genre"`
	Description string  `json:"description"`
	AvgGrade    float32 `json:"-"`
	Format      string  `json:"format"`
	S3Id        int     `json:"-"`
} */
