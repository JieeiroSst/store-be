package middleware

import (
	"strings"

	"github.com/JIeeiroSst/store/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type AccessJWTs struct {
	jwt jwt.TokenUser
}

type AccessJWT interface {
	Authenticate() gin.HandlerFunc
}

func NewAccessController(jwt jwt.TokenUser) *AccessJWTs {
	return &AccessJWTs{
		jwt: jwt,
	}
}

func (a *AccessJWTs) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		if len(strings.TrimSpace(bearToken)) == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Authentication failure: Token not provided",
			})
			return
		}
		tokenSlice := strings.Split(bearToken, " ")
		token, err := a.jwt.ParseToken(tokenSlice[1])
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"message": token,
			})
			return
		}
		c.Next()
	}
}
