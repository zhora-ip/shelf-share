CREATE TABLE IF NOT EXISTS discussion (
    id serial primary key,
    user_id int REFERENCES users(id),
    title varchar unique not null,
    description text not null 
);

CREATE TABLE IF NOT EXISTS message ( 
    id serial primary key,
    user_id int REFERENCES users(id),
    discussion_id int REFERENCES discussion(id),
    message text not null
);
