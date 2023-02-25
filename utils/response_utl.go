package utils

import "github.com/gin-gonic/gin"

const (
	accepted = "Accepted"
	xml      = "application/xml"
)

func Response(c *gin.Context, code int, obj any) {
	switch c.GetHeader(accepted) {
	case xml:
		c.XML(code, obj)
	default:
		c.JSON(code, obj)
	}
}

func ResponseError(c *gin.Context, apiError *ApplicationError) {
	switch c.GetHeader(accepted) {
	case xml:
		c.XML(apiError.StatusCode, apiError)
	default:
		c.JSON(apiError.StatusCode, apiError)
	}
}
