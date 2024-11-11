-- Create role 'postgres' if it doesn't exist
DO
$$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_catalog.pg_roles WHERE rolname = 'postgres'
    ) THEN
        CREATE ROLE postgres WITH LOGIN PASSWORD 'postgres';
    END IF;
END
$$;

-- Create database 'db' if it doesn't exist
DO
$$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_database WHERE datname = 'db'
    ) THEN
        CREATE DATABASE db;
    END IF;
END
$$;

-- Grant all privileges to the 'postgres' role on the 'db' database
GRANT ALL PRIVILEGES ON DATABASE db TO postgres;

SET search_path TO public;

CREATE TABLE IF NOT EXISTS shortened_url (
    code VARCHAR(50) PRIMARY KEY,
    long_url VARCHAR(255) NOT NULL,
    created_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS counter (
    id BIGINT PRIMARY KEY,
    current_value BIGINT NOT NULL DEFAULT 1
);

INSERT INTO counter (id, current_value) 
VALUES (1, 1)
ON CONFLICT (id) DO NOTHING;