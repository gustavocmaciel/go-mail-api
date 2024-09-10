-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL
);

-- Create emails table
CREATE TABLE IF NOT EXISTS emails (
    id VARCHAR(255) UNIQUE NOT NULL,
    sender VARCHAR(100) NOT NULL,
    recipients text[] NOT NULL,
    subject VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    email_read BOOLEAN NOT NULL,
    archived BOOLEAN NOT NULL
);