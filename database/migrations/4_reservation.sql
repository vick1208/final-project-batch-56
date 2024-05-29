-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE reservation(
    id SERIAL NOT NULL PRIMARY KEY,
    lodger_id INT REFERENCES lodger(id),
    room_id INT REFERENCES room(id),
    date_start DATE NOT NULL
)
-- +migrate StatementEnd