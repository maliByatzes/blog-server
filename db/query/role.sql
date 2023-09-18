-- name: CreateRole :one
INSERT INTO roles (
  role_name
) VALUES (
  $1
) RETURNING *;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;
