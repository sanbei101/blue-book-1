package seed

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sanbei101/blue-book/internal/db"
)

// seedLikes 创建种子点赞(帖子点赞 + 评论点赞)
func (s *Seeder) seedLikes(ctx context.Context, users []db.User, posts []db.Post, comments []db.Comment) (int, error) {
	count := 0
	seen := make(map[string]bool)

	// 帖子点赞:每个帖子 3-8 个赞
	for _, post := range posts {
		likeCount := s.rng.Intn(6) + 3
		for range likeCount {
			user := users[s.rng.Intn(len(users))]
			key := user.ID.String() + post.ID.String() + "1"
			if seen[key] {
				continue
			}
			seen[key] = true

			err := s.store.ToggleLike(ctx, db.ToggleLikeParams{
				ID:         uuid.Must(uuid.NewV7()),
				UserID:     user.ID,
				TargetID:   post.ID,
				TargetType: 1, // 帖子
			})
			if err != nil {
				return 0, fmt.Errorf("like post: %w", err)
			}
			count++
		}
	}

	// 评论点赞:每个评论 0-3 个赞
	for _, comment := range comments {
		likeCount := s.rng.Intn(4)
		for range likeCount {
			user := users[s.rng.Intn(len(users))]
			key := user.ID.String() + comment.ID.String() + "2"
			if seen[key] {
				continue
			}
			seen[key] = true

			err := s.store.ToggleLike(ctx, db.ToggleLikeParams{
				ID:         uuid.Must(uuid.NewV7()),
				UserID:     user.ID,
				TargetID:   comment.ID,
				TargetType: 2, // 评论
			})
			if err != nil {
				return 0, fmt.Errorf("like comment: %w", err)
			}
			count++
		}
	}

	return count, nil
}

// seedFollows 创建种子关注关系
func (s *Seeder) seedFollows(ctx context.Context, users []db.User) (int, error) {
	count := 0
	seen := make(map[string]bool)

	// 每个用户关注 3-8 个其他用户
	for i := range users {
		follower := &users[i]
		followCount := s.rng.Intn(6) + 3
		for range followCount {
			following := &users[s.rng.Intn(len(users))]
			if follower.ID == following.ID {
				continue
			}

			key := follower.ID.String() + following.ID.String()
			if seen[key] {
				continue
			}
			seen[key] = true

			err := s.store.ToggleFollow(ctx, db.ToggleFollowParams{
				FollowerID:  follower.ID,
				FollowingID: following.ID,
			})
			if err != nil {
				return 0, fmt.Errorf("follow: %w", err)
			}
			count++
		}
	}

	return count, nil
}
