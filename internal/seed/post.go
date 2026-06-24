package seed

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sanbei101/blue-book/internal/db"
)

// seedPosts 创建种子帖子和媒体
func (s *Seeder) seedPosts(ctx context.Context, users []db.User) ([]db.Post, error) {
	posts := make([]db.Post, 0, len(postSeeds))

	for i, p := range postSeeds {
		author := users[s.rng.Intn(len(users))]

		post, err := s.store.CreatePost(ctx, db.CreatePostParams{
			ID:      uuid.New(),
			UserID:  author.ID,
			Title:   p.Title,
			Content: p.Content,
		})
		if err != nil {
			return nil, fmt.Errorf("create post %d: %w", i, err)
		}

		// 随机添加 1-3 张图片
		mediaCount := s.rng.Intn(3) + 1
		mediaParams := make([]db.CreatePostMediaParams, 0, mediaCount)
		for j := range mediaCount {
			mediaParams = append(mediaParams, db.CreatePostMediaParams{
				ID:        uuid.New(),
				PostID:    post.ID,
				MediaURL:  mediaURL(s.rng.Intn(1000) + i*100 + j),
				MediaType: db.MediaTypeEnumImage,
				SortOrder: int16(j),
			})
		}

		if len(mediaParams) > 0 {
			if _, err := s.store.CreatePostMedia(ctx, mediaParams); err != nil {
				return nil, fmt.Errorf("create post media for post %d: %w", i, err)
			}
		}

		posts = append(posts, post)
	}

	return posts, nil
}
