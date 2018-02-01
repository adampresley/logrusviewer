package parsers

import (
	"fmt"
)

// ErrInvalidFormat is used when a log file format is invalid
var ErrInvalidFormat = fmt.Errorf("Invalid log file format")

// ErrOpeningFile is used when there is an error opening the log file
var ErrOpeningFile = fmt.Errorf("Unable to open log file")
