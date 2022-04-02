package middleware

import (
	"errors"
	"fmt"
	"net/http"

	cacheErr "github.com/allegro/bigcache"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"

	"github.com/JIeeiroSst/store/pkg/bigcache"
	"github.com/gin-gonic/gin"
)

type AuthorizationCasbin interface {
	Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc
	Authenticate() gin.HandlerFunc
	enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error)
}

type AuthorizationCasbins struct {
	cache bigcache.Cache
}

func NewAuthorization(cache bigcache.Cache) *AuthorizationCasbins {
	return &AuthorizationCasbins{
		cache: cache,
	}
}

func (a *AuthorizationCasbins) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, _ := c.Cookie("current_subject")
		sub, err := a.cache.Get(sessionId)
		if errors.Is(err, cacheErr.ErrEntryNotFound) {
			c.JSON(401, map[string]interface{}{"message": "user hasn't logged in yet", "status": 401})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("current_subject", string(sub))
		c.Next()
	}
}

func (a *AuthorizationCasbins) Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("current_subject")
		val, existed := a.cache.Get(cookie)
		if existed != nil {
			c.JSON(401, map[string]interface{}{"message": "user hasn't logged in yet", "status": 401})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ok, err := a.enforce(string(val), obj, act, adapter)
		if err != nil {
			c.JSON(403, map[string]interface{}{"message": "error occurred when authorizing user", "status": 403})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !ok {
			c.JSON(403, map[string]interface{}{"message": "forbidden", "status": 403})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func (a *AuthorizationCasbins) enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer("pkg/conf/rbac_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	return ok, err
}
