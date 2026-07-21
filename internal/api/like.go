package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/phuslu/log"

	"github.com/sanbei101/blue-book/internal/db"
	"github.com/sanbei101/blue-book/internal/pkg/jwt"
	"github.com/sanbei101/blue-book/internal/pkg/render"
)

type LikeHandler struct {
	store *db.Store
}

func NewLikeHandler(store *db.Store) *LikeHandler {
	return &LikeHandler{store: store}
}

type toggleLikeRequest struct {
	TargetID   uuid.UUID `json:"target_id"   validate:"required"`
	TargetType int16     `json:"target_type" validate:"required,oneof=1 2"`
}

type toggleLikeResponse struct {
	OK bool `json:"ok"`
}

// 切换点赞状态
//
//	@Summary	切换点赞状态
//	@Tags		likes
//	@Security	BearerAuth
//	@Param		body	body		toggleLikeRequest	true	"点赞信息"
//	@Success	200		{object}	render.Response[toggleLikeResponse]
//	@Failure	400		{object}	render.errorResponse
//	@Failure	500		{object}	render.errorResponse
//	@Router		/likes [post]
func (h *LikeHandler) Toggle(w http.ResponseWriter, r *http.Request) {
	body, err := render.ReadBody[toggleLikeRequest](w, r)
	if err != nil {
		return
	}
	currentUserID := jwt.GetUserIDFromContext(r)

	err = h.store.ToggleLike(r.Context(), db.ToggleLikeParams{
		ID:         uuid.Must(uuid.NewV7()),
		UserID:     currentUserID,
		TargetID:   body.TargetID,
		TargetType: body.TargetType,
	})
	if err != nil {
		log.Error().Err(err).Msg("点赞失败")
		render.Error(w, http.StatusInternalServerError, "点赞失败")
		return
	}

	render.Success(w, "点赞成功", toggleLikeResponse{OK: true})
}
