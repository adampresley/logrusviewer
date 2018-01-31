package controllers

import (
	"net/http"

	"github.com/adampresley/logrusviewer/pkg/ui"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

/*
ViewerController provides methods for viewing and working with
logrus entries
*/
type ViewerController struct {
	logger *logrus.Entry
}

/*
NewViewerController creates a new controller
*/
func NewViewerController(logger *logrus.Entry) *ViewerController {
	return &ViewerController{
		logger: logger,
	}
}

/*
ViewEntries renders the view logrus entries page. This page shows
all entries for the selected log file.

	GET: /viewer
*/
func (c *ViewerController) ViewEntries(ctx echo.Context) error {
	viewState := &ui.ViewState{
		Title: "View Entries",
	}

	return ctx.Render(http.StatusOK, "mainLayout:viewer", viewState)
}

/*
SelectLogFile presents the user with a form to select a log file
to parse

	GET: /selectlogfile
*/
func (c *ViewerController) SelectLogFile(ctx echo.Context) error {
	viewState := &ui.ViewState{
		Title: "Select Log File",
	}

	return ctx.Render(http.StatusOK, "mainLayout:selectLogFile", viewState)
}
