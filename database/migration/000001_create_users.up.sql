CREATE TABLE IF NOT EXISTS users (
    id serial not null primary key,
    nickname varchar not null unique,
    encrypted_password varchar not null,
    email varchar not null unique
);

