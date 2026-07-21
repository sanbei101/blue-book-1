package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/phuslu/log"

	"github.com/sanbei101/blue-book/internal/db"
	"github.com/sanbei101/blue-book/internal/pkg/jwt"
	"github.com/sanbei101/blue-book/internal/pkg/render"
)

type UserHandler struct {
	store *db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{store: store}
}

// ---- 注册 ----

type registerRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type authResponse struct {
	Token string       `json:"token"`
	User  userResponse `json:"user"`
}

type userResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	Bio       string    `json:"bio,omitempty"`
}

func toUserResponse(u *db.User) userResponse {
	resp := userResponse{
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

// 用户注册
//
//	@Summary	用户注册
//	@Tags		users
//	@Param		body	body		registerRequest	true	"注册信息"
//	@Success	200		{object}	render.Response[authResponse]
//	@Failure	409		{object}	render.errorResponse
//	@Failure	500		{object}	render.errorResponse
//	@Router		/users/register     [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := render.ReadBody[registerRequest](w, r)
	if err != nil {
		return
	}

	// TODO: 使用 bcrypt 哈希密码
	user, err := h.store.CreateUser(r.Context(), db.CreateUserParams{
		ID:           uuid.Must(uuid.NewV7()),
		Username:     body.Username,
		PasswordHash: body.Password,
		AvatarURL:    pgtype.Text{},
		Bio:          pgtype.Text{},
	})
	if err != nil {
		log.Error().Err(err).Msg("注册用户失败")
		render.Error(w, http.StatusConflict, "用户名已存在")
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Error().Err(err).Msg("生成 token 失败")
		render.Error(w, http.StatusInternalServerError, "生成 token 失败")
		return
	}

	render.Success(w, "注册成功", authResponse{
		Token: token,
		User:  toUserResponse(&user),
	})
}

// ---- 登录 ----

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// 用户登录
//
//	@Summary	用户登录
//	@Tags		users
//	@Param		body	body		loginRequest	true	"登录信息"
//	@Success	200		{object}	render.Response[authResponse]
//	@Failure	401		{object}	render.errorResponse
//	@Failure	500		{object}	render.errorResponse
//	@Router		/users/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := render.ReadBody[loginRequest](w, r)
	if err != nil {
		return
	}

	user, err := h.store.GetUserByUsername(r.Context(), body.Username)
	if err != nil {
		render.Error(w, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// TODO: 使用 bcrypt 校验密码
	if user.PasswordHash != body.Password {
		render.Error(w, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Error().Err(err).Msg("生成 token 失败")
		render.Error(w, http.StatusInternalServerError, "生成 token 失败")
		return
	}

	render.Success(w, "登录成功", authResponse{
		Token: token,
		User:  toUserResponse(&user),
	})
}

// 获取用户资料
//
//	@Summary	获取用户资料
//	@Tags		users
//	@Param		id	path		string	true	"用户 ID"
//	@Success	200	{object}	render.Response[userResponse]
//	@Failure	400	{object}	render.errorResponse
//	@Failure	404	{object}	render.errorResponse
//	@Router		/users/{id} [get]
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "无效的用户 ID")
		return
	}

	user, err := h.store.GetUserByID(r.Context(), id)
	if err != nil {
		render.Error(w, http.StatusNotFound, "用户不存在")
		return
	}

	render.Success(w, "查询成功", toUserResponse(&user))
}

// ---- 更新资料 ----

type updateProfileRequest struct {
	Username  string `json:"username"   validate:"required,min=3,max=32"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"        validate:"max=200"`
}

// 更新用户资料
//
//	@Summary	更新用户资料
//	@Tags		users
//	@Security	BearerAuth
//	@Param		body	body		updateProfileRequest	true	"更新信息"
//	@Success	200		{object}	render.Response[userResponse]
//	@Failure	400		{object}	render.errorResponse
//	@Failure	500		{object}	render.errorResponse
//	@Router		/users/profile [put]
func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	currentUserID := jwt.GetUserIDFromContext(r)

	body, err := render.ReadBody[updateProfileRequest](w, r)
	if err != nil {
		return
	}

	avatar := pgtype.Text{}
	if body.AvatarURL != "" {
		avatar = pgtype.Text{String: body.AvatarURL, Valid: true}
	}
	bio := pgtype.Text{}
	if body.Bio != "" {
		bio = pgtype.Text{String: body.Bio, Valid: true}
	}

	user, err := h.store.UpdateUser(r.Context(), db.UpdateUserParams{
		ID:        currentUserID,
		Username:  body.Username,
		AvatarURL: avatar,
		Bio:       bio,
	})
	if err != nil {
		log.Error().Err(err).Msg("更新用户资料失败")
		render.Error(w, http.StatusInternalServerError, "更新失败")
		return
	}

	render.Success(w, "更新成功", toUserResponse(&user))
}
