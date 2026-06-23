package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/phuslu/log"

	"github.com/sanbei101/blue-book/internal/db"
	"github.com/sanbei101/blue-book/internal/pkg/jwt"
	"github.com/sanbei101/blue-book/internal/pkg/render"
)

func Pagination(r *http.Request, defaultPage, defaultPageSize, maxPageSize int) (
	offset, limit int,
) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = defaultPage
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	return (page - 1) * pageSize, pageSize
}

type PostHandler struct {
	store *db.Store
}

func NewPostHandler(store *db.Store) *PostHandler {
	return &PostHandler{store: store}
}

type createPostRequest struct {
	Title   string            `json:"title"   validate:"required,max=200"`
	Content string            `json:"content" validate:"required"`
	Media   []createMediaItem `json:"media"`
}
type createMediaItem struct {
	MediaUrl  string `json:"media_url"  validate:"required"`
	MediaType string `json:"media_type" validate:"required,oneof=image video"`
	SortOrder int16  `json:"sort_order"`
}

type createPostResponse struct {
	ID uuid.UUID `json:"id"`
}

type listPostsResponse struct {
	ID        uuid.UUID       `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	ViewCount int64           `json:"view_count"`
	Author    authorResponse  `json:"author"`
	Media     []mediaResponse `json:"media,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
}

type authorResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
}

type mediaResponse struct {
	ID        uuid.UUID `json:"id"`
	MediaUrl  string    `json:"media_url"`
	MediaType string    `json:"media_type"`
	SortOrder int16     `json:"sort_order"`
}

func toAuthorFromFeed(authorID uuid.UUID, authorUsername string, authorAvatar pgtype.Text) authorResponse {
	a := authorResponse{ID: authorID, Username: authorUsername}
	if authorAvatar.Valid {
		a.AvatarUrl = authorAvatar.String
	}
	return a
}

func toMediaResponse(m db.PostMedium) mediaResponse {
	return mediaResponse{
		ID:        m.ID,
		MediaUrl:  m.MediaUrl,
		MediaType: string(m.MediaType),
		SortOrder: m.SortOrder,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := render.ReadBody[createPostRequest](w, r)
	if err != nil {
		return
	}
	currentUserID := jwt.GetUserIDFromContext(r)

	var created db.Post
	err = h.store.ExecTx(r.Context(), func(q *db.Queries) error {
		post, err := q.CreatePost(r.Context(), db.CreatePostParams{
			ID:      uuid.New(),
			UserID:  currentUserID,
			Title:   body.Title,
			Content: body.Content,
		})
		if err != nil {
			log.Error().Err(err).Msg("创建帖子失败")
			return err
		}
		created = post
		if len(body.Media) > 0 {
			params := make([]db.CreatePostMediaParams, len(body.Media))
			for i, m := range body.Media {
				mediaType := db.MediaTypeEnumImage
				if m.MediaType == "video" {
					mediaType = db.MediaTypeEnumVideo
				}

				params[i] = db.CreatePostMediaParams{
					ID:        uuid.New(),
					PostID:    post.ID,
					MediaUrl:  m.MediaUrl,
					MediaType: mediaType,
					SortOrder: m.SortOrder,
				}
			}
			_, err := q.CreatePostMedia(r.Context(), params)
			if err != nil {
				log.Error().Err(err).Msg("创建帖子媒体失败")
				return err
			}
		}
		return nil
	})
	if err != nil {
		render.Error(w, http.StatusInternalServerError, "创建帖子失败")
		return
	}
	render.Success(w, "创建成功", createPostResponse{ID: created.ID})
}

// ---- 首页信息流 ----
func (h *PostHandler) ListFeed(w http.ResponseWriter, r *http.Request) {
	offset, pageSize := Pagination(r, 1, 20, 50)
	rows, err := h.store.ListPostsFeed(r.Context(), db.ListPostsFeedParams{
		OffsetCount: int32(offset),
		LimitCount:  int32(pageSize),
	})
	if err != nil {
		render.Error(w, http.StatusInternalServerError, "获取信息流失败")
		return
	}

	posts := make([]listPostsResponse, 0, len(rows))
	for _, row := range rows {
		posts = append(posts, listPostsResponse{
			ID:        row.ID,
			Title:     row.Title,
			Content:   row.Content,
			ViewCount: row.ViewCount,
			CreatedAt: row.CreatedAt,
			Author:    toAuthorFromFeed(row.AuthorID, row.AuthorUsername, row.AuthorAvatar),
		})
	}

	render.Success(w, "查询成功", posts)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的帖子 ID")
		return
	}

	err = h.store.IncrementViewCount(r.Context(), id)
	if err != nil {
		log.Error().Err(err).Msg("增加帖子浏览量失败")
	}

	row, err := h.store.GetPostByID(r.Context(), id)
	if err != nil {
		render.Error(w, http.StatusNotFound, "帖子不存在")
		return
	}

	media, err := h.store.GetPostMediaByPostID(r.Context(), row.ID)
	if err != nil {
		log.Error().Err(err).Msg("获取帖子媒体失败")
		render.Error(w, http.StatusInternalServerError, "获取帖子媒体失败")
		return
	}
	mediaList := make([]mediaResponse, 0, len(media))
	for _, m := range media {
		mediaList = append(mediaList, toMediaResponse(m))
	}

	render.Success(w, "查询成功", listPostsResponse{
		ID:        row.ID,
		Title:     row.Title,
		Content:   row.Content,
		ViewCount: row.ViewCount,
		CreatedAt: row.CreatedAt,
		Author:    toAuthorFromFeed(row.AuthorID, row.AuthorUsername, row.AuthorAvatar),
		Media:     mediaList,
	})
}

// ---- 某用户的帖子列表 ----
func (h *PostHandler) ListByUser(w http.ResponseWriter, r *http.Request) {
	userID := jwt.GetUserIDFromContext(r)
	offset, pageSize := Pagination(r, 1, 20, 50)
	rows, err := h.store.ListPostsByUser(r.Context(), db.ListPostsByUserParams{
		UserID:      userID,
		OffsetCount: int32(offset),
		LimitCount:  int32(pageSize),
	})
	if err != nil {
		render.Error(w, http.StatusInternalServerError, "获取帖子列表失败")
		return
	}

	posts := make([]listPostsResponse, 0, len(rows))
	for _, row := range rows {
		posts = append(posts, listPostsResponse{
			ID:        row.ID,
			Title:     row.Title,
			Content:   row.Content,
			ViewCount: row.ViewCount,
			CreatedAt: row.CreatedAt,
			Author:    toAuthorFromFeed(row.AuthorID, row.AuthorUsername, row.AuthorAvatar),
		})
	}

	render.Success(w, "查询成功", posts)
}

// ---- 删除帖子 ----
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的帖子 ID")
		return
	}
	currentUserID := jwt.GetUserIDFromContext(r)

	err = h.store.DeletePost(r.Context(), db.DeletePostParams{
		ID:     id,
		UserID: currentUserID,
	})
	if err != nil {
		render.Error(w, http.StatusInternalServerError, "删除失败")
		return
	}

	render.SuccessNoData(w, http.StatusNoContent, "删除成功")
}
