CREATE TABLE books(
    id bigserial primary key,
    author varchar not null,
    title varchar not null unique,
    genre varchar not null,
    description varchar not null,
    avg_grade int,
    format varchar,
    s3_id int
);
