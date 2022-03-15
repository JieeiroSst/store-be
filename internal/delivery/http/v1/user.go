package v1

import (
	"context"
	"strconv"
	"strings"

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

	Reponse(ctx, 201, tokenDetail)
}

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

func (h *Handler) IsUserLogin(ctx *gin.Context) {
	id := ctx.Param("access-uuid")
	accessId, err := h.redis.FetchAuth(context.Background(), id)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
	}
	if len(strings.TrimSpace(accessId)) == 0 {
		message := ""
		ReponseError(ctx, 500, message)
	}
	res := map[string]interface{}{
		"access-uuid": id,
		"message":     "user is logged in",
	}
	Reponse(ctx, 200, res)
}
