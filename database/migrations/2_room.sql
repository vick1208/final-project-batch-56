-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE room(
    id SERIAL NOT NULL PRIMARY KEY,
    room_number VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    room_status VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now()
)
-- +migrate StatementEnd