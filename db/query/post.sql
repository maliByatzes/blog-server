-- name: CreatePost :one
INSERT INTO posts (
  username,
  title,
  image_url,
  content,
  category_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET 
  title = $2,
  image_url = $3,
  content = $4,
  category_id = $5
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;