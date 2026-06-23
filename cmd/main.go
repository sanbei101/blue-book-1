package cmd

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phuslu/log"
	"github.com/sanbei101/blue-book/internal/api"
	"github.com/sanbei101/blue-book/internal/db"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:password@localhost:5432/blue_book?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("无法连接数据库")
	}
	defer pool.Close()

	store := db.NewStore(pool)

	router := api.RegisterRoutes(store)

	log.Info().Msg("小蓝书后端服务启动成功,正在监听 :8080...")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal().Err(err).Msg("服务异常关闭")
	}
}
