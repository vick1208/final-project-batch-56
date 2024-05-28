
-- +migrate Up
CREATE TABLE users(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) unique not null,
    password varchar(255) not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +migrate Down
DROP TABLE IF EXISTS users;