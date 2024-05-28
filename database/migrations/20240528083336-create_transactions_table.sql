
-- +migrate Up
CREATE TABLE transactions(
    id serial primary key,
    reservation_id int references reservations(id),
    bank varchar(255) not null,
    payment_type varchar(255) not null,
    additional_fee int not null,
    payment_type int not null
    
);
-- +migrate Down
DROP TABLE IF EXISTS transactions;