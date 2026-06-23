-- name: ToggleLike :exec
INSERT INTO likes (
    id, user_id, target_id, target_type
) VALUES (
    @id, @user_id, @target_id, @target_type
) ON CONFLICT (user_id, target_id, target_type) DO NOTHING;
