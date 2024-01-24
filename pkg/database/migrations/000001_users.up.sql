CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    full_name VARCHAR(100),
    password VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);