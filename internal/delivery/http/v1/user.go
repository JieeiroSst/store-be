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
	userGroup := api.Group("/accounts")

	userGroup.POST("/login", h.Login)
	userGroup.POST("/register", h.SignUp)
	userGroup.PUT("/:user-id", h.Update)
	userGroup.GET("/:user-id", h.UserById)
	userGroup.GET("/", h.Users)
	userGroup.POST("/:access-uuid", h.LogOut)
	userGroup.PATCH("/:user-id", h.LockUser)
	userGroup.OPTIONS("/:access-uuid", h.IsUserLogin)
	userGroup.PUT("/:access-uuid", h.RefreshStoreToken)
}

// Login godoc
// @Summary      Login an account
// @Description  Login an account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.User
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/accounts/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBind(&user); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	token, err := h.usecase.Users.Login(user)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}

	var tokenDetail model.TokenResult

	if err := marshalJson(token, tokenDetail); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}

	u, err := uuid.NewRandom()
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), user.Username)
	if err := h.cache.Set(sessionId, []byte(user.Username)); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}

	ctx.SetCookie("current_subject", sessionId, 30*60, "/api", "", false, true)
	reponse(ctx, 201, tokenDetail)
}

// SignUp godoc
// @Summary      ShowSignUp an account
// @Description  ShowSignUp an account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Account
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/accounts/{id} [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBind(&user); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	if err := h.usecase.Users.SignUp(user); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}

	reponse(ctx, 201, user)
}

// Update godoc
// @Summary      Update an account
// @Description  Update an account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user-id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/accounts/{user-id} [put]
func (h *Handler) Update(ctx *gin.Context) {
	id := ctx.Param("user-id")
	var user domain.InputUser
	if err := ctx.ShouldBind(&user); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	if err := h.usecase.Users.Update(id, user); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	req := map[string]interface{}{
		"id":   id,
		"data": user,
	}
	reponse(ctx, 200, req)
}

// ShowAccount By ID godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user-id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      500  {string}  string
// @Router       /api/v1/accounts/{user-id} [get]
func (h *Handler) UserById(ctx *gin.Context) {
	id := ctx.Param("user-id")
	user, err := h.usecase.Users.UserById(id)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}

	reponse(ctx, 200, user)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        limit   path      int  true  "Account ID"
// @Param        offset   path      int  true  "Account ID"
// @Success      200  {array}  model.Account
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/accounts [get]
func (h *Handler) Users(ctx *gin.Context) {
	limmit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	pagination := domain.Pagination{
		Limit:  limmit,
		Offset: offset,
	}

	users, err := h.usecase.Users.UserByAll(pagination)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	reponse(ctx, 200, users)
}

// LockUser godoc
// @Summary      Lock an account
// @Description  lock string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {string}  string
// @Router       /api/v1/accounts/{user-id} [patch]
func (h *Handler) LockUser(ctx *gin.Context) {
	id := ctx.Param("user-id")
	if err := h.usecase.Users.LockUser(id); err != nil {
		reponseError(ctx, 500, err.Error())
	}

	res := map[string]interface{}{
		"id":      id,
		"message": "lock user by admin ",
	}
	reponse(ctx, 200, res)
}

// LogOut godoc
// @Summary      LogOut an account
// @Description  LogOut string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        access-uuid   path      string  true  "access uuid"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {string}  string
// @Router       /api/v1/accounts/{access-uuid} [post]
func (h *Handler) LogOut(ctx *gin.Context) {
	id := ctx.Param("access-uuid")

	deletedId, err := h.redis.DeleteAuth(context.Background(), id)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return 
	}

	res := map[string]interface{}{
		"deleteId": deletedId,
		"message":  "log out user by id = " + id,
	}
	reponse(ctx, 200, res)
}

// IsUserLogin godoc
// @Summary      IsUserLogin an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        access-uuid   path      int  true  "access uuid"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h *Handler) IsUserLogin(ctx *gin.Context) {
	accessId := ctx.Param("access-uuid")
	userId, err := h.redis.FetchAuth(context.Background(), accessId)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return 
	}
	if len(strings.TrimSpace(userId)) == 0 {
		message := "empty user id"
		reponseError(ctx, 500, message)
		return 
	}
	res := map[string]interface{}{
		"access-uuid": accessId,
		"message":     "user is logged in",
	}
	reponse(ctx, 200, res)
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
		reponseError(ctx, 400, err.Error())
	}

	refereToken, err := h.jwt.RegenerateAccessToken(token, accessUuid)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}

	if err := h.redis.Set(context.Background(), accessUuid, refereToken); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	res := map[string]interface{}{
		"accessId": accessUuid,
		"token":    refereToken,
	}
	reponse(ctx, 200, res)

}
