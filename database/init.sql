CREATE DATABASE restaurant_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(20) NOT NULL
);

CREATE TABLE menus (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    menu_id INT REFERENCES menus(id),
    status VARCHAR(20) NOT NULL DEFAULT 'sedang diproses',
    discount DECIMAL(5, 2) DEFAULT 0,
    rating INT DEFAULT NULL
);
