-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email,
  role_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
  hashed_password = $2,
  email = $3
WHERE username = $1
RETURNING *;