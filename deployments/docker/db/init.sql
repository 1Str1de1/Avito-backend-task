CREATE DATABASE pullRequests;
GRANT ALL PRIVILEGES ON DATABASE pullRequests to postgres;
\c pullRequests

CREATE TABLE IF NOT EXISTS team (
    team_name varchar(255) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS users (
    userId VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    team_name VARCHAR(255) REFERENCES team(team_name),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TYPE status_enum AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS pull_request (
    pull_request_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author_id VARCHAR(255) REFERENCES users(userId),
    status status_enum DEFAULT 'OPEN',
    reviewers_id INTEGER[] DEFAULT ARRAY[]::INTEGER[],
    created_at TIMESTAMP DEFAULT NOW(),
    merged_at TIMESTAMP,
    need_more_reviewers BOOLEAN GENERATED ALWAYS AS (
        array_length(reviewers_id, 1) IS NULL
        OR array_length(reviewers_id, 1) < 2
        ) STORED,
    FOREIGN KEY (author_id) REFERENCES users(userId),
    CONSTRAINT reviewers_id_max2 CHECK (array_length(reviewers_id, 1) <= 2)
);

INSERT INTO team (team_name) VALUES ('Team A'), ('Team B'), ('Team C')
ON CONFLICT DO NOTHING;