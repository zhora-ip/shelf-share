package sqlstore

import (
	"database/sql"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/ZhoraIp/ShelfShare/internal/app/store"
)

type FeedbackRepository struct {
	store *Store
}

func (r *FeedbackRepository) Create(f *model.Feedback) error {
	if err := f.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO feedback (user_id, book_id, feedback, grade) VALUES($1, $2, $3, $4) RETURNING id", f.UserID, f.BookID, f.Feedback, f.Grade).Scan(&f.ID)
}

func (r *FeedbackRepository) FindByUser(id int) ([]*model.Feedback, error) {

	feedbacks := []*model.Feedback{}

	rows, err := r.store.db.Query("SELECT id, book_id, feedback, grade FROM feedback WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		f := &model.Feedback{}
		if err := rows.Scan(
			&f.ID,
			&f.BookID,
			&f.Feedback,
			&f.Grade,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		f.UserID = id
		feedbacks = append(feedbacks, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (r *FeedbackRepository) FindByBook(id int) ([]*model.Feedback, error) {

	feedbacks := []*model.Feedback{}

	rows, err := r.store.db.Query("SELECT id, user_id, feedback, grade FROM feedback WHERE book_id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		f := &model.Feedback{}
		if err := rows.Scan(
			&f.ID,
			&f.UserID,
			&f.Feedback,
			&f.Grade,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		f.BookID = id
		feedbacks = append(feedbacks, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return feedbacks, nil
}
