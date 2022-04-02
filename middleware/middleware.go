package middleware

import (
	"github.com/JIeeiroSst/store/pkg/bigcache"
	"github.com/JIeeiroSst/store/pkg/jwt"
	"github.com/JIeeiroSst/store/pkg/otp"
)

type Middleware struct {
	AccessOTP           AccessOTP
	AuthorizationCasbin AuthorizationCasbin
	AccessJWT           AccessJWT
}

type Dependency struct {
	Serect string
	Otp    otp.OTP
	Cache  bigcache.Cache
	Jwt    jwt.TokenUser
}

func NewMiddkeware(deps Dependency) *Middleware {
	otpMiddleware := NewAccessOTP(deps.Serect, deps.Otp, deps.Cache)
	casbinMiddleware := NewAuthorization(deps.Cache)
	JwtMiddleware := NewAccessController(deps.Jwt)
	return &Middleware{
		AccessOTP:           otpMiddleware,
		AuthorizationCasbin: casbinMiddleware,
		AccessJWT:           JwtMiddleware,
	}
}
