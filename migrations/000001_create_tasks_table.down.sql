CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_completed BOOLEAN DEFAULT FALSE
);
INSERT INTO tasks (name, is_completed) VALUES ('Belajar Go', FALSE), ('Bikin API', FALSE);