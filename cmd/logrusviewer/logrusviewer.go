//go:generate esc -o ./www/www.go -pkg www -ignore DS_Store|README\.md|LICENSE|www\.go -prefix /www/ ./www

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/adampresley/logrusviewer/cmd/logrusviewer/controllers"
	"github.com/adampresley/logrusviewer/pkg/logging"
	"github.com/adampresley/logrusviewer/pkg/ui"
	"github.com/sirupsen/logrus"
	kit "gitlab.appninjas.biz/appninjas/kit/logging"
)

const (
	SERVER_VERSION string = "0.1.0"

	// Set to true while developing
	DEBUG_ASSETS bool = true
)

var logger *logrus.Entry
var renderer *ui.TemplateRenderer

var viewerController *controllers.ViewerController

func main() {
	var err error
	var handlers *echo.Echo

	flag.Parse()

	logger = logging.GetLogger(*logLevel, "Logrus Viewer")
	logger.Infof("Starting server v%s", SERVER_VERSION)

	renderer = ui.NewTemplateRenderer(DEBUG_ASSETS)
	setupControllers()

	handlers = echo.New()
	handlers.Logger = kit.MiddlewareLogger{logger.Logger}
	handlers.HideBanner = true
	handlers.Use(kit.HookWithExisting(logger))
	handlers.Use(middleware.CORS())

	if *logLevel == "debug" {
		handlers.Debug = true
	}

	handlers.GET("/", viewerController.ViewEntries)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = handlers.Shutdown(ctx); err != nil {
		logger.Errorf("There was an error shutting down the server - %s", err.Error())
	}
}

func setupControllers() {
	viewerController = controllers.NewViewerController(logging.GetLogger(*logLevel, "Viewer Controller"))
}
