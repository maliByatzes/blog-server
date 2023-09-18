-- name: CreatePost :one
INSERT INTO posts (
  username,
  tag,
  title,
  image_url,
  content,
  category_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
WHERE username = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdatePost :one
UPDATE posts
SET 
  tag = $2,
  title = $3,
  image_url = $4,
  content = $5,
  category_id = $6
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;