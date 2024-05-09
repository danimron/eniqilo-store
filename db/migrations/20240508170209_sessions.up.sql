CREATE TABLE IF NOT EXISTS sessions(
    id serial PRIMARY KEY,
    token VARCHAR (300) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    staff_id INT NOT NULL,
    CONSTRAINT fk_staff_id
        FOREIGN KEY(staff_id)
            REFERENCES staffs(id)
            ON DELETE CASCADE
);

