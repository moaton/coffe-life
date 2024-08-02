package v1

import (
	"coffe-life/internal/dto"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) newAdminPanel() {
	group := h.group.Group("admin/")
	group.POST("login", h.Login)

	groupAuth := h.group.Group("admin/data/")
	groupAuth.Use(RequireAuth(h.cfg))

	user := group.Group("user/")
	user.Use(RequireAuth(h.cfg))
	user.POST("", h.CreateUser)

	groupAuth.GET("categories/", h.GetCategories)
	groupAuth.POST("categories/", h.CreateCategory)
	groupAuth.PUT("categories/:id", h.UpdateCategory)
	groupAuth.DELETE("categories/:id", h.DeleteCategory)

	groupAuth.GET("foods/", h.GetFoods)
	groupAuth.POST("foods/", h.CreateFood)
	groupAuth.PUT("foods/:id", h.UpdateFood)
	groupAuth.DELETE("foods/:id", h.DeleteFood)
}

// Login admin login.
// @Tags auth
// @Summary Login
// @Description Login
// @Param request body dto.LoginRequest true "Login"
// @Success 200 {object} dto.Response{payload=dto.AuthResponse} "login"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/login [post]
func (h *handler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBind(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	resp, err := h.usecases.Admin().Login(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: resp,
	})
}

// User admin create user.
// @Tags user
// @Summary Create user
// @Description User
// @Security ApiKeyAuth
// @Param request body dto.CreateUserRequest true "User request"
// @Success 200 {object} dto.Response{payload=string} "user"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/user [post]
func (h *handler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().CreateUser(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: "success",
	})
}

// GetCategories get categories.
// @Tags categories
// @Summary Get categories
// @Description Get all categories
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{payload=dto.Categories} "categories"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/categories [get]
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
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body dto.CategoryRequest true "Category"
// @Success 200 {object} dto.Response{payload=string} "category"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/categories [post]
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
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param request body dto.CategoryRequest true "Category"
// @Success 200 {object} dto.Response{payload=string} "category"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/categories/{id} [put]
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
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} dto.Response{payload=string} "category"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/categories/{id} [delete]
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

// GetFoods get foods.
// @Tags foods
// @Summary Get foods
// @Description Get all foods
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{payload=dto.Foods} "foods"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/foods [get]
func (h *handler) GetFoods(c *gin.Context) {

	resp, err := h.usecases.Admin().GetFoods(c.Request.Context())
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: resp,
	})
}

// CreateFood create food.
// @Tags foods
// @Summary Create food
// @Description Create food
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body dto.FoodRequest true "Food"
// @Success 200 {object} dto.Response{payload=string} "food"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/foods [post]
func (h *handler) CreateFood(c *gin.Context) {
	var req dto.FoodRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}
	log.Println("req", req)
	resp, err := h.usecases.Admin().CreateFood(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: resp,
	})
}

// UpdateFood update food.
// @Tags foods
// @Summary Update food
// @Description Update food
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param request body dto.FoodRequest true "Food"
// @Success 200 {object} dto.Response{payload=string} "food"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/foods/{id} [put]
func (h *handler) UpdateFood(c *gin.Context) {
	var pathParams dto.IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	var req dto.FoodRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().UpdateFood(c.Request.Context(), pathParams.ID, req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: &dto.Success{Message: "success"},
	})
}

// DeleteFood delete food.
// @Tags foods
// @Summary Delete food
// @Description Delete food
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} dto.Response{payload=string} "food"
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/data/foods/{id} [delete]
func (h *handler) DeleteFood(c *gin.Context) {
	var pathParams dto.IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	err := h.usecases.Admin().DeleteFood(c.Request.Context(), pathParams.ID)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Payload: &dto.Success{Message: "success"},
	})
}
