-- +migrate Up
CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    roll_no VARCHAR(50) UNIQUE,
    is_verified BOOLEAN DEFAULT FALSE,
    department VARCHAR(255),
    branch VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS students;
