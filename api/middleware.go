package api

import (
	"encoding/base64"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	config "github.com/os1-logistics/container-reference-app/configs"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, sentry-trace, X-COREOS-TID")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Authorize(c *gin.Context) {

	glog.Info("Authenticating request")
	apiKey := c.Request.Header.Get(common.ApiKeyHeaderName)

	if apiKey != "" {
		if apiKey == config.ServiceConf.APP.ApiKey {
			glog.Info("Valid API key")
			c.Next()
			return
		} else {
			c.JSON(401, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
	}

	requestAccessToken := c.Request.Header.Get(common.AcessHeaderName)

	if requestAccessToken == "" {
		c.JSON(401, gin.H{"error": "Missing access token"})
		c.Abort()
		return
	}

	glog.Info("Validating access token")
	// TODO: validate access token against coreos authorizer
	c.Next()
	return

}

func decodeJWT(token string) *JWTTokenSchema {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		glog.Info("Error in decoding token", err)
	}
	jwtToken := new(JWTTokenSchema)
	json.Unmarshal(data, &jwtToken)
	return jwtToken
}

type Userinfo struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	AppID string `json:"appID,omitempty"`
}

type JWTTokenSchema struct {
	Sub string `json:"sub,omitempty"`
	Aud string `json:"aud,omitempty"`
}
