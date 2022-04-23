package v1

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"net/http"

	ratelimmit "github.com/JIeeiroSst/store/pkg/rate_limmit"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initRateLimter(api *gin.RouterGroup) {
	api.Use(h.rateLimit)

}

func (h *Handler) clientIdentifier(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	url := ctx.Request.URL.Path
	data := fmt.Sprintf("%v-%v", ip, url)
	md5 := md5.Sum([]byte(data))
	hash := new(big.Int).SetBytes(md5[:]).Text(62)
	return hash
}

func (h *Handler) rateLimit(ctx *gin.Context) {
	var userType string
	if val, exists := ctx.Get("user-type"); exists {
		userType = val.(string)
	}
	if userType == "" {
		userType = "gen-user"
	}

	tokenBucket := ratelimmit.GetBucket(h.clientIdentifier(ctx), userType)
	if !tokenBucket.IsRequestAllowed(5) {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "try again after sometime!",
		})
		return
	}
	ctx.Next()
}
