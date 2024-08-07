package v1

import (
	"coffe-life/config"
	"coffe-life/internal/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
)

type handler struct {
	cfg      *config.Config
	group    *gin.RouterGroup
	log      logr.Logger
	usecases interfaces.Usecases
}

type Dependencies struct {
	Cfg      *config.Config
	Logger   logr.Logger
	Usecases interfaces.Usecases
}

// @title Coffe-Life API
// @version	1.0
// @description	API for processing clients data
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Register(router *gin.Engine, deps Dependencies) {
	h := &handler{
		cfg:      deps.Cfg,
		log:      deps.Logger,
		group:    router.Group("api/v1"),
		usecases: deps.Usecases,
	}

	h.newCoffeRoutes()
	h.newAdminPanel()
}
