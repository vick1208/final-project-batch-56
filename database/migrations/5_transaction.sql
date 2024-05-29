-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE transaction(
    id SERIAL NOT NULL PRIMARY KEY,
    reservation_id INT REFERENCES reservation(id),
    bank VARCHAR(255) NOT NULL,
    payment_date DATE NOT NULL DEFAULT CURRENT_DATE,
    payment_type VARCHAR(255) NOT NULL,
    main_fee INT NOT NULL DEFAULT 0,
    additional_fee INT NOT NULL,
    total_fee INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now()
)
-- +migrate StatementEnd