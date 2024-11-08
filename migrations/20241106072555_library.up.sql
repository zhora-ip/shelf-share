CREATE TABLE library (
    user_id int,
    book_id int,
    PRIMARY KEY(user_id, book_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);