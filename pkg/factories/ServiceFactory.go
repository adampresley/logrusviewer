package factories

import (
	"github.com/adampresley/logrusviewer/pkg/logfiles"
	"github.com/adampresley/logrusviewer/pkg/logfiletypes"
	"github.com/adampresley/logrusviewer/pkg/logging"
	"github.com/adampresley/logrusviewer/pkg/parsers"
)

/*
ServiceFactory describes a component that creates services
*/
type ServiceFactory struct {
	logLevel string
}

/*
NewServiceFactory creates a new factory
*/
func NewServiceFactory(logLevel string) *ServiceFactory {
	return &ServiceFactory{
		logLevel: logLevel,
	}
}

/*
UploadService creates a new service
*/
func (f *ServiceFactory) UploadService() *logfiles.UploadService {
	return logfiles.NewUploadService(logging.GetLogger(f.logLevel, "Upload Service"))
}

/*
Parser returns a log file parser based on the provided type
*/
func (f *ServiceFactory) Parser(logFileType logfiletypes.LogFileType) parsers.IParser {
	switch logFileType {
	case logfiletypes.LogrusJSON:
		return parsers.NewLogrusJSONParser()

	default:
		return nil
	}
}
