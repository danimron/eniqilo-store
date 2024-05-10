CREATE TABLE IF NOT EXISTS transaction_detail(
    id serial PRIMARY KEY,
    transaction_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
    CONSTRAINT fk_transcation_id
        FOREIGN KEY(transaction_id)
            REFERENCES transactions(id)
            ON DELETE CASCADE
    CONSTRAINT fk_product_id
        FOREIGN KEY(product_id)
            REFERENCES products(id)
            ON DELETE CASCADE
);