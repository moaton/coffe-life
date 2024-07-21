package v1

import (
	"coffe-life/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errResponse(c *gin.Context, code int, message string) {
	var response dto.Response

	response.Meta = struct{}{}
	response.ErrorCode = http.StatusText(code)
	response.Description = message
	response.Payload = struct{}{}

	c.AbortWithStatusJSON(code, &response)
}
