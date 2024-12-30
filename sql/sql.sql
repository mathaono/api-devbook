CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(255) not null,
    nickname varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(50) not null,
    createdAt timestamp default current_timestamp() 
) ENGINE=INNODB; 