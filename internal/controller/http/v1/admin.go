package v1

import (
	"coffe-life/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) newAdminPanel() {
	group := h.group.Group("admin/")
	group.GET("categories/", h.GetCategories)
}

// createPartner get categories.
// @Tags categories
// @Summary Get categories
// @Description Get all categories
// // @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response{payload=domain.Categories} "categories"
// @Failure 400 {object} domain.Response
// @Failure 401 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /admin/categories [get]
func (h *handler) GetCategories(c *gin.Context) {

	resp, err := h.usecases.Admin().GetCategories(c.Request.Context())
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Payload: resp,
	})
}
