package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errResponse(c *gin.Context, code int, message string) {
	var response Response

	response.Meta = struct{}{}
	response.ErrorCode = http.StatusText(code)
	response.Description = message
	response.Payload = struct{}{}

	c.AbortWithStatusJSON(code, &response)
}
