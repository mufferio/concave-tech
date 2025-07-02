-- data.sql
INSERT INTO "user" (firstname, lastname, date_of_birth, income)
VALUES ('John', 'Citizen', '2003-03-03', 65000);

INSERT INTO credential (username, password_hash)
VALUES ('jcitizen', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8');

INSERT INTO profile (credential_id, first_name, last_name, date_of_birth, income)
VALUES (1, 'John', 'Citizen', '2003-03-03', 65000);