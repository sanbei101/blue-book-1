package seed

import (
	"context"
	"fmt"

	"github.com/phuslu/log"

	"github.com/google/uuid"

	"github.com/sanbei101/blue-book/internal/db"
)

// seedPosts 创建种子帖子和媒体
func (s *Seeder) seedPosts(ctx context.Context, users []db.User) ([]db.Post, []db.PostMedium, error) {
	posts := make([]db.Post, 0, len(postSeeds))
	allMedia := make([]db.PostMedium, 0, len(postSeeds)*3)

	// 预先获取所有帖子的卡片 URL
	total := len(postSeeds)
	log.Printf("Fetching card URLs from card engine... (0/%d)", total)
	cardURLs := make([][]string, 0, total)
	for i, p := range postSeeds {
		log.Printf("[%d/%d] Fetching cards for %q...", i+1, total, p.Title)
		urls, err := cards.fetchCardURLs(ctx, s.rng, p.Title)
		if err != nil {
			log.Printf("[%d/%d] Warning: failed to fetch cards: %v, using fallback", i+1, total, err)
			// 使用 picsum 作为 fallback
			cardURLs = append(cardURLs, nil)
			continue
		}
		cardURLs = append(cardURLs, urls)
		log.Printf("[%d/%d] Done ✓", i+1, total)
	}
	log.Printf("Card URLs fetched successfully (%d/%d)", total, total)

	for i, p := range postSeeds {
		author := users[s.rng.IntN(len(users))]

		post, err := s.store.CreatePost(ctx, db.CreatePostParams{
			ID:      uuid.Must(uuid.NewV7()),
			UserID:  author.ID,
			Title:   p.Title,
			Content: p.Content,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("create post %d: %w", i, err)
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
				return nil, nil, fmt.Errorf("create post media for post %d: %w", i, err)
			}
			// 收集媒体数据用于导出
			for _, param := range mediaParams {
				allMedia = append(allMedia, db.PostMedium{
					ID:        param.ID,
					PostID:    param.PostID,
					MediaURL:  param.MediaURL,
					MediaType: param.MediaType,
					SortOrder: param.SortOrder,
				})
			}
		}

		posts = append(posts, post)
	}

	return posts, allMedia, nil
}
