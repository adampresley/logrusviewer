//go:generate esc -o ./www/www.go -pkg www -ignore DS_Store|README\.md|LICENSE|www\.go -prefix /www/ ./www

package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/adampresley/logrusviewer/cmd/logrusviewer/controllers"
	"github.com/adampresley/logrusviewer/cmd/logrusviewer/www"
	"github.com/adampresley/logrusviewer/pkg/factories"
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
var serviceFactory *factories.ServiceFactory

var viewerController *controllers.ViewerController

func main() {
	var err error
	var handlers *echo.Echo

	flag.Parse()

	logger = logging.GetLogger(*logLevel, "Logrus Viewer")
	logger.Infof("Starting server v%s", SERVER_VERSION)

	renderer = ui.NewTemplateRenderer(DEBUG_ASSETS)
	setupFactory()
	setupControllers()

	handlers = echo.New()
	handlers.Renderer = renderer
	handlers.Logger = kit.MiddlewareLogger{logger.Logger}
	handlers.HideBanner = true
	handlers.Use(kit.HookWithExisting(logger))
	handlers.Use(middleware.CORS())

	if *logLevel == "debug" {
		handlers.Debug = true
	}

	assetHandler := http.FileServer(www.FS(DEBUG_ASSETS))
	handlers.GET("/www/*", echo.WrapHandler(assetHandler))

	handlers.GET("/", viewerController.ViewEntries)
	handlers.GET("/selectlogfile", viewerController.SelectLogFile)
	handlers.POST("/postselectlogfile", viewerController.PostSelectLogFile)

	go func() {
		var err error

		if err = handlers.Start(*host); err != nil {
			logger.Infof("Shutting down the server...")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = handlers.Shutdown(ctx); err != nil {
		logger.Errorf("There was an error shutting down the server - %s", err.Error())
	}
}

func setupFactory() {
	serviceFactory = factories.NewServiceFactory(*logLevel)
}

func setupControllers() {
	viewerController = controllers.NewViewerController(logging.GetLogger(*logLevel, "Viewer Controller"), serviceFactory)
}
