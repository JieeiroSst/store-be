package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initRoleRoutes(api *gin.RouterGroup) {
	roleGroup := api.Group("/role")

	roleGroup.POST("/", h.createRole)
	roleGroup.PUT("/:id", h.updateRole)
	roleGroup.DELETE("/:id", h.deleteRole)
	roleGroup.GET("/", h.roles)
	roleGroup.GET("/:id", h.roleById)

}

func (h *Handler) createRole(ctx *gin.Context) {

}

func (h *Handler) updateRole(ctx *gin.Context) {

}

func (h *Handler) deleteRole(ctx *gin.Context) {

}

func (h *Handler) roles(ctx *gin.Context) {

}

func (h *Handler) roleById(ctx *gin.Context) {

}
