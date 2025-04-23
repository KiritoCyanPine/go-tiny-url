package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kiritocyanpine/go-tiny-url/config"
)

const (
	RequestID_Key    = "request_id"
	HostName_URL_Key = "host_name"
)

// add the middleware function
func AllowCrossOriginRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	}
}

func RequestContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		// add a request ID
		requestID := uuid.New()
		c.Set(RequestID_Key, requestID)

		// add informaiton about the host URL
		hostAddress := config.GetConfigurations().HostAddress
		c.Set(HostName_URL_Key, hostAddress)
	}
}
