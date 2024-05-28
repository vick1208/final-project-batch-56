-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE room(
    id SERIAL NOT NULL PRIMARY KEY,
    room_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now()
)
-- +migrate StatementEnd