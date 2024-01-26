package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrorNoBearerToken = errors.New("no bearer token supplied")
	ErrorInvalidToken  = errors.New("invalid token")
)

type ApplicationClaims struct {
	jwt.RegisteredClaims
	// TODO: extend with custom claims
}

// TODO: make this configureable
func GetJWKS(url string) (*keyfunc.JWKS, error) {

	return keyfunc.Get(url, keyfunc.Options{
		RefreshInterval:  time.Hour,
		RefreshRateLimit: time.Minute * 5,
		RefreshTimeout:   time.Second * 10,
		RefreshErrorHandler: func(err error) {
			log.Printf("could not refresh jwks: %s", err.Error())
		},
	})
}

func TokenMiddleWare(jwks *keyfunc.JWKS) gin.HandlerFunc {

	return func(c *gin.Context) {

		bearer, ok := strings.CutPrefix(
			c.GetHeader("Authorization"),
			"Bearer ",
		)

		if !ok || bearer == "" {
			//c.AbortWithError(http.StatusUnauthorized, ErrorNoBearerToken)
			return
		}

		var claims ApplicationClaims

		token, err := jwt.ParseWithClaims(bearer, &claims, jwks.Keyfunc)

		if err != nil {
			//c.AbortWithError(http.StatusUnauthorized, ErrorNoBearerToken)
			return
		}

		c.Set("token", token)
		c.Set("claims", claims)
	}
}

func GetClaims(c *gin.Context) (ApplicationClaims, error) {
	var claims ApplicationClaims
	anyClaims, ok := c.Get("claims")
	if ok {
		claims, ok = anyClaims.(ApplicationClaims)
		if ok {
			return claims, nil
		}
		return claims, errors.New("claims not valid")
	}
	return claims, errors.New("claims not present")
}
