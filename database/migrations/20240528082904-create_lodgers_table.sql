
-- +migrate Up
CREATE TABLE lodgers(
    id serial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    city varchar(255) not null,
    phone varchar(255) not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +migrate Down
DROP TABLE IF EXISTS lodgers;
