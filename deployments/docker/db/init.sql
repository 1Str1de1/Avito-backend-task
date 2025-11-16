CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    team_name VARCHAR(255) REFERENCES team(team_name),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS team (
    team_name varchar(255) PRIMARY KEY
);

CREATE TYPE status_enum AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS pull_request (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author_id INT REFERENCES users(id),
    status status_enum DEFAULT 'OPEN',
    reviewers_id INTEGER[] DEFAULT ARRAY[]::INTEGER[],
    created_at TIMESTAMP DEFAULT NOW(),
    merged_at TIMESTAMP,
    FOREIGN KEY (reviewers_id) REFERENCES users(id),
    CONSTRAINT reviewers_id_max2 CHECK (array_length(reviewers_id, 1) <= 2)
)