package app

import (
	"strings"
	"time"

	"github.com/AleksK1NG/go-cqrs-eventsourcing/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

func (a *app) runHttpServer() error {
	a.mapRoutes()

	a.echo.Server.ReadTimeout = readTimeout
	a.echo.Server.WriteTimeout = writeTimeout
	a.echo.Server.MaxHeaderBytes = maxHeaderBytes

	return a.echo.Start(a.cfg.Http.Port)
}

func (a *app) mapRoutes() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "EventSourcing CQRS Microservice"
	docs.SwaggerInfo.Description = "EventSourcing CQRS Microservice."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	a.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	a.echo.Use(a.middlewareManager.RequestLoggerMiddleware)
	a.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: false,
		DisableStackAll:   false,
	}))
	a.echo.Use(middleware.RequestID())
	a.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	a.echo.Use(middleware.BodyLimit(bodyLimit))
}
