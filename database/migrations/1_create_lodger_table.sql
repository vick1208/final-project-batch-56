-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE lodger(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now()
)
-- +migrate StatementEnd