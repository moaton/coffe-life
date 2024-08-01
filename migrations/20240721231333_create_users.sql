-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT uuid_generate_v4()  primary key,
    first_name varchar not null,
    last_name varchar not null,
    username varchar not null unique,
    password varchar not null,
    is_first bool not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
