package http

import (
	"coffe-life/config"
	"coffe-life/docs"
	v1 "coffe-life/internal/controller/http/v1"
	"coffe-life/internal/interfaces"
	"coffe-life/pkg/middlewares/ginlogr"
	customGinSwagger "coffe-life/pkg/middlewares/ginswagger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Dependencies struct {
	Cfg      *config.Config
	Logger   logr.Logger
	Usecases interfaces.Usecases
}

func NewRouter(deps Dependencies) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(ginlogr.Logger(deps.Logger, "/healthz"))
	router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	swaggerHandlerV1 := customGinSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER", ginSwagger.InstanceName("v1"))
	v1.Register(router, v1.Dependencies{
		Cfg:      deps.Cfg,
		Logger:   deps.Logger,
		Usecases: deps.Usecases,
	})
	docs.SwaggerInfov1.BasePath = "/api/v1"

	router.GET("/swagger/v1/*any", swaggerHandlerV1)
	return router
}
