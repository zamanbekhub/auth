package http

import (
	"auth/internal/common/middleware"
	"auth/internal/config"
	v1 "auth/internal/delivery/http/v1"
	"auth/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
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
	//p := ginprom.New(
	//	ginprom.Engine(app),
	//	ginprom.Subsystem("gin"),
	//	ginprom.Path("/metrics"),
	//)
	//p.AddCustomCounter("custom", "Some help text to provide", []string{"label"})
	//p.IncrementCounterValue("custom", []string{"1"})
	//p.IncrementCounterValue("custom", []string{"1"})
	//app.Use(p.Instrument())

	app.Use(middleware.Cors())
	//app.Use(middleware.JwtAuthMiddleware())
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
