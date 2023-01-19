package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	config "github.com/os1-logistics/container-reference-app/configs"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
	domain "github.com/os1-logistics/container-reference-app/internal/pkg/domain"
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
	tenantId := c.Request.Header.Get(common.TenantIDHeaderName)
	requestId := c.Request.Header.Get(common.RequestIDHeaderName)
	apiKey := c.Request.Header.Get(common.ApiKeyHeaderName)

	user := Userinfo{}
	user.AppID = config.ServiceConf.APP.AppId
	user.Name = config.ServiceConf.APP.AppName

	c.Set(common.ContextTenantId, tenantId)
	c.Set(common.ContextUserinfo, user)
	c.Set(common.ContextRequestID, requestId)

	// skip authorization if AUTH_ENABLED environment variable is set to false
	if config.ServiceConf.APP.AuthEnabled == false {
		glog.Info("App authorization is disabled")
		c.Next()
		return
	}

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
	ApiValidateTokenRequest := domain.NewAuthorizerClient(tenantId).AuthorizationApi.ValidateToken(c.Request.Context()).
		XCOREOSTID(tenantId).
		XCoreosRequestId(c.Request.Header.Get(common.RequestIDHeaderName)).
		XCoreosAccess(requestAccessToken)
	_, r, e := domain.NewAuthorizerClient(tenantId).AuthorizationApi.ValidateTokenExecute(ApiValidateTokenRequest)

	if e != nil {
		glog.Info("Access token validation failed")
		glog.Error(r)
		glog.Error(e)
		c.JSON(401, gin.H{"error": e.Error()})
		c.Abort()
		return
	}

	c.Next()
	return

}

func ContextMiddlewareToken(c *gin.Context) {
	tenantId := c.GetString(common.ContextTenantId)
	token, err := domain.GetToken(tenantId)
	if err != nil {
		glog.Error(err)
		c.JSON(500, gin.H{"error": "Error Generating middleware token"})
		c.Abort()
		return
	}
	c.Set(common.ContextMiddlewareAccessToken, token)
	c.Next()
}

type Userinfo struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	AppID string `json:"appID,omitempty"`
}
