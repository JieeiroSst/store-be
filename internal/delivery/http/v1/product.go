package v1

import (
	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initProductRoutes(api *gin.RouterGroup) {
	productGroup := api.Group("/products")
	productGroup.POST("/", h.createProduct)
	productGroup.PUT("/:id", h.updateProduct)
	productGroup.DELETE("/:id", h.deleteProduct)
	productGroup.GET("/:id", h.productByID)
	productGroup.GET("/", h.products)
}

func (h *Handler) createProduct(ctx *gin.Context) {
	var product model.InputProduct
	if err := ctx.ShouldBind(&product); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}
	if err := h.usecase.Products.Create(product); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	reponse(ctx, 200, product)
}

func (h *Handler) updateProduct(ctx *gin.Context) {
	var product model.InputProduct
	if err := ctx.ShouldBind(&product); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}
	id := ctx.Param("id")
	if err := h.usecase.Products.Update(id, product); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	reponse(ctx, 200, product)
}

func (h *Handler) deleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.usecase.Products.Delete(id); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	res := map[string]interface{}{
		"id":      id,
		"message": "delete by id product success",
	}
	reponse(ctx, 200, res)
}

func (h *Handler) productByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.usecase.Products.ProductById(id)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	reponse(ctx, 200, product)
}

func (h *Handler) products(ctx *gin.Context) {
	var pagination domain.Pagination
	if err := ctx.ShouldBind(&pagination); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}
	products, err := h.usecase.Products.Products(pagination)
	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	reponse(ctx, 200, products)
}
