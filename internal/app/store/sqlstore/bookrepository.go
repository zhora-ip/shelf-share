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

	return r.store.db.QueryRow("INSERT INTO books (author, title, genre, description, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		b.Author,
		b.Title,
		b.Genre,
		b.Description,
		b.CreatedBy,
	).Scan(&b.ID)

}

func (r *BookRepository) FindByID(id int) (*model.Book, error) {
	book := &model.Book{}

	if err := r.store.db.QueryRow("SELECT id, author, title, genre, description, COALESCE(avg_grade, 0.0), COALESCE(format, ''), created_by from books WHERE id = $1", id).Scan(
		&book.ID,
		&book.Author,
		&book.Title,
		&book.Genre,
		&book.Description,
		&book.AvgGrade,
		&book.Format,
		&book.CreatedBy,
	); err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) FindByAuthor(author string) ([]*model.Book, error) {
	books := []*model.Book{}

	rows, err := r.store.db.Query("id, author, title, genre, description, COALESCE(avg_grade, 0.0), COALESCE(format, ''), created_by from books WHERE author = $1", author)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := &model.Book{}
		if err := rows.Scan(
			&book.ID,
			&book.Author,
			&book.Title,
			&book.Genre,
			&book.Description,
			&book.AvgGrade,
			&book.Format,
			&book.CreatedBy,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, err
}

func (r *BookRepository) FindByTitle(title string) ([]*model.Book, error) {
	books := []*model.Book{}

	rows, err := r.store.db.Query("id, author, title, genre, description, COALESCE(avg_grade, 0.0), COALESCE(format, ''), created_by from books WHERE title = $1", title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := &model.Book{}
		if err := rows.Scan(
			&book.ID,
			&book.Author,
			&book.Title,
			&book.Genre,
			&book.Description,
			&book.AvgGrade,
			&book.Format,
			&book.CreatedBy,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, err
}

func (r *BookRepository) FindByGenre(genre string) ([]*model.Book, error) {
	books := []*model.Book{}

	rows, err := r.store.db.Query("id, author, title, genre, description, COALESCE(avg_grade, 0.0), COALESCE(format, ''), created_by from books WHERE genre = $1", genre)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := &model.Book{}
		if err := rows.Scan(
			&book.ID,
			&book.Author,
			&book.Title,
			&book.Genre,
			&book.Description,
			&book.AvgGrade,
			&book.Format,
			&book.CreatedBy,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, err
}

func (r *BookRepository) FindAll() ([]*model.Book, error) {
	books := []*model.Book{}

	rows, err := r.store.db.Query("SELECT id, author, title, genre, description, COALESCE(avg_grade, 0.0), COALESCE(format, ''), created_by from books")
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
			&b.ID,
			&b.Author,
			&b.Title,
			&b.Genre,
			&b.Description,
			&b.AvgGrade,
			&b.Format,
			&b.CreatedBy,
		); err != nil {
			return nil, err
		}
		books = append(books, &b)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return books, nil

}

func (r *BookRepository) UpdateFile(id int, format string) error {

	ID := id + 1000
	_, err := r.store.db.Exec("UPDATE books SET format = $1, s3_id = $2 WHERE id = $3", format, ID, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) UpdateGrade(id int) error {

	var avg_grade float32
	err := r.store.db.QueryRow("SELECT AVG(grade) as avg_grade from feedback WHERE book_id = $1", id).Scan(&avg_grade)
	if err != nil {
		return err
	}

	_, err = r.store.db.Exec("UPDATE books SET avg_grade = $1 WHERE id = $2", avg_grade, id)
	if err != nil {
		return err
	}
	return nil
}
