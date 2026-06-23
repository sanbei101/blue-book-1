-- name: CreateUser :one
INSERT INTO users (
    id, username, password_hash, avatar_url, bio
) VALUES (
    @id, @username, @password_hash, @avatar_url, @bio
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = @id LIMIT 1;

-- name: CreatePost :one
INSERT INTO posts (
    id, user_id, title, content
) VALUES (
    @id, @user_id, @title, @content
) RETURNING *;

-- name: CreatePostMedia :copyfrom
INSERT INTO post_media (
    id, post_id, media_url, media_type, sort_order
) VALUES (
    @id, @post_id, @media_url, @media_type, @sort_order
);

-- name: ListPostsFeed :many
-- 获取首页信息流
SELECT 
    p.id, p.title, p.content, p.view_count, p.created_at,
    u.id AS author_id, u.username AS author_username, u.avatar_url AS author_avatar
FROM posts p
JOIN users u ON p.user_id = u.id
ORDER BY p.created_at DESC
LIMIT @limit_count OFFSET @offset_count;

-- name: ToggleLike :exec
-- 点赞逻辑
INSERT INTO likes (
    id, user_id, target_id, target_type
) VALUES (
    @id, @user_id, @target_id, @target_type
) ON CONFLICT (user_id, target_id, target_type) DO NOTHING;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = @username LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET username = @username, avatar_url = @avatar_url, bio = @bio, updated_at = NOW() WHERE id = @id RETURNING *;

-- name: GetPostByID :one
SELECT
    p.id, p.user_id, p.title, p.content, p.view_count, p.created_at, p.updated_at,
    u.id AS author_id, u.username AS author_username, u.avatar_url AS author_avatar
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.id = @id LIMIT 1;

-- name: ListPostsByUser :many
SELECT
    p.id, p.title, p.content, p.view_count, p.created_at,
    u.id AS author_id, u.username AS author_username, u.avatar_url AS author_avatar
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.user_id = @user_id
ORDER BY p.created_at DESC
LIMIT @limit_count OFFSET @offset_count;

-- name: GetPostMediaByPostID :many
SELECT * FROM post_media WHERE post_id = @post_id ORDER BY sort_order;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = @id AND user_id = @user_id;

-- name: IncrementViewCount :exec
UPDATE posts SET view_count = view_count + 1 WHERE id = @id;

-- name: DeletePostMediaByPostID :exec
DELETE FROM post_media WHERE post_id = @post_id;

-- name: CreateComment :one
INSERT INTO comments (
    id, post_id, user_id, parent_id, content
) VALUES (
    @id, @post_id, @user_id, @parent_id, @content
) RETURNING *;

-- name: ListCommentsByPostID :many
SELECT
    c.id, c.post_id, c.user_id, c.parent_id, c.content, c.like_count, c.created_at,
    u.username AS author_username, u.avatar_url AS author_avatar
FROM comments c
JOIN users u ON c.user_id = u.id
WHERE c.post_id = @post_id
ORDER BY c.created_at ASC
LIMIT @limit_count OFFSET @offset_count;

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