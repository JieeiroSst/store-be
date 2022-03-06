package v1

import "github.com/gin-gonic/gin"

func ReponseError(ctx *gin.Context, code int, message string) {
	ctx.JSONP(code, gin.H{
		"message": message,
	})
	return
}

func Reponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, data)
}
