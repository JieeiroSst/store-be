package http

import (
	v1 "github.com/JIeeiroSst/store/internal/delivery/http/v1"
	"github.com/JIeeiroSst/store/internal/usecase"
	"github.com/JIeeiroSst/store/pkg/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase usecase.Usecase
	redis   redis.RedisDB
}

func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	return router
}

func (h *Handler) InitApi(router *gin.Engine) {
	handlerV1 := v1.NewHandler(&h.usecase, h.redis)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
