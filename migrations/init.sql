-- init.sql

CREATE TABLE IF NOT EXISTS hub (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS team (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    hub_id INT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    FOREIGN KEY (hub_id) REFERENCES hub(id)
);

CREATE TABLE IF NOT EXISTS user_info (
    id serial PRIMARY KEY,
    name VARCHAR(255)  NOT NULL,
    email VARCHAR(255)  NOT NULL,
    team_id INT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    FOREIGN KEY (team_id) REFERENCES team(id)
);
