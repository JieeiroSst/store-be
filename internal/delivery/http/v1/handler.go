package v1

import (
	"github.com/JIeeiroSst/store/internal/usecase"
	"github.com/JIeeiroSst/store/pkg/bigcache"
	"github.com/JIeeiroSst/store/pkg/jwt"
	"github.com/JIeeiroSst/store/pkg/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.Usecase
	redis   redis.RedisDB
	jwt     jwt.TokenUser
	cache   bigcache.Cache
}

func NewHandler(usecase *usecase.Usecase, redis redis.RedisDB, jwt jwt.TokenUser) *Handler {
	return &Handler{
		usecase: usecase,
		redis:   redis,
		jwt:     jwt,
	}
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initCartRoutes(v1)
		h.initCategoryRoutes(v1)
		h.initDiscountRoutes(v1)
		h.initPaymentRoutes(v1)
		h.initProductRoutes(v1)
		h.initRoleRoutes(v1)
		h.initSaleRoutes(v1)
		h.initUserRoutes(v1)
		h.initTokenRoutes(v1)
	}
}
