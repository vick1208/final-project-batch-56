
-- +migrate Up
CREATE TABLE reservations(
    id serial primary key,
    lodger_id int references lodgers(id),
    room_id int references rooms(id),
    date_start date not null
);
-- +migrate Down
DROP table if exists reservations;
