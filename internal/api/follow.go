package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/phuslu/log"

	"github.com/sanbei101/blue-book/internal/db"
	"github.com/sanbei101/blue-book/internal/pkg/jwt"
	"github.com/sanbei101/blue-book/internal/pkg/render"
)

type FollowHandler struct {
	store *db.Store
}

func NewFollowHandler(store *db.Store) *FollowHandler {
	return &FollowHandler{store: store}
}

// ---- 关注 ----

type followResponse struct {
	OK bool `json:"ok"`
}

//	@Summary	关注用户
//	@Tags		follows
//	@Security	BearerAuth
//	@Param		id	path		string	true	"目标用户 ID"
//	@Success	200	{object}	render.Response[followResponse]
//	@Failure	400	{object}	render.errorResponse
//	@Failure	500	{object}	render.errorResponse
//	@Router		/users/{id}/follow [post]
func (h *FollowHandler) Follow(w http.ResponseWriter, r *http.Request) {
	followingIDStr := chi.URLParam(r, "id")
	followingID, err := uuid.Parse(followingIDStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的用户 ID")
		return
	}
	currentUserID := jwt.GetUserIDFromContext(r)

	if currentUserID == followingID {
		render.Error(w, http.StatusBadRequest, "不能关注自己")
		return
	}

	err = h.store.ToggleFollow(r.Context(), db.ToggleFollowParams{
		FollowerID:  currentUserID,
		FollowingID: followingID,
	})
	if err != nil {
		log.Error().Err(err).Msg("关注失败")
		render.Error(w, http.StatusInternalServerError, "关注失败")
		return
	}

	render.Success(w, "关注成功", followResponse{OK: true})
}

// ---- 取消关注 ----

//	@Summary	取消关注
//	@Tags		follows
//	@Security	BearerAuth
//	@Param		id	path		string	true	"目标用户 ID"
//	@Success	200	{object}	render.Response[followResponse]
//	@Failure	400	{object}	render.errorResponse
//	@Failure	500	{object}	render.errorResponse
//	@Router		/users/{id}/follow [delete]
func (h *FollowHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	followingIDStr := chi.URLParam(r, "id")
	followingID, err := uuid.Parse(followingIDStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的用户 ID")
		return
	}
	currentUserID := jwt.GetUserIDFromContext(r)

	err = h.store.Unfollow(r.Context(), db.UnfollowParams{
		FollowerID:  currentUserID,
		FollowingID: followingID,
	})
	if err != nil {
		log.Error().Err(err).Msg("取消关注失败")
		render.Error(w, http.StatusInternalServerError, "取消关注失败")
		return
	}

	render.Success(w, "取消关注成功", followResponse{OK: true})
}

// ---- 粉丝列表 ----

type followUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	Bio       string    `json:"bio,omitempty"`
}

func toFollowUserResponse(u *db.ListFollowersRow) followUserResponse {
	resp := followUserResponse{
		ID:       u.ID,
		Username: u.Username,
	}
	if u.AvatarURL.Valid {
		resp.AvatarURL = u.AvatarURL.String
	}
	if u.Bio.Valid {
		resp.Bio = u.Bio.String
	}
	return resp
}

//	@Summary	获取粉丝列表
//	@Tags		follows
//	@Param		id			path		string	true	"用户 ID"
//	@Param		page		query		int		false	"页码"	default(1)
//	@Param		page_size	query		int		false	"每页数量"	default(20)
//	@Success	200			{object}	render.Response[[]followUserResponse]
//	@Failure	400			{object}	render.errorResponse
//	@Failure	500			{object}	render.errorResponse
//	@Router		/users/{id}/followers [get]
func (h *FollowHandler) ListFollowers(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的用户 ID")
		return
	}

	offset, pageSize := Pagination(r, 1, 20, 50)

	rows, err := h.store.ListFollowers(r.Context(), db.ListFollowersParams{
		FollowingID: userID,
		OffsetCount: int32(offset),
		LimitCount:  int32(pageSize),
	})
	if err != nil {
		log.Error().Err(err).Msg("获取粉丝列表失败")
		render.Error(w, http.StatusInternalServerError, "获取粉丝列表失败")
		return
	}

	users := make([]followUserResponse, 0, len(rows))
	for _, row := range rows {
		users = append(users, toFollowUserResponse(&row))
	}

	render.Success(w, "查询成功", users)
}

// ---- 关注列表 ----

//	@Summary	获取关注列表
//	@Tags		follows
//	@Param		id			path		string	true	"用户 ID"
//	@Param		page		query		int		false	"页码"	default(1)
//	@Param		page_size	query		int		false	"每页数量"	default(20)
//	@Success	200			{object}	render.Response[[]followUserResponse]
//	@Failure	400			{object}	render.errorResponse
//	@Failure	500			{object}	render.errorResponse
//	@Router		/users/{id}/following [get]
func (h *FollowHandler) ListFollowing(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的用户 ID")
		return
	}

	offset, pageSize := Pagination(r, 1, 20, 50)

	rows, err := h.store.ListFollowing(r.Context(), db.ListFollowingParams{
		FollowerID:  userID,
		OffsetCount: int32(offset),
		LimitCount:  int32(pageSize),
	})
	if err != nil {
		log.Error().Err(err).Msg("获取关注列表失败")
		render.Error(w, http.StatusInternalServerError, "获取关注列表失败")
		return
	}

	users := make([]followUserResponse, 0, len(rows))
	for _, row := range rows {
		resp := followUserResponse{
			ID:       row.ID,
			Username: row.Username,
		}
		if row.AvatarURL.Valid {
			resp.AvatarURL = row.AvatarURL.String
		}
		if row.Bio.Valid {
			resp.Bio = row.Bio.String
		}
		users = append(users, resp)
	}

	render.Success(w, "查询成功", users)
}
