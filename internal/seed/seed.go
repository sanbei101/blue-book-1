package seed

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phuslu/log"

	"github.com/sanbei101/blue-book/internal/db"
)

type Seeder struct {
	store *db.Store
	pool  *pgxpool.Pool
	rng   *rand.Rand
}

func NewSeeder(pool *pgxpool.Pool) *Seeder {
	return &Seeder{
		store: db.NewStore(pool),
		pool:  pool,
		rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Run 执行种子数据注入
func (s *Seeder) Run(ctx context.Context) error {
	start := time.Now()

	// 清空所有表(按外键依赖顺序)
	if err := s.truncateAll(ctx); err != nil {
		return fmt.Errorf("truncate tables: %w", err)
	}

	// 创建用户
	users, err := s.seedUsers(ctx)
	if err != nil {
		return fmt.Errorf("seed users: %w", err)
	}

	// 创建帖子 + 媒体
	posts, err := s.seedPosts(ctx, users)
	if err != nil {
		return fmt.Errorf("seed posts: %w", err)
	}

	// 创建评论
	comments, err := s.seedComments(ctx, users, posts)
	if err != nil {
		return fmt.Errorf("seed comments: %w", err)
	}

	// 创建点赞
	likes, err := s.seedLikes(ctx, users, posts, comments)
	if err != nil {
		return fmt.Errorf("seed likes: %w", err)
	}

	// 创建关注关系
	follows, err := s.seedFollows(ctx, users)
	if err != nil {
		return fmt.Errorf("seed follows: %w", err)
	}

	log.Info().
		Int("users", len(users)).
		Int("posts", len(posts)).
		Int("comments", len(comments)).
		Int("likes", likes).
		Int("follows", follows).
		Dur("elapsed", time.Since(start)).
		Msg("Seed complete")

	return nil
}

// truncateAll 按外键依赖顺序清空所有表
func (s *Seeder) truncateAll(ctx context.Context) error {
	tables := []string{
		"likes",
		"follows",
		"comments",
		"post_media",
		"posts",
		"users",
	}
	for _, table := range tables {
		if _, err := s.pool.Exec(ctx, "TRUNCATE TABLE "+table+" CASCADE"); err != nil {
			return fmt.Errorf("truncate %s: %w", table, err)
		}
	}
	return nil
}
