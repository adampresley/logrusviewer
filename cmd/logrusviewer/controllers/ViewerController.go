package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/adampresley/logrusviewer/pkg/logfiles"
	"github.com/adampresley/logrusviewer/pkg/logfiletypes"

	"github.com/adampresley/logrusviewer/pkg/factories"
	"github.com/adampresley/logrusviewer/pkg/ui"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

/*
ViewerController provides methods for viewing and working with
logrus entries
*/
type ViewerController struct {
	logger         *logrus.Entry
	serviceFactory factories.IServiceFactory
}

/*
NewViewerController creates a new controller
*/
func NewViewerController(logger *logrus.Entry, serviceFactory factories.IServiceFactory) *ViewerController {
	return &ViewerController{
		logger:         logger,
		serviceFactory: serviceFactory,
	}
}

/*
PostSelectLogFile processes a user's selection of a log file. Basically
this just validates the file exists, then redirects to the view entries
page with a URL variable attached

	POST: /postselectlogfile
*/
func (c *ViewerController) PostSelectLogFile(ctx echo.Context) error {
	var err error
	var fileHeader *multipart.FileHeader
	var bytesWritten int64

	uploadService := c.serviceFactory.UploadService()

	/*
	 * Read the upload file
	 */
	if fileHeader, err = ctx.FormFile("logFile"); err != nil {
		c.logger.Errorf("Error parsing form file in PostSelectLogFile: %s", err.Error())
		return ctx.Redirect(302, "/selectlogfile?errormessage=Error processing form file")
	}

	if bytesWritten, err = uploadService.Upload(fileHeader); err != nil {
		c.logger.Errorf("Error uploading file in PostSelectLogFile: %s", err.Error())
		return ctx.Redirect(302, "/selectlogfile?errormessage=Error uploading file. Check logs for more information")
	}

	c.logger.Infof("%d bytes written to ./uploads/%s", bytesWritten, fileHeader.Filename)
	return ctx.Redirect(302, "/?logfile="+fileHeader.Filename)
}

/*
SelectLogFile presents the user with a form to select a log file
to parse

	GET: /selectlogfile
*/
func (c *ViewerController) SelectLogFile(ctx echo.Context) error {
	viewState := ui.NewViewState("Select Log File")

	if ctx.QueryParam("errormessage") != "" {
		viewState.IsError = true
		viewState.Message = ctx.QueryParam("errormessage")
	}

	return ctx.Render(http.StatusOK, "mainLayout:selectLogFile", viewState)
}

/*
ViewEntries renders the view logrus entries page. This page shows
all entries for the selected log file.

	GET: /
*/
func (c *ViewerController) ViewEntries(ctx echo.Context) error {
	var err error
	var f io.Reader

	viewState := ui.NewViewState("View Entries")

	if ctx.QueryParam("logfile") != "" {
		viewState.LogFile = ctx.QueryParam("logfile")
		parser := c.serviceFactory.Parser(logfiletypes.LogrusJSON)

		if f, err = parser.Open(filepath.Join(logfiles.UPLOADS_FOLDER, viewState.LogFile)); err != nil {
			c.logger.Errorf("Error opening file %s with parser in ViewEntries: %s", viewState.LogFile, err.Error())

			viewState.IsError = true
			viewState.Message = fmt.Sprintf("Error opening log file %s", viewState.LogFile)
		} else {
			if viewState.LogEntries, err = parser.Parse(f); err != nil {
				c.logger.Errorf("Error parsing log file %s in ViewEntries: %s", viewState.LogFile, err.Error())
				viewState.IsError = true
				viewState.Message = fmt.Sprintf("Error parsing %s. See log for more details", viewState.LogFile)
			}
		}
	}

	return ctx.Render(http.StatusOK, "mainLayout:viewer", viewState)
}

/*
ViewEntry gets an individual line.

	GET: /entry?logfile=file&lineNumber=1
*/
func (c *ViewerController) ViewEntry(ctx echo.Context) error {
	var err error
	var f io.Reader
	var logEntries logfiles.ParsedLogFile

	fileName := ctx.QueryParam("logfile")
	parser := c.serviceFactory.Parser(logfiletypes.LogrusJSON)

	lineNumber, _ := strconv.Atoi(ctx.QueryParam("lineNumber"))

	if f, err = parser.Open(filepath.Join(logfiles.UPLOADS_FOLDER, fileName)); err != nil {
		c.logger.Errorf("Error opening file %s with parser in ViewEntry: %s", fileName, err.Error())
		return ctx.String(http.StatusInternalServerError, "Error opening log file")
	}

	if logEntries, err = parser.Parse(f); err != nil {
		c.logger.Errorf("Error parsing log file %s in ViewEntry: %s", fileName, err.Error())
		return ctx.String(http.StatusInternalServerError, "Error parsing log file")
	}

	return ctx.JSON(http.StatusOK, logEntries.GetLine(lineNumber))
}
