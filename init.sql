CREATE DATABASE questions;

CREATE TABLE IF NOT EXISTS theory (
    id int PRIMARY KEY,
    question text,
    answer text
);

INSERT INTO theory(id, question, answer) VALUES(1, 'Q1', 'A1');
INSERT INTO theory(id, question, answer) VALUES(2, 'Q2', 'A2');
INSERT INTO theory(id, question, answer) VALUES(3, 'Q3', 'A3');
INSERT INTO theory(id, question, answer) VALUES(4, 'Q4', 'A4');