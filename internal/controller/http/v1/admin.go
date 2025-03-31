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
	user.GET("", h.GetUser)
	user.GET(":id", h.GetUserById)
	user.PUT(":id", h.UpdateUser)

	groupAuth.GET("categories/", h.GetCategories)
	groupAuth.POST("categories/", h.CreateCategory)
	groupAuth.PUT("categories/:id", h.UpdateCategory)
	groupAuth.DELETE("categories/:id", h.DeleteCategory)

	groupAuth.GET("foods/", h.GetFoods)
	groupAuth.POST("foods/", h.CreateFood)
	groupAuth.PUT("foods/:id", h.UpdateFood)
	groupAuth.DELETE("foods/:id", h.DeleteFood)

	groupAuth.GET("translates/", h.GetTranslates)
	groupAuth.GET("translates/:id", h.GetTranslateById)
	groupAuth.POST("translates/", h.CreateTranslates)
	groupAuth.PUT("translates/:id", h.UpdateTranslate)
}

// Login admin login.
// @Tags auth
// @Summary Login
// @Description Login
// @Param request body dto.LoginRequest true "Login"
// @Success 200 {object} Response{payload=dto.AuthResponse} "login"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/login [post]
func (h *handler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBind(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	resp, err := h.usecases.Admin().Users().Login(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: resp,
	})
}

// User admin create user.
// @Tags user
// @Summary Create user
// @Description User
// @Security ApiKeyAuth
// @Param request body dto.CreateUserRequest true "User request"
// @Success 200 {object} Response{payload=string} "user"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/user [post]
func (h *handler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().Users().CreateUser(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: "success",
	})
}

// User admin get users.
// @Tags user
// @Summary Get users
// @Description Users
// @Security ApiKeyAuth
// @Param search query string false "search"
// @Param page query string false "page"
// @Param size query string false "size"
// @Success 200 {object} Response{payload=[]dto.User} "user"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/user [get]
func (h *handler) GetUser(c *gin.Context) {
	var req dto.GetUsersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}
	log.Println("req", req)
	resp, err := h.usecases.Admin().Users().GetUsers(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: resp,
	})
}

// User admin get user by id.
// @Tags user
// @Summary Get user by id
// @Description User
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} Response{payload=dto.User} "user"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/user/{id} [get]
func (h *handler) GetUserById(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	resp, err := h.usecases.Admin().Users().GetUserById(c.Request.Context(), pathParams.ID)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: resp,
	})
}

// User admin get user by id.
// @Tags user
// @Summary Get user by id
// @Description User
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Param request body dto.UpdateUserRequest true "User"
// @Success 200 {object} Response{payload=string} "user"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/user/{id} [put]
func (h *handler) UpdateUser(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	var req dto.UpdateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().Users().UpdateUser(c.Request.Context(), pathParams.ID, req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
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
// @Success 200 {object} Response{payload=dto.Categories} "categories"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/categories [get]
func (h *handler) GetCategories(c *gin.Context) {

	resp, err := h.usecases.Admin().Categories().GetCategories(c.Request.Context())
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
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
// @Param request body dto.Category true "Category"
// @Success 200 {object} Response{payload=string} "category"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/categories [post]
func (h *handler) CreateCategory(c *gin.Context) {
	var req dto.Category

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	resp, err := h.usecases.Admin().Categories().CreateCategory(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
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
// @Param request body dto.Category true "Category"
// @Success 200 {object} Response{payload=string} "category"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/categories/{id} [put]
func (h *handler) UpdateCategory(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	var req dto.Category

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().Categories().UpdateCategory(c.Request.Context(), pathParams.ID, req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: &Success{Message: "success"},
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
// @Success 200 {object} Response{payload=string} "category"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/categories/{id} [delete]
func (h *handler) DeleteCategory(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	err := h.usecases.Admin().Categories().DeleteCategory(c.Request.Context(), pathParams.ID)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: &Success{Message: "success"},
	})
}

// GetFoods get foods.
// @Tags foods
// @Summary Get foods
// @Description Get all foods
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} Response{payload=dto.Foods} "foods"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/foods [get]
func (h *handler) GetFoods(c *gin.Context) {

	resp, err := h.usecases.Admin().Foods().GetFoods(c.Request.Context())
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
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
// @Param request body dto.Food true "Food"
// @Success 200 {object} Response{payload=string} "food"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/foods [post]
func (h *handler) CreateFood(c *gin.Context) {
	var req dto.Food

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}
	log.Println("req", req)
	resp, err := h.usecases.Admin().Foods().CreateFood(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
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
// @Param request body dto.Food true "Food"
// @Success 200 {object} Response{payload=string} "food"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/foods/{id} [put]
func (h *handler) UpdateFood(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	var req dto.Food

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().Foods().UpdateFood(c.Request.Context(), pathParams.ID, req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: &Success{Message: "success"},
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
// @Success 200 {object} Response{payload=string} "food"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/foods/{id} [delete]
func (h *handler) DeleteFood(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	err := h.usecases.Admin().Foods().DeleteFood(c.Request.Context(), pathParams.ID)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: &Success{Message: "success"},
	})
}

// GetTranslates get translates.
// @Tags translates
// @Summary Get translates
// @Description Get all translates
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} Response{payload=[]dto.Translate} "translates"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/translates [get]
func (h *handler) GetTranslates(c *gin.Context) {

	resp, err := h.usecases.Admin().Translates().GetTranslates(c.Request.Context())
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: resp,
	})
}

// CreateTranslates create translate.
// @Tags translates
// @Summary Create translate
// @Description Create translate
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body dto.Translate true "Translate"
// @Success 200 {object} Response{payload=string} "translate"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/translates [post]
func (h *handler) CreateTranslates(c *gin.Context) {
	var req dto.Translate

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().Translates().CreateTranslate(c.Request.Context(), req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: "success",
	})
}

// UpdateTranslate update translate.
// @Tags translates
// @Summary Update translate
// @Description Update translate
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param request body dto.Translate true "Translate"
// @Success 200 {object} Response{payload=string} "translate"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/translates/{id} [put]
func (h *handler) UpdateTranslate(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	var req dto.Translate

	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
		return
	}

	err := h.usecases.Admin().Translates().UpdateTranslate(c.Request.Context(), pathParams.ID, req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: &Success{Message: "success"},
	})
}

// GetTranslateById get translate by id.
// @Tags translates
// @Summary Get translate by id
// @Description Get translate by id
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id" format(uuid)
// @Success 200 {object} Response{payload=dto.Translate} "translate"
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 403 {object} Response
// @Failure 500 {object} Response
// @Router /admin/data/translates/{id} [get]
func (h *handler) GetTranslateById(c *gin.Context) {
	var pathParams IdPathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		errResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to parse path: %v", err))
		return
	}

	resp, err := h.usecases.Admin().Translates().GetTranslateById(c.Request.Context(), pathParams.ID)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Payload: resp,
	})
}
