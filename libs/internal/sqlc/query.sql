-- name: GetHomeFeed :many
SELECT b.* FROM blogs b JOIN follow f ON b.author = f.followee 
WHERE f.follower = $1 AND b.created_at < $2 ORDER BY created_at LIMIT $3;
