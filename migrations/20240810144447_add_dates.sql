-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS foods ADD COLUMN IF NOT EXISTS created_at timestamp;
ALTER TABLE IF EXISTS foods ADD COLUMN IF NOT EXISTS updated_at timestamp;
ALTER TABLE IF EXISTS foods ADD COLUMN IF NOT EXISTS deleted_at timestamp;

ALTER TABLE IF EXISTS categories ADD COLUMN IF NOT EXISTS created_at timestamp;
ALTER TABLE IF EXISTS categories ADD COLUMN IF NOT EXISTS updated_at timestamp;
ALTER TABLE IF EXISTS categories ADD COLUMN IF NOT EXISTS deleted_at timestamp;

ALTER TABLE IF EXISTS orders ADD COLUMN IF NOT EXISTS created_at timestamp;
ALTER TABLE IF EXISTS orders ADD COLUMN IF NOT EXISTS updated_at timestamp;
ALTER TABLE IF EXISTS orders ADD COLUMN IF NOT EXISTS deleted_at timestamp;

ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS created_at timestamp;
ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS updated_at timestamp;
ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS deleted_at timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS foods DROP COLUMN IF EXISTS created_at;
ALTER TABLE IF EXISTS foods DROP COLUMN IF EXISTS updated_at;
ALTER TABLE IF EXISTS foods DROP COLUMN IF EXISTS deleted_at;

ALTER TABLE IF EXISTS categories DROP COLUMN IF EXISTS created_at;
ALTER TABLE IF EXISTS categories DROP COLUMN IF EXISTS updated_at;
ALTER TABLE IF EXISTS categories DROP COLUMN IF EXISTS deleted_at;

ALTER TABLE IF EXISTS orders DROP COLUMN IF EXISTS created_at;
ALTER TABLE IF EXISTS orders DROP COLUMN IF EXISTS updated_at;
ALTER TABLE IF EXISTS orders DROP COLUMN IF EXISTS deleted_at;

ALTER TABLE IF EXISTS users DROP COLUMN IF EXISTS created_at;
ALTER TABLE IF EXISTS users DROP COLUMN IF EXISTS updated_at;
ALTER TABLE IF EXISTS users DROP COLUMN IF EXISTS deleted_at;
-- +goose StatementEnd
