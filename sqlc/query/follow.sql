-- name: ToggleFollow :exec
INSERT INTO follows (
    follower_id, following_id
) VALUES (
    @follower_id, @following_id
) ON CONFLICT (follower_id, following_id) DO NOTHING;

-- name: Unfollow :exec
DELETE FROM follows WHERE follower_id = @follower_id AND following_id = @following_id;

-- name: ListFollowers :many
SELECT
    u.id, u.username, u.avatar_url, u.bio
FROM follows f
JOIN users u ON f.follower_id = u.id
WHERE f.following_id = @following_id
ORDER BY f.created_at DESC
LIMIT @limit_count OFFSET @offset_count;

-- name: ListFollowing :many
SELECT
    u.id, u.username, u.avatar_url, u.bio
FROM follows f
JOIN users u ON f.following_id = u.id
WHERE f.follower_id = @follower_id
ORDER BY f.created_at DESC
LIMIT @limit_count OFFSET @offset_count;
