-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS categories(
    id          uuid    DEFAULT uuid_generate_v4()  primary key,
    name        varchar,
    description text
);

INSERT INTO categories(id, name, description)
VALUES 
    ('c931106f-8000-46d2-a139-64daf43c5e0c'::uuid, 'Мясо и мясопродукты', ''),
    ('c29f0686-c4ae-4cc4-98e2-2f760ffc941a'::uuid, 'Рыба и рыбопродукты', ''),
    ('90e85ce7-11b0-4664-81b0-5e4f83a3ad38'::uuid, 'Молоко и молочные продукты', ''),
    ('a93ed320-2fe2-457c-8878-9c1142add87c'::uuid, 'Хлеб и хлебобулочные изделия', ''),
    ('7d73049e-12ec-494c-8674-31ae883c0bf7'::uuid, 'Крупы', ''),
    ('1cfd278d-92cf-4bfb-815f-9b1fd5577651'::uuid, 'Бобовые', ''),
    ('2c9498fb-25d9-468a-be57-d35b7884c42b'::uuid, 'Овощи', ''),
    ('89da364a-2ff0-4b5c-8f3b-e9b6b2b8f2fb'::uuid, 'Фрукты и ягоды', ''),
    ('71f16115-836b-4730-a25f-193df03db128'::uuid, 'Орехи', ''),
    ('5a981acf-28fd-4312-b01e-7d393f9f04ff'::uuid, 'Грибы', ''),
    ('e03b10c6-ab37-443d-a9d0-11b9353afdeb'::uuid, 'Кондитерские изделия', ''),
    ('a5302a6b-512b-46a3-835e-781a8c56a8b1'::uuid, 'Макаронные изделия', ''),
    ('4a5fc04c-6902-4dfe-a8ac-ccd8a42452f9'::uuid, 'Яйца', ''),
    ('95ee9317-f619-435c-93ac-a9f10d44f0c6'::uuid, 'Напитки', '');

CREATE TABLE IF NOT EXISTS translates(
    id  uuid DEFAULT uuid_generate_v4()  primary key,
    ru  varchar,
    kz  varchar,
    eng varchar
);

CREATE TABLE IF NOT EXISTS foods(
    id          uuid    DEFAULT uuid_generate_v4()  primary key,
    name        uuid    not null REFERENCES translates(id),
    type        varchar not null,
    category    uuid    not null REFERENCES categories(id),
    price       integer,
    is_new      bool,
    is_spicy    bool,
    description text,
    composition jsonb
);

CREATE TABLE IF NOT EXISTS orders(
    id          uuid    DEFAULT uuid_generate_v4()  primary key,
    fullname    varchar not null,
    phone       varchar not null,
    email       varchar,
    address     text,
    description text
);

CREATE TABLE IF NOT EXISTS orders_foods(
    order_id    uuid    REFERENCES orders(id),
    food_id     uuid    REFERENCES foods(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders_foods;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS foods;
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
