-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccountById :one
SELECT * 
FROM accounts AS A 
WHERE A.id = $1
LIMIT 1;

-- name: GetAccounts :many
SELECT *
FROM accounts AS A
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateAccount :execrows
UPDATE accounts
SET balance = $2
WHERE id = $1;

-- name: AddAccountBalance :execrows
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = $1; 

-- name: DeleteAccount :execrows
DELETE FROM accounts WHERE id = $1;