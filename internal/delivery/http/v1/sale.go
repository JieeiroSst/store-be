package v1

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initSaleRoutes(api *gin.RouterGroup) {
	saleGroup := api.Group("/sales")

	saleGroup.POST("/", h.createSale)
	saleGroup.PUT("/:id", h.updateSale)
	saleGroup.DELETE("/:id", h.deleteSale)
	saleGroup.GET("/", h.sales)
	saleGroup.GET("/:id", h.SaleById)
	saleGroup.PATCH("/:id", h.IsExpireById)

}

// createSale godoc
// @Summary      create Sale
// @Description  create Sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        sale  body model.InputSale true  "sale"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  v1.ErroMessage
// @Failure      404  {object}  v1.ErroMessage
// @Failure      500  {object}  v1.ErroMessage
// @Router       /api/v1/sales [post]
func (h *Handler) createSale(ctx *gin.Context) {
	var sale model.InputSale
	if err := ctx.ShouldBind(&sale); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	if err := h.usecase.Sales.Create(sale); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	reponse(ctx, 200, sale)
}

// updateSale godoc
// @Summary      update Sale
// @Description  update Sale by ID
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Account ID"
// @Param        sale  body model.InputSale true  "sale"
// @Success      200  {object}  model.InputSale
// @Failure      400  {object}  v1.ErroMessage
// @Failure      404  {object}  v1.ErroMessage
// @Failure      500  {object}  v1.ErroMessage
// @Router       /api/v1/sales/{id} [put]
func (h *Handler) updateSale(ctx *gin.Context) {
	id := ctx.Param("id")
	var sale model.InputSale
	if err := ctx.ShouldBind(&sale); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	if err := h.usecase.Sales.Update(id, sale); err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}
	reponse(ctx, 200, sale)
}

// deleteSale godoc
// @Summary      delete Sale
// @Description  delete Sale by ID
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Account ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  v1.ErroMessage
// @Failure      404  {object}  v1.ErroMessage
// @Failure      500  {object}  v1.ErroMessage
// @Router       /api/v1/sales/{id} [delete]
func (h *Handler) deleteSale(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.usecase.Sales.Delete(id); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	res := map[string]interface{}{
		"id":      id,
		"message": "delete sale success",
	}
	reponse(ctx, 200, res)
}

// sales godoc
// @Summary      Show an sales
// @Description  get sales
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        pagination   body      domain.Pagination  true  "Pagination"
// @Success      200  {array}   domain.Sale
// @Failure      400  {object}  v1.ErroMessage
// @Failure      404  {object}  v1.ErroMessage
// @Failure      500  {object}  v1.ErroMessage
// @Router       /api/v1/sales [get]
func (h *Handler) sales(ctx *gin.Context) {
	var pagination domain.Pagination
	if err := ctx.ShouldBind(&pagination); err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}

	sales, err := h.usecase.Sales.Sales(pagination)
	if err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}
	reponse(ctx, 200, sales)
}

// SaleById godoc
// @Summary      Show an SaleById
// @Description  get sale by ID
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Sales ID"
// @Success      200  {object}  domain.Sale
// @Failure      400  {object}  v1.ErroMessage
// @Failure      404  {object}  v1.ErroMessage
// @Failure      500  {object}  v1.ErroMessage
// @Router       /api/v1/sales/{id} [get]
func (h *Handler) SaleById(ctx *gin.Context) {
	id := ctx.Param("id")
	sale, err := h.usecase.Sales.SaleById(id)
	if err != nil {
		reponseError(ctx, 400, err.Error())
		return
	}
	reponse(ctx, 200, sale)
}

// IsExpireById godoc
// @Summary      Is Expire By Id
// @Description  Is Expire By Id
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Sales ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  v1.ErroMessage
// @Failure      404  {object}  v1.ErroMessage
// @Failure      500  {object}  v1.ErroMessage
// @Router       /api/v1/sales/{id} [patch]
func (h *Handler) IsExpireById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.usecase.Sales.IsExpireById(id)

	if errors.Is(err, errors.New("sale time is over")) {
		reponseError(ctx, 400, err.Error())
		return
	}

	if err != nil {
		reponseError(ctx, 500, err.Error())
		return
	}

	res := map[string]interface{}{
		"message": "sale time is over",
	}

	reponse(ctx, 200, res)
}
