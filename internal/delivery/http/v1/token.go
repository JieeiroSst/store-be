package v1

import (
	"github.com/JIeeiroSst/store/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initTokenRoutes(api *gin.RouterGroup) {

}

func (h *Handler) RegenerateAccessToken(ctx *gin.Context) {
	var token model.Token
	if err := ctx.ShouldBind(&token); err != nil {
		ReponseError(ctx, 400, err.Error())
	}
	accessUuid := ctx.Param("access-uuid")
	accessToken, err := h.usecase.Tokens.RegenerateAccessToken(token, accessUuid)
	if err != nil {
		ReponseError(ctx, 500, err.Error())
	}
	Reponse(ctx, 200, accessToken)
}