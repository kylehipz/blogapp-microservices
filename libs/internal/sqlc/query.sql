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
-- name: FollowUser :one
INSERT INTO follow (follower, followee) VALUES ($1, $2) RETURNING *;

-- name: UnfollowUser :exec
DELETE FROM follow WHERE follower = $1 and followee = $2;

-- Blogs

-- name: CreateBlog :one
INSERT INTO blogs (author, title, content) VALUES ($1, $2, $3) RETURNING *;

-- name: FindBlog :one
SELECT * FROM blogs WHERE id = $1 LIMIT 1;

-- name: UpdateBlog :one
UPDATE blogs SET content = $2 WHERE id = $1 RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs WHERE id = $1;
