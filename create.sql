-- create.sql
CREATE TABLE IF NOT EXISTS "user" (
    id            SERIAL PRIMARY KEY,
    firstname     VARCHAR(50)  NOT NULL,
    lastname      VARCHAR(50)  NOT NULL,
    date_of_birth DATE         NOT NULL,
    income        NUMERIC(12,2) NOT NULL
);


CREATE TABLE IF NOT EXISTS credential (
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(50) UNIQUE NOT NULL,
    password_hash CHAR(64)          NOT NULL,
    created_at    TIMESTAMPTZ       DEFAULT now()
);


CREATE TABLE IF NOT EXISTS profile (
    id             SERIAL PRIMARY KEY,
    credential_id  INT UNIQUE REFERENCES credential(id) ON DELETE CASCADE,
    first_name     VARCHAR(50) NOT NULL,
    last_name      VARCHAR(50) NOT NULL,
    date_of_birth  DATE        NOT NULL,
    income         NUMERIC(12,2) NOT NULL
);