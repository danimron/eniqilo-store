CREATE TABLE IF NOT EXISTS products(
    id serial PRIMARY KEY,
    name VARCHAR (30) NOT NULL,
    sku VARCHAR (30) NOT NULL,
    category VARCHAR (11) NOT NULL,
    image_url VARCHAR (255) NOT NULL,
    notes VARCHAR (200) NOT NULL,
    price BIGINT NOT NULL,
    STOCK INT NOT NULL,
    location VARCHAR (200) NOT NULL,
    is_available BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);