CREATE DATABASE lzblogs;

\c lzblogs;

CREATE TABLE IF NOT EXISTS users (
    id uint(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    password varchar(256),
    avatar varchar(256),
    sign varchar(256),
    phone varchar(20),
    email varchar(256),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    deleted_at timestamp,
    UNIQUE(name)
)