package seed

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	"github.com/sanbei101/blue-book/internal/db"
)

// seedPosts 创建种子帖子和媒体
func (s *Seeder) seedPosts(ctx context.Context, users []db.User) ([]db.Post, error) {
	posts := make([]db.Post, 0, len(postSeeds))

	// 预先获取所有帖子的卡片 URL
	log.Println("Fetching card URLs from card engine...")
	cardURLs := make([][]string, 0, len(postSeeds))
	for _, p := range postSeeds {
		urls, err := cards.fetchCardURLs(ctx, s.rng, p.Title)
		if err != nil {
			log.Printf("Warning: failed to fetch cards for %q: %v, using fallback", p.Title, err)
			// 使用默认的 picsum 作为 fallback
			cardURLs = append(cardURLs, nil)
			continue
		}
		cardURLs = append(cardURLs, urls)
	}

	for i, p := range postSeeds {
		author := users[s.rng.IntN(len(users))]

		post, err := s.store.CreatePost(ctx, db.CreatePostParams{
			ID:      uuid.Must(uuid.NewV7()),
			UserID:  author.ID,
			Title:   p.Title,
			Content: p.Content,
		})
		if err != nil {
			return nil, fmt.Errorf("create post %d: %w", i, err)
		}

		// 添加媒体图片
		urls := cardURLs[i]
		mediaCount := 1
		if len(urls) > 0 {
			mediaCount = min(len(urls), 3) // 最多使用 3 张卡片图片
		}

		mediaParams := make([]db.CreatePostMediaParams, 0, mediaCount)
		for j := range mediaCount {
			var mediaURL string
			if len(urls) > 0 {
				mediaURL = urls[j]
			} else {
				// fallback: 使用 picsum
				mediaURL = fmt.Sprintf("https://picsum.photos/seed/%d/800/600", s.rng.IntN(1000)+i*100+j)
			}
			mediaParams = append(mediaParams, db.CreatePostMediaParams{
				ID:        uuid.Must(uuid.NewV7()),
				PostID:    post.ID,
				MediaURL:  mediaURL,
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
