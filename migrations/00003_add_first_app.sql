-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
INSERT INTO
    apps (name, secret)
VALUES
    ("gateway", "my_secret");

-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd
DELETE FROM apps
WHERE
    name = 'gateway';
