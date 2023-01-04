package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/os1-logistics/container-reference-app/api/v1"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
)

// ValidateHeader func export
func ValidateHeaders(c *gin.Context) {

	if len(c.Request.Header[common.TenantIDHeaderName]) == 0 {
		glog.Info("X-COREOS-TID Header not set")

		response := api.ErrorSchema{
			Description: "Missing required header",
			AdditionalInfo: &map[string]interface{}{
				"error": "X-COREOS-TID not set. Set the header with valid tenant name.",
			},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if len(c.Request.Header[common.RequestIDHeaderName]) == 0 {
		glog.Info("X-COREOS-REQUEST-ID Header not set")
		response := api.ErrorSchema{
			Description: "Missing required header",
			AdditionalInfo: &map[string]interface{}{
				"error": "X-COREOS-REQUEST-ID not set. Set the header with valid request Id.",
			},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.Next()
	return
}

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
