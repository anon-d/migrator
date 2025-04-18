-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
INSERT INTO
    apps (name)
VALUES
    ('gateway');

-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd
DELETE FROM apps
WHERE
    name = 'gateway';
