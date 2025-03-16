package sqlstore

import "github.com/zhora-ip/shelf-share/internal/app/model"

type DiscussionRepository struct {
	store *Store
}

func (r *DiscussionRepository) Create(d *model.Discussion) error {

	if err := d.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO discussion(user_id, title, description) VALUES($1, $2, $3) RETURNING id",
		d.UserID,
		d.Title,
		d.Description,
	).Scan(&d.ID)
}

func (r *DiscussionRepository) NewMessage(m *model.Message) error {

	if err := m.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO message(user_id, discussion_id, message) VALUES($1, $2, $3) RETURNING id",
		m.UserID,
		m.DiscussionID,
		m.Message,
	).Scan(&m.ID)
}

func (r *DiscussionRepository) FindByID(id int) (*model.Discussion, []*model.Message, error) {

	d := &model.Discussion{}
	if err := r.store.db.QueryRow("SELECT id, user_id, title, description FROM discussion WHERE id = $1", id).Scan(
		&d.ID,
		&d.UserID,
		&d.Title,
		&d.Description,
	); err != nil {
		return nil, nil, err
	}

	rows, err := r.store.db.Query("SELECT id, user_id, discussion_id, message FROM message WHERE discussion_id = $1", id)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	ms := []*model.Message{}
	for rows.Next() {
		m := &model.Message{}
		if err := rows.Scan(
			&m.ID,
			&m.UserID,
			&m.DiscussionID,
			&m.Message,
		); err != nil {
			return nil, nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return d, ms, nil
}
