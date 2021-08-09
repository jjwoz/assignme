-- +goose Up
-- +goose StatementBegin
CREATE TABLE ticketit_statuses
(
    id    INT          NOT NULL PRIMARY KEY,
    name  VARCHAR(255) NOT NULL,
    color BIGINT       NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS ticketit_settings;
-- +goose StatementEnd

