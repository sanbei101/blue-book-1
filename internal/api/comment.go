package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/phuslu/log"

	"github.com/sanbei101/blue-book/internal/db"
	"github.com/sanbei101/blue-book/internal/pkg/jwt"
	"github.com/sanbei101/blue-book/internal/pkg/render"
)

type CommentHandler struct {
	store *db.Store
}

func NewCommentHandler(store *db.Store) *CommentHandler {
	return &CommentHandler{store: store}
}

// ---- 创建评论 ----

type createCommentRequest struct {
	PostID   uuid.UUID  `json:"post_id"   validate:"required"`
	ParentID *uuid.UUID `json:"parent_id"`
	Content  string     `json:"content"   validate:"required,max=1000"`
}

type createCommentResponse struct {
	ID uuid.UUID `json:"id"`
}

//	@Summary	创建评论
//	@Tags		comments
//	@Security	BearerAuth
//	@Param		body	body		createCommentRequest	true	"评论内容"
//	@Success	200		{object}	render.Response[createCommentResponse]
//	@Failure	400		{object}	render.errorResponse
//	@Failure	500		{object}	render.errorResponse
//	@Router		/comments [post]
func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := render.ReadBody[createCommentRequest](w, r)
	if err != nil {
		return
	}
	currentUserID := jwt.GetUserIDFromContext(r)

	comment, err := h.store.CreateComment(r.Context(), db.CreateCommentParams{
		ID:       uuid.New(),
		PostID:   body.PostID,
		UserID:   currentUserID,
		ParentID: body.ParentID,
		Content:  body.Content,
	})
	if err != nil {
		log.Error().Err(err).Msg("创建评论失败")
		render.Error(w, http.StatusInternalServerError, "创建评论失败")
		return
	}

	render.Success(w, "评论成功", createCommentResponse{ID: comment.ID})
}

// ---- 帖子评论列表 ----
type commentResponse struct {
	ID             uuid.UUID  `json:"id"`
	PostID         uuid.UUID  `json:"post_id"`
	UserID         uuid.UUID  `json:"user_id"`
	ParentID       *uuid.UUID `json:"parent_id,omitempty"`
	Content        string     `json:"content"`
	LikeCount      int32      `json:"like_count"`
	AuthorUsername string     `json:"author_username"`
	AuthorAvatar   string     `json:"author_avatar,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

//	@Summary	获取帖子评论列表
//	@Tags		comments
//	@Param		post_id		query		string	true	"帖子 ID"
//	@Param		page		query		int		false	"页码"	default(1)
//	@Param		page_size	query		int		false	"每页数量"	default(20)
//	@Success	200			{object}	render.Response[[]commentResponse]
//	@Failure	400			{object}	render.errorResponse
//	@Failure	500			{object}	render.errorResponse
//	@Router		/posts/{id}/comments [get]
func (h *CommentHandler) ListByPost(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.URL.Query().Get("post_id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的帖子 ID")
		return
	}

	offset, pageSize := Pagination(r, 1, 20, 50)

	rows, err := h.store.ListCommentsByPostID(r.Context(), db.ListCommentsByPostIDParams{
		PostID:      postID,
		OffsetCount: int32(offset),
		LimitCount:  int32(pageSize),
	})
	if err != nil {
		log.Error().Err(err).Msg("获取评论列表失败")
		render.Error(w, http.StatusInternalServerError, "获取评论列表失败")
		return
	}

	comments := make([]commentResponse, 0, len(rows))
	for i := range rows {
		c := commentResponse{
			ID:             rows[i].ID,
			PostID:         rows[i].PostID,
			UserID:         rows[i].UserID,
			Content:        rows[i].Content,
			LikeCount:      rows[i].LikeCount,
			AuthorUsername: rows[i].AuthorUsername,
			CreatedAt:      rows[i].CreatedAt,
		}
		if rows[i].ParentID != nil {
			c.ParentID = rows[i].ParentID
		}
		if rows[i].AuthorAvatar.Valid {
			c.AuthorAvatar = rows[i].AuthorAvatar.String
		}
		comments = append(comments, c)
	}

	render.Success(w, "查询成功", comments)
}
