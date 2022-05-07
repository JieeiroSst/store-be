package v1

import (
	"github.com/JIeeiroSst/store/internal/usecase"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initMediaRoutes(api *gin.RouterGroup) {
	uploadGroup := api.Group("/upload")

	uploadGroup.POST("/", h.Upload)
}

func (h *Handler) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}
	if err := h.usecase.Medias.CreateMedia(ctx, &usecase.CreateRequest{
		FileHeader: file,
	}); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	res := map[string]interface{}{
		"message": "upload file success",
	}
	reponse(ctx, 200, res)
}
