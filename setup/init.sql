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