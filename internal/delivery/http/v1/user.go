package v1

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	userGroup := api.Group("/user")

	userGroup.POST("/login", h.Login)
	userGroup.POST("/register", h.SignUp)
	userGroup.POST("/:user-id", h.Update)

}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *Handler) Login(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBind(&user); err != nil {
		ReponseError(ctx, 400, err.Error())
	}

	token, err := h.usecase.Users.Login(user)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
	}

	var tokenDetail model.TokenResult

	if err := marshalJson(token, tokenDetail); err != nil {
		ReponseError(ctx, 500, err.Error())
	}

	u, err := uuid.NewRandom()
	if err != nil {
		Reponse(ctx, 500, map[string]interface{}{"data": "no data", "status": 500})
		return
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), user.Username)
	if err := h.cache.Set(sessionId, []byte(user.Username)); err != nil {
		Reponse(ctx, 500, map[string]interface{}{"data": "no data", "status": 500})
		return
	}

	ctx.SetCookie("current_subject", sessionId, 30*60, "/api", "", false, true)
	Reponse(ctx, 201, tokenDetail)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *Handler) SignUp(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBind(&user); err != nil {
		ReponseError(ctx, 400, err.Error())
	}

	if err := h.usecase.Users.SignUp(user); err != nil {
		ReponseError(ctx, 500, err.Error())
	}

	Reponse(ctx, 201, user)
}

func (h *Handler) Update(ctx *gin.Context) {
	id := ctx.Param("user-id")
	var user domain.InputUser
	if err := ctx.ShouldBind(&user); err != nil {
		ReponseError(ctx, 400, err.Error())
	}

	if err := h.usecase.Users.Update(id, user); err != nil {
		ReponseError(ctx, 400, err.Error())
	}

	reponse := map[string]interface{}{
		"id":   id,
		"data": user,
	}
	Reponse(ctx, 200, reponse)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *Handler) UserById(ctx *gin.Context) {
	id := ctx.Param("user-id")
	user, err := h.usecase.Users.UserById(id)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
	}

	Reponse(ctx, 200, user)
}

func (h *Handler) Users(ctx *gin.Context) {
	limmit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ReponseError(ctx, 400, err.Error())
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ReponseError(ctx, 400, err.Error())
	}

	pagination := domain.Pagination{
		Limit:  limmit,
		Offset: offset,
	}

	users, err := h.usecase.Users.UserByAll(pagination)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
	}
	Reponse(ctx, 200, users)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *Handler) LockUser(ctx *gin.Context) {
	id := ctx.Param("user-id")
	if err := h.usecase.Users.LockUser(id); err != nil {
		ReponseError(ctx, 500, err.Error())
	}

	res := map[string]interface{}{
		"id":      id,
		"message": "lock user by admin ",
	}
	Reponse(ctx, 200, res)
}

func (h *Handler) LogOut(ctx *gin.Context) {
	id := ctx.Param("access-uuid")

	deletedId, err := h.redis.DeleteAuth(context.Background(), id)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
	}

	res := map[string]interface{}{
		"deleteId": deletedId,
		"message":  "log out user by id = " + id,
	}
	Reponse(ctx, 200, res)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *Handler) IsUserLogin(ctx *gin.Context) {
	accessId := ctx.Param("access-uuid")
	userId, err := h.redis.FetchAuth(context.Background(), accessId)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
	}
	if len(strings.TrimSpace(userId)) == 0 {
		message := ""
		ReponseError(ctx, 500, message)
	}
	res := map[string]interface{}{
		"access-uuid": accessId,
		"message":     "user is logged in",
	}
	Reponse(ctx, 200, res)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *Handler) RefreshStoreToken(ctx *gin.Context) {
	accessUuid := ctx.Param("access-uuid")
	var token model.Token
	if err := ctx.ShouldBind(&token); err != nil {
		ReponseError(ctx, 400, err.Error())
	}

	refereToken, err := h.jwt.RegenerateAccessToken(token, accessUuid)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
		return
	}

	if err := h.redis.Set(context.Background(), accessUuid, refereToken); err != nil {
		ReponseError(ctx, 500, err.Error())
		return
	}
	res := map[string]interface{}{
		"accessId": accessUuid,
		"token":    refereToken,
	}
	Reponse(ctx, 200, res)

}
