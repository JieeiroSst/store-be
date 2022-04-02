package middleware

import (
	"net/http"

	"github.com/JIeeiroSst/store/pkg/bigcache"
	"github.com/JIeeiroSst/store/pkg/otp"
	"github.com/gin-gonic/gin"
)

type AccessOtps struct {
	serect string
	otp    otp.OTP
	cache  bigcache.Cache
}

type AccessOTP interface {
	Authenticate() gin.HandlerFunc
}

func NewAccessOTP(serect string, otp otp.OTP, cache bigcache.Cache) *AccessOtps {
	return &AccessOtps{
		serect: serect,
		otp:    otp,
		cache:  cache,
	}
}

func (a *AccessOtps) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("current_subject")
		val, existed := a.cache.Get(cookie)
		if existed != nil {
			c.JSON(401, map[string]interface{}{"message": "user hasn't logged in yet", "status": 401})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		otp := c.Request.Header.Get("otp")
		if err := a.otp.Authorize(otp, string(val)); err != nil {
			c.JSON(401, map[string]interface{}{"message": "otp expire in yet", "status": 403})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
