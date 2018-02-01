package parsers

import (
	"bufio"
	"io"
	"os"

	"github.com/adampresley/logrusviewer/pkg/logfiles"
	"github.com/pkg/errors"
)

/*
LogrusJSONParser is an implementation of IParser which can
process Logrus JSON log files
*/
type LogrusJSONParser struct {
}

/*
NewLogrusJSONParser creates a new struct
*/
func NewLogrusJSONParser() *LogrusJSONParser {
	return &LogrusJSONParser{}
}

/*
Open opens a reader to the supplied log file
*/
func (p *LogrusJSONParser) Open(fileName string) (io.Reader, error) {
	var err error
	var f *os.File

	if f, err = os.Open(fileName); err != nil {
		return nil, errors.Wrapf(ErrOpeningFile, "Error opening log file in LogrusJSONParser")
	}

	return f, nil
}

/*
Parse processes the contents of the log file
*/
func (p *LogrusJSONParser) Parse(contents io.Reader) (logfiles.ParsedLogFile, error) {
	var err error
	var scanner *bufio.Scanner
	var logEntry logfiles.ParsedLogEntry

	result := logfiles.NewParsedLogFile()
	scanner = bufio.NewScanner(contents)
	line := 0

	for scanner.Scan() {
		line++

		if logEntry, err = logfiles.NewParsedLogEntryFromJSON(scanner.Bytes()); err != nil {
			return result, errors.Wrapf(ErrInvalidFormat, "Error parsing line %d from log file in LogrusJSONParser", line)
		}

		if !p.isValidEntry(logEntry) {
			return result, errors.Wrapf(ErrInvalidFormat, "Line %d has an invalid format", line)
		}

		result = append(result, logEntry)
	}

	return result, nil
}

func (p *LogrusJSONParser) isValidEntry(logEntry logfiles.ParsedLogEntry) bool {
	var ok bool

	if _, ok = logEntry["level"]; !ok {
		return false
	}

	if _, ok = logEntry["msg"]; !ok {
		return false
	}

	if _, ok = logEntry["time"]; !ok {
		return false
	}

	return true
}
