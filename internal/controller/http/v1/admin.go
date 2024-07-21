package v1

import (
	"coffe-life/internal/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) newAdminPanel() {
	group := h.group.Group("admin/")
	group.GET("categories/", h.GetCategories)
	group.POST("categories/", h.CreateCategory)
	group.PUT("categories/:id", h.UpdateCategory)
	group.DELETE("categories/:id", h.DeleteCategory)
}

// GetCategories get categories.
// @Tags categories
// @Summary Get categories
// @Description Get all categories
// // @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{payload=dto.Categories} "categories"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/categories [get]
func (h *handler) GetCategories(c *gin.Context) {

	resp, err := h.usecases.Admin().GetCategories(c.Request.Context())
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: resp,
	})
}

// CreateCategory create category.
// @Tags categories
// @Summary Create category
// @Description Create category
// // @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body dto.CategoryRequest true "Category"
// @Success 200 {object} dto.Response{payload=string} "category"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/categories [post]
func (h *handler) CreateCategory(c *gin.Context) {
	var req dto.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	resp, err := h.usecases.Admin().CreateCategory(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: resp,
	})
}

// UpdateCategory update category.
// @Tags categories
// @Summary Update category
// @Description Update category
// // @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param request body dto.CategoryRequest true "Category"
// @Success 200 {object} dto.Response{payload=string} "category"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/categories/{id} [put]
func (h *handler) UpdateCategory(c *gin.Context) {
	var pathParams dto.IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	var req dto.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().UpdateCategory(c.Request.Context(), pathParams.ID, req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: &dto.Success{Message: "success"},
	})
}

// DeleteCategory delete category.
// @Tags categories
// @Summary Delete category
// @Description Delete category
// // @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} dto.Response{payload=string} "category"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/categories/{id} [delete]
func (h *handler) DeleteCategory(c *gin.Context) {
	var pathParams dto.IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	err := h.usecases.Admin().DeleteCategory(c.Request.Context(), pathParams.ID)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: &dto.Success{Message: "success"},
	})
}
