package parsers

import (
	"io"

	"github.com/adampresley/logrusviewer/pkg/logfiles"
)

/*
IParser describes an interface for implementers to open, parse
and work with log files
*/
type IParser interface {
	Open(fileName string) (io.Reader, error)
	Parse(contents io.Reader) (logfiles.ParsedLogFile, error)
}
