-- Home feed

-- name: GetHomeFeed :many
SELECT b.* FROM blogs b JOIN follow f ON b.author = f.followee 
WHERE f.follower = $1 AND b.created_at < $2 ORDER BY created_at LIMIT $3;

-- Users

-- name: CreateUser :one
INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *;

-- name: FindUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;
