CREATE TABLE IF NOT EXISTS transactions(
    id serial PRIMARY KEY,
    amount BIGINT NOT NULL,
    change BIGINT NOT NULL,
    total_item INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);