-- +goose Up
CREATE TABLE IF NOT EXISTS dummy (
    "id" uuid DEFAULT uuid_generate_v4(),
    "data" JSONB NOT NULL,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE dummy;