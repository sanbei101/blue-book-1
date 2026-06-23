-- name: CreateUser :one
INSERT INTO users (
    id, username, password_hash, avatar_url, bio
) VALUES (
    @id, @username, @password_hash, @avatar_url, @bio
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = @id LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = @username LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET username = @username, avatar_url = @avatar_url, bio = @bio, updated_at = NOW() WHERE id = @id RETURNING *;
