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