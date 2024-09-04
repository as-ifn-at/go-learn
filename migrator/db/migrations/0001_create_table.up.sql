-- CREATE DATABASE test_db WITH ENCODING = 'UTF8';

-- \c test_db;

-- CREATE TABLE employees (
--     id SERIAL PRIMARY KEY,
--     name TEXT NOT NULL,
--     email TEXT UNIQUE,
--     hire_date DATE
-- );

-- CREATE TABLE asif (
--     id SERIAL PRIMARY KEY,
--     name TEXT NOT NULL,
--     email TEXT UNIQUE,
--     hire_date DATE
-- );

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

