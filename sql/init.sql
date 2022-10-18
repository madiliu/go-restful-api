-- create table
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL
);

-- insert data
INSERT INTO users (name)
VALUES ('peter'),('amy'),('bob'),('catherine'),('david')
