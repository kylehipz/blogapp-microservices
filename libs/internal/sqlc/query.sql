-- Home feed

-- name: GetHomeFeed :many
SELECT b.* FROM blogs b JOIN follow f ON b.author = f.followee 
WHERE f.follower = $1 AND b.created_at < $2 ORDER BY created_at DESC LIMIT $3;

-- Users

-- name: CreateUser :one
INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *;

-- name: FindUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;


-- Follow
-- name: CreateFollow :one
INSERT INTO follow (follower, followee) VALUES ($1, $2) RETURNING *;

-- name: FindFollowers :many
SELECT follower FROM follow WHERE followee = $1;

-- name: DeleteFollow :exec
DELETE FROM follow WHERE follower = $1 and followee = $2;

-- Blogs

-- name: CreateBlog :one
INSERT INTO blogs (author, title, content) VALUES ($1, $2, $3) RETURNING *;

-- name: FindBlog :one
SELECT * FROM blogs WHERE id = $1 LIMIT 1;

-- name: UpdateBlog :one
UPDATE blogs SET title = $2, content = $3 WHERE id = $1 RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs WHERE id = $1;
