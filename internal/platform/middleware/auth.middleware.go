package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const basicAuthPrefix = "Basic "

		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, basicAuthPrefix) {
			encodedCredentials := strings.TrimPrefix(authHeader, basicAuthPrefix)
			decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)

			if err == nil {
				credentials := strings.SplitN(string(decodedBytes), ":", 2)
				if len(credentials) == 2 {
					username := credentials[0]
					password := credentials[1]

					if username == "adminUser" && password == "qwerty123" {
						c.Set("username", username)
						c.Next()
						return
					} else {

						fmt.Println("Invalid credentials")
					}
				}
			}
		}

		c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
