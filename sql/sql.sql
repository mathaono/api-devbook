CREATE DATABASE IF NOT EXISTS devbook;

DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS publications;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 

CREATE TABLE followers (
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follow_id INT NOT NULL,
    FOREIGN KEY (follow_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    PRIMARY KEY (user_id, follow_id)
);

CREATE TABLE publications (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content VARCHAR(300) NOT NULL,

    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);