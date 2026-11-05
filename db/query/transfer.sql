-- name: CreateTransfer :one
INSERT INTO transfers (sender, receiver, amount)
VALUES ($1, $2, $3)
RETURNING *;