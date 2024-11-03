-- +goose Up
INSERT INTO person (login, password, role) VALUES
('admin', '1234', 'admin'),
('prizon', '1379', 'common');

INSERT INTO component (type, description, price) VALUES
('processor', 'amd ryzen 5 7600', 15000.0),
('processor', 'intel core i5-12400', 15000.0),
('motherboard', 'asus b250k', 4000.0),
('motherboard', 'gigabyte h110m', 6000.0),
('video_card', 'amd rx 7700 xt', 50000.0),
('video_card', 'nvidia gtx 4060 ti', 50000.0),
('ram', 'adata 16gb', 7500.0),
('ram', 'kingston 16gb', 9000.0),
('disk', 'adata ssd 1000gb', 4000.0),
('disk', 'toshiba hdd 2000gb', 6000.0),
('power_unit', 'xilence 600w', 4000.0),
('power_unit', 'powercool 700w', 6000.0),
('frame', 'ardor b15', 1500.0),
('frame', 'deepcool k25', 3500.0);

INSERT INTO customer (name, phone, email, address) VALUES
('Ivan', '+75440110980', 'ivan1@main.ru', 'Pushkina 3'),
('Sasha', '+79562675302', 'sasha2@main.ru', 'Svercovo 36'),
('Igor', '+75278712414', 'igor3@main.ru', 'Ovalnoe 19'),
('Nikita', '+72343828470', 'nikita4@main.ru', 'Proezdnoe 109'),
('Andrey', '+76451844022', 'andrey5@main.ru', 'Shishkina 85');

INSERT INTO sale (component_id, customer_id, count) VALUES
(1, 1, 1),
(3, 2, 1),
(5, 3, 1),
(7, 4, 2),
(9, 5, 4),
(2, 1, 1),
(4, 2, 1),
(6, 3, 1),
(8, 4, 2),
(10, 5, 4);

-- +goose Down
DELETE FROM sale;
DELETE FROM customer;
DELETE FROM component;
DELETE FROM person;

ALTER SEQUENCE sale_id_seq RESTART WITH 1;
ALTER SEQUENCE customer_id_seq RESTART WITH 1;
ALTER SEQUENCE component_id_seq RESTART WITH 1;
ALTER SEQUENCE person_id_seq RESTART WITH 1;