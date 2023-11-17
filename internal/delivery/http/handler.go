package http

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"service/internal/common/middleware"
	"service/internal/config"
	v1 "service/internal/delivery/http/v1"
	"service/internal/service"
)

type Handler struct {
	logger   *log.Logger
	services *service.Services
	baseUrl  string
}

func NewHandlerDelivery(
	logger *log.Logger,
	services *service.Services,
	baseUrl string,
) *Handler {
	return &Handler{
		logger:   logger,
		services: services,
		baseUrl:  baseUrl,
	}
}

func (h *Handler) Init(cfg *config.Config) (*gin.Engine, error) {
	app := gin.New()
	p := ginprom.New(
		ginprom.Engine(app),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	app.Use(middleware.Cors(), p.Instrument())
	h.initAPI(app)
	return app, nil
}

func (h *Handler) initAPI(router *gin.Engine) {
	baseUrl := router.Group(h.baseUrl)

	router.GET(h.baseUrl+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	handlerV1 := v1.NewHandler(h.services)
	api := baseUrl.Group("/api")
	{
		handlerV1.Init(api)
	}
}
