-- +goose Up
CREATE TABLE IF NOT EXISTS people (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    age INTEGER,
    gender VARCHAR(10),
    nationality VARCHAR(2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_people_last_name ON people(last_name);

-- +goose Down
DROP INDEX IF EXISTS idx_people_last_name;
DROP TABLE IF EXISTS people;
