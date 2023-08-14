package main

import (
	"context"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/brpaz/echozap"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"smart-door-opener/handler"
	"time"
)

var CLI struct {
	Server struct {
		IFTTTServer string `help:"IFTTT server for webhook." env:"SERVER_IFTTT_SERVER" default:"maker.ifttt.com" required:""`
		IFTTTServerKey string `help:"IFTTT server key for webhook." env:"SERVER_IFTTT_SERVER_KEY" required:""`
		IFTTTWebHookEventName string `help:"IFTTT event name to trigger door opening." env:"SERVER_IFTTT_WEB_HOOK_EVENT_NAME" required:""`

		AccessSecretCode string `help:"Secret code used to access the opening page." env:"SERVER_ACCESS_SECRET_CODE" required:""`
	} `cmd:"" help:"Start the server."`
}

func main() {
	// logger
	zapLogger, _ := zap.NewProduction()

	// command loading
	ctxCmd := kong.Parse(&CLI,
		kong.Name("server"),
		kong.Description("A web server app."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	switch ctxCmd.Command() {
	case "server":
	default:
		zapLogger.Fatal(fmt.Sprintf("Unkown command: %v", ctxCmd.Command()))
	}

	e := echo.New()

	// load middlewares
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	e.Use(echozap.ZapLogger(zapLogger))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Static("/static", "static")

	// Set Renderer
	e.Renderer = echoview.Default()

	// Handlers
	doorHandler := handler.NewHandlerDoor(
		CLI.Server.AccessSecretCode,
		CLI.Server.IFTTTServer,
		CLI.Server.IFTTTServerKey,
		CLI.Server.IFTTTWebHookEventName,
		)

	e.GET("/", doorHandler.GetOpenDoor)
	e.GET("/:accessCode", doorHandler.GetOpenDoor)
	e.POST("/:accessCode", doorHandler.PostOpenDoor)
	e.GET("/_health", func(c echo.Context) error {
		return c.String(http.StatusOK, "RUNNING")
	})

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
