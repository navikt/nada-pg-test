-- name: InsertData :one
INSERT INTO dummy (
    "data"
) VALUES (
    @data
)
RETURNING *;
