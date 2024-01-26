package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/go-api-oauth/pkg/auth"
)

// Everyone can view this
func GetPublic(c *gin.Context) {

	identity := "anonymous"

	// if there are valid claims, get the subject
	claims, err := auth.GetClaims(c)
	if err == nil {
		identity = claims.Subject
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":      "this is public information",
		"identity": identity,
	})
}

// only loggedin users can view this
func GetAuthenticated(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":      "you are only viewing this page because you are authenticated",
		"identity": claims.Subject,
	})

}

// only authorized users can view this
func GetAuthorized(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	id := c.Param("id")
	if id != claims.Subject {
		c.Status(http.StatusForbidden)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":      "you are only viewing this page because you are the resource owner",
		"identity": claims.Subject,
	})
}
