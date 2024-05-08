CREATE TABLE IF NOT EXISTS staffs(
    id serial PRIMARY KEY,
    phone_number VARCHAR (16) NOT NULL,
    password VARCHAR (300) NOT NULL,
    name VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);