INSERT INTO personal_info ("name", phone_number, email) VALUES
('Дэйв Мастейн', '+7(495)000-00-00', 'mustaine@mail.ru'),
('Марти Фридман', '+7(495)111-11-11', 'fridman@mail.ru'),
('Алекс Скольник', '+7(495)222-22-22', 'skolnik@mail.ru'),
('Джефф Хэннеман', '+7(495)333-33-33', 'hanneman@mail.ru'),
('Джефф Уотерс', '+7(495)444-44-44', 'waters@mail.ru'),
('Джеймс Хэтфилд', '+7(495)555-55-55', 'hetfield@mail.ru'),

('Гарик Холтов', '+7(499)666-66-66', 'holtov@mail.ru'),
('Андреас Киссэрян', '+7(499)777-77-77', 'kisseryan@mail.ru'),
('Кирк Хэмметталиев', '+7(499)888-88-88', 'hammettaliev@mail.ru'),
('Керрин Кингаев', '+7(499)999-99-99', 'kingaev@mail.ru');

INSERT INTO customer_rank ("name") VALUES
('Начальный'),
('Бронзовый'),
('Серебряный'),
('Золотой');

INSERT INTO customer (personal_info_id) VALUES
(1),
(2),
(3),
(4),
(5),
(6);

INSERT INTO master_rank ("name", salary) VALUES
('Стажер', 60000),
('Младший', 100000),
('Старший', 150000),
('Главный', 200000);

INSERT INTO "master" (personal_info_id, master_rank_id) VALUES
(7, 1),
(8, 2),
(9, 3),
(10, 4);

INSERT INTO guitar ("name", condition, serial_number) VALUES
('Gibson Les Paul черная', 'Новая, деффекты не обнаружены', 'AT7NASQ'),
('Gibson Flying V красная', 'Потертости, царапины, трещина на грифе', 'A777BOP'),
('ESP LTD GH-600 белая с красным', 'Незначительные следы пользования', ''),
('Yamaha f370', 'Новая, деффекты не обнаружены', 'F6HM73GN6S0');


INSERT INTO component (type, manufacturer, "name", quantity) VALUES
('Бридж', 'Floyd Rose', 'FRT500', 10),
('Накладка', 'Первый древесный поставщик', 'Черное дерево', 10),
('Корпус', 'Первый древесный поставщик', 'Красное дерево', 10),
('Гриф', 'Первый древесный поставщик', 'Красное дерево', 10),
('Струны', 'Elixir', '19102', 10),
('Колки', 'Gotoh', 'SD510-SL-GG', 10),
('Звукосниматель', 'Fishman', 'Fishman Fluence Modern Humb', 10);

INSERT INTO order_type (description) VALUES
('Новая гитара'),
('Старая гитара');

INSERT INTO order_status (description) VALUES
('Создан'),
('Оплачен'),
('В работе'),
('Готов, на хранении'),
('Завершен'),
('Отменён, ожидает возврата средств и гитары'),
('Отменён, ожидает возврата средств'),
('Отменён, ожидает возврата гитары'),
('Отменён');

INSERT INTO "order" (customer_id, master_id, order_status_id, guitar_id, order_type_id, price, description, created_at) VALUES
(1, 1, 1, 1, 2, 3000, 'Базовая отстройка, замена струн', current_timestamp),
(2, 2, 2, 2, 2, 38000, 'Замена звукоснимателей, замена струн, базовая отстройка', current_timestamp),
(3, 2, 3, NULL, 1, 200000, 'Запилить новую гитару формы Explorer под бюджет', current_timestamp),
(3, 2, 3, NULL, 1, 230000, 'Запилить новую гитару формы Jaguar под бюджет', current_timestamp),
(4, 4, 3, NULL, 1, 400000, 'Запилить новую гитару формы Les Paul под бюджет', current_timestamp);

INSERT INTO required_components (component_id, order_id) VALUES
(5, 1),
(7, 2),
(5, 2),
(1, 3),
(2, 3),
(3, 3),
(4, 3),
(5, 3),
(6, 3),
(7, 3),
(1, 4),
(2, 4),
(3, 4),
(4, 4),
(5, 4),
(6, 4),
(7, 4),
(7, 4),
(1, 5),
(2, 5),
(3, 5),
(4, 5),
(5, 5),
(6, 5),
(7, 5),
(7, 5),
(7, 5);

