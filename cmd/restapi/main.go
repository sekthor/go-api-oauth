package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/go-api-oauth/pkg/auth"
)

func main() {

	jwks, err := auth.GetJWKS("http://localhost:8080/realms/test/protocol/openid-connect/certs")
	if err != nil {
		log.Fatalf("could not get jwks: %s", err.Error())
	}

	router := gin.New()
	router.Use(auth.TokenMiddleWare(jwks))

	router.GET("/public", GetPublic)
	router.GET("/authenticated", GetAuthenticated)
	router.GET("/authorized/:id", GetAuthorized)

	if err := router.Run(":8081"); err != nil {
		log.Fatal("could not start server")
	}
}
