package factories

import (
	"github.com/adampresley/logrusviewer/pkg/logfiles"
	"github.com/adampresley/logrusviewer/pkg/logfiletypes"
	"github.com/adampresley/logrusviewer/pkg/parsers"
)

/*
IServiceFactory describes a component that creates services
*/
type IServiceFactory interface {
	UploadService() *logfiles.UploadService
	Parser(logFileType logfiletypes.LogFileType) parsers.IParser
}
