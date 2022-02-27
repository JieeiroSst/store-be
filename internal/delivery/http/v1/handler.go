package v1

import (
	"github.com/JIeeiroSst/store/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

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
	}
}
