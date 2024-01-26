package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/go-api-oauth/pkg/auth"
)

func main() {

	jwksUrl := os.Getenv("JWKSURL")
	host := os.Getenv("HOST")

	jwks, err := auth.GetJWKS(jwksUrl)
	if err != nil {
		log.Fatalf("could not get jwks: %s", err.Error())
	}

	router := gin.New()
	router.Use(auth.TokenMiddleWare(jwks))

	router.GET("/public", GetPublic)
	router.GET("/authenticated", GetAuthenticated)
	router.GET("/authorized/:id", GetAuthorized)

	if err := router.Run(host); err != nil {
		log.Fatal("could not start server")
	}
}
