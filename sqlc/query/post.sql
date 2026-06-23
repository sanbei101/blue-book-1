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
SELECT
    p.id, p.title, p.content, p.view_count, p.created_at,
    u.id AS author_id, u.username AS author_username, u.avatar_url AS author_avatar
FROM posts p
JOIN users u ON p.user_id = u.id
ORDER BY p.created_at DESC
LIMIT @limit_count OFFSET @offset_count;

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
