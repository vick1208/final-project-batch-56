-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE rental(
    id SERIAL NOT NULL PRIMARY KEY,
    lodger_id INT REFERENCES lodger(id),
    room_id INT REFERENCES room(id),
    date_start DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now()
)
-- +migrate StatementEnd