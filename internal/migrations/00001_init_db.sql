-- +goose Up
CREATE TYPE person_role AS ENUM (
    'admin',
    'common'
);

CREATE TYPE pc_component AS ENUM (
    'processor',
    'motherboard',
    'video_card',
    'ram',
    'disk',
    'power_unit',
    'frame'
);

CREATE TABLE person (
    id SERIAL PRIMARY KEY,
    login VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    role person_role NOT NULL
);

CREATE TABLE component (
    id SERIAL PRIMARY KEY,
    type pc_component NOT NULL,
    description VARCHAR(256) NOT NULL,
    price INTEGER NOT NULL
);

CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(256) NOT NULL,
    address VARCHAR(100) NOT NULL
);

CREATE TABLE sale (
    id SERIAL PRIMARY KEY,
    component_id INTEGER REFERENCES component(id) ON DELETE CASCADE NOT NULL,
    customer_id INTEGER REFERENCES customer(id) ON DELETE CASCADE NOT NULL,
    count INTEGER NOT NULL
);

-- +goose Down
DROP TABLE sale;
DROP TABLE customer;
DROP TABLE component;
DROP TABLE person;

DROP TYPE pc_component;
DROP TYPE person_role;