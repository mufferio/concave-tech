-- create.sql
-- One simple table that mirrors the Person struct in Go

CREATE TABLE IF NOT EXISTS persons (
    id            SERIAL PRIMARY KEY,
    firstname     VARCHAR(50)  NOT NULL,
    lastname      VARCHAR(50)  NOT NULL,
    date_of_birth DATE         NOT NULL,
    income        NUMERIC(12,2) NOT NULL   -- supports up to 999,999,999.99
);