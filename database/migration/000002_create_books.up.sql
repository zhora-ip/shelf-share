CREATE TABLE IF NOT EXISTS books(
    id serial primary key,
    author varchar not null,
    title varchar not null unique,
    genre varchar not null,
    description varchar not null,
    avg_grade numeric,
    format varchar,
    s3_id int,
    created_by int,
    FOREIGN KEY (created_by) REFERENCES users(id)
);
