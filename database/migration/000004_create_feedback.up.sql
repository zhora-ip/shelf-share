CREATE TABLE IF NOT EXISTS feedback(
    id serial primary key,
    user_id int references users(id),
    book_id int references books(id),
    feedback varchar not null,
    grade int not null
);
