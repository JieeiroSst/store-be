package http

import (
	v1 "github.com/JIeeiroSst/store/internal/delivery/http/v1"
	"github.com/JIeeiroSst/store/internal/usecase"
	"github.com/JIeeiroSst/store/pkg/jwt"
	"github.com/JIeeiroSst/store/pkg/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase usecase.Usecase
	redis   redis.RedisDB
	jwt     jwt.TokenUser
}

func NewHandler(usecase usecase.Usecase, jwt jwt.TokenUser, redis redis.RedisDB) *Handler {
	return &Handler{
		usecase: usecase,
		redis:   redis,
		jwt:     jwt,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	handlerV1 := v1.NewHandler(&h.usecase, h.redis, h.jwt)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
