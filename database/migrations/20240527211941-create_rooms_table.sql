
-- +migrate Up
CREATE TABLE rooms(
    id SERIAL PRIMARY KEY,
    room_number VARCHAR(255) NOT NULL,
    bed_qty INT NOT NULL DEFAULT 0,
    description TEXT NULL,
    price VARCHAR(255) NOT NULL,
    room_status VARCHAR(255) NOT NULL
); 
-- +migrate Down
DROP TABLE IF EXISTS rooms;
